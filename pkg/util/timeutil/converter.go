package timeutil

import "time"

func FromString(str string) (time.Time, error) {
	return time.Parse("01-02-2006 15:04:05", str)
}

func ToString(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

func ToStringDateOnly(date time.Time) string {
	return date.Format("2006-01-02")
}

func DateTimeFromString(str string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05Z", str)
}

func ToPointer(value time.Time) *time.Time {
	return &value
}
