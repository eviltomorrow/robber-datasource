package model

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/eviltomorrow/robber-core/pkg/zlog"
	"go.uber.org/zap"
)

// Metadata trade data
type Metadata struct {
	ObjectID        string  `json:"_id" bson:"_id"`
	Code            string  `json:"code" bson:"code"`
	Name            string  `json:"name" bson:"name"`                         // 0 股票简称
	Open            float64 `json:"open" bson:"open"`                         // 1 今日开盘价格
	YesterdayClosed float64 `json:"yesterday_closed" bson:"yesterday_closed"` // 2 昨日收盘价格
	Latest          float64 `json:"latest" bson:"latest"`                     // 3 最近成交价格
	High            float64 `json:"high" bson:"high"`                         // 4 最高成交价
	Low             float64 `json:"low" bson:"low"`                           // 5 最低成交价
	Volume          uint64  `json:"volume" bson:"volume"`                     // 8 成交数量（股）
	Account         float64 `json:"account" bson:"account"`                   // 9 成交金额（元）
	Date            string  `json:"date" bson:"date"`                         // 30 日期
	Time            string  `json:"time" bson:"time"`                         // 31 时间
	Suspend         string  `json:"suspend" bson:"suspend"`                   // 32 停牌状态
}

func (m *Metadata) String() string {
	buf, err := json.Marshal(m)
	if err != nil {
		return fmt.Sprintf("Trade marshal json failure, nest error: %v", err)
	}
	return string(buf)
}

func atof64(name string, loc int, val string) float64 {
	f64, err := strconv.ParseFloat(val, 64)
	if err != nil {
		zlog.Error("ParseFloat64 failure", zap.String("name", name), zap.Int("loc", loc), zap.String("val", val))
		return 0
	}
	return f64
}

func atou64(name string, loc int, val string) uint64 {
	u64, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		zlog.Error("ParseUint64 failure", zap.String("name", name), zap.Int("loc", loc), zap.String("val", val))
		return 0
	}
	return u64
}

//
const (
	suspendNormal    = "正常"
	suspendOneHour   = "停牌一小时"
	suspendOneDay    = "停牌一天"
	suspendKeep      = "连续停牌"
	suspendMid       = "盘中停牌"
	suspendHalfOfDay = "停牌半天"
	suspendPause     = "暂停"
	suspendNoRecord  = "无该记录"
	suspendUnlisted  = "未上市"
	suspendDelist    = "退市"
	suspendUnknown   = "未知"
)

func (m *Metadata) Assign(loc int, val string) {
	switch loc {
	case 0:
		m.Name = val
	case 1:
		m.Open = atof64(m.Name, loc, val)
	case 2:
		m.YesterdayClosed = atof64(m.Name, loc, val)
	case 3:
		m.Latest = atof64(m.Name, loc, val)
	case 4:
		m.High = atof64(m.Name, loc, val)
	case 5:
		m.Low = atof64(m.Name, loc, val)
	case 8:
		m.Volume = atou64(m.Name, loc, val)
	case 9:
		m.Account = atof64(m.Name, loc, val)
	case 30:
		m.Date = val
	case 31:
		m.Time = val
	case 32:
		m.Suspend = getSuspendDesc(val)
	default:
	}
}

// getSuspendDesc get suspend desc
func getSuspendDesc(val string) string {
	switch {
	case val == "00":
		return suspendNormal
	case val == "01":
		return suspendOneHour
	case val == "02":
		return suspendOneDay
	case val == "03":
		return suspendKeep
	case val == "04":
		return suspendMid
	case val == "05":
		return suspendHalfOfDay
	case val == "07":
		return suspendPause
	case val == "-1":
		return suspendNoRecord
	case val == "-2":
		return suspendUnlisted
	case val == "-3":
		return suspendDelist
	default:
		return suspendUnknown
	}
}
