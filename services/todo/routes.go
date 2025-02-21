package todo

import (
	"net/http"

	"github.com/OmprakashD20/go-todo-api/middlewares"
	"github.com/OmprakashD20/go-todo-api/models"
	"github.com/OmprakashD20/go-todo-api/types"
	"github.com/OmprakashD20/go-todo-api/utils"
	"github.com/OmprakashD20/go-todo-api/validations"

	"github.com/gofiber/fiber/v2"
)

type Service struct {
	todoStore types.TodoStore
	userStore types.UserStore
}

func NewService(todoStore types.TodoStore, userStore types.UserStore) *Service {
	return &Service{todoStore, userStore}
}

func (s *Service) SetupTodoRoutes(api fiber.Router) {
	api.Use(middlewares.VerifyToken(s.userStore))

	api.Post("/create", validator.ValidateSchema[types.CreateTodoPayload](*validator.CreateTodoSchema), s.CreateTodoHandler)
	api.Get("/all", s.GetTodosByUserIdHandler)
	api.Get("/:id", s.GetTodoHandler)
	api.Put("/update/:id", validator.ValidateSchema[types.UpdateTodoPayload](*validator.UpdateTodoSchema), s.UpdateTodoHandler)
	// api.Delete("/delete/:id", s.DeleteTodoHandler)
}

func (s *Service) CreateTodoHandler(ctx *fiber.Ctx) error {
	// Get validated todo data from Fiber context locals
	todo := ctx.Locals("validatedData").(*types.CreateTodoPayload)

	// Get user data from Fiber context locals
	user := utils.MapToStruct[models.User](ctx.Locals("user").(map[string]interface{}))

	data := models.Todo{
		Title:       todo.Title,
		Description: todo.Description,
		Priority:    todo.Priority,
		DueDate:     todo.DueDate,
		UserID:      user.ID,
	}

	err := s.todoStore.CreateTodo(&data)
	if err != nil {
		return utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": "Your todo has been created successfully!!",
	})

}

func (s *Service) GetTodoHandler(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return utils.SendErrorResponse(ctx, http.StatusBadRequest, "Todo ID is required")
	}

	todo, err := s.todoStore.GetTodoById(uint(id))

	if err != nil {
		return utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
	}
	if todo == nil {
		return utils.SendErrorResponse(ctx, http.StatusNotFound, "No todo found")
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"todo": *todo,
	})
}

func (s *Service) GetTodosByUserIdHandler(ctx *fiber.Ctx) error {
	// Get user data from Fiber context locals
	user := utils.MapToStruct[models.User](ctx.Locals("user").(map[string]interface{}))

	todos, err := s.todoStore.GetTodosByUserId(user.ID)
	if err != nil {
		return utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"todos": todos,
	})

}

func (s *Service) UpdateTodoHandler(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return utils.SendErrorResponse(ctx, http.StatusBadRequest, "Todo ID is required")
	}

	// Get validated todo data from Fiber context locals
	updatedTodo := ctx.Locals("validatedData").(*types.UpdateTodoPayload)

	// Check if todo exists
	if todo, err := s.todoStore.GetTodoById(uint(id)); err != nil {
		return utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
	} else if todo == nil {
		return utils.SendErrorResponse(ctx, http.StatusNotFound, "Todo not found")
	}

	data := models.Todo{
		Title:       updatedTodo.Title,
		Description: updatedTodo.Description,
		Priority:    updatedTodo.Priority,
		DueDate:     updatedTodo.DueDate,
	}

	err = s.todoStore.UpdateTodoById(uint(id), &data)
	if err != nil {
		return utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Todo has been updated successfully",
	})
}

// func (s *Service) DeleteTodoHandler(ctx *fiber.Ctx) error {}
