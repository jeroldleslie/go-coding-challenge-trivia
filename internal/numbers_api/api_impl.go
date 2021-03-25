package numbers_api

import "context"

type APIImpl interface {
	GetDetails(ctx context.Context, number string, text string) ([]Number, error)
}

