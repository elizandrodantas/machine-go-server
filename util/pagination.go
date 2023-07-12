package util

const (
	PAGINATION_LIMIT = 50

	PAGINATION_PAGE_MIN = 1
)

func ResolvePagination(queryPage string) int {
	var page int

	if len(queryPage) == 0 {
		return PAGINATION_PAGE_MIN - 1
	}

	page, ok := StringToInt(queryPage)

	if !ok {
		return PAGINATION_PAGE_MIN - 1
	}

	if page < PAGINATION_PAGE_MIN {
		return PAGINATION_PAGE_MIN - 1
	}

	return page - 1
}
