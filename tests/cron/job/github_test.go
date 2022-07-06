package job

import (
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
