package service

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/eviltomorrow/robber-core/pkg/zlog"
	"go.uber.org/zap"
)

var CodeList = []string{
	"sh688***",
	"sh605***",
	"sh603***",
	"sh601***",
	"sh600***",
	"sz300***",
	"sz0030**",
	"sz002***",
	"sz000***",
}

func BuildRangeCode() chan string {
	var data = make(chan string, 64)
	go func() {
		for _, code := range CodeList {
			result, err := buildRangeCode(code)
			if err != nil {
				zlog.Error("Build range code failure", zap.Error(err))
				continue
			}
			for _, r := range result {
				data <- r
			}
		}
		close(data)
	}()
	return data
}

func buildRangeCode(baseCode string) ([]string, error) {
	if len(baseCode) != 8 {
		return nil, fmt.Errorf("code length must be 8, code is [%s]", baseCode)
	}
	if !strings.HasPrefix(baseCode, "sh") && !strings.HasPrefix(baseCode, "sz") {
		return nil, fmt.Errorf("code must be start with [sh/sz], code is [%s]", baseCode)
	}

	if !strings.Contains(baseCode, "*") {
		return []string{baseCode}, nil
	}

	var (
		n      = strings.Index(baseCode, "*")
		prefix = baseCode[:n]
		codes  = make([]string, 0, int(math.Pow10(8-n)))
	)

	var builder strings.Builder
	builder.Grow(8)

	var next = int(math.Pow10(8-n)) - 1
	var mid = ""
	var count = -1
	var changed = false
	for i := next; i >= 0; i-- {
		if i == next && i != 0 {
			next = i / 10
			count++
			changed = true
			mid = ""
		} else {
			changed = false
		}

		if changed {
			for j := 0; j < count; j++ {
				mid += "0"
			}
		}

		builder.WriteString(prefix)
		builder.WriteString(mid)
		builder.WriteString(strconv.Itoa(i))
		codes = append(codes, builder.String())
		builder.Reset()
	}
	return codes, nil
}
