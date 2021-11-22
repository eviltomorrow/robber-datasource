package command

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/eviltomorrow/robber-core/pkg/mongodb"
	"github.com/eviltomorrow/robber-core/pkg/zlog"
	"github.com/eviltomorrow/robber-datasource/internal/config"
	"github.com/eviltomorrow/robber-datasource/internal/server"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use:   "robber-datasource",
	Short: "",
	Long:  "  \r\nrobber-datasource server running",
	Run: func(cmd *cobra.Command, args []string) {
		setupCfg()
		setupVars()
		if err := mongodb.Build(); err != nil {
			zlog.Fatal("Build mongodb connection failure", zap.Error(err))
		}

		if err := server.StartupGRPC(); err != nil {
			zlog.Fatal("Startup GRPC service failure", zap.Error(err))
		}
		registerCleanFuncs()
		blockingUntilTermination()
	},
}

var (
	cleanFuncs []func() error
	cfg        = config.GlobalConfig
	cfgPath    = ""
)

func init() {
	rootCmd.CompletionOptions = cobra.CompletionOptions{
		DisableDefaultCmd: true,
	}
	rootCmd.Flags().StringVarP(&cfgPath, "config", "c", "config.toml", "robber-datasource's config file")
	rootCmd.MarkFlagRequired("config")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func blockingUntilTermination() {
	var ch = make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	switch <-ch {
	case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
	case syscall.SIGUSR1:
	case syscall.SIGUSR2:
	default:
	}

	for _, f := range cleanFuncs {
		f()
	}
}

func registerCleanFuncs() {
	cleanFuncs = append(cleanFuncs, server.RevokeEtcdConn)
	cleanFuncs = append(cleanFuncs, server.ShutdownGRPC)
}

func setupCfg() {
	if err := cfg.Load(cfgPath, nil); err != nil {
		log.Fatalf("[Fatal] Load config file failure, nest error: %v\r\n", err)
	}

	global, prop, err := zlog.InitLogger(&zlog.Config{
		Level:            cfg.Log.Level,
		Format:           cfg.Log.Format,
		DisableTimestamp: cfg.Log.DisableTimestamp,
		File: zlog.FileLogConfig{
			Filename:   cfg.Log.FileName,
			MaxSize:    cfg.Log.MaxSize,
			MaxDays:    30,
			MaxBackups: 30,
		},
		DisableStacktrace: true,
	})
	if err != nil {
		log.Fatalf("[Fatal] Setup log config failure, nest error: %v\r\n", err)
	}
	zlog.ReplaceGlobals(global, prop)
	zlog.Info("Global Config info", zap.String("cfg", cfg.String()))
}

func setupVars() {
	server.Host = cfg.Server.Host
	server.Port = cfg.Server.Port
	server.Endpoints = cfg.Etcd.Endpoints

	mongodb.DSN = cfg.MongoDB.DSN

}
