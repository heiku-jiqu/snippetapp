package main

import (
	"testing"
	"time"

	"github.com/heiku-jiqu/snippetapp/internal/assert"
)

func TestHumanDate(t *testing.T) {
	testCases := []struct {
		testname string
		input    time.Time
		expected string
	}{
		{
			"Check humanDate",
			time.Date(2020, 3, 17, 10, 15, 0, 0, time.UTC),
			"17 Mar 2020 at 10:15",
		},
		{
			"Check zerotime is empty string",
			time.Time{},
			"",
		},
		{
			"Always use UTC timezone",
			time.Date(2020, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			"17 Mar 2020 at 09:15",
		},
	}
	for _, c := range testCases {
		t.Run(c.testname, func(t *testing.T) {
			output := humanDate(c.input)
			assert.Equal(t, output, c.expected)
		})
	}
}
