package models

import (
	"errors"
)

type Product struct {
	ID                uint    `json:"id"`
	Title             string  `json:"title"`
	Category          string  `json:"category"`
	Price             float64 `json:"price"`
	Store             string  `json:"store"`
	ImageURL          string  `json:"imageURL"`
	AffiliateURL      string  `json:"affiliateURL"`
	ProductURL        string  `json:"productURL"`
	Priority          float64 `json:"priority"`
	CommissionPercent float32 `json:"commissionPercent"`
}

func (p *Product) CalcPriority() error {
	if p.Price == 0 {
		return errors.New("product must have Price")
	}
	if p.CommissionPercent != 0 {
		p.Priority = (p.Price / 100) * float64(p.CommissionPercent)
		return nil
	} else if p.Priority != 0 {
		p.CommissionPercent = float32((p.Priority * 100) / p.Price)
		return nil
	}

	return errors.New("product must have either CommissionPercent or Priority")
}
