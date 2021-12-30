package command

import (
	"log"
	"time"

	"github.com/eviltomorrow/robber-core/pkg/mongodb"
	"github.com/eviltomorrow/robber-datasource/internal/service"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch new data to mongodb",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		setupCfg()
		setupVars()
		if err := mongodb.Build(); err != nil {
			log.Fatalf("Build mongodb connection failure, nest error: %v\r\n", err)
		}

		date, fetchCount, err := service.FetchMetadataFromSina(time.Now(), false)
		if err != nil {
			log.Fatalf("FetchMetadataFromSina failure, nest error: %v\r\n", err)
		}
		log.Printf("[%s]Fetch success, date: %v, count: %v\r\n", color.GreenString("success"), date, fetchCount)
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
