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

type PullRequest struct {
	Id             string          `json:"id"`
	IsDraft        bool            `json:"isDraft"`
	Number         uint64          `json:"number"`
	State          string          `json:"state"`
	Comments       []Comment       `json:"comments"`
	Commits        []Commit        `json:"commits"`
	ReviewRequests []ReviewRequest `json:"reviewRequests"`
}

type Repository struct {
	Url string `json:"url"`
}
