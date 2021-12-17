package command

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

$ source <(robber-datasource completion bash)

# To load completions for each session, execute once:
Linux:
  $ robber-datasource completion bash > /etc/bash_completion.d/robber-datasource
MacOS:
  $ robber-datasource completion bash > /usr/local/etc/bash_completion.d/robber-datasource

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ robber-datasource completion zsh > "${fpath[1]}/_robber-datasource"

# You will need to start a new shell for this setup to take effect.

Fish:

$ robber-datasource completion fish | source

# To load completions for each session, execute once:
$ robber-datasource completion fish > ~/.config/fish/completions/robber-datasource.fish
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
