package handler

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJsonError(t *testing.T) {
	msg := "test error message"
	expected := []byte(`{"error":"test error message"}`)

	result := jsonError(msg)

	var actual map[string]interface{}
	err := json.Unmarshal(result, &actual)
	require.Nil(t, err)
	require.Equal(t, expected, result)
}
