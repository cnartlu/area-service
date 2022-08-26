package country

type CountryRepo interface{}

type Country struct {
	repo CountryRepo
}

func (r *Country) FindList() {

}

func NewCountry(
	repo CountryRepo,
) *Country {
	return &Country{
		repo: repo,
	}
}
