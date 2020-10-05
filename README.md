# GOalie

Golang CRUD users REST api built with sqlite3, gofiber & gorm

##### Start api #####
cd GOalie
go run main.go

Test api with curl in second termial.
Example:

curl -X POST -H "Content-Type: application/json" --data "{\"name\": \"Neil deGrasse Tyson\", \"age\": 62, \"amount_spent\": 2500.75}"  http://localhost:3000/api/v1/users