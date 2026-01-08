package entity

import "time"

type Product struct {
	ID    		int64
	Name  		string
	Code  		string
	Price 		float64
	CreatedAt 	time.Time
	CreatedBy 	int
}