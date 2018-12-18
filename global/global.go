package global

import (
	"github.com/openzipkin/zipkin-go/reporter"
	models "myservices/books/models"
	"os"
	"strings"
)

var Logger logger

var BOOK_DB *db

var zipkinReporter reporter.Reporter

var Redis *redisClient

func init() {

	for _, v := range os.Args {
		if strings.Contains(v, "-test.v=true") {
			return
		}
	}

	Logger = newLogger()

	loadConf()

	BOOK_DB = newBookDB()

	Redis = newRedisClient()

	models.Migrate(BOOK_DB.DB)

	zipkinReporter = NewZipkinReporter()
}
