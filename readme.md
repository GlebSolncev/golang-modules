![alt text](https://sagacitysoftware.co.in/wp-content/uploads/2020/07/RPA3.jpg)
Simple CRUD

http endpoint: localhost:8081
You can create, update, read and delete todo items.
You can update method for save data (StructMethod and FileMethod). Path pkg/database/base_database.go

set:
```text
    DefaultMethod DBMethods = &StructMethod{}
```
or

```text
    DefaultMethod DBMethods = &FileMethod{}
```

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

