package utils

import (
	"fmt"
	"testing"

	"github.com/RomainC75/network/cases"
	"github.com/stretchr/testify/require"
)

func TestIsMinBeforeMax(t *testing.T) {
	for _, test := range cases.Compare {
		res := IsMinBeforeMax(test.Min, test.Max)
		fmt.Println("-> ", test, res)
		require.Equal(t, test.Expected, res)
	}
}
