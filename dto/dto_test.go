package dto

import (
	"encoding/json"
	"github.com/gookit/validate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStart(t *testing.T) {
	tests := []struct {
		name string
		json string
	}{
		{"", `{"user":"dedede","avatar":"data:image/png;base64,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"}`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			st := &UserRequest{}
			err := json.Unmarshal([]byte(test.json), st)
			assert.NoError(t, err)

			v := validate.Struct(st)
			if !v.Validate() {
				t.Error(v.Errors)

			}
		})
	}
}
