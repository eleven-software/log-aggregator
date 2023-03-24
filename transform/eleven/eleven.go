package eleven

import (
	"os"

	"log-aggregator/transform"
	"log-aggregator/types"
)

const (
	EnvProduct   = "ELEVEN_PRODUCT"
	EnvComponent = "ELEVEN_COMPONENT"
	EnvLogFormat = "LOG_FORMAT"
)

func New() transform.Transformer {

	product := os.Getenv(EnvProduct)
	component := os.Getenv(EnvComponent)

	logFormat := os.Getenv(EnvLogFormat)
	if logFormat == "" {
		logFormat = "2006-01-02T15:04:05.000"
	}

	return func(rec *types.Record) (*types.Record, error) {
		rec.Fields["product"] = product
		rec.Fields["component"] = component

		formattedTime := rec.Time.Format(logFormat)
		rec.Fields["when"] = formattedTime
		return rec, nil
	}
}
