Test Project
===
How to run the application
---

1. Create the Mysql database named test_db and use the test_db.sql to create the tables and insert some test data.

2. Once the database is created, change the Connection String as per your mysql username, password and database name on line No. 12 in the config.go.

There are number of dependencies which need to be imported before running the application. Please get the dependencies through the following commands -

    go get "github.com/go-sql-driver/mysql"
    go get "github.com/jeanphorn/log4go"
    go get "github.com/labstack/echo"
    go get "github.com/stretchr/testify/assert"
    go get "github.com/dgrijalva/jwt-go"
To run the application, please use the following command -

    go run main.go
Note: The default port number is 80.

Endpoints Description:
===
Login And Get JWT (Json Web Token)
---
Note: The JWT is valid for 1 hour

    URL - *http://localhost/login*
    Method - POST
    Body - (content-type = application/json)
    {
        "username":"admin",
        "password":"admin"
    }
GET User By ID
---
Note: Attach JWT you got when you logged in to authorization of request header.

    URL - *http://localhost/user/1*
    Method - GET
GET All Users
---
Note: Attach JWT you got when you logged in to authorization of request header.

    URL - *http://localhost/users*
    Method - GET
CREAT User Information
---
Note: Attach JWT you got when you logged in to authorization of request header.

    URL - *http://localhost/user*
    Method - POST
UPDATE User Information
---
Note: Attach JWT you got when you logged in to authorization of request header.

    URL - *http://localhost/user*
    Method - PUT
DELETE User Information
---
Note: Attach JWT you got when you logged in to authorization of request header.

    URL - *http://localhost/user/7*
    Method - DELETE
Test Driven Development Description
===
To run all the unit test cases, please do the following -

    go test -v

Hope everything works out. Thank you!
===