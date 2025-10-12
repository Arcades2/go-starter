package errors

import (
	"app/internal/domain/errors"
	"log"
)

func EnsureAllErrorsMapped(all []errors.ErrorCode, mapping map[errors.ErrorCode]int) {
	for _, code := range all {
		if _, ok := mapping[code]; !ok {
			log.Panicf("missing HTTP status mapping for error code: %s", code)
		}
	}
}
