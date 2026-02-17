package models

import (
	"database/sql"
	"errors"
	"log"
)

type Snippet struct {
	Id       int64
	Name     string `binding:"required" json:"name"`
	Code     string `json:"code"`
	Language string `binding:"required" json:"language"`
	Category string `binding:"required" json:"category"`
	Userid   int    `binding:"required" json:"userid"`
}

func GetSnippets() ([]Snippet, error) {
	var snippets []Snippet
	query := `select id, name, code, language, category, userid from snippets`

	res, err := db.Query(query)
	if err != nil {
		log.Println("error fetching snippets from database")
		return nil, err
	}

	for res.Next() {
		var snippet Snippet
		err := res.Scan(&snippet.Id, &snippet.Name, &snippet.Code, &snippet.Language, &snippet.Category, &snippet.Userid)
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

	query := `select id, name, code, language, category, userid from snippets where id = $1`

	err := db.QueryRow(query, snippetid).Scan(&snippet.Id, &snippet.Name, &snippet.Code, &snippet.Language, &snippet.Category, &snippet.Userid)
	if err != nil {
		log.Println("error fetching user from database")
		return nil, err
	}

	return &snippet, nil
}

func (snippet *Snippet) AddSnippet() error {
	query := `insert into snippet(name, code, language, category, userid) values ($1, $2, $3, $4, $5)`

	err := db.QueryRow(query, snippet.Name, snippet.Code, snippet.Language, snippet.Category, snippet.Userid).Scan(&snippet.Id)
	if err != nil {
		log.Println("error adding snippet to database")
		return err
	}

	return nil
}

func (snippet *Snippet) UpdateSnippet() error {
	query := `update snippets set name=$1, code=$2, language=$3, category=$4, userid=$5 where id=$6`

	err := db.QueryRow(query, snippet.Name, snippet.Code, snippet.Language, snippet.Category, snippet.Userid).Scan(&snippet.Id)
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

func GetSnippetsByField(query, field string) ([]Snippet, error) {
	var snippets []Snippet

	res, err := db.Query(query, field)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("no rows found")
			return nil, err
		}
		log.Println("error fetching snippets from database")
		return nil, err
	}

	for res.Next() {
		var snippet Snippet
		err = res.Scan(&snippet.Id, &snippet.Name, &snippet.Code, &snippet.Category, &snippet.Language, &snippet.Userid)
		snippets = append(snippets, snippet)
	}

	return snippets, nil
}

func GetSnippetByCategory(category string) ([]Snippet, error) {
	query := `select id, name, code, language, category, userid from snippets where category = $1`
	res, err := GetSnippetsByField(query, category)
	if err != nil {
		log.Println("error fetching snippets from database")
		return nil, err
	}

	return res, nil
}

func GetSnippetByLanguage(language string) ([]Snippet, error) {
	query := `select id, name, code, language, category, userid from snippets where language = $1`
	res, err := GetSnippetsByField(query, language)
	if err != nil {
		log.Println("error fetching snippets from database")
		return nil, err
	}

	return res, nil
}
