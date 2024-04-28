package usecase

import (
	"strconv"
	"strings"

	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/spf13/cast"
)

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
