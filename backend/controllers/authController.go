package controllers

import (
	"fmt"
	"go-fiber/database"
	"go-fiber/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

)

const SecretKey = "secret"

func AddUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Print("error")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Password: password,
		Email:    data["email"],
	}
	database.DB.Create(&user)
	return c.JSON(fiber.Map{
		"message": "Add data complete",
	})
}


// Dellete User

// func Delete(c *fiber.Ctx) error {
// 	var data map[string]string

// 	if err := c.BodyParser(&data); err != nil {
// 		fmt.Print(err)
// 	}

// 	var user models.User

// 	database.DB.Where("name = ?", data["name"]).Delete(&user)

// 	return c.JSON(fiber.Map{
// 		"message": "Delete complete",
// 	})
// }

func Delete(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Print(err)
	}

	fmt.Println(data["Id"])
	
	var user models.User

	database.DB.Where("Id = ?", data["Id"]).Delete(&user)
	return c.JSON(fiber.Map{
		"message": "Delete completed",
	})
}



//Update User
func UpdateUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Print("error")
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Password: password,
		Email:    data["email"],
	}
	database.DB.Where("name = ?", data["name"]).Updates(&user)
	return c.JSON(fiber.Map{
		"message": "Update data complete",
	})
}

//Login

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Print(err)
	}

	var user models.User

	database.DB.Where("name = ?", data["name"]).First(&user)

	if user.Name == "" {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User Not Found",
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}
	return c.JSON(user)
}

