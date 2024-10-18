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

type ContactInfos []ContactInfo

func (a ContactInfos) Value() (driver.Value, error) {
	if a == nil {
		return []byte("[]"), nil
	}

	return json.Marshal(a)
}

func (a *ContactInfos) Scan(value interface{}) error {
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
	TaxCode    string
	Address    string
	AvatarURL  string
	Emails     ContactInfos
	Phones     ContactInfos
	Messangers ContactInfos
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
