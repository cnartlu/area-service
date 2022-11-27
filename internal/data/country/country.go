package country

import (
	"context"

	"github.com/cnartlu/area-service/internal/data/data"
)

type Country struct {
	data *data.Data
}

func (r *Country) FindList(ctx context.Context) {

}

func NewCountry(
	d *data.Data,
) *Country {
	return &Country{
		data: d,
	}
}
