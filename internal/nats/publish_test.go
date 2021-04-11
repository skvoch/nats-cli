package natsc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateJson(t *testing.T) {
	cases := []struct {
		Name    string
		JSON    []byte
		IsValid bool
	}{
		{
			Name:    "#1 object",
			JSON:    []byte(`{"filed":1234}`),
			IsValid: true,
		},
		{
			Name:    "#2 array",
			JSON:    []byte(`[{"filed":1234}]`),
			IsValid: true,
		},
		{
			Name:    "#3 invalid object",
			JSON:    []byte(`{"filed":1234}`),
			IsValid: true,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			client := Client{}

			err := client.Validate(c.JSON)
			if c.IsValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
