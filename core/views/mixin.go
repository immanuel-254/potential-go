package views

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
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

func GetId(route string, w http.ResponseWriter, r *http.Request) (int64, error) {
	idStr := strings.TrimPrefix(r.URL.Path, route)
	idStr = strings.TrimLeft(idStr, "/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return 0, err
	}

	return int64(id), nil
}
