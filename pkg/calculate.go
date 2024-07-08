package pkg

import (
	"time"
)

func CalculateExpiration(expired int64) time.Time {
	return time.Unix(expired, 0)
}
