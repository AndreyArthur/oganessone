package verifier

import (
	"regexp"
	"time"
)

func IsUuid(uuid string) bool {
	regex := regexp.MustCompile("^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$")
	matches := regex.Match([]byte(uuid))
	return matches
}

func IsAccountUsername(username string) bool {
	var isValid bool
	isValid = !(len(username) < 4 || len(username) > 16)
	if !isValid {
		return false
	}
	regex := regexp.MustCompile(`^.*\s.*$`)
	usernameHasWhitespaces := regex.Match([]byte(username))
	if usernameHasWhitespaces {
		isValid = false
	} else {
		isValid = true
	}
	return isValid
}

func IsEmail(email string) bool {
	regex := regexp.MustCompile("^[-!#$%&'*+/0-9=?A-Z^_a-z`{|}~](\\.?[-!#$%&'*+/0-9=?A-Z^_a-z`{|}~])*@[a-zA-Z0-9](-*\\.?[a-zA-Z0-9])*\\.[a-zA-Z](-?[a-zA-Z0-9])+$")
	matches := regex.Match([]byte(email))
	return matches
}

func IsBcrypt(bcrypt string) bool {
	regex := regexp.MustCompile(`^\$2[aby]?\$\d{1,2}\$[.\/A-Za-z0-9]{53}$`)
	matches := regex.Match([]byte(bcrypt))
	return matches
}

func IsISO8601(date string) bool {
	regex := regexp.MustCompile(`^\d{4}-\d\d-\d\dT\d\d:\d\d:\d\d(\.\d+)?(([+-]\d\d:\d\d)|Z)?$`)
	matches := regex.Match([]byte(date))
	return matches
}

func IsTime(date *time.Time) bool {
	iso := date.Format(time.RFC3339)
	valid := IsISO8601(iso)
	return valid
}
