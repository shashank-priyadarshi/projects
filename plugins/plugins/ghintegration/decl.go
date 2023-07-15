package ghintegration

import (
	config "plugins/common"
	"time"
)

var (
	authToken   = config.FetchConfig().GITHUBTOKEN
	userName    = config.FetchConfig().GITHUBUSERNAME
	days        = int(time.Now().AddDate(0, 0, -31).Unix() / (24 * 60 * 60)) // adding -31d to today's date, converting to days, gives days elapsed from Jan 1, 1970
	queryParams = []string{"per_page 100", "page 1"}
)

type RepoResponse struct {
	Name  string `json:"name"`
	URL   string `json:"html_url"`
	Owner struct {
		Login string `json:"login"`
	} `json:"owner"`
	Forked     bool   `json:"fork"`
	CommitsURL string `json:"commits_url"`
	PRURL      string `json:"pulls_url"`
	PushedAt   string `json:"pushed_at"`
}

type CommitResponse struct {
	Commit struct {
		Committer struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Date  string `json:"date"`
		} `json:"committer"`
	} `json:"commit"`
}

type PullRequestResponse struct {
	State string `json:"state"`
	User  struct {
		Login string `json:"login"`
	} `json:"user"`
	CreatedAt string `json:"created_at"`
}

type IssueRequest struct {
	URL       string `json:"url"`
	Title     string `json:"title"`
	State     string `json:"state"`
	CreatedAt string `json:"created_at"`
	ClosedAt  string `json:"closed_at"`
	Assignee  struct {
		Login string `json:"login"`
	} `json:"assignee"`
}

type GitHubData struct {
	Time         time.Time     `json:"execution_time"`
	StarredRepos Repo          `json:"starredrepos"`
	WeekData     []SCMActivity `json:"weekdata"`
}

type GraphData struct {
	WeekData []SCMActivity `json:"weekdata"`
}

type Repo struct {
	Count int      `json:"count"`
	List  []string `json:"list"`
}

type SCMActivity struct {
	PR     int    `json:"pr"`
	LOC    int    `json:"loc"`
	Date   string `json:"date"`
	Commit int    `json:"commit"`
}
