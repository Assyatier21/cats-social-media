package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/backend-magang/cats-social-media/models/entity"
)

func (r *repository) InsertUser(ctx context.Context, req entity.User) (err error) {

	query := `INSERT INTO users (email, name, password, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING *`

	_, err = r.db.ExecContext(ctx,
		query,
		req.Email,
		req.Name,
		req.Password,
		req.CreatedAt,
		req.UpdatedAt,
	)

	if err != nil {
		log.Println("[Repository][User][InsertUser] failed to insert new user, err: ", err.Error())
		return
	}

	return
}

func (r *repository) FindUserByEmail(ctx context.Context, email string) (result entity.User, err error) {
	query := `
		SELECT * FROM users WHERE 
		email = $1
	`

	err = r.db.QueryRowxContext(ctx, query, email).StructScan(&result)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[Repository][User][InsertUser] failed to find user by email "+email+", err: ", err.Error())
		return
	}

	return
}
