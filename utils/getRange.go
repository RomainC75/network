package utils

import (
	"errors"
	"strconv"
	"strings"
)

// type Addresses struct {
// 	addresses []string
// 	m         sync.Mutex
// }

// func New() *Addresses {
// 	return &Addresses{
// 		addresses: []string{},
// 	}
// }

// func (a *Addresses) GetEveryAddresses(min []int, max []int) []int {
// 	for
// }

func IsMinBeforeMax(min []int, max []int) bool {
	for i := 0; i < 4; i++ {
		if min[i] < max[i] {
			return true
		} else if min[i] > max[i] {
			return false
		}
	}
	return true
}

func GetAddresses(min []int, max []int) []string {
	res := []string{}
	for IsMinBeforeMax(min, max) {
		strAddress, err := getStringAddress(min)
		if err == nil {
			res = append(res, strAddress)
		}
		increment(min, max)
	}
	return res
}

func increment(min []int, max []int) {
	isMax := true
	index := 3
	for isMax && index >= 0 {
		if min[index] == 255 {
			min[index] = 0
		} else {
			min[index]++
			isMax = false
		}
		index--
	}
}

func getStringAddress(addr []int) (string, error) {
	strAddr := []string{}
	if len(addr) != 4 {
		return "", errors.New("wrong length")
	}
	for _, intByte := range addr {
		if intByte < 0 || intByte > 255 {
			return "", errors.New("numbers should be between 0 and 255")
		}
		strByte := strconv.Itoa(intByte)
		strAddr = append(strAddr, strByte)
	}
	return strings.Join(strAddr, "."), nil
}

// func GetIntAddresses(min []int, max []int) []int {
// 	for i, val := range max {
// 		if min[i] < val && i < 3 {
// 			cpy := make([]int, len(min))
// 			copy(cpy, min)
// 		}
// 	}

// }
