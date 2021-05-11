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
package peers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/karasz/gomesh/keys"
)

// Peers is a map containing all
// registered peers
type Peers []Peer

//Peer is a Wireguard Peer
type Peer struct {
	Name       string
	PrivateKey string
	Address    []string
	ListenPort int
	Endpoint   string
	AllowedIPs []string
	FwMark     int
	DNS        string
	MTU        int
	Table      string
	PreUp      string
	PostUp     string
	PreDown    string
	PostDown   string
	SaveConfig bool
}

// LoadPeers will load the register with Peers
// from the specified JSON file
func LoadPeers(peersPath string) (Peers, error) {

	var p Peers

	peersFile, err := os.Open(peersPath)
	if err != nil {
		return nil, err
	}

	defer peersFile.Close()

	jsonParser := json.NewDecoder(peersFile)
	err = jsonParser.Decode(&p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p Peers) peerExists(pr Peer) bool {
	result := false

	for _, k := range p {
		if strings.EqualFold(pr.Name, k.Name) {
			result = true
		}
	}

	return result

}

//AddPeer will add a Peer to the register
func (p Peers) AddPeer(pr Peer) error {
	if pr.PrivateKey == "" {
		k, err := keys.GenerateKey()
		if err != nil {
			return err
		}
		pr.PrivateKey = k
	}

	if pr.ListenPort == 0 {
		pr.ListenPort = 51820
	}
	p = append(p, pr)
	err := p.DumpPeers("", true)
	fmt.Println(err)
	return err
}

//DeletePeer will delete the named Peer from the register
func (p Peers) DeletePeer(pr string) {
	index := 0
	for i := range p {
		if p[i].Name == pr {
			index = i
		}
	}

	p = append(p[:index], p[index+1:]...)
	p.DumpPeers("", true)
}

//GenerateConfigs will generate the Wireguard mesh
//configs in the specified folder
func (p Peers) GenerateConfigs(folder string) {
	fmt.Println("not yet implemented")
}

// DumpPeers will generate a JSON file at the provided location with the
// registered Peers
func (p Peers) DumpPeers(peersPath string, overwrite bool) error {
	if peersPath == "" {
		peersPath = "database.json"
	}
	// we chose to represent that a file exists with value 2 (linux read ACL)
	// and that we want to overwrite with value 4 (linux write ACL) so that
	// we can check all the possibilities with one if statement
	var err error
	var exists int = 2
	var overwritebits = 0

	if overwrite {
		overwritebits = 4
	}

	peersJson, _ := json.MarshalIndent(p, "", "    ")

	if _, err = os.Stat(peersPath); os.IsNotExist(err) {
		exists = 0
		err = os.MkdirAll(filepath.Dir(peersPath), 0755)
		if err != nil {
			return err
		}
	}

	if overwritebits+exists != 2 {
		err = os.WriteFile(peersPath, peersJson, 0644)
		if err != nil {
			return err
		}
	} else {
		err = fmt.Errorf("Peers database exists and I am not allowed to overwrite")
	}
	return err
}

// PrettyPrint will print a table with the peers
func (p Peers) PrettyPrint(brief bool) {
	const padding = 3
	fmt.Println(brief)
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.AlignRight|tabwriter.Debug)

	if !brief {
		fmt.Fprintln(tw, "NAME\t", "PRIVATEKEY\t", "ADDRESS\t", "LISTENPORT\t", "ENDPOINT\t", "ALLOWEDIPS\t", "FWMARK\t", "DNS\t", "MTU\t", "TABLE\t", "PREUP\t", "POSTUP\t", "PREDOWN\t", "POSTDOWN\t", "SAVE\t")
		for _, v := range p {
			fmt.Fprintln(tw, v.Name+"\t", v.PrivateKey+"\t", strings.Join(v.Address, ",")+"\t", strconv.Itoa(v.ListenPort)+"\t", v.Endpoint+"\t", strings.Join(v.AllowedIPs, ",")+"\t", strconv.Itoa(v.FwMark)+"\t", v.DNS+"\t", strconv.Itoa(v.MTU)+"\t", v.Table+"\t", v.PreUp+"\t", v.PostUp+"\t", v.PreDown+"\t", v.PostDown+"\t", strconv.FormatBool(v.SaveConfig)+"\t")

		}
	} else {
		fmt.Fprintln(tw, "NAME\t", "PRIVATEKEY\t", "ADDRESS\t", "LISTENPORT\t", "ENDPOINT\t")
		for _, v := range p {
			fmt.Fprintln(tw, v.Name+"\t", v.PrivateKey+"\t", "["+strings.Join(v.Address, ",")+"]\t", strconv.Itoa(v.ListenPort)+"\t", v.Endpoint+"\t")
		}
	}

	tw.Flush()

}
