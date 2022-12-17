# Gokomodo Rest API
This is assignment project for Golang Developer in GoKomodo.

## Prerequisites
- Go *language*
- Mysql 
- GNU Make ( for windows you can install make via [chocolatey](https://community.chocolatey.org/packages/make) )
- Docker & Docker compose ( optional )

## Installation

First clone the project, then open terminal from project directory.

Then you need to generate setup yml file ( env.yml, etc ) from example folder, using command :

```bash
make setup
```

### *Using Docker*
If you're using docker, get golang image from [dockerhub](https://hub.docker.com/layers/library/golang/1.18.3-alpine/images/sha256-725f8fd50191209a4c4a00def1d93c4193c4d0a1c2900139daf8f742480f3367?context=explore) 
for create go image ( 1.18.3-alpine ) .

Then make container for the app using command :

```bash
docker-compose up -d
```

That's it, you're ready to go.

*if you want to change ip host or port, change in env/env.yml and in docker & docker-compose files before create app container*

### *Without Docker*

- First make your gokomodo database in MySQL DBMS.
- Then adjust ip host or port for database config in env.yml
```yaml
database:
  gokomododb:
    mysql:
      host: 172.23.0.3
      port: 3306
      username: root
      password: password
      dbname: gokomodo
      adapter: mysql

```
- Run this project, use this command on project directory:

```bash
make run
```

Program will run at port :8801, if you need to change the port, you can edit in main.go file and change it.

This project has executable API that you can import in [Postman](https://www.postman.com/) application. You can import **gokomodo-api.postman_collection.json** file to your postman.

## Unit test

To run unit test go to app/usecase/seller or app/usecase/buyer in terminal then run this command :
```bash
go test -v -cover
```
