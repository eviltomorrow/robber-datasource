package service

import (
	"log"
	"testing"
	"time"

	"github.com/eviltomorrow/robber-core/pkg/mongodb"
	"github.com/stretchr/testify/assert"
)

func TestFetchMetadataFromSina(t *testing.T) {
	_assert := assert.New(t)
	var codes = []string{
		"sh601012", "sz300002", "sz000001", "sh688009",
	}
	data, err := FetchMetadataFromSina(codes)
	_assert.Nil(err)
	t.Logf("data: %s\r\n", data)
}

func TestFetchMetadataLatest(t *testing.T) {
	var (
		now        = time.Now()
		retrytimes = 0
		count      int64
		timeout    = 10 * time.Second
		size       = 30
		codes      = make([]string, 0, size)
	)
	for code := range BuildRangeCode() {
		codes = append(codes, code)
		if len(codes) == size {
		retry_1:
			metadata, err := FetchMetadataFromSina(codes)
			if err != nil {
				retrytimes++
				if retrytimes == 10 {
					log.Fatalf("FetchMetadataFromSina failure, nest codes: %v, err: %v\r\n", codes, err)
				} else {
					time.Sleep(30 * time.Second)
					goto retry_1
				}
			}
			retrytimes = 0
			codes = codes[:0]

			if len(metadata) == 0 {
				continue
			}

			for _, md := range metadata {
				_, err := DeleteMetadataByDate(mongodb.DB, md.Code, md.Date, timeout)
				if err != nil {
					log.Fatalf("DeleteMetadataByDate failure, nest codes: %v, err: %v\r\n", codes, err)
				}
			}
			affected, err := InsertMetadataMany(mongodb.DB, metadata, timeout)
			if err != nil {
				log.Fatalf("InsertMetadataMany failure, nest error: %v\r\n", err)
			}
			count += affected
		}
	}

	if len(codes) != 0 {
	retry_2:
		metadata, err := FetchMetadataFromSina(codes)
		if err != nil {
			retrytimes++
			if retrytimes == 10 {
				log.Fatalf("FetchMetadataFromSina failure, nest codes: %v, err: %v\r\n", codes, err)
			} else {
				time.Sleep(30 * time.Second)
				goto retry_2
			}
		}

		if len(metadata) != 0 {
			for _, md := range metadata {
				_, err := DeleteMetadataByDate(mongodb.DB, md.Code, md.Date, timeout)
				if err != nil {
					log.Fatalf("DeleteMetadataByDate failure, nest codes: %v, err: %v\r\n", codes, err)
				}
			}
			affected, err := InsertMetadataMany(mongodb.DB, metadata, timeout)
			if err != nil {
				log.Fatalf("InsertMetadataMany failure, nest error: %v\r\n", err)
			}
			count += affected
		}
	}
	log.Printf("Finish sync metadata, date: %v, count: %v, cost: %v\r\n", date, count, time.Since(now))

}
