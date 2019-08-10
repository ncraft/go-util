package base

import (
	"strconv"
)

// MustAtoi is like strconv.Atoi but panics on errors.
func MustAtoi(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return i
}

func MustBeSet(value string) string {
	if value == "" {
		panic("missing required value")
	}
	return value
}

// ParseBoolOptional accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False. Pass a defaultResult for the case that value is empty or cannot be parsed without error.
func ParseBoolOptional(value string, defaultResult bool) bool {
	if value == "" {
		return defaultResult
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		return defaultResult
	}
	return b
}

// MustParseBoolOptional accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False. Pass a defaultResult for the case that value is empty. Any other value leads to panic.
func MustParseBoolOptional(value string, defaultResult bool) bool {
	if value == "" {
		return defaultResult
	}
	return MustParseBool(value)
}

// MustParseBool accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False. Any other value leads to panic.
func MustParseBool(value string) bool {
	b, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}
	return b
}

// AnyStringEmpty checks if any of the given strings is empty.
func AnyStringEmpty(strings []string) bool {
	for _, item := range strings {
		if item == "" {
			return true
		}
	}
	return false
}
