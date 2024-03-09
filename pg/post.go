package pg

import (
	"context"
	"el-tech-bloggers/api/types"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (pg *postgres) GetPosts(ctx context.Context) ([]types.Post, error) {
	query := "SELECT * FROM post LIMIT 25"
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("couldn't get posts: %w", err)
	}
	defer rows.Close()
	posts := []types.Post{}
	for rows.Next() {
		post := types.Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.Body)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (pg *postgres) GetPost(ctx context.Context, id string) (*types.Post, error) {
	// Using named args instead of positional ex: $1, $2
	query := "SELECT * FROM post WHERE id = @id"
	args := pgx.NamedArgs{
		"id": id,
	}
	post := types.Post{}
	err := pg.db.QueryRow(ctx, query, args).Scan(&post.ID, &post.Title, &post.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to scan row: %w", err)
	}
	return &post, nil
}

func (pg *postgres) CreatePost(ctx context.Context, post types.Post) error {
	query := `INSERT INTO post (title,body) VALUES (@title, @body)`
	args := pgx.NamedArgs{
		"title": post.Title,
		"body":  post.Body,
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}
	return nil
}
