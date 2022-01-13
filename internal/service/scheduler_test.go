package service

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOneCron(t *testing.T) {
	_assert := assert.New(t)
	err := initMongodb()
	_assert.Nil(err)

	var (
		now = time.Now()
	)

	date, fetchCount, err := FetchMetadataFromSina(now, false)
	if err != nil {
		log.Fatal(err)
	}
	_ = date

	date = "2021-12-17"
	metadata, stock, day, week, err := PushMetadataToRepository(date)
	if err != nil {
		log.Fatal(err)
	}

	err = CompleteTaskToRepository(date, metadata, stock, day, week)
	if err != nil {
		log.Printf("complete failure, nest error: %v\r\n", err)
	}
	t.Logf("fetch: %d, push: %d, stock: %d, day: %d, week: %d\r\n", fetchCount, metadata, stock, day, week)
}
