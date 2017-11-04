package model

// WorkingHistory represents a company that the user work(s/ed)
type WorkingHistory struct {
	ID        uint64 `db:"id"`
	ProfileID uint64 `db:"profile_id"`
	Company   string `db:"company"`
}
