package models

import (
    "time"
)

// Each day may define perDay, perHour, or be null
type SpecialPayRule struct {
    PerDay  *float64 `json:"perDay,omitempty"`
    PerHour *float64 `json:"perHour,omitempty"`
} /* special pay json example:
	{
	    "mon": {
	        "perDay": 123.40
	    },
	    "tue": {
	        "perDay": 100.00
	        "perHour": 15.00
	    },
	    "wed": {
	        "perHour": 50.00
	    },
	    "thu": null
	    "fri": {}
	}
	in the above example,
	on monday this employee is paid $123.40 if they work on that day
	on tuesday, they're paid $100 if they're working that day plus $15.00 per hour
	on wednesday, they're paid 50.00 per-hour
	this employee won't appear in the menu where hours/days are recorded on thursday or friday
*/

// Map of weekday string to pay rule or nil
type SpecialPay map[string]*SpecialPayRule

type Employee struct {
    ID         int         `db:"id" json:"id"`
    Name       string      `db:"name" json:"name"`
    Type       string      `db:"type" json:"type"` // "FLOOR" or "KITCHEN"
    Wage       *float64    `db:"wage" json:"wage"` // can be nil if only special pay
    SpecialPay SpecialPay  `db:"special_pay" json:"special_pay"`
    CreatedAt  time.Time   `db:"created_at" json:"created_at"`
    UpdatedAt  time.Time   `db:"updated_at" json:"updated_at"`
}
