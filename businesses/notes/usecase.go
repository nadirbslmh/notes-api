package notes

import "context"

type noteUsecase struct {
	noteRepository Repository
}

func NewNoteUsecase(nr Repository) Usecase {
	return &noteUsecase{
		noteRepository: nr,
	}
}

func (nu *noteUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return nu.noteRepository.GetAll(ctx)
}

func (nu *noteUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return nu.noteRepository.GetByID(ctx, id)
}

func (nu *noteUsecase) Create(ctx context.Context, noteDomain *Domain) (Domain, error) {
	return nu.noteRepository.Create(ctx, noteDomain)
}

func (nu *noteUsecase) Update(ctx context.Context, id string, noteDomain *Domain) (Domain, error) {
	return nu.noteRepository.Update(ctx, id, noteDomain)
}

func (nu *noteUsecase) Delete(ctx context.Context, id string) error {
	return nu.noteRepository.Delete(ctx, id)
}

func (nu *noteUsecase) Restore(ctx context.Context, id string) (Domain, error) {
	return nu.noteRepository.Restore(ctx, id)
}

func (nu *noteUsecase) ForceDelete(ctx context.Context, id string) error {
	return nu.noteRepository.ForceDelete(ctx, id)
}
