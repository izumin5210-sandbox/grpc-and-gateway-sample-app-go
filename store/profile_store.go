package store

import (
	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/type/model"
)

// ProfileStore is an interface for accessing profile resources
type ProfileStore interface {
	GetByUserID(userID int64) (*model.Profile, error)
}
