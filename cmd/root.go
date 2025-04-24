package cmd

import (
	"fmt"
	"os"

	"github.com/Pratam-Kalligudda/github-user-activity-go/github"
)

func Excecute() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("Usage : github-activity <username>")
	}
	return github.GetUserActivity(os.Args[1])
}
