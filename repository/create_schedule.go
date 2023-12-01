package repository

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/seigaalghi/e-library/model"
)

var (
	createScheduleQuery = `
	INSERT INTO book_schedules (
		book_title,
		book_edition_number,
		book_authors,
		pickup_date,
		dropoff_date,
		created_at,
		updated_at
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		CURRENT_TIMESTAMP,
		CURRENT_TIMESTAMP
	);`

	bookAvailability = `
	SELECT COUNT(*) FROM book_schedules WHERE book_edition_number = $1 AND dropoff_date > CURRENT_TIMESTAMP;`
)

func (r *repository) CreateSchedule(pickup, dropoff time.Time, book *model.Books) (bool, error) {
	var count int
	err := r.db.QueryRow(bookAvailability, book.EditionNumber).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return false, nil
	}

	authorsStr, _ := json.Marshal(book.Authors)
	res, err := r.db.Exec(createScheduleQuery, book.Title, book.EditionNumber, authorsStr, pickup, dropoff)
	if err != nil {
		return false, err
	}

	if af, _ := res.RowsAffected(); af == 0 {
		return false, sql.ErrNoRows
	}

	return true, nil
}
