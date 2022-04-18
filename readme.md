Simple CRUD

http endpoint: localhost:8081/todo
You can create, update, read and delete pages.

Also you can edit save method. I have i methods: StructMethod and FileMethod
Please, switch in services/database/base_database.go

set:
```go
    defaultMethod DBMethods = &StructMethod{}
```
or

```go
    defaultMethod DBMethods = &FileMethod{}
```

If setup FileMethod:
- Create new file in storage dir: page.json and save data about pages.

If setup StrctMethod:
- Save data in memory 


install guide:
```shell
    go run .
```