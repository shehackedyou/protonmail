package api

import "encoding/json"

type Boolean bool

func (boolean *Boolean) UnmarshalJSON(b []byte) error {
	var value int
	err := json.Unmarshal(b, &value)
	if err != nil {
		return err
	}

	*boolean = Boolean(value == 1)
	return nil
}

func (boolean Boolean) MarshalJSON() ([]byte, error) {
	var value int
	if boolean {
		value = 1
	}
	return json.Marshal(value)
}
