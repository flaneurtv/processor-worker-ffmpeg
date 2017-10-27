package utils

import (
	"bytes"
	"encoding/json"
)

func JSON(v interface{}) string {
	switch o := v.(type) {
	case string:
		var out bytes.Buffer
		if err := json.Indent(&out, []byte(o), "", "\t"); err != nil {
			return o
		}
		res, _ := out.ReadString(0)
		return res
	default:
		if barr, err := json.MarshalIndent(o, "", "\t"); err != nil {
			return ""
		} else {
			return string(barr)
		}
	}
}
