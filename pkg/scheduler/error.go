package scheduler

import (
	"errors"
)

var (
	ErrEventNotExist = errors.New("finding a scheduled event requires a existent cron id")
)
