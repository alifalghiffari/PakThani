package domain

import "time"

type Image struct {
	Id         int
	ProductId  int
	Image      string
	Created_at time.Time
	Updated_at time.Time
}
