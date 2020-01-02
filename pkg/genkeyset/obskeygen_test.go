/*
Copyright 2020 The genkeyset Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
			fmt.Sprintf("%d%02d%02d", dt.Year()-2000, int(dt.Month()), dt.Day()),
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
