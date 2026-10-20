package storage

import (
	"encoding/json"
	"os"
)

func LoadLibraryFromJSON(filePath string) (interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lib interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&lib); err != nil {
		return nil, err
	}

	return lib, nil
}

func SaveLibraryToJSON(filePath string, lib interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(lib)
}
