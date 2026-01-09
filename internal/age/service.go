package age

import "time"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// GetAge calculates the age based on the provided birthdate string in "YYYY-MM-DD" format.
// It returns the age as an integer and a boolean indicating if the person is an adult (18 or older).
func (s *Service) GetAge(birthdate string) (int32, bool) {

	birthTime := parseDate(birthdate)
	if birthTime.IsZero() {
		return 0, false
	}
	now := time.Now()
	age := now.Year() - birthTime.Year()
	if now.YearDay() < birthTime.YearDay() {
		age--
	}

	isAdult := age >= 18
	return int32(age), isAdult
}

func parseDate(birthdate string) (t time.Time) {
	t, err := time.Parse("2006-01-02", birthdate)
	if err != nil {
		return time.Time{}
	}
	return
}
