package db

import "time"

type DbTime struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type IdType struct {
	Id uint64 `json:"id,omitempty"`
}

type User struct {
	IdType
	Name      string `json:"name"`
	Timestamp uint64 `json:"timestamp"`
}
