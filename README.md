# Example-Blog-Server

Basic blog server written in go for CRUD operation

# [Database Models](./model/readme.md)

# To-do

- [x] Write Database schema
- [x] Create Database Trigger for like_count(posts), total_post(tags)
- [ ] Create Model Function in Golang
- [ ] Refactoring Code
- [ ] Write Unit Test of model
- [ ] Connect Database with GO
- [ ] Design All API end-points
- [ ] Implement API in REST
  - [ ] Write open API end-points
  - [ ] Implement JWT Authentication
  - [ ] Write Authentication Middleware
- [ ] Write Docker File

# TMP API end-points

- Comment
  - POST a Comment on Post
  - GET Comments By Post id
  - DELETE a Comment on post
- Tag
  - GET Tags(By length)
  - GET Tag detail By id
  - GET Tag Post(Top By given data<length of tag_post)
