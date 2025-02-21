package middlewares

import (
	"net/http"
	"strconv"

	"github.com/OmprakashD20/go-todo-api/types"
	"github.com/OmprakashD20/go-todo-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(store types.UserStore) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Get("Authorization")

		if tokenString == "" {
			return utils.SendErrorResponse(ctx, http.StatusUnauthorized, "Missing Authorization Token")
		}

		token, err := utils.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			return utils.SendErrorResponse(ctx, http.StatusUnauthorized, "You are not authorized")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return utils.SendErrorResponse(ctx, http.StatusUnauthorized, "You are not authorized")
		}

		userIdStr, _ := claims["userId"].(string)

		userId, _ := strconv.Atoi(userIdStr)

		user, err := store.GetUserById(uint(userId))
		if err != nil {
			return utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
		}
		if user == nil {
			return utils.SendErrorResponse(ctx, http.StatusUnauthorized, "You are not authorized")
		} 

		ctx.Locals("user", utils.Omit(*user, "HashedPassword", "PasswordSalt", "CreatedAt", "UpdatedAt"))

		return ctx.Next()
	}
}
