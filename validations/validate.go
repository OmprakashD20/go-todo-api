package validator

import (
	"github.com/OmprakashD20/go-todo-api/utils"
	z "github.com/Oudwins/zog"
	"github.com/gofiber/fiber/v2"
)

func ValidateSchema[T any](schema z.StructSchema) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		data := new(T)

		// Parse request body
		body := new(T)

		if err := ctx.BodyParser(body); err != nil {
			return utils.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid Data")
		}

		// Convert the struct to map since Zog parses only map[string]interface{}
		result := utils.StructToMap(body)

		// Validate request body
		if errs := schema.Parse(result, data); errs != nil {
			errors := utils.MapToArray(z.Errors.SanitizeMap(errs))

			return utils.SendErrorResponse(ctx, fiber.StatusUnprocessableEntity, errors[0][0])
		}

		// Store validated data in Fiber context
		ctx.Locals("validatedData", data)

		return ctx.Next()
	}
}
