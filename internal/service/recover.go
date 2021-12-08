package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eviltomorrow/robber-datasource/internal/model"
)

func CollectMetadataFromLog(path string) (chan *model.Metadata, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	var data = make(chan *model.Metadata, 128)
	go func() {
		var scanner = bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			var text = scanner.Text()
			if !strings.Contains(text, "Invalid trade data") {
				continue
			}
			var (
				result = parseLogLine(text)
			)
			for _, r := range result {
				if strings.HasPrefix(r, `data="var `) {
					r = strings.Replace(r, `\"`, `"`, -1)
					r = strings.Replace(r, "data=", "", -1)
					r = strings.TrimPrefix(r, `"`)
					r = strings.TrimSuffix(r, `"`)
					if !strings.HasPrefix(r, "var") || !strings.HasSuffix(r, ";") {
						continue
					}
					var n = strings.Index(r, "=")
					if n == -1 {
						continue
					}

					var code = strings.Replace(r[:n], "var hq_str_", "", -1)
					metadata, err := parseDataToTradeNolog(code, r)
					if err == nil {
						data <- metadata
					}
				}
			}
		}
		close(data)
		file.Close()
	}()
	return data, nil
}

func parseLogLine(line string) []string {
	var (
		result     = make([]string, 0, 8)
		begin, end int
	)
	for i, b := range line {
		if b == '[' {
			begin = i
		}
		if b == ']' {
			end = i
			result = append(result, line[begin+1:end])
		}
	}
	return result
}

func parseDataToTradeNolog(code, data string) (*model.Metadata, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("no data received from Sina API")
	}

	var begin = strings.Index(strings.TrimSpace(data), `"`)
	var end = strings.LastIndex(strings.TrimSpace(data), `"`)

	if begin == -1 || end == -1 || begin == end {
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
	case strings.HasPrefix(code, "sh68"):
		if len(attr) != 34 {
			return nil, ErrNoStockData
		}
	case strings.HasPrefix(code, "sh60"):
		if len(attr) != 33 {
			return nil, ErrNoStockData
		}
	case strings.HasPrefix(code, "sz0"):
		if len(attr) != 33 {
			return nil, ErrNoStockData
		}
	case strings.HasPrefix(code, "sz3"):
		if len(attr) != 33 {
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
