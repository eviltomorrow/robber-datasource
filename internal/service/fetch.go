package service

import (
	"bufio"
	"fmt"
	"strings"
	"time"

	"github.com/eviltomorrow/robber-core/pkg/httpclient"
	"github.com/eviltomorrow/robber-core/pkg/zlog"
	"github.com/eviltomorrow/robber-datasource/internal/model"
	"go.uber.org/zap"
)

var (
	ErrNoStockCode = fmt.Errorf("no exist stock code in Sina API")
	ErrNoStockData = fmt.Errorf("no stock data")
)

// FetchMetadataFromSina fetch data from sina
func FetchMetadataFromSina(codes []string) ([]*model.Metadata, error) {
	data, err := httpclient.GetHTTP(fmt.Sprintf("http://hq.sinajs.cn/list=%s", strings.Join(codes, ",")), 20*time.Second, httpclient.DefaultHeader)
	if err != nil {
		return nil, err
	}

	var result = make([]*model.Metadata, 0, len(codes))
	for key, val := range parseDataToMap(data) {
		metadata, err := parseDataToTrade(key, val)
		if err != nil {
			continue
		}
		result = append(result, metadata)
	}
	return result, nil
}

func parseDataToMap(data string) map[string]string {
	var result = make(map[string]string)

	var scanner = bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		var text = strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		if !strings.HasPrefix(text, "var") || !strings.HasSuffix(text, ";") {
			continue
		}

		var n = strings.Index(text, "=")
		if n == -1 {
			continue
		}

		var code = strings.Replace(text[:n], "var hq_str_", "", -1)
		result[code] = text
	}
	return result
}

func parseDataToTrade(code, data string) (*model.Metadata, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("no data received from Sina API")
	}

	var begin = strings.Index(strings.TrimSpace(data), `"`)
	var end = strings.LastIndex(strings.TrimSpace(data), `"`)

	if begin == -1 || end == -1 || begin == end {
		zlog.Error("Invalid trade data", zap.String("code", code), zap.String("data", data))
		return nil, fmt.Errorf("invalid data received from Sina API")
	}

	var attr = strings.Split(data[begin+1:end], ",")
	if len(attr) == 1 {
		return nil, ErrNoStockCode
	}
	if len(attr) >= 2 && attr[len(attr)-1] == "" {
		attr = attr[:len(attr)-1]
	}
	switch {
	case strings.HasPrefix(code, "sh6"):
		if len(attr) != 33 {
			zlog.Warn("Invalid trade data", zap.String("code", code), zap.String("data", data), zap.Int("len", len(attr)))
			return nil, ErrNoStockData
		}
	case strings.HasPrefix(code, "sz0"):
		if len(attr) != 33 {
			zlog.Warn("Invalid trade data", zap.String("code", code), zap.String("data", data), zap.Int("len", len(attr)))
			return nil, ErrNoStockData
		}
	case strings.HasPrefix(code, "sz3"):
		if len(attr) != 33 {
			zlog.Warn("Invalid trade data", zap.String("code", code), zap.String("data", data), zap.Int("len", len(attr)))
			return nil, ErrNoStockData
		}
	default:
		return nil, fmt.Errorf("no support code[%v]", code)
	}

	var md = &model.Metadata{
		Code: code,
	}
	for i, val := range attr {
		md.Assign(i, val)
	}
	return md, nil
}
