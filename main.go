package main

import (
	"flag"
	"log"

	"github.com/DexterHD/dip-visa-app/visa"
)

func main() {
	var id int
	flag.IntVar(&id, "id", 0, "Specify application id you want to check")
	flag.Parse()

	err := visa.CheckVisaConfirmation(id)
	if err != nil {
		log.Fatalln(err)
	}
}
