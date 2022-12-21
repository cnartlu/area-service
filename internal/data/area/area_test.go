package area

import (
	"context"
	"testing"

	"github.com/cnartlu/area-service/component/database"
	"github.com/cnartlu/area-service/component/log"
	"github.com/go-redis/redis/v8"

	"github.com/cnartlu/area-service/internal/data/data"
)

func TestReplaceParentListPrefix(t *testing.T) {
	logger, _ := log.New(nil)
	dataData, cleanup2, err := data.NewData(
		logger,
		redis.NewClient(&redis.Options{}),
		&database.Config{
			Address:  "localhost:3306",
			Username: "root",
			Password: "xingyun",
			Database: "ad_area_service",
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Cleanup(cleanup2)
	repo := NewAreaRepo(dataData)
	_, err = repo.ReplaceParentListPrefix(context.TODO(), "1", "2")
	if err != nil {
		t.Error(err)
		return
	}
}
