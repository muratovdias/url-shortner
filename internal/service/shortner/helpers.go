package shortner

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"unicode/utf8"
)

func (u *urlShortenerImpl) validateURL(input string) error {
	if utf8.RuneCountInString(input) == 0 {
		return errors.New("URL cannot be empty")
	}

	if !u.regexp.MatchString(input) {
		return errors.New("invalid URL format")
	}

	parsedURL, err := url.ParseRequestURI(input)
	if err != nil {
		return errors.New("invalid URL structure")
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return errors.New("URL must start with http:// or https://")
	}

	return nil
}

func generateShortURL(url string) (string, error) {
	const aliasLength = 6

	var sb strings.Builder
	for i := 0; i < aliasLength; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(url))))
		if err != nil {
			return "", fmt.Errorf("error generating random number: %v", err)
		}

		sb.WriteByte(url[index.Int64()])
	}

	return sb.String(), nil
}
