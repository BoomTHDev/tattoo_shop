package custom

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

func IsRecordFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func IsDupicateKeyError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key value")
}
