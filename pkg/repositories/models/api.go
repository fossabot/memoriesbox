package models

// Memory is used for response to the view
type Memory struct {
	ID      int    `json:"id"`
	Author  string `json:"author"`
	Content string `json:"content"`
	Append  string `json:"append"`
}
