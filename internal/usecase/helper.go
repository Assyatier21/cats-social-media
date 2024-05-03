package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/backend-magang/cats-social-media/utils/constant"
	"github.com/backend-magang/cats-social-media/utils/helper"
	"github.com/spf13/cast"
)

func (u *usecase) validateCatsWillBeMatched(ctx context.Context, issuerID int, targetCat, userCat entity.Cat) (err error) {
	if issuerID != userCat.UserID {
		return errors.New(constant.FAILED_CAT_USER_UNAUTHORIZED)
	}

	if targetCat.UserID == userCat.UserID {
		return errors.New(constant.FAILED_CAT_USER_IDENCTIC)
	}

	if targetCat.Sex == userCat.Sex {
		return errors.New(constant.FAILED_CAT_GENDER_IDENTIC)
	}

	if targetCat.IsAlreadyMatched || userCat.IsAlreadyMatched {
		return errors.New(constant.FAILED_CAT_MATCHED)
	}

	return nil
}

func (u *usecase) validateMatchCat(ctx context.Context, targetUserID int, matchCat entity.MatchCat) (err error) {
	if matchCat.Status != entity.MatchCatStatusPending || matchCat.DeletedAt.Valid {
		return errors.New(constant.FAILED_MATCH_NOT_VALID)
	}

	if matchCat.TargetUserID != targetUserID {
		return errors.New(constant.FAILED_CAN_NOT_APPROVE)
	}

	return nil
}

func builFilterAgeRequest(req *entity.GetListCatRequest) (err error) {
	if req.Age == "" {
		return
	}

	operatorAge, valueAge, err := parseFilterAge(req.Age)
	if err != nil {
		return
	}

	req.AgeOperator = operatorAge
	req.AgeValue = valueAge

	return
}

func parseFilterAge(age string) (operator string, value string, err error) {
	operators := []string{">", "<", "="}

	for _, op := range operators {
		if strings.HasPrefix(age, op) {
			operator = op
			value = strings.TrimPrefix(age, op)
			break
		}
	}

	if operator == "" || !helper.IsNumeric(value) {
		return operator, value, errors.New(constant.FAILED_WRONG_AGE_FORMAT)
	}

	return
}

func buildResponseListMatchCat(match_cat entity.GetListMatchCatQueryResponse) entity.GetListMatchCatResponse {
	return entity.GetListMatchCatResponse{
		ID: cast.ToString(match_cat.ID),
		IssuedBy: entity.IssuedByData{
			Name:      match_cat.IssuedByName,
			Email:     match_cat.IssuedByEmail,
			CreatedAt: match_cat.CreatedAt,
		},
		MatchCatDetail: entity.CatDetail{
			ID:          cast.ToString(match_cat.MatchCatID),
			Name:        match_cat.MatchCatName,
			Race:        match_cat.MatchCatRace,
			Sex:         match_cat.MatchCatSex,
			Description: match_cat.MatchCatDescription,
			AgeInMonth:  match_cat.MatchCatAge,
			ImageUrls:   match_cat.MatchCatImages,
			HasMatched:  match_cat.MatchCatHasMatched,
			CreatedAt:   match_cat.MatchCatCreatedAt,
		},
		UserCatDetail: entity.CatDetail{
			ID:          cast.ToString(match_cat.UserCatID),
			Name:        match_cat.UserCatName,
			Race:        match_cat.UserCatRace,
			Sex:         match_cat.UserCatSex,
			Description: match_cat.UserCatDescription,
			AgeInMonth:  match_cat.UserCatAge,
			ImageUrls:   match_cat.UserCatImages,
			HasMatched:  match_cat.UserCatHasMatched,
			CreatedAt:   match_cat.UserCatCreatedAt,
		},
		Message:   match_cat.Message,
		CreatedAt: match_cat.CreatedAt,
	}
}

func buildResponseCat(cat entity.Cat) entity.GetListCatResponse {
	return entity.GetListCatResponse{
		ID:               cast.ToString(cat.ID),
		UserID:           cat.UserID,
		Name:             cat.Name,
		Race:             cat.Race,
		Sex:              cat.Sex,
		Age:              cat.Age,
		Description:      cat.Description,
		Images:           cat.Images,
		IsAlreadyMatched: cat.IsAlreadyMatched,
		CreatedAt:        cat.CreatedAt,
		UpdatedAt:        cat.UpdatedAt,
	}
}
