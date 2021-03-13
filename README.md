# Example-Blog-Server

Basic blog server written in go for CRUD operation

# [Database Models](./model/readme.md)

# To-do

- [x] Write Database schema
- [x] Create Database Trigger for like_count(posts), total_post(tags)
- [x] Connect Database with GO
- [x] Create Model Function in Golang
- [x] Refactoring Code
- [ ] Write Unit Test of model
- [ ] Design All API end-points
- [ ] Implement API in REST
  - [ ] Write open API end-points
  - [ ] Implement JWT Authentication
  - [ ] Write Authentication Middleware
- [ ] Write Docker File

# TMP API end-points

- Post delete, delete its comment
- Tag
  - Create Tag
  - Add tag on Post
  - Delete tag from Post
  - Get tags (Sorted)
  - GET Tag detail By id
  - GET Tag Post(Top By given data<length of tag_post)
