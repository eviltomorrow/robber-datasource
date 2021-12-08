package command

import (
	"log"
	"time"

	"github.com/eviltomorrow/robber-datasource/internal/service"
	"github.com/spf13/cobra"
)

var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive data with specify date to repository",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
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

		for begin.Before(end) {
			count, err := service.PushMetadataToRepository(begin.Format("2006-01-02"))
			if err != nil {
				log.Fatalf("[PushMetadataToRepository failure], nest error: %v, date: %v\r\n", err, begin.Format("2006-01-02"))
			} else {
				log.Printf("[PushMetadataToRepository success], date: %v, count: %v\r\n", begin.Format("2006-01-02"), count)
			}
			begin = begin.AddDate(0, 0, 1)
		}
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

	rootCmd.AddCommand(archiveCmd)
}
