package categories

import "context"

type categoryUsecase struct {
	categoryRepository Repository
}

func NewCategoryUsecase(cr Repository) Usecase {
	return &categoryUsecase{
		categoryRepository: cr,
	}
}

func (cu *categoryUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return cu.categoryRepository.GetAll(ctx)
}

func (cu *categoryUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return cu.categoryRepository.GetByID(ctx, id)
}

func (cu *categoryUsecase) Create(ctx context.Context, categoryDomain *Domain) (Domain, error) {
	return cu.categoryRepository.Create(ctx, categoryDomain)
}

func (cu *categoryUsecase) Update(ctx context.Context, id string, categoryDomain *Domain) (Domain, error) {
	return cu.categoryRepository.Update(ctx, id, categoryDomain)
}

func (cu *categoryUsecase) Delete(ctx context.Context, id string) error {
	return cu.categoryRepository.Delete(ctx, id)
}
