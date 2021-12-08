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

	pushCount, err := PushMetadataToRepository(date)
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("fetch: %d, push: %d\r\n", fetchCount, pushCount)
}
