package eleven

import (
	"os"

	"log-aggregator/transform"
	"log-aggregator/types"
)

const (
	EnvProduct   = "ELEVEN_PRODUCT"
	EnvComponent = "ELEVEN_COMPONENT"
	EnvTimeFormat = "TIME_FORMAT"
)

func New() transform.Transformer {

	product := os.Getenv(EnvProduct)
	component := os.Getenv(EnvComponent)

	timeFormat := os.Getenv(EnvTimeFormat)
	if timeFormat == "" {
		timeFormat = "2006-01-02T15:04:05.000Z07:00"
	}

	return func(rec *types.Record) (*types.Record, error) {
		rec.Fields["product"] = product
		rec.Fields["component"] = component

		formattedTime := rec.Time.UTC().Format(timeFormat)
		rec.Fields["when"] = formattedTime
		return rec, nil
	}
}
