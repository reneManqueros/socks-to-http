package cmd

import (
	"github.com/spf13/cobra"
	"math/rand"
	"sockstohttp/internal"
	"strconv"
	"strings"
	"time"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run socks to http proxy",
	Run: func(cmd *cobra.Command, args []string) {

		rand.Seed(time.Now().UnixMilli())

		// Defaults
		socksAddress := ":8081"
		listenAddress := ":8083"
		timeout := 10
		// Overrides
		for _, v := range args {
			argumentParts := strings.Split(v, "=")
			if len(argumentParts) == 2 {
				if argumentParts[0] == "socks" {
					socksAddress = argumentParts[1]
				}
				if argumentParts[0] == "listen" {
					listenAddress = argumentParts[1]
				}
				if argumentParts[0] == "timeout" {
					timeout, _ = strconv.Atoi(argumentParts[1])
				}
			}
		}

		internal.Server{
			ListenAddress: listenAddress,
			SOCKSAddress:  socksAddress,
			Timeout:       timeout,
		}.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags()
}
