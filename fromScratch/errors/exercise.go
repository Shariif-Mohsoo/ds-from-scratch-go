package errors

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

// Example time string: 14:07:33
func ParseTime(input string) (Time, error) {
	components := strings.Split(input, ":")
	if len(components) != 3 {
		return Time{}, &TimeParseError{"Invalid number of time components", input}
	} else {
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hour: %v", err), input}
		}
		min, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing minute: %v", err), input}
		}

		sec, err := strconv.Atoi(components[2])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing second: %v", err), input}
		}

		if hour > 23 || hour < 0 {
			return Time{}, &TimeParseError{"Hour out of range: 0 <= hour <= 23", fmt.Sprintf("%v", hour)}
		}
		if min > 59 || min < 0 {
			return Time{}, &TimeParseError{"Minute out of range: 0 <= minute <= 59", fmt.Sprintf("%v", min)}
		}
		if sec > 59 || sec < 0 {
			return Time{}, &TimeParseError{"Second out of range: 0 <= second <= 59", fmt.Sprintf("%v", sec)}
		}
		return Time{hour: hour, minute: min, second: sec}, nil
	}

}

func Exercise() {
	table := []struct {
		time string
		ok   bool
	}{
		{"19:00:12", true},
		{"1:03:12", true},
		{"bad", false},
		{"0:00:12", true},
		{":00:12", false},
		{"19:00:", false},
		{"19::12", false},
	}
	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			fmt.Println(fmt.Errorf("%v: %v, error should be nil", data.time, err))
		}
	}

}
