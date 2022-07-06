package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"wairimuian.com/GoReddit"
)

type CommentStore struct {
	*sqlx.DB
}

func (c CommentStore) Comment(id uuid.UUID) (GoReddit.Comment, error) {
	var t GoReddit.Comment
	if err := c.Get(&t, `SELECT * FROM comments WHERE id = $1`, id); err != nil {
		return GoReddit.Comment{}, fmt.Errorf("error getting comment by id: %w", err)
	}
	return t, nil
}

func (c CommentStore) CommentsByPostId(postId uuid.UUID) ([]GoReddit.Comment, error) {
	var t []GoReddit.Comment
	if err := c.Select(&t, `SELECT * FROM comments WHERE id = 1`, postId); err != nil {
		return []GoReddit.Comment{}, fmt.Errorf("error getting all comments: %w", err)
	}
	return t, nil
}

func (c CommentStore) CreateComment(t *GoReddit.Comment) error {
	if err := c.Get(t, `INSERT INTO comments VALUES ($1, $2, $3, $4) RETURN *`,
		t.ID,
		t.Content,
		t.Votes); err != nil {
		return fmt.Errorf("error creating the post: %w", err)
	}
	return nil
}

func (c CommentStore) UpdateComment(t *GoReddit.Comment) error {
	if err := c.Get(t, `UPDATE comments SET post_id = $1 content = $2, votes = $3 WHERE id = $4 RETURNING *`,
		t.PostID,
		t.Content,
		t.Votes,
		t.ID); err != nil {
		return fmt.Errorf("error updating the post: %w", err)
	}
	return nil
}

func (c CommentStore) DeleteComment(id uuid.UUID) error {
	if _, err := c.Exec(`DELETE FROM comments WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting comments: %w", err)
	}
	return nil
}
