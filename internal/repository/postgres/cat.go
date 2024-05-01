package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/backend-magang/cats-social-media/models/entity"
)

func (r *repository) GetListCat(ctx context.Context, req entity.GetListCatRequest) ([]entity.Cat, error) {
	result := []entity.Cat{}

	query, args := buildQueryGetListCats(req)
	query = r.db.Rebind(query)

	err := r.db.SelectContext(ctx, &result, query, args...)
	if err != nil {
		r.logger.Errorf("[Repository][Cat][GetList] failed to query, err: %s", err.Error())
		return result, err
	}

	return result, err
}

func (r *repository) FindCatByID(ctx context.Context, id int) (entity.Cat, error) {
	result := entity.Cat{}
	query := `SELECT * FROM cats WHERE id = $1`

	err := r.db.QueryRowxContext(ctx, query, id).StructScan(&result)
	if err != nil && err != sql.ErrNoRows {
		r.logger.Errorf("[Repository][Cat][FindByID] failed to query, err: %s", err.Error())
		return result, err
	}

	return result, err
}

func (r *repository) InsertCat(ctx context.Context, req entity.Cat) (result entity.Cat, err error) {
	query := `INSERT INTO cats (user_id, name, race, sex, age, description, images, is_already_matched, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
		RETURNING *`

	err = r.db.QueryRowxContext(ctx,
		query,
		req.UserID,
		req.Name,
		req.Race,
		req.Sex,
		req.Age,
		req.Description,
		req.Images,
		req.IsAlreadyMatched,
		req.CreatedAt,
		req.UpdatedAt,
	).StructScan(&result)

	if err != nil {
		log.Println("[Repository][Cat][InsertCat] failed to insert new cat, err: ", err.Error())
		return
	}

	return
}

func (r *repository) UpdateCat(ctx context.Context, req entity.Cat) (result entity.Cat, err error) {
	query := `UPDATE cats 
		SET name = $1, race = $2, sex = $3, age = $4, description = $5, images = $6, updated_at = $7 
		WHERE id = $8 AND user_id = $9 RETURNING *`

	err = r.db.QueryRowxContext(ctx,
		query,
		req.Name,
		req.Race,
		req.Sex,
		req.Age,
		req.Description,
		req.Images,
		req.UpdatedAt,
		req.ID,
		req.UserID,
	).StructScan(&result)

	if err != nil {
		log.Println("[Repository][Cat][UpdateCat] failed to update cat, err: ", err.Error())
		return
	}

	return
}

func (r *repository) FindUserCatByID(ctx context.Context, userId int, catId int) (result entity.Cat, err error) {
	query := `
		SELECT * FROM cats 
		WHERE user_id = $1 AND id = $2 AND deleted_at IS NULL
	`

	err = r.db.QueryRowxContext(ctx, query, userId, catId).StructScan(&result)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[Repository][Cat][FindUserCatByID] failed to find user cat, err: ", err.Error())
		return
	}

	return
}

func (r *repository) FindRequestedMatch(ctx context.Context, catId int) (result []entity.MatchCat, err error) {
	query := `
		SELECT * FROM match_cats 
		WHERE (match_cat_id = $1 OR user_cat_id = $1) AND deleted_at IS NULL
	`

	err = r.db.SelectContext(ctx, &result, query, catId)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[Repository][Cat][FindRequestedMatch] failed to find requested match, err: ", err.Error())
		return
	}

	return
}
