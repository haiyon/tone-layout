package utils

import "time"

// TimeToString - time to string
func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
