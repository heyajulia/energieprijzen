package internal

import (
	"net/url"
	"time"
)

func PrepareQueryParameters() string {
	// like time.RFC3339Nano but with millisecond precision
	layout := "2006-01-02T15:04:05.999Z07:00"

	t := time.Now().UTC()
	fromDate := time.Date(t.Year(), t.Month(), t.Day(), 22, 0, 0, 0, t.Location())
	tillDate := fromDate.AddDate(0, 0, 1).Add(-1 * time.Millisecond)

	params := url.Values{
		"fromDate":  {fromDate.Format(layout)},
		"tillDate":  {tillDate.Format(layout)},
		"interval":  {"4"},
		"usageType": {"1"},
		"inclBtw":   {"true"},
	}

	return params.Encode()
}
