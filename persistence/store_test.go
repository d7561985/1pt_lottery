package persistence

import (
	"github.com/iris-contrib/go.uuid"
	"testing"
)

func TestClean(t *testing.T) {
	for i := 0; i < 10; i++ {
		id, _ := uuid.NewV4()
		S.Store(id.String(), id.String())
	}

	S.Clean()

	counter := 0
	S.Range(func(_, _ interface{}) bool {
		counter++
		return true
	})
	if counter > 0 {
		t.Errorf("counter %d", counter)
	}
}
