package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sciffr/backend"
	"sciffr/frontend"
)

func main() {
	r := gin.Default()

	backend.New().Register(r)
	frontend.New().Register(r)

	err := r.Run()
	if err != nil {
		fmt.Printf("could not start sciffr: %v", err)
	}
	fmt.Printf("started sciffr")
}
