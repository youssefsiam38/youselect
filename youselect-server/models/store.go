package models

import ()

type Store struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	AffiliateURL  string `json:"affiliateURL"`
	ImageURL      string `json:"imageURL"`
	QuerySelector string `json:"querySelector"`
}
