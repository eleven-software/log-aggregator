package eleven

import (
	"math/big"
	"os"
	"strings"

	"log-aggregator/transform"
	"log-aggregator/types"
)

const (
	EnvProduct    = "ELEVEN_PRODUCT"
	EnvComponent  = "ELEVEN_COMPONENT"
	EnvTimeFormat = "TIME_FORMAT"
)

func hexToDecimal(hex string) *big.Int {
	// Convert hexadecimal to decimal
	decimalValue, success := new(big.Int).SetString(hex, 16)
	if !success {
		return nil
	}
	return decimalValue
}

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
		seq_no := strings.Split(string(rec.Cursor), ";")[1]
		seq_no = strings.Split(seq_no, "=")[1]
		rec.Fields["sequence_number"] = hexToDecimal(seq_no)

		formattedTime := rec.Time.UTC().Format(timeFormat)
		rec.Fields["when"] = formattedTime
		return rec, nil
	}
}
