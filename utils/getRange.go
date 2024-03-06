package utils

import "sync"

type Addresses struct {
	addresses []string
	m         sync.Mutex
}

func New() *Addresses {
	return &Addresses{
		addresses: []string{},
	}
}

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

// func GetAddresses(min string, max string) []int {

// }

// func GetIntAddresses(min []int, max []int) []int {
// 	for i, val := range max {
// 		if min[i] < val && i < 3 {
// 			cpy := make([]int, len(min))
// 			copy(cpy, min)
// 		}
// 	}

// }
