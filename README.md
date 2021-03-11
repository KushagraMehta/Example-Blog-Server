# Example-Blog-Server

Basic blog server written in go for CRUD operation

# [Database Models](./model/readme.md)

# To-do

- [x] Write Database schema
- [x] Create Database Trigger for like_count(posts), total_post(tags)
- [ ] Connect Database with GO
- [ ] Design All API end-points
- [ ] Implement API in REST
  - [ ] Write open API end-points
  - [ ] Implement JWT Authentication
  - [ ] Write Authentication Middleware
- [ ] Write Docker File

# TMP API end-points

user
POST SignUp
POST login

PATCH update account -> update users set last_login=current_timestamp where username='abc';
PUT new password token -> update users set password_hashed='password', last_login=current_timestamp where username='abc';

    GET liked posts

    POST a like on post
    DELETE a like on post
    POST a Comment on Post

GET Comments By Post id
DELETE a Comment on post

    GET User Data By id

Post

get/create draft Post -> insert into posts(author_id,title) values(7,'Thats bad');
PATCH Save draft Post
PUT Post Content
PATCH Publish Post

GET finished Post by id

GET finished Posts(By Length)

Tag
GET Tags(By length)
GET Tag detail By id
GET Tag Post(Top By given data<length of tag_post)
