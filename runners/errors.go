package runners

import "fmt"

type UnsupportedLanguageError struct {
	lang string
}

func (u *UnsupportedLanguageError) Error() string {
	return fmt.Sprintf("%s is not a supported langauge", u.lang)
}

func (u *UnsupportedLanguageError) Is(err error) bool {
	_, ok := err.(*UnsupportedLanguageError)
	return ok
}
