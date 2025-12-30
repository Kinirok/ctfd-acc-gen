package ctfdgen

import "errors"

var (
	ErrNoClient         = errors.New("CTFDClient is required")
	ErrNoDB             = errors.New("DB is required")
	ErrNoSqlDB          = errors.New("cannot get sqlDB")
	ErrLostConnectionDB = errors.New("lost connection to DB")
	ErrCreateAccount    = errors.New("account have not been created, stopping execution")
	ErrCreateTeam       = errors.New("team have not been created, stopping execution")
)
