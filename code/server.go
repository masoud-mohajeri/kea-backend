package main

import (
	"fmt"

	"github.com/masoud-mohajeri/kea-backend/app"
)

func main() {

	// TODO: add logger
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()

	app.StartApplication()
}
