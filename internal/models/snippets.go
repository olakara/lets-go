package models

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Snippet struct {
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}

func (s *Snippet) ToString() string {
	return "ID: " + strconv.Itoa(s.ID) + ", Title: " + s.Title + ", Content: " + s.Content + ", Created: " + s.Created.Format(time.RFC3339) + ", Expires: " + s.Expires.Format(time.RFC3339)
}

type SnippetModel struct {
	DB *pgxpool.Pool
}

func (m *SnippetModel) Insert(title, content string, expires int) (uuid.UUID, error) {
	id := uuid.New()
	stmt := `INSERT INTO snippets (id, title, content, created, expires) VALUES ($1, $2, $3, $4, $5)`
	_, err := m.DB.Exec(context.Background(), stmt, id, title, content, time.Now(), time.Now().AddDate(0, 0, expires))
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE id = $1`
	row := m.DB.QueryRow(context.Background(), stmt, id)
	snippet := &Snippet{}
	err := row.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return snippet, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*Snippet{}
	for rows.Next() {
		snippet := &Snippet{}
		err := rows.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, snippet)
	}
	return snippets, nil
}