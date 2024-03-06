package main

import (
	"fmt"
	"log"

	"github.com/RomainC75/network/utils"
)

// const target1 = "google.com"
const target_min = "212.0.2.43"
const target_max = "212.0.3.43"

func main() {
	// for {
	// 	time.Sleep(time.Second * 1)
	// 	ping.Ping(target)
	// }
	min, err := utils.ParseAdress(target_min)
	if err != nil {
		log.Fatalf("error :  %s\n", err.Error())
	}
	max, err := utils.ParseAdress(target_max)
	if err != nil {
		log.Fatalf("error :  %s\n", err.Error())
	}
	res := utils.GetAddresses(min, max)

	fmt.Println(res)
	fmt.Println("len : ", len(res))
}
