package routes

import (
	"fiber-crud/database"
	"fiber-crud/models"
	"github.com/gofiber/fiber/v3"

)
type User struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Username string `json:"username"`
    Email    string `json:"email"`
    // other fields...
}
func CreateUser(c fiber.Ctx) error { // Pointer hata diya hai
    user := new(models.User)
     if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	 }
	database.DB.Create(&user)  
    return c.JSON(user)
}



func GetUsers(c fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

func DeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	// ✅ Correct way to find the user
	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	database.DB.Delete(&user) // ✅ Corrected GORM Delete usage
	return c.JSON(fiber.Map{"message": "User deleted"})
}
