package models

import "time"

type Weather struct {
	Address     string `json:"resolvedAddress"`
	Timezone    string `json:"timezone"`
	Description string `json:"description"`
	Timestamp   time.Time
}
