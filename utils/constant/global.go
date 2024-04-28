package constant

import "time"

var (
	jakartaLoc, _ = time.LoadLocation("Asia/Jakarta")
	TimeNow       = time.Now().In(jakartaLoc).Format("2006-01-02T15:04:05Z")
)
