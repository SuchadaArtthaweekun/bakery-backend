package controllers

import (

	"fmt"
	"go-fiber/database"
	"go-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func AddProduct(c *fiber.Ctx) error {
    data := new(models.Product)

    err := c.BodyParser(data)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
    }
    err = database.DB.Create(&data).Error
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
    }
	return c.JSON(fiber.Map{"message": "Save Completed"})
}



func ProductAll(c *fiber.Ctx) error {
	var product []models.Product

	database.DB.Find(&product)
	return c.JSON(product)
}


func EditProduct(c *fiber.Ctx) error {
	type UpProduct struct {
		Pname   string  `json:"pname"`
		Descript string `json:"descript"`
		Price   string `json:"price"`
		Picture string  `json:"picture"`
	}

	var data models.Product
	id := c.Params("id")
	database.DB.Find(&data, "Pid = ?",id)

	var newData UpProduct
	err := c.BodyParser(&newData)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
    }
	
	data.Pname = newData.Pname
	data.Descript = newData.Descript
	data.Price = newData.Price
	data.Picture = newData.Picture

	// save 
	database.DB.Save(&data)

	return c.JSON(fiber.Map{"message": "Update data complete"})
}

func FineProduct(c *fiber.Ctx) error {
	var data []models.Product
	id := c.Params("id")
	database.DB.Find(&data,"Pid = ?",id)

	return c.JSON(data)
}


func ProductDelete(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Print(err)
	}

	fmt.Println(data["pid"])
	
	var product models.Product

	database.DB.Where("pid = ?", data["pid"]).Delete(&product)
	return c.JSON(fiber.Map{
		"message": "Delete completed",
	})
}

func ImgUp (c *fiber.Ctx) error{
	file,err := c.FormFile("file")
	if err != nil {
		fmt.Print(err)
	}
	c.SaveFile(file, fmt.Sprintf("./images/%s", file.Filename))
	return c.JSON(fiber.Map{
		"message": "Save completed",
	})
}