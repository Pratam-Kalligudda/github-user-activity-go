package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Pratam-Kalligudda/github-user-activity-go/model"
)

func GetUserActivity(uName string) error {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", uName))
	if err != nil {
		return fmt.Errorf("couldnt Connect to backend : %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("couldnt Read data : %v", err)
	}
	var events []model.Event

	if err = json.Unmarshal(body, &events); err != nil {
		return fmt.Errorf("couldnt unmarshal data : %v", err)
	}
	if len(events) == 0 {
		return fmt.Errorf("no data in events")
	}

	for _, event := range events {
		GetTypeBasedActivity(event)
	}

	// result, _ := json.MarshalIndent(events, "", "  ")
	return nil
	// fmt.Println(string(result))
}

func GetTypeBasedActivity(event model.Event) {
	switch event.Type {
	case "PushEvent":
		GetPushEvent(event)
	case "CreateEvent":
		GetCreateEvent(event)
	case "IssuesEvent":
		GetIssueEvent(event)
	case "PullRequestEvent":
		GetPulledRequestEvent(event)
	}
}

func GetPushEvent(event model.Event) {
	var repo_name string
	if event.Repo.ID == *event.Payload.Repo_id {
		repo_name = event.Repo.Name
	}
	fmt.Printf("Pushed %d commits to %s\n", len(*event.Payload.Commits), repo_name)
}

func GetIssueEvent(event model.Event) {
	fmt.Printf("Issue %s by %s on %s\n", *event.Payload.Action, event.Actor.Display_Login, event.Repo.Name)
}

func GetCreateEvent(event model.Event) {
	if *event.Payload.Ref_Type == "branch" {
		fmt.Printf("Branch `%s` created by %s on %s\n", *event.Payload.Ref, event.Actor.Display_Login, event.Repo.Name)
	} else if *event.Payload.Ref_Type == "repository" {
		fmt.Printf("Repository `%s` created by %s \n", event.Repo.Name, event.Actor.Display_Login)
	}
}

func GetPulledRequestEvent(event model.Event) {
	fmt.Printf("PR %s by %s on %s\n", *event.Payload.Action, event.Payload.Pull_Request.User.Login, event.Repo.Name)
}
