
-- Create Blog DataBase
CREATE DATABASE BLOG;

-- Select Blog database

\c blog;

-- Create User Table
 CREATE TABLE IF NOT EXISTS USERS(
 ID              SERIAL   PRIMARY KEY,
 USERNAME        VARCHAR(20) NOT NULL UNIQUE,
 EMAIL           VARCHAR(20) NOT NULL UNIQUE,
 PASSWORD_HASHED VARCHAR(50) NOT NULL,
 CREATED_ON      timestamp DEFAULT CURRENT_TIMESTAMP,
 UPDATED_ON      timestamp DEFAULT CURRENT_TIMESTAMP,
 LAST_LOGIN      timestamp DEFAULT CURRENT_TIMESTAMP
 );

-- Create Posts table

 CREATE TABLE IF NOT EXISTS POSTS(
 ID              SERIAL   PRIMARY KEY,
 AUTHOR_ID       INTEGER NOT NULL,
 TITLE           VARCHAR(100) NOT NULL,
 SUMMARY         VARCHAR(500),
 PUBLISHED       BOOLEAN DEFAULT FALSE,
 CREATED_ON      timestamp DEFAULT CURRENT_TIMESTAMP,
 UPDATED_ON      timestamp DEFAULT CURRENT_TIMESTAMP,
 LIKE_COUNT      INTEGER DEFAULT 0,
 VIEWS           INTEGER DEFAULT 0,
 BODY            TEXT,
 FOREIGN KEY (AUTHOR_ID)
     REFERENCES USERS (ID)
     ON DELETE NO ACTION
     ON UPDATE NO ACTION

 );

-- Create Comments Table
 CREATE TABLE IF NOT EXISTS COMMENTS(
    ID              SERIAL   PRIMARY KEY,
    AUTHOR_ID       INTEGER NOT NULL,
    CREATED_ON      timestamp DEFAULT CURRENT_TIMESTAMP,
    UPDATED_ON      timestamp DEFAULT CURRENT_TIMESTAMP,
    BODY            TEXT,
    FOREIGN KEY (AUTHOR_ID)
    REFERENCES USERS (ID)
    ON DELETE NO ACTION
    );

-- Create Tags Table

    CREATE TABLE IF NOT EXISTS TAGS(
    ID              SERIAL   PRIMARY KEY,
    TITLE           VARCHAR(100) NOT NULL,
    SUMMARY         VARCHAR(500),
    TOTAL_POST      INTEGER DEFAULT 0
    );

-- Create Post Tags Table

    CREATE TABLE IF NOT EXISTS POST_TAGS(
    POST_ID     INTEGER NOT NULL,
    TAG_ID      INTEGER NOT NULL,
    PRIMARY KEY (POST_ID,TAG_ID),
    FOREIGN KEY (POST_ID)
        REFERENCES POSTS (ID)
        ON DELETE CASCADE,
    FOREIGN KEY (TAG_ID)
        REFERENCES TAGS (ID)
        ON DELETE CASCADE
    );

-- Create Post Comments Table

    CREATE TABLE IF NOT EXISTS POST_COMMENTS(
    COMMENT_ID      INTEGER NOT NULL,
    POST_ID         INTEGER NOT NULL,
    PRIMARY KEY (COMMENT_ID,POST_ID),
    FOREIGN KEY (COMMENT_ID)
        REFERENCES COMMENTS (ID)
        ON DELETE CASCADE,
    FOREIGN KEY (POST_ID)
        REFERENCES POSTS (ID)
        ON DELETE CASCADE
    );

-- Create Post Likes Table

    CREATE TABLE IF NOT EXISTS POST_LIKES(
    AUTHOR_ID       INTEGER NOT NULL,
    POST_ID         INTEGER NOT NULL,
    PRIMARY KEY (AUTHOR_ID,POST_ID),
    FOREIGN KEY (AUTHOR_ID)
        REFERENCES USERS (ID)
        ON DELETE NO ACTION,
    FOREIGN KEY (POST_ID)
        REFERENCES POSTS (ID)
        ON DELETE CASCADE
    );