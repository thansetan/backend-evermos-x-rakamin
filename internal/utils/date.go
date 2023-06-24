package utils

// @TODO : make function parsing date
import "time"

func ParseDate(date string) (time.Time, error) {
	time, err := time.Parse("02/01/2006", date)
	if err != nil {
		return time, err
	}
	return time, nil
}

func DateToString(date time.Time) string {
	return date.Format("02/01/2006")
}
