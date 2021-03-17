package model

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// Hash Funtion do the Hashing of given String.
func hash(password string) (string, error) {
	hashedValue, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedValue), err
}

//Validate Function check if the user Data data is filled or not. For smooth database entry.
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "signup":
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email ")
		}
		if u.UserName == "" {
			return errors.New("required username ")
		}
		if u.Email != "" {
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				return errors.New("invalid email")
			}
		}
		return nil
	case "update":
		if u.ID == 0 {
			return errors.New("id is required")
		}
		return nil
	default:
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" && u.UserName == "" {
			return errors.New("required email or username")
		}
		if u.Email != "" {
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				return errors.New("invalid email")
			}
		}
		return nil
	}
}

// init will initiate user object value.
func (u *User) Init(username, email, password string) error {
	var err error
	u.UserName = username
	u.Email = email
	u.Password, err = hash(password)
	return err
}

// SignUp will save user detail into database. REQUIRE: User Object init.
func (u *User) SignUp(db *pgxpool.Pool) (int, error) {
	if err := u.Validate("signup"); err != nil {
		return 0, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.QueryRow(ctx, "INSERT INTO users(username,email,password_hashed) VALUES ($1,$2,$3) RETURNING id, created_on,updated_on,last_login;", u.UserName, u.Email, u.Password).Scan(&u.ID, &u.CreatedOn, &u.UpdatedOn, &u.LastLogin); err != nil {
		return 0, err
	}
	return u.ID, nil
}

// Login will check the user detail and send the UID  REQUIRE: username|email, Password.
func (u *User) Login(db *pgxpool.Pool) (int, error) {

	if err := u.Validate(""); err != nil {
		return 0, err
	}

	if u.UserName != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := db.QueryRow(ctx, "SELECT ID,email,created_on,updated_on,last_login FROM users WHERE username=$1 AND password_hashed=$2;", u.UserName, u.Password).Scan(&u.ID, &u.Email, &u.CreatedOn, &u.UpdatedOn, &u.LastLogin); err != nil {
			return 0, err
		}
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := db.QueryRow(ctx, "SELECT ID FROM users WHERE email=$1 AND password_hashed=$2;", u.Email, u.Password).Scan(&u.ID); err != nil {
			return 0, err
		}

	}
	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := db.Exec(ctx2, "UPDATE users SET last_login=current_timestamp WHERE ID=$1;", u.ID); err != nil {
		return 0, err
	}

	return u.ID, nil

}

// PutNewPassword will update the password. REQUIRE: ID
func (u *User) PutNewPassword(db *pgxpool.Pool, newPassword string) error {
	var err error
	if err := u.Validate("update"); err != nil {
		return err
	}
	if u.Password, err = hash(newPassword); err != nil {
		return err
	}
	if u.UserName != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if _, err := db.Exec(ctx, "UPDATE users SET password_hashed=$1, UPDATED_ON=current_timestamp WHERE id=$2;", u.Password, u.ID); err != nil {
			return err
		}
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if _, err := db.Exec(ctx, "UPDATE users SET password_hashed=$1,UPDATED_ON=current_timestamp WHERE id=$2;", u.Password, u.ID); err != nil {
			return err
		}
	}
	return nil
}

// GetLikedPost return the array of postID liked by user. REQUIRE: ID
func (u *User) GetLikedPost(db *pgxpool.Pool) ([]int, error) {

	if err := u.Validate("update"); err != nil {
		return []int{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result []int
	if rows, err := db.Query(ctx, "select post_id from post_likes where author_id=$1;", u.ID); err != nil {
		return []int{}, err
	} else {
		defer rows.Close()
		var tmp int
		for rows.Next() {
			rows.Scan(&tmp)
			result = append(result, tmp)
		}

		if rows.Err() != nil {
			return []int{}, err
		}
	}
	return result, nil
}

// PatchLike will can put like/Remove like from a post. REQUIRE: ID
func (u *User) PatchLike(db *pgxpool.Pool, postID int) error {
	if err := u.Validate("update"); err != nil {
		return err
	}
	var doesLike int
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.QueryRow(ctx, " select COUNT(*) as count from post_likes where author_id=$1 AND post_id=$2;", u.ID, postID).Scan(&doesLike); err != nil {
		return err
	}

	if doesLike == 0 {
		if _, err := db.Exec(ctx, "INSERT INTO POST_LIKES(AUTHOR_ID,POST_ID) VALUES($1,$2);", u.ID, postID); err != nil {
			return err
		}
	} else {
		//If User likes it then dis-like the post
		if _, err := db.Exec(ctx, "DELETE FROM POST_LIKES WHERE AUTHOR_ID=$1 AND POST_ID=$2;", u.ID, postID); err != nil {
			return err
		}
	}
	return nil
}

// FindUserByID will find a user with specific UID
func FindUserByID(db *pgxpool.Pool, uid int) (*User, error) {
	newUser := &User{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.QueryRow(ctx, "SELECT username,email,last_login FROM users WHERE ID=$1;", uid).Scan(&newUser.UserName, &newUser.Email, &newUser.LastLogin); err != nil {
		return &User{}, err
	}
	return newUser, nil
}
