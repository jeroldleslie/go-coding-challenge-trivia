#Go coding challenge

###Technologies used:
- GoLang
- MySQL

###What are all done?
- Developed based on dockerized microservice architecture
- few unit tests covered
- dockerized the whole project so that it can be easily shipped to any cloud instances
- prodution ready code(70%)

###Endpoint Document
https://docs.google.com/document/d/14Q2KV9nLY-vCVCfLz6g_xj5tizOZeMOkwGG__LYnQng/edit?usp=sharing
TODO: Swagger API document need to be added

##How to run and check?
###Prerequisites
- install docker https://docs.docker.com/engine/install/
- install MySQL

###Set Environment variables
- APP_DB_CONNECTION_STR (ex: APP_DB_CONNECTION_STR="root:trivia@tcp(localhost:3306)/triviadb")
- MYSQL_ROOT_PASSWORD (this is only if running DB in docker container)

###How to run and check using docker?
- open the terminal or command prompt
- `git clone https://github.com/jeroldleslie/go-coding-challenge-trivia`
- `cd <to the checkedout project folder>`
- run `docker-compose up`
- wait for database service and api service to start
- you will see like below if the service started successfully
```
   ____    __
  / __/___/ /  ___
  / _// __/ _ \/ _ \
  /___/\__/_//_/\___/ v3.3.10-dev
  High performance, minimalist Go web framework
  https://echo.labstack.com
  ____________________________________O/_______
  O\
  ⇨ http server started on [::]:8000
```
- test using any http client (ex: postman client)

###How to run and check locally?
run the below commands in terminal
- `git clone https://github.com/jeroldleslie/go-coding-challenge-trivia`
- run the SQL scripts to initiate DB schema and sample data. Run the below scripts in order
  1. <project folder>/scripts/schema.sql
  1. <project folder>/scripts/sample_data.sql
- `cd <to the checkedout project folder>/cmd/api_trivia/`
- `go build`
- `./api_trivia`
- you will see like below if the service started successfully 
```
   ____    __
  / __/___/ /  ___
  / _// __/ _ \/ _ \
  /___/\__/_//_/\___/ v3.3.10-dev
  High performance, minimalist Go web framework
  https://echo.labstack.com
  ____________________________________O/_______
  O\
  ⇨ http server started on [::]:8000
```
NOTE: 
1. The above instructions were very high level. Hope you can understand. Let me know if you have any doubts on running the services.
2. Almost covered tha main things that showcasing essential things. Due to time restriction I cannot covered a lot of functionalities.
