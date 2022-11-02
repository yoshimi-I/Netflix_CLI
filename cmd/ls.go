package cmd

import (
	"github.com/spf13/cobra"
	"netflix/src"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long: `You can show your videos log:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	//以下に処理を書く
	// 履歴の画面からスクレイピング
	Run: func(cmd *cobra.Command, args []string) {
		webPage := ("https://www.netflix.com/settings/viewed/BPTKRBAK5ZC6PFLO6OTBD2EF6A")
		src.Scraip(webPage)

	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
