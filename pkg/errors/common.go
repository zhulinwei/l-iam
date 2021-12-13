package errors

import "strings"

func IsMysqlError1062(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "Error 1062: Duplicate entry")
}
