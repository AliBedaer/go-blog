package helpers

import (
	"fmt"
	"log"
)

func ValidateError(err error) {

	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}

}
