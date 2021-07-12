/*
Copyright © 2021 JPI Technologies Ltd <oss@jpi.io>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/karasz/gomesh/wireguard"
)

var thePeers wireguard.Peers

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gomesh",
	Short:   "Generate Wireguard Mesh VPN configurations",
	Long:    "This little tool will generate and manage configuration files for Wireguard Mesh VPNs.",
	Version: "0.1.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(0)
	}
}

func init() {
	var dbFile string
	rootCmd.PersistentFlags().StringVarP(&dbFile, "database", "d", "database.json", "registry file")
	initDatabase(dbFile)
}

func initDatabase(filePath string) {
	var err error
	thePeers, err = wireguard.LoadPeers(filePath)

	if err != nil {
		fmt.Println(err)
	}
}
