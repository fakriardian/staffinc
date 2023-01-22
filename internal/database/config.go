package database

import (
	"github.com/fakriardian/staffinc/internal/utils"
)

var (
	DbAddress = utils.GoDotEnvVariable("DB_ADDRESS")
)
