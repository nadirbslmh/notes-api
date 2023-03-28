package notes_test

import (
	"context"
	"errors"
	"notes-api/businesses/categories"
	"notes-api/businesses/notes"
	_noteMock "notes-api/businesses/notes/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	noteRepository _noteMock.Repository
	noteService    notes.Usecase

	noteDomain notes.Domain
	ctx        context.Context
)

func TestMain(m *testing.M) {
	noteService = notes.NewNoteUsecase(&noteRepository)

	categoryDomain := categories.Domain{
		Name: "test category",
	}

	noteDomain = notes.Domain{
		Title:      "title",
		Content:    "my content",
		CategoryID: categoryDomain.ID,
	}

	ctx = context.TODO()

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		noteRepository.On("GetAll", ctx).Return([]notes.Domain{noteDomain}, nil).Once()

		result, err := noteService.GetAll(ctx)

		assert.Equal(t, 1, len(result))
		assert.Nil(t, err)
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		noteRepository.On("GetAll", ctx).Return([]notes.Domain{}, nil).Once()

		result, err := noteService.GetAll(ctx)

		assert.Equal(t, 0, len(result))
		assert.Nil(t, err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Get By ID | Valid", func(t *testing.T) {
		noteRepository.On("GetByID", ctx, "1").Return(noteDomain, nil).Once()

		result, err := noteService.GetByID(ctx, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get By ID | InValid", func(t *testing.T) {
		noteRepository.On("GetByID", ctx, "-1").Return(notes.Domain{}, errors.New("failed")).Once()

		result, err := noteService.GetByID(ctx, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		noteRepository.On("Create", ctx, &noteDomain).Return(noteDomain, nil).Once()

		result, err := noteService.Create(ctx, &noteDomain)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		noteRepository.On("Create", ctx, &notes.Domain{}).Return(notes.Domain{}, errors.New("failed")).Once()

		result, err := noteService.Create(ctx, &notes.Domain{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		noteRepository.On("Update", ctx, "1", &noteDomain).Return(noteDomain, nil).Once()

		result, err := noteService.Update(ctx, "1", &noteDomain)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		noteRepository.On("Update", ctx, "1", &notes.Domain{}).Return(notes.Domain{}, errors.New("failed")).Once()

		result, err := noteService.Update(ctx, "1", &notes.Domain{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		noteRepository.On("Delete", ctx, "1").Return(nil).Once()

		err := noteService.Delete(ctx, "1")

		assert.Nil(t, err)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		noteRepository.On("Delete", ctx, "-1").Return(errors.New("failed")).Once()

		err := noteService.Delete(ctx, "-1")

		assert.NotNil(t, err)
	})
}

func TestRestore(t *testing.T) {
	t.Run("Restore | Valid", func(t *testing.T) {
		noteRepository.On("Restore", ctx, "1").Return(noteDomain, nil).Once()

		result, err := noteService.Restore(ctx, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Restore | InValid", func(t *testing.T) {
		noteRepository.On("Restore", ctx, "-1").Return(notes.Domain{}, errors.New("failed")).Once()

		result, err := noteService.Restore(ctx, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestForceDelete(t *testing.T) {
	t.Run("ForceDelete | Valid", func(t *testing.T) {
		noteRepository.On("ForceDelete", ctx, "1").Return(nil).Once()

		err := noteService.ForceDelete(ctx, "1")

		assert.Nil(t, err)
	})

	t.Run("ForceDelete | InValid", func(t *testing.T) {
		noteRepository.On("ForceDelete", ctx, "-1").Return(errors.New("failed")).Once()

		err := noteService.ForceDelete(ctx, "-1")

		assert.NotNil(t, err)
	})
}
