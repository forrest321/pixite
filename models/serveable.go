package models

import (
	"github.com/jpillora/velox"
	"image"
)

type Serveable struct {
	velox.State
	Id          int64
	Name        string
	Description string
	MainImage   image.Image
}