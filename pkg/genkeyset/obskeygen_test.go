package genkeyset

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGenerateKeyID(t *testing.T) {
	dt := time.Now()

	var inputs = []struct {
		name   string
		input  time.Time
		output string
	}{
		{
			"04-April-2018",
			time.Date(2018, 4, 4, 0, 0, 0, 0, time.UTC),
			"180404",
		},
		{
			"now",
			dt,
			fmt.Sprintf("%d%d%d", dt.Year()-2000, int(dt.Month()), dt.Day()),
		},
	}

	for _, tc := range inputs {
		t.Run(tc.name, func(t *testing.T) {
			actual := GenerateKeyID(tc.input)
			assert.Equal(t, tc.output, actual, "Date format fails to meet expectations.")
		})
	}
}

func TestGenerateKeyset(t *testing.T) {
	keyset, err := GenerateKeySet(3)
	if err != nil {
		fmt.Printf("Error generating JWK: %s", err)
		t.Fail()
	}

	assert.True(t, len(keyset.Keys) == 3, "Failed to generate 3 keys")
}
