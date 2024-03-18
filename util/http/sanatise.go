package http

import (
	"errors"
	"net/http"
	"strconv"
)

func SanitizePageableParams(w http.ResponseWriter, pageNoStr, limitStr string) (int, int, error) {
	if pageNoStr == "" || limitStr == "" {
		return 0, 0, errors.New("missing param")
	}

	pageNo, err := strconv.Atoi(pageNoStr)
	if err != nil {
		return 0, 0, errors.New("invalid pageNo")
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return 0, 0, errors.New("invalid limit")
	}

	if pageNo < 1 || limit < 1 {
		return 0, 0, errors.New("invalid amount of pageNo or limit")
	}

	if limit > 300 {
		return 0, 0, errors.New("limit cannot be greater than 300")
	}

	return pageNo, limit, nil
}
