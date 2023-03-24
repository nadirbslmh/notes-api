package notes

import (
	"net/http"
	"notes-api/businesses/notes"
	"notes-api/controllers"
	"notes-api/controllers/notes/request"
	"notes-api/controllers/notes/response"

	"github.com/labstack/echo/v4"
)

type NoteController struct {
	noteUsecase notes.Usecase
}

func NewNoteController(noteUC notes.Usecase) *NoteController {
	return &NoteController{
		noteUsecase: noteUC,
	}
}

func (ctrl *NoteController) GetAllNotes(c echo.Context) error {
	ctx := c.Request().Context()
	notesData, err := ctrl.noteUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "error when fetching data", "")
	}

	notes := []response.Note{}

	for _, note := range notesData {
		notes = append(notes, response.FromDomain(note))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all notes", notes)
}

func (ctrl *NoteController) GetNoteByID(c echo.Context) error {
	var id string = c.Param("id")
	ctx := c.Request().Context()

	note, err := ctrl.noteUsecase.GetByID(ctx, id)

	isNotFound := err != nil || note.ID == 0

	if isNotFound {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "note not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "note found", response.FromDomain(note))
}

func (ctrl *NoteController) CreateNote(c echo.Context) error {
	input := request.Note{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	note, err := ctrl.noteUsecase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "error when inserting data", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "note created", response.FromDomain(note))
}

func (ctrl *NoteController) UpdateNote(c echo.Context) error {
	input := request.Note{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	var noteId string = c.Param("id")

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	note, err := ctrl.noteUsecase.Update(ctx, noteId, input.ToDomain())

	isFailed := err != nil || note.ID == 0

	if isFailed {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "error when updating data", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "note updated", response.FromDomain(note))
}

func (ctrl *NoteController) DeleteNote(c echo.Context) error {
	var noteId string = c.Param("id")
	ctx := c.Request().Context()

	err := ctrl.noteUsecase.Delete(ctx, noteId)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "error when deleting data", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "note deleted", "")
}

func (ctrl *NoteController) RestoreNote(c echo.Context) error {
	var noteId string = c.Param("id")
	ctx := c.Request().Context()

	note, err := ctrl.noteUsecase.Restore(ctx, noteId)

	isFailed := err != nil || note.ID == 0

	if isFailed {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "error when restoring data", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "note restored", response.FromDomain(note))
}

func (ctrl *NoteController) ForceDeleteNote(c echo.Context) error {
	var noteId string = c.Param("id")
	ctx := c.Request().Context()

	err := ctrl.noteUsecase.ForceDelete(ctx, noteId)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "error when force deleting data", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "note deleted permanently", "")
}
