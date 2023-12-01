package repository

import (
	"time"

	"github.com/seigaalghi/e-library/model"
)

type Repository interface {
	CreateSchedule(pickup, dropoff time.Time, book *model.Books) (bool, error)
}
