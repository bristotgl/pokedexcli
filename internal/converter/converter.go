package converter

import (
	"bytes"
	"encoding/gob"
)

func ToBytes[T any](data T) ([]byte, error) {
	var buf bytes.Buffer

	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func FromBytes[T any](dataBytes []byte) (T, error) {
	var data T

	decoder := gob.NewDecoder(bytes.NewReader(dataBytes))
	if err := decoder.Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}
