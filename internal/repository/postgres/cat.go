package postgres

import (
	"context"
	"database/sql"

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
