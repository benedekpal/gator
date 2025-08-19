package main

import (
	"fmt"
	"log"

	"github.com/benedekpal/gator/internal/config"
)

func main() {
	cnfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Read config: %+v\n", cnfg)

	cnfg.CurrentUserName = "benedekpal"
	serr := cnfg.SetUser("benedekpal")
	if serr != nil {
		log.Fatalf("error setting username: %v", serr)
	}

	cnfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cnfg)
}
