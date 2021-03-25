package main

import (
	"encoding/json"
	"github.com/jeroldleslie/go-coding-challenge-trivia/internal/numbers_api"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDetail(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/api/v1/trivia", nil)
	q := req.URL.Query()
	q.Add("number", "17")
	req.URL.RawQuery = q.Encode()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := &numbers_api.APIHandler{&numbers_api.APIMock{
		DBMapper: &numbers_api.DBMapper{
			DB: nil,
		},
	}}

	var expectedJSONStr = `[{"text":"5300 is the number of gum wrappers that Steve Fletcher has, the record for the largest gum wrapper collection.","number":"5300","found":true,"type":"trivia"}]`
	var actual []numbers_api.Number
	var expected []numbers_api.Number

	err := json.Unmarshal([]byte(expectedJSONStr), &expected)
	if err != nil {
		t.Error(err)
	}
	
	if assert.NoError(t, h.GetDetails(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		err := json.Unmarshal([]byte(rec.Body.String()), &actual)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, expected, actual)
	}
}
