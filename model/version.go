package model

import (
)

type Version struct {
	Number    int64  `json:"number"`
}

func NewVersion (number int64, ) *Version {
	return &Version{
		Number:    number,
	}
}
