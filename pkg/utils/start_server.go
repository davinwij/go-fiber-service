package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func ConnectionURLBuilder(n string) (string, error) {
	var url string

	switch n {
	case "fiber":
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		return "", fmt.Errorf("connection name %s is not supported", n)
	}

	return url, nil
}

func StartServer(a *fiber.App) {
	fiberConnUrl, _ := ConnectionURLBuilder("fiber")

	if err := a.Listen(fiberConnUrl); err != nil {
		log.Printf("Oops.. Server is not running! Because, %v", err)
	}
}
