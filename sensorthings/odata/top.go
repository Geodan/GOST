package odata

import (
	"net/http"
	"strconv"
)

// QueryTop is a filter for limiting the number of returns to be returned.
// Specifies a non-negative integer that limits the number of entities
// returned within a collection. The service must return the number of
// available entities up to, but not exceeding, the specified value.
type QueryTop struct {
	QueryBase
	Limit int
}

// Parse $top values in QueryTop, returns error if the supplied value is
// invalid (non-integer or < 0)
func (q *QueryTop) Parse(value string) error {
	q.RawQuery = value
	i, err := strconv.Atoi(value)
	if err != nil || i < 0 {
		return CreateQueryError(QueryTopInvalid, http.StatusBadRequest, value)
	}

	q.Limit = i
	return nil
}

// IsValid always returns true, errors are already filtered out by parse
func (q *QueryTop) IsValid() (bool, error) {
	return true, nil
}

// GetQueryOptionType returns the QueryOptionType for QueryTop
func (q *QueryTop) GetQueryOptionType() QueryOptionType {
	return QueryOptionTop
}

// IsNil checks if *QueryTop is nil
func (q *QueryTop) IsNil() bool {
	if q == nil {
		return true
	}

	return false
}
