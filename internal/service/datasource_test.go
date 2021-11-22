package service

import (
	"testing"
	"time"

	"github.com/eviltomorrow/robber-core/pkg/mongodb"
	"github.com/eviltomorrow/robber-datasource/internal/model"
	"github.com/stretchr/testify/assert"
)

const (
	date = "2021-04-21"
)

func init() {
	mongodb.DSN = "mongodb://localhost:27017"
	mongodb.Build()
}
func TestInsertMetadataMany(t *testing.T) {
	_assert := assert.New(t)

	var md = &model.Metadata{
		Code:            "sz000001",
		Name:            "平安银行",
		Open:            10.3,
		YesterdayClosed: 10.67,
		Latest:          10.67,
		High:            10.88,
		Low:             10.11,
		Volume:          29322303,
		Account:         10.72 * 29322303,
		Date:            date,
		Time:            "15:03:00",
		Suspend:         "正常",
	}

	affected, err := InsertMetadataMany(mongodb.DB, []*model.Metadata{md}, 10*time.Second)
	_assert.Nil(err)
	_assert.Equal(int64(1), affected)

	var count = 5000
	var data = make([]*model.Metadata, 0, count)
	for i := 0; i < count; i++ {
		data = append(data, md)
	}
	affected, err = InsertMetadataMany(mongodb.DB, data, 10*time.Second)
	_assert.Nil(err)
	_assert.Equal(int64(count), affected)

	affected, err = DeleteMetadataByDate(mongodb.DB, md.Code, date, 10*time.Second)
	_assert.Nil(err)
	_assert.Equal(int64(count+1), affected)
}

func TestDeleteMetadataByDate(t *testing.T) {
	_assert := assert.New(t)

	affected, err := DeleteMetadataByDate(mongodb.DB, "", "2030-04-05", 10*time.Second)
	_assert.Nil(err)
	_assert.Equal(int64(0), affected)

	var md = &model.Metadata{
		Code:            "sz000001",
		Name:            "平安银行",
		Open:            10.3,
		YesterdayClosed: 10.67,
		Latest:          10.67,
		High:            10.88,
		Low:             10.11,
		Volume:          29322303,
		Account:         10.72 * 29322303,
		Date:            date,
		Time:            "15:03:00",
		Suspend:         "正常",
	}

	affected, err = InsertMetadataMany(mongodb.DB, []*model.Metadata{md}, 10*time.Second)
	_assert.Nil(err)
	_assert.Equal(int64(1), affected)

	affected, err = DeleteMetadataByDate(mongodb.DB, md.Code, date, 10*time.Second)
	_assert.Nil(err)
	_assert.Equal(int64(1), affected)
}

func BenchmarkInsertMetadataMany(b *testing.B) {
	var md = &model.Metadata{
		Code:            "sz000001",
		Name:            "平安银行",
		Open:            10.3,
		YesterdayClosed: 10.67,
		Latest:          10.67,
		High:            10.88,
		Low:             10.11,
		Volume:          29322303,
		Account:         10.72 * 29322303,
		Date:            date,
		Time:            "15:03:00",
		Suspend:         "正常",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertMetadataMany(mongodb.DB, []*model.Metadata{md}, 10*time.Second)
	}
}
