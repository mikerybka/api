package api

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func kebab2pascal(s string) string {
	parts := strings.Split(s, "-")
	for i := range parts {
		parts[i] = cases.Title(language.English).String(parts[i])
	}
	return strings.Join(parts, "")
}
