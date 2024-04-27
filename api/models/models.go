package models

import "encoding/json"

// Converter function
func JSONModelConverter(input, output any) error {
	dataBytes, err := json.Marshal(input)
	if err != nil {
		return err
	}
	err = json.Unmarshal(dataBytes, &output)
	if err != nil {
		return err
	}
	return nil
}
