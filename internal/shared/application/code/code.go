package code

import (
	"fmt"
	"strings"

	gonanoid "github.com/matoous/go-nanoid"
	"github.com/techforge-lat/errortrace/v2"
)

var alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Generate(prefix string, length int) (string, error) {
	id, err := gonanoid.Generate(alphabet, length)
	if err != nil {
		fmt.Println("Error generating NanoID:", err)
		return "", errortrace.OnError(err)
	}

	return fmt.Sprintf("%s-%s", strings.ReplaceAll(strings.ToLower(prefix), " ", "-"), id), nil
}
