package numbers_api

import (
	"context"
)

type API struct {
	DBMapper *DBMapper
}

//GetDetails function used to fetch numbers
//Args number and text
func (a *API) GetDetails(ctx context.Context, number string, text string) ([]Number, error) {
	numbers, err := a.DBMapper.fetchDetails(ctx, number, text)
	return numbers, err
}
