package usecase

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/backend-magang/cats-social-media/utils/constant"
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

func builFilterAgeRequest(req *entity.GetListCatRequest) {
	operatorAge, valueAge := parseFilterAge(req.Age)
	req.AgeOperator = operatorAge
	req.AgeValue = cast.ToString(valueAge)
}

func parseFilterAge(filter string) (operator string, value int) {
	parts := strings.Split(filter, "=")
	if len(parts) != 2 {
		return
	}

	operator = string(parts[0][len(parts[0])-1])
	value, _ = strconv.Atoi(parts[1])

	return operator, value
}
