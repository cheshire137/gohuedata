package pagination

import (
	"net/url"
	"strconv"
)

type PageInfo struct {
	Page    int
	PerPage int
}

func GetPageInfoParams(url *url.URL) (*PageInfo, error) {
	pageStr := url.Query().Get("page")
	var page int
	var err error
	if pageStr == "" {
		page = 1
	} else {
		if page, err = strconv.Atoi(pageStr); err != nil {
			return nil, err
		}
	}
	perPageStr := url.Query().Get("per_page")
	var perPage int
	if perPageStr == "" {
		perPage = 10
	} else {
		if perPage, err = strconv.Atoi(perPageStr); err != nil {
			return nil, err
		}
	}
	return &PageInfo{Page: page, PerPage: perPage}, nil
}
