package validator

import (
	z "github.com/Oudwins/zog"
)

var RegisterUserSchema = z.Struct(
	z.Schema{
		"FirstName": z.String().Min(3, z.Message("First name must be at least 3 characters long")).Required(z.Message("First name is required")),
		"LastName":  z.String().Optional(),
		"Email":     z.String().Email(z.Message("Invalid email address")).Required(z.Message("Email is required")),
		"Password":  z.String().Min(8, z.Message("Password must be at least 8 characters long")).Required(z.Message("Password is required")),
	},
)

var LoginUserSchema = z.Struct(
	z.Schema{
		"Email":    z.String().Email(z.Message("Invalid email address")).Required(z.Message("Email is required")),
		"Password": z.String().Optional(),
	},
)
