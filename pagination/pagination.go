package pagination

import (
	"fmt"
	"strconv"

	"github.com/bongochat/utils/resterrors"
)

func GetNextPage(router string, count int, currentPage int64, limit int64, SITE_URL string) string {
	servedData := currentPage * limit

	if int64(count) > servedData {
		nextPage := currentPage + 1
		return fmt.Sprintf("%s%s/?page=%d&limit=%d", SITE_URL, router, nextPage, limit)
	}
	return ""
}

func GetPreviousPage(router string, count int, currentPage int64, limit int64, SITE_URL string) string {
	servedData := currentPage * limit

	if int64(count) <= servedData {
		nextPage := currentPage - 1
		if nextPage <= 0 {
			return ""
		}
		return fmt.Sprintf("%s%s/?page=%d&limit=%d", SITE_URL, router, nextPage, limit)
	}
	return ""
}

func GetPageNumber(pageParam string) (int64, resterrors.RestError) {
	pageNo, pageErr := strconv.ParseInt(pageParam, 10, 64)
	if pageErr != nil {
		return 1, resterrors.NewBadRequestError("page number should be a number", "")
	}
	if pageNo <= 0 {
		return 1, resterrors.NewBadRequestError("page number should be greater than 0", "")
	}
	return pageNo, nil
}

func GetLimitNumber(limitParam string) (int64, resterrors.RestError) {
	limit, limitErr := strconv.ParseInt(limitParam, 10, 64)
	if limitErr != nil {
		return 1, resterrors.NewBadRequestError("limit number should be a number", "")
	}
	if limit <= 0 {
		return 1, resterrors.NewBadRequestError("limit number should be greater than 0", "")
	}
	return limit, nil
}
