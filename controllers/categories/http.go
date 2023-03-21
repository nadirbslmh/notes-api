package categories

import (
	"net/http"
	"notes-api/businesses/categories"
	"notes-api/controllers"
	"notes-api/controllers/categories/request"
	"notes-api/controllers/categories/response"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryUsecase categories.Usecase
}

func NewCategoryController(categoryUC categories.Usecase) *CategoryController {
	return &CategoryController{
		categoryUsecase: categoryUC,
	}
}

func (ctrl *CategoryController) GetAllCategories(c echo.Context) error {
	ctx := c.Request().Context()

	categoriesData, err := ctrl.categoryUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "error when fetching data", "")
	}

	categories := []response.Category{}

	for _, category := range categoriesData {
		categories = append(categories, response.FromDomain(category))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all categories", categories)
}

func (ctrl *CategoryController) CreateCategory(c echo.Context) error {
	input := request.Category{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	category, err := ctrl.categoryUsecase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "error when inserting data", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "category created", response.FromDomain(category))
}

func (ctrl *CategoryController) UpdateCategory(c echo.Context) error {
	input := request.Category{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	var id string = c.Param("id")

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	category, err := ctrl.categoryUsecase.Update(ctx, id, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "error when updating data", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "category updated", response.FromDomain(category))
}

func (ctrl *CategoryController) DeleteCategory(c echo.Context) error {
	var id string = c.Param("id")
	ctx := c.Request().Context()

	err := ctrl.categoryUsecase.Delete(ctx, id)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "error when deleting data", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "category deleted", "")
}
