package config

import "os"

var jwt_secret = os.Getenv("jwt_secret")

var Key = []byte(jwt_secret)
