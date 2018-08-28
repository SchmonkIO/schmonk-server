package models

import "encoding/json"

type BaseAction struct {
	Action string `json:"action"`
}

func (ba *BaseAction) Marshal() ([]byte, error) {
	bytes, err := json.Marshal(ba)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (ba *BaseAction) Unmarshal(bytes []byte) error {
	err := json.Unmarshal(bytes, ba)
	return err
}

func (ba *BaseAction) Check(action string) bool {
	if ba.Action == action {
		return true
	}
	return false
}
