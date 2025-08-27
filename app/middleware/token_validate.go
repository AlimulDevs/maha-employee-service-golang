package middleware

import (
	"api/app/lib"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// TokenValidator middleware
func TokenValidator() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get(viper.GetString("HEADER_TOKEN_KEY"))
		if token != viper.GetString("VALUE_TOKEN_KEY") {
			return lib.ErrorUnauthorized(c, "Wrong x-Token header")
		}
		return c.Next()
	}
}
