package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"wairimuian.com/GoReddit"
)

func NewThreadStore(db *sqlx.DB) *ThreadStore {
	return &ThreadStore{
		DB: db,
	}
}

type ThreadStore struct {
	*sqlx.DB
}

func (t *ThreadStore) Thread(id uuid.UUID) (GoReddit.Thread, error) {
	var s GoReddit.Thread
	err := t.Get(&s, `SELECT * FROM threads WHERE id=$1`, id)
	if err != nil {
		return GoReddit.Thread{}, fmt.Errorf("error getting thread: %w", err)
	}
	return s, nil
}

func (t *ThreadStore) Threads() ([]GoReddit.Thread, error) {
	var tt []GoReddit.Thread
	err := t.Select(&tt, `SELECT * FROM threads`)
	if err != nil {
		return []GoReddit.Thread{}, fmt.Errorf("error getting threads: %w", err)
	}
	return tt, nil
}

func (s *ThreadStore) CreateThread(t *GoReddit.Thread) error {
	err := s.Get(t, `INSERT INTO threads VALUES ($1, $2, $3) RETURN *`,
		t.ID,
		t.Title,
		t.Content)
	if err != nil {
		return fmt.Errorf("error creating threads: %w", err)
	}
	return nil
}

func (s *ThreadStore) UpdateThread(t *GoReddit.Thread) error {
	err := s.Get(t, `UPDATE threads SET title = $1, content = $2 WHERE id = $3 RETURNING *`,
		t.Title,
		t.Content,
		t.ID)
	if err != nil {
		return fmt.Errorf("error updating thread: %w", err)
	}
	return nil
}

func (s *ThreadStore) DeleteThread(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM threads WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting thread: %w", err)
	}
	return nil
}
