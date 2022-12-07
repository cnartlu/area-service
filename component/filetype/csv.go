package filetype

import "github.com/h2non/filetype/types"

var TypeCsv = types.NewType("csv", "application/csv")

func Csv(b []byte) bool {
	return false
}
