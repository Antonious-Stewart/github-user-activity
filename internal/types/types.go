package types

import (
	"encoding/json"
	"time"
)

type TypeEvent int

const (
	PushEvent TypeEvent = iota
	CreateEvent
	DeleteEvent
	ForkEvent
	GollumEvent
	IssueCommentEvent
	IssuesEvent
	MemberEvent
	PublicEvent
	PullRequestEvent
	PullRequestReviewEvent
	PullRequestReviewCommentEvent
	PullRequestReviewThreadEvent
	ReleaseEvent
	SponsorshipEvent
	WatchEvent
)

func (e TypeEvent) String() string {
	return [...]string{
		"PushEvent", "CreateEvent", "DeleteEvent", "ForkEvent", "GollumEvent", "IssueCommentEvent",
		"IssuesEvent", "MemberEvent", "PublicEvent", "PullRequestEvent", "PullRequestReviewEvent",
		"PullRequestReviewCommentEvent", "PullRequestReviewThreadEvent", "ReleaseEvent", "SponsorshipEvent", "WatchEvent",
	}[e]
}

func (e *TypeEvent) UnmarshalJSON(b []byte) error {
	var t string
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	switch t {
	case "PushEvent":
		*e = PushEvent
	case "CreateEvent":
		*e = CreateEvent
	case "DeleteEvent":
		*e = DeleteEvent
	case "ForkEvent":
		*e = ForkEvent
	case "GollumEvent":
		*e = GollumEvent
	case "IssueCommentEvent":
		*e = IssueCommentEvent
	case "IssuesEvent":
		*e = IssuesEvent
	case "MemberEvent":
		*e = MemberEvent
	case "PublicEvent":
		*e = PublicEvent
	case "PullRequestEvent":
		*e = PullRequestEvent
	case "PullRequestReviewEvent":
		*e = PullRequestReviewEvent
	case "PullRequestReviewCommentEvent":
		*e = PullRequestReviewCommentEvent
	case "PullRequestReviewThreadEvent":
		*e = PullRequestReviewThreadEvent
	case "ReleaseEvent":
		*e = ReleaseEvent
	case "SponsorshipEvent":
		*e = SponsorshipEvent
	case "WatchEvent":
		*e = WatchEvent
	}
	return nil
}

type ApiResponse struct {
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	CreatedAt time.Time `json:"created_at"`
	EventType TypeEvent `json:"type"`
	Payload   struct {
		Commits []struct {
			Message string `json:"message"`
		} `json:"commits"`
	} `json:"payload"`
}
