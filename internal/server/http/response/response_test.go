package response

import (
	"encoding/json"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("response", func(t *testing.T) {
		p := NewResponse(1000, "操作失败，请稍后再试")
		bs, err := json.Marshal(p)
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("%s", bs)
	})

	t.Run("dataResponse", func(t *testing.T) {
		p := NewDataResponse(1000, "操作失败，请稍后再试", map[string]interface{}{
			"items": []string{"a", "b", "c", "d", "e"},
		})
		bs, err := json.Marshal(p)
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("%s", bs)
	})
}
