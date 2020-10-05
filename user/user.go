package user

import (
	"fmt"
	"github.com/bcarrol2/GOalie/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Age int `json:"age`
	AmountSpent float64 `json:"amountSpent"`
}
 
func GetAllUsers(c *fiber.Ctx) error{
	db := database.DBConn
	var users []User
	db.Find(&users)
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var user User
	db.Find(&user, id)
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DBConn
	// var user User
	// user.Name = "Lytics Client"
	// user.Age = 50
	// user.AmountSpent = 500.50
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		fmt.Println("error", err)
		return c.Status(503).SendString("err")
	}

	db.Create(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var user User
	db.First(&user, id)
	if user.Name == "" {
		return c.Status(500).SendString("User not found")
	}
	
	db.Delete(&user)
	return c.SendString("User deleted successfully")
}