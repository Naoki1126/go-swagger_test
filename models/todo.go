package models

// Todo ...
type Todo struct {
	ID    int    `json:"id" example:"1"`
	Title string `json:"title" example:"title1"`
	Body  string `json:"body" example:"body1"`
}

func PostTodo() [2]string {
	ar := [2]string{"goLang", "postTodo"}
	return ar
}
