package cmd

import (
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ceelei",
	Short: "Ceelei is a CLI helper inside the CLI to help you remember your commands.",
	Long: `Ceelei will help you to remember your favourite CLI commands without the need to search for it in the Web or go through your notes. It will be saved into a local database within your home directory called ceelei.db.
Stay inside the CLI, stay focussed!`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(pDB *bolt.DB) {
	db = pDB
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	cobra.CheckErr(rootCmd.Execute())
}