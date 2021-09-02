package utils

import (
	"apus-sample/common/constant"
	"os"
	"strings"
)

func GetApplicationEnv() string {
	env := os.Getenv("ENV")
	if env == ""{
		return constant.EnvProd
	}
	return strings.ToLower(env)
}
