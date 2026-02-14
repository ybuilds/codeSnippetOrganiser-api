package models

type Snippet struct {
	Id          int
	Name        string `binding:"required" json:"name"`
	Description string `binding:"required" json:"description"`
	Language    string `binding:"required" json:"language"`
	Category    string `binding:"required" json:"category"`
	Userid      int    `binding:"required" json:"userid"`
}
