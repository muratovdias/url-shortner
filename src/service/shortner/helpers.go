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

var (
	ErrInvalidUrl = errors.New("invalid url")
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

func generateShortURL() (string, error) {
	const aliasLength = 6
	const alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var sb strings.Builder
	for i := 0; i < aliasLength; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphanumeric))))
		if err != nil {
			return "", fmt.Errorf("error generating random number: %v", err)
		}

		sb.WriteByte(alphanumeric[index.Int64()])
	}

	return sb.String(), nil
}
