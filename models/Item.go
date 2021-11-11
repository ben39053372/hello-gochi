package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
  ID     uint   `json:"id" gorm:"primary_key"`
  Title  string `json:"title"`
  Author string `json:"author"`
}
