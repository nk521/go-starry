package cmd

import (
	"fmt"
	"os"

	config "github.com/nk521/go-starry/config"
	log "github.com/nk521/go-starry/log"
	"github.com/nk521/go-starry/tui"

	"github.com/nk521/go-starry/util"
	ytm "github.com/nk521/go-starry/youtube_music"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "starry",
	Short: "Youtube Music but for your terminal!",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Starting starry...")
		tui.RednerTUI()
	},
}

var headerCmd = &cobra.Command{
	Use:   "headers",
	Short: "Set headers in config!",
	Run: func(cmd *cobra.Command, args []string) {
		ytm.GetHeaders()
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show configs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("config.GetRawConfig().ConfigFileUsed(): %v\n", config.GetRawConfig().ConfigFileUsed())
		err := util.OpenEditor(config.GetRawConfig().ConfigFileUsed())
		if err != nil {
			log.Panicln(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	rootCmd.AddCommand(headerCmd)
	rootCmd.AddCommand(configCmd)

}
