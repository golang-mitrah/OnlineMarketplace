package main

import (
	"fmt"
	"onlinemarketplace/driver"
	"onlinemarketplace/router"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in panic", r)
		}
	}()
	driver.NewDBConn(false)
	r := router.NewRouter()
	r.Run()
}
