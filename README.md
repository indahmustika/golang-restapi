# Simple CRUD Rest API Using Golang
- Create folder inside <code>GOPATH/scr</code>
- Create folder named rest-crud
- Project and dependencies initialization inside rest-crud folder
  - MySQL Database: <code>go get github.com/go-sql-driver/mysql</code>
  - Mux Router: <code>go get github.com/gorilla/mux</code>
- Create main.go to set router and server
- Create mysql.go to configure database
- Create model.go to set data structure
- Create user.go to make function for each route
- Run project
  - <code>go build</code>
  - <code>rest-crud.exe</code>
- Test using postman (post, put, delete type use x-www-form-urlencoded)
#### References 
 - <url>https://kiddyxyz.medium.com/tutorial-golang-rest-api-mysql-part-1-45cd9f4e75a6</url>
 - <url>https://kiddyxyz.medium.com/tutorial-golang-rest-api-mysql-part-2-6a9c732a42ca</url> 
