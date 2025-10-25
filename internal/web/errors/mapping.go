package errors

import (
	"log"

	"app/internal/domain/errors"
)

func EnsureAllErrorsMapped(all []errors.ErrorCode, mapping map[errors.ErrorCode]int) {
	for _, code := range all {
		if _, ok := mapping[code]; !ok {
			log.Panicf("missing HTTP status mapping for error code: %s", code)
		}
	}
}
