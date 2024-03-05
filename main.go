package main

import (
	"fmt"
	"log"

	"github.com/RomainC75/network/utils"
)

// const target1 = "google.com"
const target_min = "212.0.2.43"
const target_max = "212.23.0.45"

func main() {
	// for {
	// 	time.Sleep(time.Second * 1)
	// 	ping.Ping(target)
	// }
	res, err := utils.ParseAdress(target_max)
	if err != nil {
		log.Fatalf("error :  %s\n", err.Error())
	}
	fmt.Println(res)
}
