package importer

import (
	"context"
	"testing"
)

func Test_Import(t *testing.T) {
	r := NewRepository(nil)
	ctx := context.Background()
	r.ImportWithRow(ctx, Row{
		ID:  "1",
		Pid: "200",
	})

}
