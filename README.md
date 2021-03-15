# Example-Blog-Server

Basic blog server written in go for CRUD operation

# [Database Models](https://github.com/KushagraMehta/Example-Blog-Server/blob/main/model/README.md)

# To-do

- [x] Write Database schema
- [x] Create Database Trigger for like_count(posts), total_post(tags)
- [x] Connect Database with GO
- [x] Create Model Function in Golang
- [x] Refactoring Code
- [x] Write Unit Test of model
- [ ] Write Mock Test on model
- [ ] Design All API end-points
- [ ] Change Error formate
- [ ] Implement API in REST
  - [ ] Write open API end-points
  - [ ] Implement JWT Authentication
  - [ ] Write Authentication Middleware
- [ ] Write Docker File

# Folder Structure

```
.
├── cmd                  main applications of the project
│   └── server           the API server application
├── pkg                  public library code
│   ├── auth
│   ├── error
│   ├── middleware      access log middleware
│   ├── model
│   ├── response
│   └── util
```

[📂Folder structure Inspiration](https://github.com/qiangxue/go-rest-api)
