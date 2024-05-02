package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/backend-magang/cats-social-media/utils/constant"
	"github.com/backend-magang/cats-social-media/utils/helper"
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
	operators := []string{"=>", "<=", "="}

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