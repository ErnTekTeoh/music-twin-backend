package config

import "os"

var env string

var (
	DEV  = "dev"
	TEST = "test"
	UAT  = "uat"
	LIVE = "live"
)
var validEnv = map[string]bool{
	DEV:  true,
	TEST: true,
	UAT:  true,
	LIVE: true,
}

func IsLIVEEnv() bool {
	return env == LIVE
}
func GetEnv() string {
	// init project env
	if env == "" {
		env = os.Getenv("ENV_NAME")

		if !isValidEnv(env) {
			env = DEV
		}
	}
	return env
}

func isValidEnv(env string) bool {
	_, isValid := validEnv[env]
	return isValid
}
