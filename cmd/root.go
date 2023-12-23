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

var cookieCmd = &cobra.Command{
	Use:   "cookie",
	Short: "Set cookie if not already set!",
	Run: func(cmd *cobra.Command, args []string) {
		ytm.GetCookie()
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

	rootCmd.AddCommand(cookieCmd)
	rootCmd.AddCommand(configCmd)

}
