package lib

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Hashing Functions

// MD5FromInt generates an MD5 hash from an integer value.
func MD5FromInt(value int) string {
	str := strconv.Itoa(value)
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

// MD5FromString generates an MD5 hash from a string value.
func MD5FromString(value string) string {
	hasher := md5.New()
	hasher.Write([]byte(value))
	return hex.EncodeToString(hasher.Sum(nil))
}

// SHA1FromString generates a SHA-1 hash from a string value.
func SHA1FromString(value string) string {
	sha := sha1.New()
	sha.Write([]byte(value))
	return fmt.Sprintf("%x", sha.Sum(nil))
}

// SHA256FromString generates a SHA-256 hash from a string value.
func SHA256FromString(value string) string {
	hash := sha256.Sum256([]byte(value))
	return fmt.Sprintf("%x", hash)
}

// Conversion Functions

// IntToStr converts an integer to a string.
func IntToStr(value int) string {
	return strconv.Itoa(value)
}

// Int64ToStr converts a 64-bit integer to a string.
func Int64ToStr(value int64) string {
	return strconv.FormatInt(value, 10)
}

// StrToInt converts a string to an integer. Returns 0 on error.
func StrToInt(value string) int {
	result, _ := strconv.Atoi(value)
	return result
}

// StrToInt64 converts a string to a 64-bit integer. Returns 0 on error.
func StrToInt64(value string) int64 {
	result, _ := strconv.ParseInt(value, 10, 64)
	return result
}
func StrToUInt64(value string) uint64 {
	result, _ := strconv.ParseUint(value, 10, 64)
	return result
}

// StrToFloat converts a string to a float64. Returns 0.0 on error.
func StrToFloat(value string) float64 {
	result, _ := strconv.ParseFloat(value, 64)
	return result
}

// StrToBool converts a string to a boolean. Returns false on error.
func StrToBool(value string) bool {
	result, _ := strconv.ParseBool(value)
	return result
}

// FloatToStr converts a float64 to a string with optional precision.
func FloatToStr(value float64, precision ...int) string {
	prec := 0 // Default precision
	if len(precision) > 0 {
		prec = precision[0]
	}
	return strconv.FormatFloat(value, 'f', prec, 64)
}

// FloatToFormattedStr converts a float64 to a formatted string without decimals.
func FloatToFormattedStr(value float64) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%.0f", Round(value))
}

// JSON Conversion Functions

// String Conversion Functions

// StrToTime parses a string to a time.Time object. Uses the default layout if none is provided.
//
// Usage:
//
//	StrToTime("2021-05-19 11:56:30") // Using default layout "2006-01-02 15:04:05"
//	StrToTime("2021-05-19", "2006-01-02") // Using custom layout "2006-01-02"
func StrToTime(value string, layout ...string) time.Time {
	var l string
	if len(layout) > 0 {
		l = layout[0]
	} else {
		l = "2006-01-02"
	}
	result, _ := time.Parse(l, value)
	return result
}

// StrToStrfmtDate parses a string to a strfmt.Date object.
//
// Usage:
//
//	StrToStrfmtDate("2023-08-15")
func StrToStrfmtDate(value string) strfmt.Date {
	res, _ := time.Parse("2006-01-02", value)
	return strfmt.Date(res)
}

// StrToStrfmtDateTime parses a string to a strfmt.DateTime object.
//
// Usage:
//
//	StrToStrfmtDateTime("2023-08-15 10:08:15")
func StrToStrfmtDateTime(value string) strfmt.DateTime {
	res, _ := time.Parse("2006-01-02 15:04:05", value)
	return strfmt.DateTime(res)
}

// Slice Conversion Functions

// IntSliceToStr converts a slice of integers to a delimited string.
func IntSliceToStr(values []int, delimiter string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(values), " ", delimiter, -1), "[]")
}

// StrSliceToStr converts a slice of strings to a delimited string.
func StrSliceToStr(values []string, delimiter string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(values), " ", delimiter, -1), "[]")
}

// UUIDSliceToStr converts a slice of UUIDs to a delimited string.
func UUIDSliceToStr(values []uuid.UUID, delimiter string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(values), " ", delimiter, -1), "[]")
}

// UUIDSliceToStrSlice converts a slice of UUIDs to a slice of strings.
func UUIDSliceToStrSlice(values []uuid.UUID) []string {
	var result []string
	for _, uuid := range values {
		result = append(result, uuid.String())
	}
	return result
}

// Utility Functions

// RemoveLeadingZeros removes leading zeros from a string.
func RemoveLeadingZeros(value string) string {
	return strings.TrimLeft(value, "0")
}

// Force Pointer and Fallback Functions

// ForceStr returns the string value of a pointer or an empty string if nil.
func ForceStr(input *string) string {
	if input == nil {
		return ""
	}
	return *input
}

// ForceInt returns the integer value of a pointer or 0 if nil.
func ForceInt(input *int) int {
	if input == nil {
		return 0
	}
	return *input
}

// ForceInt64 returns the int64 value of a pointer or 0 if nil.
func ForceInt64(input *int64) int64 {
	if input == nil {
		return 0
	}
	return *input
}

// ForceBool returns the boolean value of a pointer or false if nil.
func ForceBool(input *bool) bool {
	if input == nil {
		return false
	}
	return *input
}

// ForceFloat64 returns the float64 value of a pointer or 0.0 if nil.
func ForceFloat64(input *float64) float64 {
	if input == nil {
		return 0.0
	}
	return *input
}

// ForceTime returns the time.Time value of a pointer or zero time if nil.
func ForceTime(input *time.Time) time.Time {
	if input == nil {
		return time.Time{}
	}
	return *input
}

// ForceUUID returns the UUID value of a pointer or uuid.Nil if nil.
func ForceUUID(input *uuid.UUID) uuid.UUID {
	if input == nil {
		return uuid.Nil
	}
	return *input
}

// ReplaceSpecialCharacters replaces all non-alphanumeric characters in a string with a specified replacement.
func ReplaceSpecialCharacters(value, replacement string) string {
	return regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(value, replacement)
}
