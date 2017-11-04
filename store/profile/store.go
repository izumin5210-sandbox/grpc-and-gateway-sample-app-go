package profilestore

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/type/model"
	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/type/system"
)

type storeImpl struct {
	*sqlx.DB
}

func (s *storeImpl) GetByUserID(userID int64) (*model.Profile, error) {
	prof := &model.Profile{}
	err := s.DB.Get(prof, "SELECT * FROM profiles WHERE user_id = $1", userID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, system.ErrorNotFound.Wrap(
				errors.Wrap(err, "profile was not found"),
			)
		default:
			return nil, system.ErrorFailedToReadDB.WithReport(
				errors.Wrap(err, "failed to read profile"),
			)
		}
	}
	return prof, nil
}
