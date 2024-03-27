package state

import (
	"github.com/jinzhu/gorm"
)

type s struct {
	gorm.Model
	AllStates []states `json:"States"`
}
type states struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Capital     string `json:"capital"`
	Population  string `json:"population"`
	RulingParty string `json:"ruling_party"`
	Website     string `json:"official_website"`
	Description string `json:"description"`
}
