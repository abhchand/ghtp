package main

type PullRequest struct {
	Id      int    `json:"number"`
	HtmlUrl string `json:"html_url"`
	Title   string `json:"title"`
}

type PullRequestList []PullRequest

func (pr *PullRequest) ShouldSync() bool {
	// Is PR & Title regex
	return true
}
