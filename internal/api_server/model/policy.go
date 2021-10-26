package model

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/ory/ladon"
	"gorm.io/gorm"
)

type Policy struct {
	gorm.DB
	Policy AuthzPolicy
}

type AuthzPolicy struct {
	ladon.Policy
}

func (a *AuthzPolicy) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, &a)
}

func (a AuthzPolicy) Value() (value driver.Value, err error) {
	bytes, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}

func (Policy) TableName() string {
	return "policy"
}
