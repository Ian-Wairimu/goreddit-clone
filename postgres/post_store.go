package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"wairimuian.com/GoReddit"
)

func NewPostStore(db *sqlx.DB) *PostStore {
	return &PostStore{
		DB: db,
	}
}

type PostStore struct {
	*sqlx.DB
}

func (p PostStore) Post(id uuid.UUID) (GoReddit.Post, error) {
	var t GoReddit.Post
	if err := p.Get(&t, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return GoReddit.Post{}, fmt.Errorf("error getting post by id: %w", err)
	}
	return t, nil
}

func (p PostStore) PostByThread(threadId uuid.UUID) ([]GoReddit.Post, error) {
	var t []GoReddit.Post
	if err := p.Select(&t, `SELECT * FROM posts WHERE id = $1`, threadId); err != nil {
		return []GoReddit.Post{}, fmt.Errorf("error getting all: %w", err)
	}
	return t, nil
}

func (p PostStore) CreatePost(t *GoReddit.Post) error {
	err := p.Get(t, `INSERT INTO posts VALUES ($1, $2, $3, $4) RETURN *`,
		t.ID,
		t.Title,
		t.Content,
		t.Votes)
	if err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}
	return nil
}

func (p PostStore) UpdatePost(t *GoReddit.Post) error {
	err := p.Get(t, `UPDATE posts SET thread_id = $1 title = $2, content =  $3, votes = $4 WHERE id = $5 RETURNING *`,
		t.ThreadID,
		t.Title,
		t.Content,
		t.Votes,
		t.ID)
	if err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}
	return nil
}

func (p PostStore) DeletePost(id uuid.UUID) error {
	_, err := p.Exec(`DELETE FROM posts WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
