package postgres

import (
	"context"
	"database/sql"

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
		r.logger.Errorf("[Repository][User][InsertUser] failed to insert new user, err: %s", err.Error())
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
		r.logger.Errorf("[Repository][User][FindByEmail] failed to find user by email %s, err: %s", email, err.Error())
		return
	}

	return
}
