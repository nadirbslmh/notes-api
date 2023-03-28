package routes

import (
	"notes-api/app/middlewares"
	"notes-api/controllers/categories"
	"notes-api/controllers/notes"
	"notes-api/controllers/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware   echo.MiddlewareFunc
	JWTMiddleware      echojwt.Config
	AuthController     users.AuthController
	NoteController     notes.NoteController
	CategoryController categories.CategoryController
}

func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	users := e.Group("/api/v1/users")

	users.POST("/register", cl.AuthController.Register)
	users.POST("/login", cl.AuthController.Login)

	note := e.Group("/api/v1/notes", echojwt.WithConfig(cl.JWTMiddleware))
	note.Use(middlewares.VerifyToken)

	note.GET("", cl.NoteController.GetAllNotes)
	note.GET("/:id", cl.NoteController.GetNoteByID)
	note.POST("", cl.NoteController.CreateNote)
	note.PATCH("/:id", cl.NoteController.UpdateNote)
	note.DELETE("/:id", cl.NoteController.DeleteNote)
	note.POST("/:id", cl.NoteController.RestoreNote)
	note.DELETE("/force/:id", cl.NoteController.ForceDeleteNote)

	category := e.Group("/api/v1/categories", echojwt.WithConfig(cl.JWTMiddleware))
	category.Use(middlewares.VerifyToken)

	category.GET("", cl.CategoryController.GetAllCategories)
	category.POST("", cl.CategoryController.CreateCategory)
	category.PATCH("/:id", cl.CategoryController.UpdateCategory)
	category.DELETE("/:id", cl.CategoryController.DeleteCategory)

	auth := e.Group("/api/v1/users", echojwt.WithConfig(cl.JWTMiddleware))

	auth.POST("/logout", cl.AuthController.Logout)
}
