package utl

import (
	"fmt"
)

func ErrorLog(err error) {
	if err != nil {
		fmt.Println("Error", err)
	}
}
