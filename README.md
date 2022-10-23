# efishery

I use existing libs :

 - Ozzo Validation, for input request validation
 - Godotenv, for env loader
 - jmoiron/sqlx for postgres driver


# For setup after cloning the repo:
- Auth App
> cd auth <br />
> go mod tidy <br />

- Fetch App

# to do a unit test :
> go to the each usecase package you want to testing then run a command "go test"
>> you can see the coverage testing in each usecase package by open the project with vscode, choose the testing file, right click then choose "Go:Toogle Test Coverage in Current Package"

# summary of unit test 
# Auth app
I have done the unit test and here are the results :
>Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-gosNxhO2/go-code-cover efishery/auth/src/app/use_cases/user
>>ok  	efishery/auth/src/app/use_cases/user	0.667s	coverage: 100.0% of statements

# Auth app
# for db table :
- in folder db, there is a .sql file with the create table command and insert command. I use postgresql for this case. you can run the command in your sql editor page.
- if you running this project without docker, just make a connection into your localhost
- then make new schema in db "projek", named "efishery", after that run all the command in .sql file
- if you running this project with docker, make a connection into 0.0.0.0 and make a database e.g I use projek
- then make new schema in db "projek", named "efishery", after that run all the command in .sql file 

# the endpoint
here is the curl for the endpoint :
>curl --location --request POST 'http://localhost:8080/api/user' \
--header 'x-api-key: efishery' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "test6",
    "phone": "1234567",
    "role": "admin"
}'

>curl --location --request POST 'http://localhost:8080/api/user/login' \
--header 'x-api-key: efishery' \
--header 'Content-Type: application/json' \
--data-raw '{
    "phone": "1234567",
    "password": "XVlB"
}'

>curl --location --request POST 'http://localhost:8080/api/user/info' \
--header 'x-api-key: efishery' \
--header 'Content-Type: application/json' \
--data-raw '{
    "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoidGVzdCIsIlBob25lIjoiMTIzNDU2NyIsIlJvbGUiOiJhZG1pbiIsIlRpbWUiOiIyMDIyLTEwLTIzVDE2OjA0OjE1LjU2NjIyNSswNzowMCJ9.JlqXwvSJDmVHdo-W1tsq8GNhItNbA-V4oqJpy2hcaeI"
}'

> here is the postman link if you want to use postman instead : 
>> https://www.getpostman.com/collections/7f2f753f84e4783dfd00

# to run the project
if you're not using docker, just set the .env file with your database credential, then cd into each app (auth or fetch).
> for Auth app do "go run main.go"

# to run the project with docker
after clone and do some set up that explained before, do this following actions :
- set database credential in .env

- in this part :
>> DB_HOST=database (recommend to literally use "database" according to the docker-compose.yaml) <br />
>> DB_PORT=5432   <br />
>> DB_NAME=projek/your_db_name <br />
>> DB_USERNAME=your_db_user <br />
>> DB_PASSWORD=your_db_password <br />
>> DB_SCHEMA=warung_pintar <br />
>> DB_SSL_MODE=disable <br />

- in this part :
>> POSTGRES_USER=postgres <br />
>> POSTGRES_PASSWORD=your_db_password <br />
>> POSTGRES_DB=projek <br />
 
>> cd efishery, docker-compose up <br />
>> go to you postgresql db editor (pgAdmin, etc) <br />
>> make a new connection to 0.0.0.0 <br />
>> make a new database, in this project I make a db named "projek" <br />
>> in db "projek" make a new schema named "efishery" <br />
>> do all command to make the table and insert, you can see the command in db/account.sql and db/customer.sql <br />
>> project ready to use <br />
