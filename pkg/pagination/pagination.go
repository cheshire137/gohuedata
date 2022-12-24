package pagination

const MaxPerPage = 500

func ConstrainPageAndPerPage(page int, perPage int) (int, int) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 1
	}
	if perPage > MaxPerPage {
		perPage = MaxPerPage
	}
	return page, perPage
}

func GetLimitAndOffset(page int, perPage int) (int, int) {
	limit := perPage
	offset := (page - 1) * perPage
	return limit, offset
}
