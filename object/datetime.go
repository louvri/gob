package object

import "time"

func ConvertDateTimeTextWithTimezone(dateTimeText, sourceFormat, targetFormat string, timezone *time.Location) string {
	if dateTimeText != "" {
		orderTime, _ := time.Parse(sourceFormat, dateTimeText)
		return orderTime.In(timezone).Format(targetFormat)
	}

	return ""
}

func GetTimeVariableValue(dateTimeText, sourceFormat, targetFormat string, timezone *time.Location) string {
	if dateTimeText != "" {
		timeVariableValue, _ := time.Parse(sourceFormat, dateTimeText)
		return timeVariableValue.In(timezone).Format(targetFormat)
	}

	return ""
}

func ConvertDateWithTimezone(dateTimeText, targetFormat string, timezone *time.Location) string {
	if dateTimeText != "" {
		orderTime, _ := time.Parse(targetFormat, dateTimeText)
		return orderTime.In(timezone).Format(targetFormat)
	}

	return ""
}
