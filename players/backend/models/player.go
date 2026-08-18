package models

import "time"

type Player struct {
	PlayerID     string     `json:"playerID"`
	BirthYear    *int32     `json:"birthYear,omitempty"`
	BirthMonth   *int32     `json:"birthMonth,omitempty"`
	BirthDay     *int32     `json:"birthDay,omitempty"`
	BirthCountry *string    `json:"birthCountry,omitempty"`
	BirthState   *string    `json:"birthState,omitempty"`
	BirthCity    *string    `json:"birthCity,omitempty"`
	DeathYear    *int32     `json:"deathYear,omitempty"`
	DeathMonth   *int32     `json:"deathMonth,omitempty"`
	DeathDay     *int32     `json:"deathDay,omitempty"`
	DeathCountry *string    `json:"deathCountry,omitempty"`
	DeathState   *string    `json:"deathState,omitempty"`
	DeathCity    *string    `json:"deathCity,omitempty"`
	NameFirst    *string    `json:"nameFirst,omitempty"`
	NameLast     *string    `json:"nameLast,omitempty"`
	NameGiven    *string    `json:"nameGiven,omitempty"`
	Weight       *int32     `json:"weight,omitempty"`
	Height       *int32     `json:"height,omitempty"`
	Bats         *string    `json:"bats,omitempty"`
	Throws       *string    `json:"throws,omitempty"`
	Debut        *time.Time `json:"debut,omitempty"`
	FinalGame    *time.Time `json:"finalGame,omitempty"`
	RetroID      *string    `json:"retroID,omitempty"`
	BbrefID      *string    `json:"bbrefID,omitempty"`
}
