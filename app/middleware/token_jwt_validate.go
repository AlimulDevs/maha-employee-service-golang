package middleware

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyTokenJwt() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"status":  "error",
				"code":    403,
				"message": "Unauthorized, token missing",
			})
		}

		// Harus pakai format Bearer token
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"status":  "error",
				"code":    403,
				"message": "Invalid Authorization header format, must start with Bearer",
			})
		}

		// Ambil token string
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse dan validasi token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			secretBase64 := "VCPmUSBSV2pgrZwli4hkbsc5EFxhwFC5qIHAqzA5n6ZB4Mi1HesoNTVap3XWEJOm"
			secretBytes, err := base64.StdEncoding.DecodeString(secretBase64)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "error",
					"code":    500,
					"message": "Invalid JWT_SECRET encoding",
				}), nil
			}
			return secretBytes, nil

		})

		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"status":  "error",
				"code":    403,
				"message": fmt.Sprintf("Token invalid: %v", err.Error()),
			})
		}

		if !token.Valid {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"status":  "error",
				"code":    403,
				"message": "Token not valid",
			})
		}

		// Ambil claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if exp, ok := claims["exp"].(float64); ok {
				if int64(exp) < time.Now().Unix() {
					return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
						"status":  "error",
						"code":    403,
						"message": "Token expired",
					})
				}
			}
			c.Locals("token_payload", claims["data"])
		}

		return c.Next()
	}
}
