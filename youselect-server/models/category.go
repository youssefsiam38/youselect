package models

import ()

type Category struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	ImageURL string `json:"imageURL"`
}
