package string_formater

import (
	"errors"
	"fmt"
	"strings"
)

func Format(ss string, args ...interface{}) (string, error) {
	aux := ""
	compose := strings.Split(ss, "{}")
	if len(compose)-1 != len(args) {
		println(len(compose), len(args))
		return "", errors.New(
			"parameters not the same as the format string " + fmt.Sprintf("%v != %v", len(compose)-1, len(args)))
	}
	for ct, i := range compose {
		aux += i
		if ct <= len(args)-1 {
			aux += fmt.Sprintf("%v", args[ct])

		}
		// fmt.Sprintf("%v", i)

	}
	return aux, nil
}
