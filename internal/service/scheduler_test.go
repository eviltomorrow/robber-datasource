package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_fetchMetadataFromSina(t *testing.T) {
	_assert := assert.New(t)
	err := initMongodb()
	_assert.Nil(err)

	date, err := fetchMetadataFromSina(false)
	_assert.Nil(err)
	_assert.Equal(time.Now().Format("2006-01-02"), date)
}
