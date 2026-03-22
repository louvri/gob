package object

import (
	"fmt"
	"time"
)

func ConvertDateTimeTextWithTimezone(dateTimeText, sourceFormat, targetFormat string, timezone *time.Location) (string, error) {
	if dateTimeText == "" {
		return "", nil
	}
	t, err := time.Parse(sourceFormat, dateTimeText)
	if err != nil {
		return "", fmt.Errorf("failed to parse %q with format %q: %w", dateTimeText, sourceFormat, err)
	}
	return t.In(timezone).Format(targetFormat), nil
}

// Deprecated: GetTimeVariableValue is identical to ConvertDateTimeTextWithTimezone. Use that instead.
func GetTimeVariableValue(dateTimeText, sourceFormat, targetFormat string, timezone *time.Location) (string, error) {
	return ConvertDateTimeTextWithTimezone(dateTimeText, sourceFormat, targetFormat, timezone)
}

func ConvertDateWithTimezone(dateTimeText, targetFormat string, timezone *time.Location) (string, error) {
	if dateTimeText == "" {
		return "", nil
	}
	t, err := time.Parse(targetFormat, dateTimeText)
	if err != nil {
		return "", fmt.Errorf("failed to parse %q with format %q: %w", dateTimeText, targetFormat, err)
	}
	return t.In(timezone).Format(targetFormat), nil
}
