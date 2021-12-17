package command

import (
	"log"
	"time"

	"github.com/eviltomorrow/robber-core/pkg/mongodb"
	"github.com/eviltomorrow/robber-datasource/internal/model"
	"github.com/eviltomorrow/robber-datasource/internal/service"
	"github.com/spf13/cobra"
)

var recoverCmd = &cobra.Command{
	Use:   "recover",
	Short: "Recover data from log",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		setupCfg()
		setupVars()
		mongodb.Build()

		if logPath == "" {
			log.Fatalf("[Fatal] invalid data log file, logPath: [%s]", logPath)
		}
		data, err := service.CollectMetadataFromLog(logPath)
		if err != nil {
			log.Fatalf("[Fatal] CollectMetadataFromLog failure, nest error: %v\r\n", err)
		}

		var (
			count    int64
			timeout  = 10 * time.Second
			size     = 30
			metadata = make([]*model.Metadata, 0, size)
			s        int64
		)
		for d := range data {
			if d.Volume != 0 {
				metadata = append(metadata, d)
			}

			if len(metadata) == size {
				for _, md := range metadata {
					_, err := service.DeleteMetadataByDate(mongodb.DB, md.Code, md.Date, timeout)
					if err != nil {
						log.Fatalf("[Fatal] DeleteMetadataByDate failure, code: %s, date: %s, nest error: %v\r\n", md.Code, md.Date, err)
					}
				}
				affected, err := service.InsertMetadataMany(mongodb.DB, metadata, timeout)
				if err != nil {
					log.Fatalf("[Fatal] InsertMetadataMany failure, nest error: %v\r\n", err)
				}
				count += affected
				metadata = metadata[:0]
			}
			if s != count/1000 {
				s = count / 1000
				log.Printf("[Info] Recovering data from log file[%s]: %d\r\n", logPath, s*1000)
			}
		}
		if len(metadata) != 0 {
			for _, md := range metadata {
				_, err := service.DeleteMetadataByDate(mongodb.DB, md.Code, md.Date, timeout)
				if err != nil {
					log.Fatalf("[Fatal] DeleteMetadataByDate failure, code: %s, date: %s, nest error: %v\r\n", md.Code, md.Date, err)
				}
			}
			affected, err := service.InsertMetadataMany(mongodb.DB, metadata, timeout)
			if err != nil {
				log.Fatalf("[Fatalf] InsertMetadataMany failure, nest error: %v\r\n", err)
			}
			count += affected
		}
		log.Printf("[Info] Recover data from log file[%s] complete, total: %d\r\n", logPath, count)
	},
}

var (
	logPath = ""
)

func init() {
	recoverCmd.Flags().StringVarP(&logPath, "log", "l", "", "robber-datasource's data log file")
	recoverCmd.MarkFlagRequired("log")

	recoverCmd.Flags().StringVarP(&cfgPath, "config", "c", "config.toml", "robber-datasource's config file")

	rootCmd.AddCommand(recoverCmd)
}
