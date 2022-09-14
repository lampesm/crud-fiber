package jwt

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/lampesm/crud-fiber/serializers"
)

// JWT godoc
// @Summary create a JWT
// @Description create a JWT
// @TAGs authenticate
// @Accept  json
// @Produce  json
// @Param payload body serializers.Login true "Login"
// @Success 200
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	env := godotenv.Load(".env")

	if env != nil {
		panic("Failed to load .env file")
	}

	payload := new(serializers.Login)

	err := c.BodyParser(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if payload.User != os.Getenv("JWT_USER") || payload.Pass != os.Getenv("JWT_PASS") {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  os.Getenv("JWT_USER"),
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
