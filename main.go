package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/bcarrol2/GOalie/user"
	"github.com/bcarrol2/GOalie/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App){
	app.Get("/api/v1/users", user.GetAllUsers)
	app.Get("/api/v1/user/:id", user.GetUser)
	app.Get("/api/v1/users/under40", user.GetUsersUnderForty)
	app.Get("/api/v1/users/highpaying", user.GetHighPayingUsers)
	app.Get("/api/v1/users/brokeunder21", user.GetYoungerCheaperUsers)
	app.Post("/api/v1/user", user.CreateUser)
	app.Delete("/api/v1/user/:id", user.DeleteUser)
}

func initalizeDB(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connected")

	database.DBConn.AutoMigrate(&user.User{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initalizeDB()
	defer database.DBConn.Close()
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello World")
	// })
	setupRoutes(app)
	
	app.Listen(":3000")
}