package service

import (
	"time"

	"github.com/eviltomorrow/robber-core/pkg/mongodb"
	"github.com/eviltomorrow/robber-core/pkg/zlog"
	"github.com/eviltomorrow/robber-core/pkg/zmath"
	"github.com/robfig/cron"
	"go.uber.org/zap"
)

var DefaultCronSpec = "05 17 * * MON,TUE,WED,THU,FRI"

func RunSchedulerBackground() {
	var c = cron.New()
	_, err := c.AddFunc(DefaultCronSpec, func() {
		date, err := fetchMetadataFromSina(true)
		if err != nil {
			zlog.Error("Cron fetchMetadataFromSina failure", zap.Error(err))
		}

		_ = date
	})
	if err != nil {
		zlog.Fatal("Cron add func failure", zap.Error(err))
	}
	c.Start()
}

func fetchMetadataFromSina(delay bool) (string, error) {
	var (
		now  = time.Now()
		date = now.Format("2006-01-02")

		retrytimes = 0
		count      int64
		timeout    = 10 * time.Second
		size       = 30
		codes      = make([]string, 0, size)
	)
	zlog.Info("Begin sync metadata", zap.String("date", date))

	for code := range BuildRangeCode() {
		codes = append(codes, code)
		if len(codes) == size {
		retry_1:
			metadata, err := FetchMetadataFromSina(codes)
			if err != nil {
				retrytimes++
				if retrytimes == 10 {
					zlog.Error("FetchMetadataFromSina failure", zap.Strings("codes", codes), zap.Error(err))
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
					zlog.Error("DeleteMetadataByDate failure", zap.Strings("codes", codes), zap.Error(err))
				}
			}
			affected, err := InsertMetadataMany(mongodb.DB, metadata, timeout)
			if err != nil {
				zlog.Error("InsertMetadataMany failure", zap.Error(err))
			}
			count += affected
			if delay {
				time.Sleep(time.Duration(zmath.GenRandInt(10, 30)) * time.Second)
			}
		}
	}

	if len(codes) != 0 {
	retry_2:
		metadata, err := FetchMetadataFromSina(codes)
		if err != nil {
			retrytimes++
			if retrytimes == 10 {
				zlog.Error("FetchMetadataFromSina failure", zap.Strings("codes", codes), zap.Error(err))
			} else {
				time.Sleep(30 * time.Second)
				goto retry_2
			}
		}

		if len(metadata) != 0 {
			for _, md := range metadata {
				_, err := DeleteMetadataByDate(mongodb.DB, md.Code, md.Date, timeout)
				if err != nil {
					zlog.Error("DeleteMetadataByDate failure", zap.Strings("codes", codes), zap.Error(err))
				}
			}
			affected, err := InsertMetadataMany(mongodb.DB, metadata, timeout)
			if err != nil {
				zlog.Error("InsertMetadataMany failure", zap.Error(err))
			}
			count += affected
		}
	}
	zlog.Info("Finish sync metadata", zap.String("date", date), zap.Int64("count", count), zap.Duration("cost", time.Since(now)))
	return date, nil
}
