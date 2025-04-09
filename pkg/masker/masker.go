package masker

import (
	"regexp"
	"strings"
)

func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return email
	}
	name := parts[0]
	domain := parts[1]
	if len(name) <= 1 {
		return "***@" + domain
	}
	return string(name[0]) + strings.Repeat("*", len(name)-1) + "@" + domain
}

func MaskPhone(phone string) string {
	re := regexp.MustCompile(`\d`)
	digits := re.FindAllString(phone, -1)
	if len(digits) < 4 {
		return phone
	}
	masked := strings.Repeat("*", len(digits)-4) + strings.Join(digits[len(digits)-4:], "")
	return re.ReplaceAllStringFunc(phone, func(d string) string {
		ch := masked[0]
		masked = masked[1:]
		return string(ch)
	})
}

func MaskSSN(ssn string) string {
	re := regexp.MustCompile(`\d`)
	digits := re.FindAllString(ssn, -1)
	if len(digits) < 4 {
		return ssn
	}
	masked := strings.Repeat("*", len(digits)-4) + strings.Join(digits[len(digits)-4:], "")
	return re.ReplaceAllStringFunc(ssn, func(d string) string {
		ch := masked[0]
		masked = masked[1:]
		return string(ch)
	})
}
