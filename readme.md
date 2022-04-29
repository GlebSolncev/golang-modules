![alt text](https://sagacitysoftware.co.in/wp-content/uploads/2020/07/RPA3.jpg)
Simple golang-modules

http endpoint: localhost:8081
You can create, update, read and delete todo items.
You can update method for save data (StructMethod and FileMethod). Path pkg/database/methods.go Line: 15

Set to Struct(var) as memory:
```text
    defaultMethod contracts.DBMethods = &memory.Method{}
```
Or u can set to File as memory:

```text
    defaultMethod contracts.DBMethods = &memory.Method{}
```
Or u can set to SQLITE as memory
```text
    defaultMethod contracts.DBMethods = &sqlite.Method{}
```

## Lib list:
- [Labstack echo](https://echo.labstack.com/)
- [ent.](https://entgo.io/)
- godotenv
- go-sqlite3
-

If setup FileMethod:
- Create new file in storage dir: todo.json and save data about pages.

If setup StructMethod:
- Save data in memory 


Runing:
```shell
     make run
```

Check info about commands:
```shell
    make help
```

