package main

type Author struct {
	Login string `json:"login"`
}

type Comment struct {
	Id        string `json:"id"`
	Author    Author `json:"author"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
}

type Commit struct {
	Oid string `json:"oid"`
}

type ReviewRequest struct {
	Typename string `json:"__typename"`
	Login    string `json:"login,omitempty"`
}

type HeadRepositoryOwner struct {
	Login string `json:"login"`
}

type HeadRepository struct {
	Name string `json:"name"`
}

type PullRequest struct {
	Id             string              `json:"id"`
	IsDraft        bool                `json:"isDraft"`
	Number         uint64              `json:"number"`
	State          string              `json:"state"`
	Comments       []Comment           `json:"comments"`
	Commits        []Commit            `json:"commits"`
	ReviewRequests []ReviewRequest     `json:"reviewRequests"`
	Repository     HeadRepository      `json:"headRepository"`
	Owner          HeadRepositoryOwner `json:"headRepositoryOwner"`
}

type Repository struct {
	Url string `json:"url"`
}

type ApiRequestedReviewer struct {
	Login string `json:"login"`
	Type  string `json:"type"`
}

type ApiRequestedTeam struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ApiPullRequest struct {
	Id            uint64                 `json:"id"`
	Number        uint64                 `json:"number"`
	Url           string                 `json:"url"`
	State         string                 `json:"state"`
	Locked        bool                   `json:"locked"`
	Reviewers     []ApiRequestedReviewer `json:"requested_reviewers"`
	ReviewerTeams []ApiRequestedTeam     `json:"requested_teams"`
}
