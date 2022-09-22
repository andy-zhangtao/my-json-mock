package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func Bind(r *http.Request, o interface{}) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, o)
	if err != nil {
		return err
	}

	return nil
}
