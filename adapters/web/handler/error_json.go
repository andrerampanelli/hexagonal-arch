package handler

import "encoding/json"

func jsonError(msg string) []byte {
	strErr := struct {
		Error string `json:"error"`
	}{
		Error: msg,
	}

	r, err := json.Marshal(strErr)
	if err != nil {
		return []byte(`{"error": "internal server error"}`)
	}
	return r
}
