package rdb

import (
	"time"
)

func PersistPeriodically() error {
	for range time.Tick(time.Second * 30) {
		err := Persist()
		if err != nil {
			return err
		}
	}

	return nil
}
