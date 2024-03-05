package utils

import (
	"testing"

	"github.com/RomainC75/network/cases"
	"github.com/stretchr/testify/require"
)

func TestParseAddress(t *testing.T) {

	for _, cas := range cases.Addresses {
		res, err := ParseAdress(cas.Address)
		if cas.Err {
			require.Error(t, err)
		} else {
			require.Equal(t, res, cas.Expected)
		}

	}
}
