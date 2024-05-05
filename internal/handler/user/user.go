package userHandler

import (
	fiber "github.com/gofiber/fiber/v2"
	"rest-api/database"
	"rest-api/internal/model"
)

func Get(c *fiber.Ctx) error {
	{
		var users []model.User
		// Find all users in database
		result := database.DB.Find(&users)
		// Check for errors during query execution
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return list of users
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "user Berhasil Ditampilkan!",
			"data":    users,
		})
	}
}

func Create(c *fiber.Ctx) error {
	// Parse request body
	{
		var user model.User
		if err := c.BodyParser(&user); err != nil {
			return err
		}
		// Insert new user into database
		result := database.DB.Create(&user)
		// Check for errors during insertion
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "user Berhasil Ditambahkan!",
			"data":    user,
		})
	}
}
func GetUser(c *fiber.Ctx) error {
	{
		// Get id_user parameter from request url
		id := c.Params("id_user")
		// Find user by id_user in database
		var user model.User
		result := database.DB.First(&user, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "user Tidak Ditemukan!",
			})
		}
		// Check for errors during query
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return user
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Success",
			"data":    user,
		})
	}
}
func Update(c *fiber.Ctx) error {
	{
		// Get id_user parameter from request url
		id := c.Params("id_user")
		// Find user by id_user in database
		var user model.User
		result := database.DB.First(&user, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "user Tidak Ditemukan",
			})
		}
		// Parse request body
		var newuser model.User
		if err := c.BodyParser(&newuser); err != nil {
			return err
		}
		// Update user in database
		result = database.DB.Model(&user).Updates(newuser)
		// Check for errors during update
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "user Berhasil Diperbarui!",
			"data":    user,
		})
	}
}
func Delete(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	{
		id := c.Params("id_user")
		// Find user by id_user in database
		var user model.User
		result := database.DB.First(&user, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "user Tidak Ditemukan",
			})
		}
		// Delete user from database
		result = database.DB.Delete(&user)
		// Check for errors during deletion
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "user Berhasil Dihapus!",
			"data":    user,
		})
	}
}
