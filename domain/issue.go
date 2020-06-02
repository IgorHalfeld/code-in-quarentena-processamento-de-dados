package domain

type Issue struct {
	ID    string `json:"id" bson:"id"`
	Owner string `json:"owner" bson:"owner"`
	Repo  string `json:"repo" bson:"repo"`
	Url   string `json:"url" bson:"url"`
	Body  string `json:"body" bson:"body"`
}
