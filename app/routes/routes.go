package routes

import (
	"notes-api/controllers/categories"
	"notes-api/controllers/notes"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	NoteController     notes.NoteController
	CategoryController categories.CategoryController
}

func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {

	note := e.Group("/api/v1/notes")

	note.GET("", cl.NoteController.GetAllNotes)
	note.GET("/:id", cl.NoteController.GetNoteByID)
	note.POST("", cl.NoteController.CreateNote)
	note.PUT("/:id", cl.NoteController.UpdateNote)
	note.DELETE("/:id", cl.NoteController.DeleteNote)
	note.POST("/:id", cl.NoteController.RestoreNote)
	note.DELETE("/force/:id", cl.NoteController.ForceDeleteNote)

	category := e.Group("/api/v1/categories")

	category.GET("", cl.CategoryController.GetAllCategories)
	category.POST("", cl.CategoryController.CreateCategory)
	category.PUT("/:id", cl.CategoryController.UpdateCategory)
	category.DELETE("/:id", cl.CategoryController.DeleteCategory)
}
