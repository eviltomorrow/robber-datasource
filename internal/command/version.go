package command

import (
	"bytes"
	"fmt"

	"github.com/eviltomorrow/robber-core/pkg/system"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version about robber-datasource",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		printClientVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func printClientVersion() {
	var buf bytes.Buffer
	buf.WriteString("Client: \r\n")
	buf.WriteString(fmt.Sprintf("   robber-datasource Version (Current): %s\r\n", system.MainVersion))
	buf.WriteString(fmt.Sprintf("   Go Version: %v\r\n", system.GoVersion))
	buf.WriteString(fmt.Sprintf("   Go OS/Arch: %v\r\n", system.GoOSArch))
	buf.WriteString(fmt.Sprintf("   Git Sha: %v\r\n", system.GitSha))
	buf.WriteString(fmt.Sprintf("   Git Tag: %v\r\n", system.GitTag))
	buf.WriteString(fmt.Sprintf("   Git Branch: %v\r\n", system.GitBranch))
	buf.WriteString(fmt.Sprintf("   Build Time: %v\r\n", system.BuildTime))
	fmt.Println(buf.String())
}
