package responses

// Todo ...
type Todo struct {
	ID    int    `json:"id" example:"1"`
	Title string `json:"title" example:"title1"`
	Body  string `json:"body" example:"body1"`
}
