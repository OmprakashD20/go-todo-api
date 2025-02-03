package utils

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func SendErrorResponse(ctx *fiber.Ctx, statusCode int, message string) error {
	return ctx.Status(statusCode).JSON(&fiber.Map{
		"error": message,
	})
}

func MapToArray(input map[string][]string) [][]string {
	result := make([][]string, 0, len(input))

	for _, values := range input {
		result = append(result, values)
	}

	return result
}

func StructToMap(data interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(data)

	result := make(map[string]interface{})
	t := v.Type().Elem()
	for i := 0; i < v.Elem().NumField(); i++ {
		fieldValue := v.Elem().Field(i)
		fieldName := t.Field(i).Name
		if fieldValue.CanInterface() {
			result[fieldName] = fieldValue.Interface()
		}
	}

	return result, nil
}

func Pick[T any](input T, pickFields ...string) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(input)
	t := reflect.TypeOf(input)

	pickMap := make(map[string]bool)
	for _, field := range pickFields {
		pickMap[field] = true
	}

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		if pickMap[fieldName] {
			result[fieldName] = v.Field(i).Interface()
		}
	}

	return result
}

func Omit[T any](input T, omitFields ...string) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(input)
	t := reflect.TypeOf(input)

	omitMap := make(map[string]bool)
	for _, field := range omitFields {
		omitMap[field] = true
	}

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		if !omitMap[fieldName] {
			result[fieldName] = v.Field(i).Interface()
		}
	}

	return result
}
