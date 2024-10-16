package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type ContactInfo struct {
	Contact string `json:"contact"`
	Info    string `json:"info"`
}

func (a ContactInfo) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *ContactInfo) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type Organization struct {
	ID         int
	BrandName  string
	FullName   string
	ShortName  string
	INN        string
	OKPO       string
	ORGN       string
	KPP        string
	TaxCode    int
	Address    string
	AvatarURL  string
	Emails     []ContactInfo
	Phones     []ContactInfo
	Messangers []ContactInfo
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
