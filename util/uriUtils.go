package util

import (
	"encoding/json"
	"io"
	"net/http"
)

/*
 * URI Utils
 *
 * Copyright (C) 2023 Louis Trevino, Torino Consulting, Ltd.
 */

func GetUriContent(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func GetMapFromUri(uri string) (map[string]string, error) {
	myMap := make(map[string]string)
	bytes, err := GetUriContent(uri)
	if err != nil {
		return myMap, err
	}
	err = json.Unmarshal(bytes, &myMap)
	return myMap, err
}
