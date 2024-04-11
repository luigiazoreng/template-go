package utils

import (
	"encoding/json"
	"io"
)

func DecodeIO(reader io.Reader, s interface{}) error {
	e := json.NewDecoder(reader)
	return e.Decode(s)
}

func EncodeIO(writer io.Writer, s interface{}) error {
	e := json.NewEncoder(writer)
	return e.Encode(s)
}