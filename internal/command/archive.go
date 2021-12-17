package command

import (
	"log"
	"time"

	"github.com/eviltomorrow/robber-core/pkg/mongodb"
	"github.com/eviltomorrow/robber-datasource/internal/service"
	"github.com/spf13/cobra"
)

var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive data with specify date to repository",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			now = time.Now()
		)
		if beginDate == "" || endDate == "" {
			log.Fatalf("invalid begin/end date param\r\n")
		}
		begin, err := time.ParseInLocation("2006-01-02", beginDate, time.Local)
		if err != nil {
			log.Fatalf("ParseInLocation begin failure, nest error: %v\r\n", err)
		}
		end, err := time.ParseInLocation("2006-01-02", endDate, time.Local)
		if err != nil {
			log.Fatalf("ParseInLocation end failure, nest error: %v\r\n", err)
		}
		end = end.Add(1 * time.Second)

		setupCfg()
		setupVars()
		if err := mongodb.Build(); err != nil {
			log.Fatalf("Build mongodb connection failure, nest error: %v\r\n", err)
		}

		var total int64 = 0
		for begin.Before(end) {
			count, err := service.PushMetadataToRepository(begin.Format("2006-01-02"))
			if err != nil {
				log.Fatalf("[failure] date: %v, error: %v\r\n", begin.Format("2006-01-02"), err)
			} else {
				log.Printf("[success] date: %v, count: %v\r\n", begin.Format("2006-01-02"), count)
			}
			begin = begin.AddDate(0, 0, 1)
			total += count
		}
		log.Printf("[complete] total count: %v, cost: %v\r\n", total, time.Since(now))
	},
}

var (
	beginDate, endDate string
)

func init() {
	archiveCmd.Flags().StringVar(&beginDate, "begin", "", "archive begin param")
	archiveCmd.Flags().StringVar(&endDate, "end", "", "archive end param")
	archiveCmd.MarkFlagRequired("begin")
	archiveCmd.MarkFlagRequired("end")
	archiveCmd.Flags().StringVarP(&cfgPath, "config", "c", "config.toml", "robber-datasource's config file")

	rootCmd.AddCommand(archiveCmd)
}
