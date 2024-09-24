package printer

import (
	"fmt"
	"log"

	"github.com/Antonious-Stewart/github-user-activity/internal/types"
)

func setEventMessage(response types.ApiResponse, message *string) {
	switch response.EventType.String() {
	case "PushEvent":
		length := len(response.Payload.Commits)
		if length == 1 {
			*message = fmt.Sprintf("- Pushed %d commit to %v", length, response.Repo.Name)
			return
		}
		*message = fmt.Sprintf("- Pushed %d commits to %v", length, response.Repo.Name)
	case "CreateEvent":
		*message = fmt.Sprintf("- Created %v on %v", response.Repo.Name, response.CreatedAt)
	case "DeleteEvent":
		*message = fmt.Sprintf("- Deleted %v on %v", response.Repo.Name, response.CreatedAt)
	case "ForkEvent":
		*message = fmt.Sprintf("- Forked %v on %v", response.Repo.Name, response.CreatedAt)
	case "GollumEvent":
		*message = fmt.Sprintf("- Wiki updated for %v", response.Repo.Name)
	case "IssueCommentEvent":
		*message = fmt.Sprintf("- Commented on issue for %v", response.Repo.Name)
	case "IssuesEvent":
		*message = fmt.Sprintf("- Opened issue for %v", response.Repo.Name)
	case "MemberEvent":
		*message = fmt.Sprintf("- Added a member to %v", response.Repo.Name)
	case "PublicEvent":
		*message = fmt.Sprintf("- Open sourced %v", response.Repo.Name)
	case "PullRequestEvent":
		*message = fmt.Sprintf("- Opened pull request for %v", response.Repo.Name)
	case "PullRequestReviewEvent":
		*message = fmt.Sprintf("- Reviewed pull request for %v", response.Repo.Name)
	case "PullRequestReviewCommentEvent":
		*message = fmt.Sprintf("- Commented on pull request for %v", response.Repo.Name)
	case "PullRequestReviewThreadEvent":
		*message = fmt.Sprintf("- Threaded comment on pull request for %v", response.Repo.Name)
	case "ReleaseEvent":
		*message = fmt.Sprintf("- Released %v", response.Repo.Name)
	case "SponsorshipEvent":
		*message = fmt.Sprintf("- Sponsored %v", response.Repo.Name)
	case "WatchEvent":
		*message = fmt.Sprintf("- Starred %v", response.Repo.Name)
	default:
		log.Fatal("Event type not found")
	}
}

func PrintMessage(data *[]types.ApiResponse) {
	var message string
	log.Println("Output:")

	for _, event := range *data {
		setEventMessage(event, &message)
		log.Println(message)
	}
}
