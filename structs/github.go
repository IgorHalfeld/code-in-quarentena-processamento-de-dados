package structs

type GithubUserResponse struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

type GithubResponse struct {
	ID   int                `json:"id"`
	User GithubUserResponse `json:"user"`
	Body string             `json:"body"`
}
