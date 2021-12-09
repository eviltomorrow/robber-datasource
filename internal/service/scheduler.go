package service

import (
	"context"
	"fmt"
	"time"

	"github.com/eviltomorrow/robber-core/pkg/mongodb"
	"github.com/eviltomorrow/robber-core/pkg/zlog"
	"github.com/eviltomorrow/robber-core/pkg/zmath"
	client_repository "github.com/eviltomorrow/robber-repository/pkg/client"
	pb_repository "github.com/eviltomorrow/robber-repository/pkg/pb"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

var DefaultCronSpec = "05 17 * * MON,TUE,WED,THU,FRI"

func RunScheduler() {
	var c = cron.New()
	_, err := c.AddFunc(DefaultCronSpec, func() {
		var (
			now = time.Now()
		)

		zlog.Info("Fetch metadata begin", zap.String("date", now.Format("2006-01-02")))
		date, fetchCount, err := FetchMetadataFromSina(now, true)
		if err != nil {
			zlog.Error("Cron: FetchMetadataFromSina failure", zap.Error(err))
		}

		pushCount, err := PushMetadataToRepository(date)
		if err != nil {
			zlog.Error("Cron: PushMetadataToRepository failure", zap.Error(err))
		}

		zlog.Info("Fetch metadata complete", zap.String("date", date), zap.Int64("fetch-count", fetchCount), zap.Int64("push-count", pushCount), zap.Duration("cost", time.Since(now)))
	})
	if err != nil {
		zlog.Fatal("Cron add func failure", zap.Error(err))
	}
	c.Start()
}

func PushMetadataToRepository(date string) (int64, error) {
	var (
		offset  int64 = 0
		limit   int64 = 100
		count   int64 = 0
		lastID  string
		timeout = 20 * time.Second
	)

	rstub, cancel, err := client_repository.NewClientForRepository()
	if err != nil {
		return count, err
	}
	defer cancel()

	req, err := rstub.PushData(context.Background())
	if err != nil {
		return count, err
	}

	for {
		data, err := SelectMetadataRange(mongodb.DB, offset, limit, date, lastID, timeout)
		if err != nil {
			return count, err
		}
		for _, d := range data {
			if d.Volume == 0 {
				continue
			}
			err := req.Send(&pb_repository.Metadata{
				Code:            d.Code,
				Name:            d.Name,
				Open:            d.Open,
				YesterdayClosed: d.YesterdayClosed,
				Latest:          d.Latest,
				High:            d.High,
				Low:             d.Low,
				Volume:          d.Volume,
				Account:         d.Account,
				Date:            d.Date,
				Time:            d.Time,
				Suspend:         d.Suspend,
			})
			if err != nil {
				_, e1 := req.CloseAndRecv()
				return count, fmt.Errorf("%v, nest error: %v", err, e1)
			}
			count++
		}

		if len(data) < int(limit) {
			break
		}
		if len(data) != 0 {
			lastID = data[len(data)-1].ObjectID
		}
		offset += limit
	}

	_, err = req.CloseAndRecv()
	return count, err
}

func FetchMetadataFromSina(today time.Time, delay bool) (string, int64, error) {
	var (
		date = today.Format("2006-01-02")

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
			metadata, err := CollectMetadataFromSina(codes)
			if err != nil {
				retrytimes++
				if retrytimes == 10 {
					zlog.Error("Fetch: CollectMetadataFromSina failure", zap.Strings("codes", codes), zap.Error(err))
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
					zlog.Error("Fetch: DeleteMetadataByDate failure", zap.Strings("codes", codes), zap.Error(err))
				}
			}
			affected, err := InsertMetadataMany(mongodb.DB, metadata, timeout)
			if err != nil {
				zlog.Error("Fetch: InsertMetadataMany failure", zap.Error(err))
			}
			count += affected
			if delay {
				time.Sleep(time.Duration(zmath.GenRandInt(10, 30)) * time.Second)
			}
		}
	}

	if len(codes) != 0 {
	retry_2:
		metadata, err := CollectMetadataFromSina(codes)
		if err != nil {
			retrytimes++
			if retrytimes == 10 {
				zlog.Error("Fetch: CollectMetadataFromSina failure", zap.Strings("codes", codes), zap.Error(err))
			} else {
				time.Sleep(30 * time.Second)
				goto retry_2
			}
		}

		if len(metadata) != 0 {
			for _, md := range metadata {
				_, err := DeleteMetadataByDate(mongodb.DB, md.Code, md.Date, timeout)
				if err != nil {
					zlog.Error("Fetch: DeleteMetadataByDate failure", zap.Strings("codes", codes), zap.Error(err))
				}
			}
			affected, err := InsertMetadataMany(mongodb.DB, metadata, timeout)
			if err != nil {
				zlog.Error("Fetch: InsertMetadataMany failure", zap.Error(err))
			}
			count += affected
		}
	}
	return date, count, nil
}
