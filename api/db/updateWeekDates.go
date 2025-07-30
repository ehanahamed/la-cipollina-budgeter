package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func updateWeekDates() error {
	_, err := Pool.Exec(
		context.Background(),
		`UPDATE weeks w
SET 
    first_date = s.first_day,
    last_date = s.last_day
FROM (
    SELECT 
        w.id AS week_id, 
        MIN(d.date) AS first_day,
        MAX(d.date) AS last_day
    FROM weeks w
    JOIN days d 
        ON d.date BETWEEN w.start_date AND w.end_date
    GROUP BY w.id
) AS s
WHERE w.id = s.week_id`,
	)
	if err != nil {
		log.Print("Error updating first & last dates for weeks: ", err)
	}
}

