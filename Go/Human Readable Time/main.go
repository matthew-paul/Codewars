package kata

import "fmt"

func HumanReadableTime(seconds int) string {
	var hours int = seconds / 3600
	seconds = seconds - hours*3600
	var minutes int = seconds / 60
	seconds = seconds - minutes*60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
