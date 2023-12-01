package model

import "time"

type LendBookRequest struct {
	EditionNumber  string    `json:"edition_number" validate:"required"`
	DropOffDate    time.Time `json:"-" validate:"required"`
	PickupDate     time.Time `json:"-" validate:"required"`
	DropOffDateStr string    `json:"dropoff_date" validate:"required"`
	PickupDateStr  string    `json:"pickup_date" validate:"required"`
}

type LendBookResponse struct {
	EditionNumber string    `json:"edition_number"`
	Title         string    `json:"title"`
	DropOffDate   time.Time `json:"dropoff_date"`
	PickupDate    time.Time `json:"pickup_date"`
	Valid         bool      `json:"valid"`
	Message       string    `json:"message"`
}
