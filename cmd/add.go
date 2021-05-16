/*
Copyright © 2021 Nagy Károly Gábriel <k@jpi.io>

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

	"github.com/karasz/gomesh/peers"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:        "add",
	Aliases:    []string{},
	SuggestFor: []string{},
	Short:      "Add or update a node in the registry",
	Long:       `Will add/update a node in the registry`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		privatekey, _ := cmd.Flags().GetString("privatekey")
		address, _ := cmd.Flags().GetStringSlice("address")
		listenport, _ := cmd.Flags().GetInt("listenport")
		endpoint, _ := cmd.Flags().GetString("endpoint")
		allowedips, _ := cmd.Flags().GetStringSlice("allowedips")
		fwmark, _ := cmd.Flags().GetInt("fwmark")
		dns, _ := cmd.Flags().GetString("dns")
		mtu, _ := cmd.Flags().GetInt("mtu")
		table, _ := cmd.Flags().GetString("table")
		preup, _ := cmd.Flags().GetString("preup")
		predown, _ := cmd.Flags().GetString("predown")
		postup, _ := cmd.Flags().GetString("postup")
		postdown, _ := cmd.Flags().GetString("postdown")
		saveconfig, _ := cmd.Flags().GetBool("saveconfig")
		p := peers.Peer{Name: name, PrivateKey: privatekey, Address: address, ListenPort: listenport, Endpoint: endpoint, AllowedIPs: allowedips, FwMark: fwmark, DNS: dns, MTU: mtu, Table: table, PreUp: preup, PostUp: postup, PreDown: predown, PostDown: postdown, SaveConfig: saveconfig}
		err := thePeers.AddPeer(p)
		return err
	},
}

func init() {
	var err error
	addCmd.Flags().BoolP("update", "u", false, "Update Peer if existing.")
	addCmd.Flags().StringP("name", "n", "", "Name of the node. (Required)")
	addCmd.Flags().StringSliceP("address", "a", []string{}, "Address of the node. (Required)")
	addCmd.Flags().StringP("endpoint", "e", "", "The node's endpoint")
	addCmd.Flags().StringSliceP("allowedips", "", []string{}, "Additional allowed IP addresses")
	addCmd.Flags().StringP("privatekey", "p", "", "Private key of server interface (if none given one will be generated")
	addCmd.Flags().IntP("listenport", "l", 51820, "Port to listen on, default 51820")
	addCmd.Flags().IntP("fwmark", "f", 0, "Mark the outgoing packets with")
	addCmd.Flags().StringP("dns", "", "", "DNS server")
	addCmd.Flags().IntP("mtu", "m", 0, "Node interface MTU")
	addCmd.Flags().StringP("routing_table", "r", "", "Node routing table")
	addCmd.Flags().StringP("preUP", "", "", "Command to run before bringing the interface UP")
	addCmd.Flags().StringP("postUP", "", "", "Command to run after bringing the interface UP")
	addCmd.Flags().StringP("preDown", "", "", "Command to run before bringing the interface DOWN")
	addCmd.Flags().StringP("postDown", "", "", "Command to run after bringing the interface DOWN")
	addCmd.Flags().BoolP("saveconfig", "s", false, "Save config between reboots")
	err = addCmd.MarkFlagRequired("name")
	if err != nil {
		fmt.Println(err)
	}
	err = addCmd.MarkFlagRequired("address")
	if err != nil {
		fmt.Println(err)
	}
	err = addCmd.MarkFlagRequired("endpoint")
	if err != nil {
		fmt.Println(err)
	}

	rootCmd.AddCommand(addCmd)
}
