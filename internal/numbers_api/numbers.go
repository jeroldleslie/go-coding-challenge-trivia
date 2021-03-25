package numbers_api

import "errors"

var (
	ErrBadInput        = errors.New("ERR_BAD_INPUT")
	ErrNoDataAvailable = errors.New("NO_DATA_AVAILABLE")
)

type Number struct {
	Text   string `json:"text" db:"text"`
	Number string `json:"number" db:"number"`
	Found  bool   `json:"found" db:"found"`
	Type   string `json:"type" db:"type"`
}
