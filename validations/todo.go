package validator

import (
	"time"

	z "github.com/Oudwins/zog"
)

var BaseTodoSchema = z.Struct(
	z.Schema{
		"Title":       z.String().Min(3, z.Message("Title must be atleast 3 characters long")).Required(z.Message("Title is required")),
		"Description": z.String().Min(15, z.Message("Description must be atleast 15 characters long")).Required(z.Message("Description is required")),
		"Priority":    z.String().OneOf([]string{"Low", "Medium", "High"}, z.Message("Priority must be Low, Medium or High")),
		"DueDate":     z.Time().After(time.Now(), z.Message("Due Date must be either today or in the future")).Required(z.Message("Due Date is required")),
	},
)

var CreateTodoSchema = BaseTodoSchema.Extend(z.Schema{})

var UpdateTodoSchema = BaseTodoSchema.Extend(
	z.Schema{
		"IsCompleted": z.Bool(),
	},
)
