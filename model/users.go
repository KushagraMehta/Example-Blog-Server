package model

import (
	"context"
	"errors"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// Hash Funtion do the Hashing of given String
func hash(password string) (string, error) {
	hashedValue, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedValue), err
}

// VerifyPassword Compare the password
func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

//Validate Function check if the user Data data is filled or not. For smooth database entry
func (u *User) Validate() error {
	if u.PasswordHashed == "" {
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

// init will initiate user object value
func (u *User) Init(username, email, password string) error {
	var err error
	u.UserName = username
	u.Email = email
	u.PasswordHashed, err = hash(password)
	return err
}

// SignUp will save user detail into database
func (u *User) SignUp(db *pgxpool.Pool) error {
	if err := u.Validate(); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := db.Exec(ctx, "INSERT INTO users(username,email,password_hashed) VALUES ($1,$2,$3);", u.UserName, u.Email, u.PasswordHashed); err != nil {
		return errors.New("internal error")
	}
	return nil
}

// Login will check the user detail and send the UID
func (u *User) Login(db *pgxpool.Pool) (int64, error) {
	var UID int64
	if err := u.Validate(); err != nil {
		return 0, err
	}

	if u.UserName != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := db.QueryRow(ctx, "SELECT ID FROM users WHERE username=$1 AND password_hashed=$2;", u.UserName, u.PasswordHashed).Scan(&UID); err != nil {
			return 0, err
		}
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := db.QueryRow(ctx, "SELECT ID FROM users WHERE email=$1 AND password_hashed=$2;", u.Email, u.PasswordHashed).Scan(&UID); err != nil {
			return 0, err
		}

	}
	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := db.Exec(ctx2, "UPDATE users SET last_login=current_timestamp WHERE ID=$1;", UID); err != nil {
		return 0, err
	}

	return UID, nil

}

// PutNewPassword will update the password
func (u *User) PutNewPassword(db *pgxpool.Pool, newPassword string) error {
	var err error
	if err := u.Validate(); err != nil {
		return err
	}
	u.PasswordHashed, err = hash(newPassword)
	if err != nil {
		return err
	}
	if u.UserName != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if _, err := db.Exec(ctx, "UPDATE users SET password_hashed=$1 WHERE username=$2;", u.PasswordHashed, u.UserName); err != nil {
			return err
		}
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if _, err := db.Exec(ctx, "UPDATE users SET password_hashed=$1 WHERE email=$2;", u.PasswordHashed, u.Email); err != nil {
			return err
		}

	}
	return nil
}

// FindByID will find a user with specific UID
func (u *User) FindByID(db *pgxpool.Pool, uid int64) (*User, error) {
	newUser := &User{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.QueryRow(ctx, "SELECT  id,username,email,last_login FROM users WHERE ID=$1;", uid).Scan(&newUser.ID, &newUser.UserName, &newUser.Email, &newUser.LastLogin); err != nil {
		return &User{}, err
	}
	return newUser, nil
}

// GetLikedPost return the array of postID liked by user
func (u *User) GetLikedPost(db *pgxpool.Pool) ([]int64, error) {

	if err := u.Validate(); err != nil {
		return []int64{}, err
	}
	var size int64
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if rows, err := db.Query(ctx, "select COUNT(post_id) as count from post_likes where author_id=1;"); err != nil {
		return []int64{}, err
	} else {
		defer rows.Close()
		rows.Next()
		rows.Scan(&size)

		if rows.Err() != nil {
			return []int64{}, err
		}
	}
	result := make([]int64, size)
	if rows, err := db.Query(ctx, "select post_id from post_likes where author_id=1;"); err != nil {
		return []int64{}, err
	} else {
		defer rows.Close()
		i := 0
		for rows.Next() {
			rows.Scan(&result[i])
			i++
		}

		if rows.Err() != nil {
			return []int64{}, err
		}
	}
	return result, nil
}

// PatchLike will can put like/Remove like from a post
func (u *User) PatchLike(db *pgxpool.Pool, postID int64) error {
	var doesLike int
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if rows, err := db.Query(ctx, " select COUNT(*) as count from post_likes where author_id=$1 AND post_id=$2;", u.ID, postID); err != nil {
		return err
	} else {
		defer rows.Close()
		rows.Next()
		rows.Scan(&doesLike)

		if rows.Err() != nil {
			return err
		}
	}

	if doesLike == 0 {
		if _, err := db.Exec(ctx, "INSERT INTO POST_LIKES(AUTHOR_ID,POST_ID) VALUES($1,$2);", u.ID, postID); err != nil {
			return err
		}
	} else {
		//If User likes it then dis-like the post
		if _, err := db.Exec(ctx, "DELETE FROM POST_LIKES WHERE AUTHOR_ID=$1 AND POST_ID=$2);", u.ID, postID); err != nil {
			return err
		}
	}
	return nil
}
