package categories

import "notes-api/businesses/categories"

type CategoryController struct {
	categoryUsecase categories.Usecase
}

func NewCategoryController(categoryUC categories.Usecase) *CategoryController {
	return &CategoryController{
		categoryUsecase: categoryUC,
	}
}

//TODO: implement this
