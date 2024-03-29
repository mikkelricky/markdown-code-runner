package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [shell]",
	Short: fmt.Sprintf("Generate completions for %[1]s", appName),
	Long: fmt.Sprintf(`Generate shell completions for %[1]s.

Examples:

%[2]s completion bash
%[2]s completion fish
%[2]s completion powershell
%[2]s completion zsh
`, appName, mainScript),
	ValidArgs: []string{"bash", "fish", "powershell", "zsh"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		shell := args[0]
		switch shell {
		case "bash":
			rootCmd.GenBashCompletionV2(os.Stdout, true)

		case "fish":
			rootCmd.GenFishCompletion(os.Stdout, true)

		case "powershell":
			rootCmd.GenPowerShellCompletion(os.Stdout)

		case "zsh":
			rootCmd.GenZshCompletion(os.Stdout)

		default:
			log.Fatalf("invalid shell: %s", shell)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
