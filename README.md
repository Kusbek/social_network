# My funny attempt to make a social network
## Prerequisites
This is the next iteration of those projects - [real time forum](https://github.com/Kusbek/Div-01/tree/master/real-time-forum) and [net cat](https://github.com/Kusbek/Div-01/tree/master/net-cat). Real time forum is a SPA application without any web framework, pure vanilla javascript with backend written on golang. Backend also does not use any framework. The main purpose of the project was learning goroutines and its practical application in making chats. This project was a mess and was made after workdays. As an outcome, from real time forum I learnt:
*   authorization, authentication
*   making posts, comments to posts
*   private messaging, websockets

## About this project
What I'm planning to realise in this project
*   frontend framework - Vuejs
*   authorization, authentication
*   Profile page with following and follower users, groups and group events
*   avatar, post images upload
*   webservers
*   making posts, comments to posts
*   private messaging, group messaging, websockets

#### How to run:
no dockerfiles or docker-composes yet
```bash
caddy run
```
```bash
go run ./pkg/db/sqlite/cmd/main.go
```
```bash
go run server.go
```

enter username and pass: kusbek and 123456