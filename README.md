## curl

- getUserById
```
$ curl localhost:8080/user/${user_id}
```

- store user
```
$ curl -X POST -H "Content-Type: application/json" -d '{"Name":"test", "Email":"test_email@gmail.com","Password":"test_password"}' localhost:8080/user
```

- update user
```
$ curl -X POST -H "Content-Type: application/json" -d '{"Name":"test", "Email":"test_email@gmail.com","Password":"test_password"}' localhost:8080/user/1
```

- delete user
```
$ curl -X DELETE  localhost:8080/user/1
```


## Makefile
- 新しいSQLファイルを作成する
```
make newsql NAME=< 作成するファイルの名前 CreateChatなど >
```

- マイグレーションをローカルからやる
```
make migrate MIGRATE=up
```


## command
- volume-data削除
```
docker volume rm $(docker volume ls -qf dangling=true)
```