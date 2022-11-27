package country

type CountryRepo interface{}

type ManageUsecase struct {
	repo CountryRepo
}

func (r *ManageUsecase) FindList() {

}

func NewCountryUsecase(
	repo CountryRepo,
) *ManageUsecase {
	return &ManageUsecase{
		repo: repo,
	}
}
