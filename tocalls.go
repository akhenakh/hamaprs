package hamaprs

import (
	"encoding/json"
	"fmt"
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

// source 	https://raw.githubusercontent.com/hessu/aprs-deviceid/master/generated/tocalls.pretty.json
var ToCallsJSONData []byte = []byte(`{"mice":{"*v":{"model":"Tracker","class":"tracker","vendor":"KissOZ"},"|3":{"model":"TinyTrak3","class":"tracker","vendor":"Byonics"},"_\"":{"model":"FTM-350","class":"rig","vendor":"Yaesu"},"_$":{"model":"FT1D","class":"ht","vendor":"Yaesu"},"_ ":{"model":"VX-8","class":"ht","vendor":"Yaesu"},"_#":{"model":"VX-8G","class":"ht","vendor":"Yaesu"},"^v":{"model":"anyfrog","vendor":"HinzTec"},"_%":{"model":"FTM-400DR","class":"rig","vendor":"Yaesu"},"|4":{"model":"TinyTrak4","class":"tracker","vendor":"Byonics"}},"classes":{"app":{"shown":"Mobile app","description":"Mobile phone or tablet app"},"rig":{"shown":"Rig","description":"Mobile or desktop radio"},"tracker":{"shown":"Tracker","description":"Tracker device"},"wx":{"shown":"Weather station","description":"Dedicated weather station"},"ht":{"shown":"HT","description":"Hand-held radio"},"software":{"shown":"Software","description":"Desktop software"},"digi":{"shown":"Digipeater","description":"Digipeater firmware"},"dstar":{"shown":"D-Star","description":"D-Star radio"},"satellite":{"shown":"Satellite","description":"Satellite-based station"}},"micelegacy":{"]=":{"features":["messaging"],"suffix":"=","model":"TM-D710","class":"rig","vendor":"Kenwood","prefix":"]"},">=":{"features":["messaging"],"suffix":"=","model":"TH-D72","class":"ht","vendor":"Kenwood","prefix":">"},"]":{"features":["messaging"],"model":"TM-D700","class":"rig","vendor":"Kenwood","prefix":"]"},">":{"features":["messaging"],"model":"TH-D7A","class":"ht","vendor":"Kenwood","prefix":">"}},"tocalls":{"APJA??":{"model":"JavAPRS","vendor":"K4HG & AE5PL"},"APSTM?":{"model":"Balloon tracker","class":"tracker","vendor":"W7QO"},"APX???":{"model":"Xastir","class":"software","os":"Linux/Unix","vendor":"Open Source"},"APMI05":{"model":"PLXTracker","os":"embedded","vendor":"Microsat"},"APRS":{"model":"Unknown","vendor":"Unknown"},"APBPQ?":{"model":"BPQ32","class":"software","os":"Windows","vendor":"John Wiseman, G8BPQ"},"APOA??":{"model":"app","class":"app","os":"ios","vendor":"OpenAPRS"},"APWA??":{"model":"APRSISCE","class":"software","os":"Android","vendor":"KJ4ERJ"},"APERXQ":{"model":"PE1RXQ APRS Tracker","class":"tracker","vendor":"PE1RXQ"},"APWM??":{"features":["messaging","item-in-msg"],"model":"APRSISCE","class":"software","os":"Windows Mobile","vendor":"KJ4ERJ"},"APMI03":{"model":"PLXDigi","os":"embedded","vendor":"Microsat"},"APSAR":{"model":"SARTrack","class":"software","os":"Windows","vendor":"ZL4FOX"},"APDPRS":{"model":"D-Star APDPRS","class":"dstar","vendor":"unknown"},"APSMS?":{"model":"SMS gateway","class":"software","vendor":"Paul Defrusne"},"APOT??":{"model":"OpenTracker","class":"tracker","vendor":"Argent Data Systems"},"APRNOW":{"model":"APRSNow","class":"app","os":"ipad","vendor":"Gregg Wonderly, W5GGW"},"APRHH?":{"model":"HamHud","class":"tracker","vendor":"Steven D. Bragg, KA9MVA"},"APCLEY":{"model":"EYTraker","class":"tracker","vendor":"ZS6EY"},"APNKMP":{"model":"KAM+","vendor":"Kantronics"},"APJID2":{"model":"D-Star APJID2","class":"dstar","vendor":"Peter Loveall, AE5PL"},"APVE??":{"model":"EchoLink","vendor":"unknown"},"APNW??":{"model":"WX3in1","os":"embedded","vendor":"SQ3FYK"},"APAM??":{"model":"AltOS","class":"tracker","vendor":"Altus Metrum"},"APAH??":{"model":"AHub"},"APUDR?":{"model":"UDR","vendor":"NW Digital Radio"},"APMI04":{"model":"WX3in1 Mini","os":"embedded","vendor":"Microsat"},"APY01D":{"model":"FT1D","class":"ht","vendor":"Yaesu"},"APAND?":{"model":"APRSdroid","class":"app","os":"Android","vendor":"Open Source"},"APHAX?":{"model":"SM2APRS SondeMonitor","class":"software","os":"Windows","vendor":"PY2UEP"},"APLM??":{"class":"software","vendor":"WA0TQG"},"APXR??":{"model":"Xrouter","vendor":"G8PZT"},"APECAN":{"model":"Pecan Pico APRS Balloon Tracker","class":"tracker","vendor":"KT5TK/DL7AD"},"APS???":{"model":"APRS+SA","class":"software","vendor":"Brent Hildebrand, KH2Z"},"APBL??":{"model":"BeeLine GPS","class":"tracker","vendor":"BigRedBee"},"APK0??":{"model":"TH-D7","class":"ht","vendor":"Kenwood"},"APT4??":{"model":"TinyTrak4","class":"tracker","vendor":"Byonics"},"APDT??":{"model":"APRStouch Tone (DTMF)","vendor":"unknown"},"APZTKP":{"model":"TrackPoint","class":"tracker","os":"embedded","vendor":"Nick Hanks, N0LP"},"APRG??":{"model":"aprsg","class":"software","os":"Linux/Unix","vendor":"OH2GVE"},"APCLWX":{"model":"EYWeather","class":"wx","vendor":"ZS6EY"},"APJI??":{"model":"jAPRSIgate","class":"software","vendor":"Peter Loveall, AE5PL"},"APNM??":{"model":"TNC","vendor":"MFJ"},"APN9??":{"model":"KPC-9612","vendor":"Kantronics"},"APNU??":{"model":"UIdigi","class":"digi","vendor":"IW3FQG"},"APN3??":{"model":"KPC-3","vendor":"Kantronics"},"APIC??":{"model":"PICiGATE","vendor":"HA9MCQ"},"APK1??":{"model":"TM-D700","class":"rig","vendor":"Kenwood"},"APnnnD":{"model":"uSmartDigi D-Gate","class":"dstar","vendor":"Painter Engineering"},"APSK63":{"model":"APRS Messenger","class":"software","os":"Windows","vendor":"Chris Moulding, G4HYG"},"APOLU?":{"model":"Oscar","class":"satellite","vendor":"AMSAT-LU"},"APKRAM":{"model":"Ham Tracker","class":"app","os":"ios","vendor":"kramstuff.com"},"APAW??":{"model":"AGWPE","class":"software","os":"Windows","vendor":"SV2AGW"},"APJY??":{"model":"YAAC","class":"software","vendor":"KA2DDO"},"APTW??":{"model":"WXTrak","class":"wx","vendor":"Byonics"},"APR8??":{"model":"APRSdos","class":"software","vendor":"Bob Bruninga, WB4APR"},"APC???":{"model":"APRS/CE","class":"app","vendor":"Rob Wittner, KZ5RW"},"APT3??":{"model":"TinyTrak3","class":"tracker","vendor":"Byonics"},"APTT*":{"model":"TinyTrak","class":"tracker","vendor":"Byonics"},"PSKAPR":{"model":"PSKmail","class":"software","vendor":"Open Source"},"APFG??":{"model":"Flood Gage","class":"software","vendor":"KP4DJT"},"APN102":{"model":"APRSNow","class":"app","os":"ipad","vendor":"Gregg Wonderly, W5GGW"},"APDS??":{"model":"dsDIGI","os":"embedded","vendor":"SP9UOB"},"APNX??":{"model":"TNC-X","vendor":"K6DBG"},"APAF??":{"model":"AFilter"},"APDST?":{"model":"dsTracker","os":"embedded","vendor":"SP9UOB"},"APJE??":{"model":"JeAPRS","vendor":"Gregg Wonderly, W5GGW"},"APRX??":{"model":"aprx","class":"software","vendor":"OH2MQK"},"APNP??":{"model":"TNC","vendor":"PacComm"},"APJS??":{"model":"javAPRSSrvr","vendor":"Peter Loveall, AE5PL"},"APDnnn":{"model":"aprsd","class":"software","os":"Linux/Unix","vendor":"Open Source"},"APGO??":{"model":"APRS-Go","class":"app","vendor":"AA3NJ"},"APZ18":{"model":"UIdigi","class":"digi","vendor":"IW3FQG"},"API???":{"model":"unknown","class":"dstar","vendor":"Icom"},"APAGW?":{"model":"AGWtracker","class":"software","os":"Windows","vendor":"SV2AGW"},"APZG??":{"model":"aprsg","class":"software","os":"Linux/Unix","vendor":"OH2GVE"},"AP1WWX":{"model":"T-238+","class":"wx","vendor":"TAPR"},"APnnnU":{"model":"uSmartDigi Digipeater","class":"digi","vendor":"Painter Engineering"},"APVR??":{"model":"IRLP","vendor":"unknown"},"APZMDR":{"model":"HaMDR","class":"tracker","os":"embedded","vendor":"Open Source"},"APOZ??":{"model":"KissOZ","class":"tracker","vendor":"OZ1EKD, OZ7HVO"},"APDR??":{"model":"APRSdroid","class":"app","os":"Android","vendor":"Open Source"},"APMI02":{"model":"WXEth","os":"embedded","vendor":"Microsat"},"APU2*":{"model":"UI-View32","class":"software","os":"Windows","vendor":"Roger Barker, G4IDE"},"APCLEZ":{"model":"Telit EZ10 GSM application","class":"tracker","vendor":"ZS6EY"},"APNT??":{"model":"TNT TNC as a digipeater","class":"digi","vendor":"SV2AGW"},"APNK80":{"model":"KAM","vendor":"Kantronics"},"APND??":{"model":"DIGI_NED","vendor":"PE1MEW"},"APNK01":{"features":["messaging"],"model":"TM-D700","class":"rig","vendor":"Kenwood"},"APZ186":{"model":"UIdigi","class":"digi","vendor":"IW3FQG"},"APMI01":{"model":"WX3in1","os":"embedded","vendor":"Microsat"},"APDI??":{"model":"DIXPRS","class":"software","vendor":"Bela, HA5DI"},"APDF??":{"model":"Automatic DF units"},"APCWP8":{"model":"WinphoneAPRS","class":"app","vendor":"GM7HHB"},"APDU??":{"model":"U2APRS","class":"app","os":"Android","vendor":"JA7UDE"},"APHH?":{"model":"HamHud","class":"tracker","vendor":"Steven D. Bragg, KA9MVA"},"APDG??":{"model":"ircDDB Gateway","class":"dstar","vendor":"Jonathan, G4KLX"},"APWW??":{"features":["messaging","item-in-msg"],"model":"APRSIS32","class":"software","os":"Windows","vendor":"KJ4ERJ"},"APAX??":{"model":"AFilterX"},"APTR??":{"model":"MotoTRBO","vendor":"Motorola"},"APSC??":{"model":"aprsc","class":"software","vendor":"OH2MQK, OH7LZB"},"APRRT?":{"model":"RTrak","class":"tracker","vendor":"RPC Electronics"},"APAG??":{"model":"AGate"},"APFI??":{"class":"app","vendor":"aprs.fi"},"APDW??":{"model":"DireWolf","vendor":"WB2OSZ"},"APE???":{"model":"Telemetry devices"},"APFII?":{"class":"app","os":"ios","vendor":"aprs.fi"},"APMI??":{"os":"embedded","vendor":"Microsat"},"APAGW":{"model":"AGWtracker","class":"software","os":"Windows","vendor":"SV2AGW"},"APMG??":{"model":"MiniGate","class":"software","os":"Netduino","vendor":"Alex, AB0TJ"},"AP4R??":{"model":"APRS4R","class":"software","vendor":"Open Source"},"APCL??":{"model":"maprs","class":"app","vendor":"maprs.org"},"APZ19":{"model":"UIdigi","class":"digi","vendor":"IW3FQG"},"APZWKR":{"model":"NetSked","class":"software","vendor":"GM1WKR"},"APK003":{"model":"TH-D72","class":"ht","vendor":"Kenwood"},"APW???":{"model":"WinAPRS","class":"software","os":"Windows","vendor":"Sproul Brothers"},"APAVT5":{"model":"AP510","class":"tracker","vendor":"SainSonic"},"APT2??":{"model":"TinyTrak2","class":"tracker","vendor":"Byonics"},"APU1??":{"model":"UI-View16","class":"software","os":"Windows","vendor":"Roger Barker, G4IDE"},"APPT??":{"model":"KetaiTracker","class":"tracker","vendor":"JF6LZE"},"APHK??":{"model":"Digipeater/tracker","vendor":"LA1BR"}}}`)

// init load the tocalls json files and fill the data structure for lookups
func init() {
	// loading tocalls file

	// initializing the exact match map
	toCalls = make(map[string]Device)

	// initiliazing the trie
	trieRoot = &Trie{root: &TrieNode{children: make([]*TrieNode, 0)}}

	var toCallsJSON ToCallsJSON
	json.Unmarshal(ToCallsJSONData, &toCallsJSON)

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
	json.Unmarshal(ToCallsJSONData, &miceJSON)

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
