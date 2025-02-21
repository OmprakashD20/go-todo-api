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

func MapToStruct[T any](data map[string]any) T {
	var result T
	v := reflect.ValueOf(&result).Elem()

	for i := range v.NumField() {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		if val, exists := data[fieldType.Name]; exists && field.CanSet() {
			valRef := reflect.ValueOf(val)

			if valRef.Type().ConvertibleTo(field.Type()) {
				field.Set(valRef.Convert(field.Type()))
			}
		}
	}

	return result
}

func StructToMap(data any) map[string]any {
	v := reflect.ValueOf(data)

	result := make(map[string]any)
	t := v.Type().Elem()
	for i := range v.Elem().NumField() {
		fieldValue := v.Elem().Field(i)
		fieldName := t.Field(i).Name
		if fieldValue.CanInterface() {
			result[fieldName] = fieldValue.Interface()
		}
	}

	return result
}

func Pick[T any](input T, pickFields ...string) map[string]any {
	result := make(map[string]any)
	v := reflect.ValueOf(input)
	t := reflect.TypeOf(input)

	pickMap := make(map[string]bool)
	for _, field := range pickFields {
		pickMap[field] = true
	}

	for i := range v.NumField() {
		fieldName := t.Field(i).Name
		if pickMap[fieldName] {
			result[fieldName] = v.Field(i).Interface()
		}
	}

	return result
}

func Omit[T any](input T, omitFields ...string) map[string]any {
	result := make(map[string]any)
	v := reflect.ValueOf(input)
	t := reflect.TypeOf(input)

	omitMap := make(map[string]bool)
	for _, field := range omitFields {
		omitMap[field] = true
	}

	for i := range v.NumField() {
		fieldName := t.Field(i).Name
		if !omitMap[fieldName] {
			result[fieldName] = v.Field(i).Interface()
		}
	}

	return result
}
