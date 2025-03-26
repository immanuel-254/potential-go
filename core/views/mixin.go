package views

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func GetData(data *map[string]string, w http.ResponseWriter, r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		if errors.Is(err, io.EOF) {
			w.WriteHeader(http.StatusBadRequest)
			return errors.New("empty request body")
		}
		if _, ok := err.(*json.SyntaxError); ok {
			w.WriteHeader(http.StatusBadRequest)
			return errors.New("invalid json syntax")
		}
		w.WriteHeader(http.StatusBadRequest)
		return errors.New(strings.ToLower(err.Error()))
	}

	// Check if the decoded data is empty
	if data == nil || len(*data) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return errors.New("no data provided")
	}

	return nil
}
