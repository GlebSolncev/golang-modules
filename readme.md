![alt text](https://sagacitysoftware.co.in/wp-content/uploads/2020/07/RPA3.jpg)
Simple golang-modules

http endpoint: localhost:8080
replace: .env.example to .env

## Start Application
1. Update .env.example to .env
2. Enter command
```bash
    make up
```

## Lib list:
- [Labstack echo](https://echo.labstack.com/)
- [ent.](https://entgo.io/)
- [Swagger](https://github.com/swaggo/swag) 
- godotenv
- go-sqlite3


### Show list commands:
```shell
    make help
```

## Dev info:
- Add model(table) -> ent init {{name}}
- Add enum const -> enumer --type={name} -json [..path]
- url://assets/... -> have files with link on site
- live update assets - true for static -> TEMPLATE_LIVE=false
- go func -> on storage 
- Create new app(ms) to parse url. Work with go func, select, channels. He create new file in dir storage and show info about links. 

