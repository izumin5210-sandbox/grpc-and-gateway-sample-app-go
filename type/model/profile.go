package model

// Profile represents an user profile
type Profile struct {
	UserID           uint64 `db:"uesr_id"`
	Name             string `db:"name"`
	Location         string `db:"location"`
	WorkingHistories []*WorkingHistory
}
