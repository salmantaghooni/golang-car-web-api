package dto

import "time"

type CreatePersianYearRequest struct {
	PersianTitle string    `json:"persianTitle" binding:"min=4,max=4"`
	Year         int       `json:"year"`
	StartAt      time.Time `json:"startAt"`
	EndAt        time.Time `json:"endAt"`
}

type UpdatePersianYearRequest struct {
	PersianTitle string    `json:"persianTitle,omitempty" binding:"min=4,max=4"`
	Year         int       `json:"year,omitempty"`
	StartAt      time.Time `json:"startAt,omitempty"`
	EndAt        time.Time `json:"endAt,omitempty"`
}

type PersianYearResponse struct {
	Id           int       `json:"id"`
	PersianTitle string    `json:"persianTitle,omitempty"`
	Year         int       `json:"year,omitempty"`
	StartAt      time.Time `json:"startAt,omitempty"`
	EndAt        time.Time `json:"endAt,omitempty"`
}

type PersianYearWithoutDateResponse struct {
	Id           int    `json:"id"`
	PersianTitle string `json:"persianTitle,omitempty"`
	Year         int    `json:"year,omitempty"`
}
