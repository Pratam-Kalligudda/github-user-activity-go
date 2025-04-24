package main

import (
	"fmt"

	"github.com/Pratam-Kalligudda/github-user-activity-go/cmd"
)

func main() {
	if err := cmd.Excecute(); err != nil {
		fmt.Printf("Error : \n %v", err)
	}
}
