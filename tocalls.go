package hamaprs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// map for exact matches callsign
var toCalls map[string]Device

// trieRoot the root of the tries for wildcard lookup
var trieRoot *Trie

// A Trie struct to handle the tocalls lookups
type Trie struct {
	root *TrieNode
}

// A Node struct to store the tocalls data
type TrieNode struct {
	value    rune
	children []*TrieNode
	Device   *Device
}

// ToCallsJSON struct to handle the tocalls file
type ToCallsJSON struct {
	Tocalls map[string]Device `json:"tocalls"`
}

// MiceJSON struct to handle the tocalls file
type MiceJSON struct {
	Mice map[string]Device `json:"mice"`
}

// Device stores the model information for a radio device
type Device struct {
	Model  string `json:"model"`
	Vendor string `json:"vendor"`
	Class  string `json:"class"`
	Prefix string `json:"prefix"`
}

// init load the tocalls json files and fill the data structure for lookups
func init() {
	// loading tocalls file
	// https://raw.githubusercontent.com/hessu/aprs-deviceid/master/generated/tocalls.pretty.json
	file, e := ioutil.ReadFile("tocalls.pretty.json")
	if e != nil {
		fmt.Printf("File error %v\n", e)
		os.Exit(1)
	}

	// initializing the exact match map
	toCalls = make(map[string]Device)

	// initiliazing the trie
	trieRoot = &Trie{root: &TrieNode{children: make([]*TrieNode, 0)}}

	var toCallsJSON ToCallsJSON
	json.Unmarshal(file, &toCallsJSON)

	// iterate over callsigns
	for c, d := range toCallsJSON.Tocalls {
		//  1st insert none wildcard entries
		if !strings.ContainsAny(c, "*n?") {
			// do not insert unknown unknown
			if c == "APRS" {
				continue
			}
			// insert the exact match
			toCalls[c] = d
		} else {
			trieRoot.Add(c, d)
		}
	}

	var miceJSON MiceJSON
	json.Unmarshal(file, &miceJSON)

	// iterate over Mice in the same map toCalls
	for c, d := range miceJSON.Mice {
		// insert the exact match
		toCalls[c] = d
	}

}

// Add a string c to the trie root
func (t *Trie) Add(c string, d Device) {
	cnode := t.root
	for pos, v := range c {
		var n *TrieNode
		for _, c := range cnode.children {
			if c.value == v {
				n = c
			}
		}

		if n == nil {
			// add it to the current node's children
			n = &TrieNode{value: v, children: make([]*TrieNode, 0)}
			switch n.value {
			case '*', '?', 'n':
				// add to the end
				cnode.children = append(cnode.children, n)
			default:
				// insert at pos 0
				cnode.children = append(cnode.children, nil)
				copy(cnode.children[0+1:], cnode.children[0:])
				cnode.children[0] = n
			}

			last := len(cnode.children) - 1

			if last > 0 {
				// ensure the wildcard is always the latest element in the children list
				if cnode.children[last-1].value == '*' {
					cnode.children[last-1], cnode.children[last] = cnode.children[last], cnode.children[last-1]
				}
			}
		}

		cnode = n

		if pos == len(c)-1 {
			cnode.Device = &d
		}
	}
}

// match a string against the devices tries
func (t *Trie) match(c string) *Device {
	cnode := t.root
	for _, v := range c {
		n := cnode.childWithRune(v)
		if n == nil {
			return nil
		}
		if n.value == '*' {
			return n.Device
		}
		cnode = n
	}
	return cnode.Device
}

// return the children node named r
func (n *TrieNode) childWithRune(r rune) *TrieNode {
	for _, c := range n.children {
		if c.value == '*' || c.value == '?' || (c.value == 'n' && isNumeric(r) || c.value == r) {
			return c
		}
	}
	return nil
}

// String display a Node, mainly for debug purpose
func (n *TrieNode) String() string {
	var l string
	for _, c := range n.children {
		l += string(c.value)
	}
	var device string
	if n.Device != nil {
		device = n.Device.Model
	}
	return fmt.Sprintf("(%s) c[%s] s<%s>", string(n.value), l, device)
}

// return true if the rune is between 0 - 9
func isNumeric(r rune) bool {
	// convert a rune to it's numerica value
	if int(r-'0') >= 0 && int(r-'0') <= 9 {
		return true
	}
	return false
}
