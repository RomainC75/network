package utils

import (
	"errors"
	"strconv"
	"strings"
)

func ParseAdress(address string) ([]int, error) {
	parsedStr := strings.Split(address, ".")
	numbers := []int{}
	counter := 0
	for _, strNum := range parsedStr {
		intNum, err := strconv.Atoi(strNum)
		if err != nil || intNum < 0 || intNum > 255 {
			return []int{}, errors.New("not parsable")
		}
		numbers = append(numbers, intNum)
		counter++
	}
	if counter != 4 {
		return []int{}, errors.New("not parsable")
	}
	return numbers, nil
}
