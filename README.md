# Example-Blog-Server

---

Basic practice blog server written in go for CRUD operation

# Getting started

Project requires having Go 1.16 or above and postgreSQL. Once you clone(or go get) you need to configure the following:

1. Change the DotEnv file according to your envirnment.
2. Install Task runner `go get -u github.com/go-task/task/v3/cmd/task`
3. `task migration` for initializing database.
4. `task test` for testing project working.
5. `task run` to start project.

---

## More About project

- ### [Database Relation](https://github.com/KushagraMehta/Example-Blog-Server/blob/main/pkg/model/README.md)

- ### [Database Methods](https://pkg.go.dev/github.com/KushagraMehta/Example-Blog-Server/pkg/model)

- ### Folder Structure

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

---

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
