package dates

import (
	"fmt"
	"testing"
	"time"
)

var cstSh, _ = time.LoadLocation("Asia/Jerusalem")

type DateTest struct {
	daysToSubtract int
	date time.Time
}

var tests = []struct{
	input DateTest
	expected time.Time
	expectedStartMonth time.Time
}{
	{
		input: DateTest{
			date: time.Date(2020, 12, 11, 18, 0, 0, 0, cstSh),
			daysToSubtract: 28,
		},
		expected: time.Date(2020, 11, 13, 18, 0, 0, 0, cstSh),
		expectedStartMonth: time.Date(2020, 12, 1, 0, 0, 0, 0, cstSh),
	},
}

func TestInterval(t *testing.T) {
	t.Run("check dates", func(t *testing.T) {
		for i, test := range tests {
			subtracted := GetTimeFromNowBackDays(test.input.date, test.input.daysToSubtract)
			fmt.Printf("subtracted days %d month %d", subtracted.Day(), subtracted.Month())
			if subtracted.Day() != test.expected.Day() {
				t.Fatalf("error test %d", i)
			}
		}
	})	
	
	t.Run("check dates", func(t *testing.T) {
		for i, test := range tests {
			subtracted := GetTimeFromMonthStart(test.input.date)
			fmt.Printf("subtracted days %d month %d", subtracted.Day(), subtracted.Month())
			if subtracted.Day() != test.expectedStartMonth.Day() {
				t.Fatalf("error test %d", i)
			}
		}
	})
}
