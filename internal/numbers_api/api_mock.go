package numbers_api

import (
	"context"
)

type APIMock struct {
	DBMapper *DBMapper
}

func (a *APIMock) GetDetails(ctx context.Context, number string, text string) ([]Number, error) {
	numbers := []Number{
		{
			Text:   "5300 is the number of gum wrappers that Steve Fletcher has, the record for the largest gum wrapper collection.",
			Number: "5300",
			Found:  true,
			Type:   "trivia",
		},
	}
	return numbers, nil
}
