package main

import (
	"compareJson/compare"
	"compareJson/validate"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		panic(errors.New("you must inform two json files for comparison"))
	}

	//check if files are valid
	err := validate.Validate(args[1], args[2])
	if err != nil {
		log.Fatal(err)
	}

	isEqual, err := compare.Apply(args[1], args[2])
	if err != nil {
		log.Fatal(err)
	}
	if isEqual {
		fmt.Println("The json files are equals")
	} else {
		fmt.Println("The json files aren't equals")
	}
}
