package models

import (
	"log"
)

type Snippet struct {
	Id          int64
	Name        string `binding:"required" json:"name"`
	Description string `json:"description"`
	Language    string `binding:"required" json:"language"`
	Category    string `binding:"required" json:"category"`
	Userid      int    `binding:"required" json:"userid"`
}

func GetSnippets() ([]Snippet, error) {
	var snippets []Snippet
	query := `select id, name, description, language, category, userid from snippets`

	res, err := db.Query(query)
	if err != nil {
		log.Println("error fetching snippets from database")
		return nil, err
	}

	for res.Next() {
		var snippet Snippet
		err := res.Scan(&snippet.Id, &snippet.Name, &snippet.Description, &snippet.Language, &snippet.Category, &snippet.Userid)
		if err != nil {
			log.Println("error parsing user data from database")
			return nil, err
		}

		snippets = append(snippets, snippet)
	}

	return snippets, nil
}

func GetSnippet(snippetid int64) (*Snippet, error) {
	var snippet Snippet

	query := `select id, name, description, language, category, userid from snippets where id = $1`

	err := db.QueryRow(query, snippetid).Scan(&snippet.Id, &snippet.Name, &snippet.Description, &snippet.Language, &snippet.Category, &snippet.Userid)
	if err != nil {
		log.Println("error fetching user from database")
		return nil, err
	}

	return &snippet, nil
}

func (snippet *Snippet) AddSnippet() error {
	query := `insert into snippet(name, description, language, category, userid) values ($1, $2, $3, $4, $5)`

	err := db.QueryRow(query, snippet.Name, snippet.Description, snippet.Language, snippet.Category, snippet.Userid).Scan(&snippet.Id)
	if err != nil {
		log.Println("error adding snippet to database")
		return err
	}

	return nil
}

func (snippet *Snippet) UpdateSnippet() error {
	query := `update snippets set name=$1, description=$2, language=$3, category=$4, userid=$5 where id=$6`

	err := db.QueryRow(query, snippet.Name, snippet.Description, snippet.Language, snippet.Category, snippet.Userid).Scan(&snippet.Id)
	if err != nil {
		log.Println("error updating snippet in database")
		return err
	}

	return nil
}

func (snippet *Snippet) DeleteSnippet() error {
	query := `delete from snippets where id=$1`

	_, err := db.Query(query, snippet.Id)
	if err != nil {
		log.Println("error deleting user from database")
		return err
	}

	return nil
}
