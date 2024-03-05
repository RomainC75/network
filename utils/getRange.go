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

func (a *Addresses) GetEveryAddresses(min string, max string) []int {

}

func GetAddresses(min string, max string) []int {

}

func GetIntAddresses(min []int, max []int) []int {
	for i, val := range max {
		if min[i] < val && i < 3 {
			cpy := make([]int, len(min))
			copy(cpy, min)
		}
	}
}
