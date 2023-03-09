package controllers

import (
	"go-fiber-tutor/app/models"
	"go-fiber-tutor/configs/database"

	"go-fiber-tutor/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User models.User

func UserSignUp(c *fiber.Ctx) error {

	user := User{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "User Register Failed",
			"err": err.Error(),
		})
	}

	if err := database.DB.First(&user, "email = ?", user.Email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			hashed, err := utils.HashPassword(user.Password)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"err": err.Error(),
				})
			}

			user.Password = hashed

			database.DB.Create(&user)

			return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
				"msg":  "Success Register User",
				"data": user,
			})
		}
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"msg": "User with email " + user.Email + " already registered!",
	})
}

func UserSignIn(c *fiber.Ctx) error {
	user := User{}
	foundedUser := User{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	if err := database.DB.First(&foundedUser, "email = ?", user.Email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"msg": "No User Found with this email " + user.Email,
			})
		}
	}

	if err := utils.CheckHashPassword(foundedUser.Password, user.Password); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"msg": "Wrong Password",
			"err": err.Error(),
			"pw": user.Password,
			"hash": foundedUser.Password,
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"msg": "You are Authenticated!",
	})

}

func GetAllUser(c *fiber.Ctx) error {
	user := []User{}

	if err := database.DB.Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"msg": "No User Found",
			})
		}
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"msg":  "User Found",
		"data": user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	user := User{}
	updateData := User{}

	id := c.Params("id")

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Bad Request",
			"err": err.Error(),
		})
	}

	if err := database.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"msg": "User %v not found!" + user.Email,
			})
		}
	}

	user.Email = updateData.Email
	user.UserRole = updateData.UserRole

	database.DB.Save(&user)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"msg": "Success Update User " + user.Email,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	user := User{}
	id := c.Params("id")

	if err := database.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"msg": "User with id " + id + " not found!",
			})
		}
	}

	if err := database.DB.Delete(&user, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"msg": "Success Delete User with id " + id,
	})
}
