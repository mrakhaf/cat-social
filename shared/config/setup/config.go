package setup

import "time"

func SetupTimezone() {
	location, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = location
}
