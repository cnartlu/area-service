package job

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/cnartlu/area-service/tests"
)

func Test_Run(t *testing.T) {
	ts, cleanup, err := tests.Init()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	ts.CronJob.Github.Run()
}

func Test_Json(t *testing.T) {
	b, _ := os.ReadFile("C:\\Users\\huanghu\\Desktop\\shop-202207191043.json")
	var k interface{}
	if err := json.Unmarshal(b, &k); err != nil {
		t.Error(err)
	}
}
