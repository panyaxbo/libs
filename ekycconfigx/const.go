package ekycconfigx

const (
	ekycDB    = "EKYCDB"
	dipchipDB = "DIPCHIPDB"

	redisHostLocal = "redisHostLocal"
	redisHostDev   = "redisHostDev"
	redisHostUat   = "redisHostUat"
	redisHostProd  = "redisHostProd"
)

var (
	redisHost = map[string]string{
		redisHostLocal: "127.0.0.1:6379",
		redisHostDev:   "10.98.3.99:6379",
		redisHostUat:   "10.98.3.91:6379",
		redisHostProd:  "10.99.1.52:6379",
	}
)
