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

// A TrieNode struct to store the tocalls data
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
var ToCallsJSONData = []byte(`{"classes":{"app":{"shown":"Mobile app","description":"Mobile phone or tablet app"},"software":{"description":"Desktop software","shown":"Software"},"tracker":{"description":"Tracker device","shown":"Tracker"},"wx":{"shown":"Weather station","description":"Dedicated weather station"},"ht":{"description":"Hand-held radio","shown":"HT"},"satellite":{"shown":"Satellite","description":"Satellite-based station"},"dstar":{"shown":"D-Star","description":"D-Star radio"},"rig":{"description":"Mobile or desktop radio","shown":"Rig"},"digi":{"shown":"Digipeater","description":"Digipeater firmware"}},"tocalls":{"APJS??":{"model":"javAPRSSrvr","vendor":"Peter Loveall, AE5PL"},"APSC??":{"class":"software","vendor":"OH2MQK, OH7LZB","model":"aprsc"},"APK0??":{"model":"TH-D7","vendor":"Kenwood","class":"ht"},"APPT??":{"class":"tracker","vendor":"JF6LZE","model":"KetaiTracker"},"APDF??":{"model":"Automatic DF units"},"APSK63":{"class":"software","os":"Windows","vendor":"Chris Moulding, G4HYG","model":"APRS Messenger"},"APW???":{"class":"software","os":"Windows","model":"WinAPRS","vendor":"Sproul Brothers"},"APDR??":{"vendor":"Open Source","model":"APRSdroid","os":"Android","class":"app"},"APUDR?":{"model":"UDR","vendor":"NW Digital Radio"},"APAGW":{"model":"AGWtracker","vendor":"SV2AGW","os":"Windows","class":"software"},"APDW??":{"vendor":"WB2OSZ","model":"DireWolf"},"APAVT5":{"class":"tracker","model":"AP510","vendor":"SainSonic"},"APCLEY":{"class":"tracker","vendor":"ZS6EY","model":"EYTraker"},"APAX??":{"model":"AFilterX"},"AP1WWX":{"class":"wx","model":"T-238+","vendor":"TAPR"},"APXR??":{"model":"Xrouter","vendor":"G8PZT"},"APCLWX":{"class":"wx","model":"EYWeather","vendor":"ZS6EY"},"APCL??":{"class":"app","model":"maprs","vendor":"maprs.org"},"APAND?":{"os":"Android","class":"app","model":"APRSdroid","vendor":"Open Source"},"APNKMP":{"vendor":"Kantronics","model":"KAM+"},"API???":{"vendor":"Icom","model":"unknown","class":"dstar"},"APMI05":{"model":"PLXTracker","vendor":"Microsat","os":"embedded"},"APU1??":{"vendor":"Roger Barker, G4IDE","model":"UI-View16","class":"software","os":"Windows"},"APNP??":{"model":"TNC","vendor":"PacComm"},"APDPRS":{"class":"dstar","model":"D-Star APDPRS","vendor":"unknown"},"APK1??":{"class":"rig","model":"TM-D700","vendor":"Kenwood"},"APX???":{"class":"software","os":"Linux/Unix","model":"Xastir","vendor":"Open Source"},"AP4R??":{"class":"software","vendor":"Open Source","model":"APRS4R"},"APIC??":{"model":"PICiGATE","vendor":"HA9MCQ"},"APN9??":{"vendor":"Kantronics","model":"KPC-9612"},"APWW??":{"class":"software","os":"Windows","features":["messaging","item-in-msg"],"model":"APRSIS32","vendor":"KJ4ERJ"},"APN3??":{"model":"KPC-3","vendor":"Kantronics"},"APMG??":{"os":"Netduino","class":"software","model":"MiniGate","vendor":"Alex, AB0TJ"},"APDST?":{"model":"dsTracker","vendor":"SP9UOB","os":"embedded"},"APZ19":{"class":"digi","model":"UIdigi","vendor":"IW3FQG"},"APnnnU":{"vendor":"Painter Engineering","model":"uSmartDigi Digipeater","class":"digi"},"APNT??":{"vendor":"SV2AGW","model":"TNT TNC as a digipeater","class":"digi"},"APGO??":{"class":"app","model":"APRS-Go","vendor":"AA3NJ"},"APAH??":{"model":"AHub"},"APNK01":{"vendor":"Kenwood","model":"TM-D700","features":["messaging"],"class":"rig"},"APT2??":{"class":"tracker","model":"TinyTrak2","vendor":"Byonics"},"APNX??":{"model":"TNC-X","vendor":"K6DBG"},"APTT*":{"class":"tracker","vendor":"Byonics","model":"TinyTrak"},"APMI04":{"model":"WX3in1 Mini","vendor":"Microsat","os":"embedded"},"APT3??":{"model":"TinyTrak3","vendor":"Byonics","class":"tracker"},"APOT??":{"model":"OpenTracker","vendor":"Argent Data Systems","class":"tracker"},"APU2*":{"class":"software","os":"Windows","vendor":"Roger Barker, G4IDE","model":"UI-View32"},"APDI??":{"class":"software","vendor":"Bela, HA5DI","model":"DIXPRS"},"APOA??":{"os":"ios","class":"app","model":"app","vendor":"OpenAPRS"},"APRRT?":{"class":"tracker","vendor":"RPC Electronics","model":"RTrak"},"PSKAPR":{"model":"PSKmail","vendor":"Open Source","class":"software"},"APZWKR":{"class":"software","vendor":"GM1WKR","model":"NetSked"},"APTR??":{"model":"MotoTRBO","vendor":"Motorola"},"APAGW?":{"model":"AGWtracker","vendor":"SV2AGW","os":"Windows","class":"software"},"APTW??":{"model":"WXTrak","vendor":"Byonics","class":"wx"},"APY01D":{"class":"ht","vendor":"Yaesu","model":"FT1D"},"APZG??":{"vendor":"OH2GVE","model":"aprsg","os":"Linux/Unix","class":"software"},"APAW??":{"vendor":"SV2AGW","model":"AGWPE","os":"Windows","class":"software"},"APJA??":{"vendor":"K4HG & AE5PL","model":"JavAPRS"},"APHK??":{"model":"Digipeater/tracker","vendor":"LA1BR"},"APRHH?":{"class":"tracker","model":"HamHud","vendor":"Steven D. Bragg, KA9MVA"},"APE???":{"model":"Telemetry devices"},"APZMDR":{"vendor":"Open Source","model":"HaMDR","os":"embedded","class":"tracker"},"APZ18":{"vendor":"IW3FQG","model":"UIdigi","class":"digi"},"APFI??":{"vendor":"aprs.fi","class":"app"},"APERXQ":{"class":"tracker","vendor":"PE1RXQ","model":"PE1RXQ APRS Tracker"},"APMI02":{"os":"embedded","model":"WXEth","vendor":"Microsat"},"APT4??":{"class":"tracker","vendor":"Byonics","model":"TinyTrak4"},"APAG??":{"model":"AGate"},"APECAN":{"model":"Pecan Pico APRS Balloon Tracker","vendor":"KT5TK/DL7AD","class":"tracker"},"APFII?":{"class":"app","os":"ios","vendor":"aprs.fi"},"APSMS?":{"model":"SMS gateway","vendor":"Paul Defrusne","class":"software"},"APDT??":{"model":"APRStouch Tone (DTMF)","vendor":"unknown"},"APCLEZ":{"class":"tracker","vendor":"ZS6EY","model":"Telit EZ10 GSM application"},"APR8??":{"class":"software","model":"APRSdos","vendor":"Bob Bruninga, WB4APR"},"APDU??":{"class":"app","os":"Android","model":"U2APRS","vendor":"JA7UDE"},"APRS":{"model":"Unknown","vendor":"Unknown"},"APFG??":{"model":"Flood Gage","vendor":"KP4DJT","class":"software"},"APNU??":{"class":"digi","model":"UIdigi","vendor":"IW3FQG"},"APAF??":{"model":"AFilter"},"APNM??":{"model":"TNC","vendor":"MFJ"},"APOLU?":{"vendor":"AMSAT-LU","model":"Oscar","class":"satellite"},"APVE??":{"vendor":"unknown","model":"EchoLink"},"APNW??":{"model":"WX3in1","vendor":"SQ3FYK","os":"embedded"},"APDnnn":{"model":"aprsd","vendor":"Open Source","class":"software","os":"Linux/Unix"},"APY400":{"class":"ht","model":"FTM-400","vendor":"Yaesu"},"APKRAM":{"vendor":"kramstuff.com","model":"Ham Tracker","class":"app","os":"ios"},"APHH?":{"vendor":"Steven D. Bragg, KA9MVA","model":"HamHud","class":"tracker"},"APBPQ?":{"vendor":"John Wiseman, G8BPQ","model":"BPQ32","os":"Windows","class":"software"},"APOZ??":{"class":"tracker","vendor":"OZ1EKD, OZ7HVO","model":"KissOZ"},"APNK80":{"vendor":"Kantronics","model":"KAM"},"APK003":{"model":"TH-D72","vendor":"Kenwood","class":"ht"},"APBL??":{"vendor":"BigRedBee","model":"BeeLine GPS","class":"tracker"},"APAM??":{"vendor":"Altus Metrum","model":"AltOS","class":"tracker"},"APS???":{"class":"software","model":"APRS+SA","vendor":"Brent Hildebrand, KH2Z"},"APRX??":{"class":"software","vendor":"OH2MQK","model":"aprx"},"APRNOW":{"vendor":"Gregg Wonderly, W5GGW","model":"APRSNow","class":"app","os":"ipad"},"APN102":{"os":"ipad","class":"app","model":"APRSNow","vendor":"Gregg Wonderly, W5GGW"},"APJID2":{"class":"dstar","vendor":"Peter Loveall, AE5PL","model":"D-Star APJID2"},"APHAX?":{"class":"software","os":"Windows","vendor":"PY2UEP","model":"SM2APRS SondeMonitor"},"APZTKP":{"os":"embedded","class":"tracker","model":"TrackPoint","vendor":"Nick Hanks, N0LP"},"APMI03":{"model":"PLXDigi","vendor":"Microsat","os":"embedded"},"APZ186":{"vendor":"IW3FQG","model":"UIdigi","class":"digi"},"APSAR":{"model":"SARTrack","vendor":"ZL4FOX","os":"Windows","class":"software"},"APSTM?":{"class":"tracker","vendor":"W7QO","model":"Balloon tracker"},"APC???":{"class":"app","vendor":"Rob Wittner, KZ5RW","model":"APRS/CE"},"APY02D":{"class":"ht","model":"FT2D","vendor":"Yaesu"},"APWM??":{"os":"Windows Mobile","class":"software","features":["messaging","item-in-msg"],"model":"APRSISCE","vendor":"KJ4ERJ"},"APMI01":{"os":"embedded","model":"WX3in1","vendor":"Microsat"},"APDG??":{"class":"dstar","model":"ircDDB Gateway","vendor":"Jonathan, G4KLX"},"APnnnD":{"model":"uSmartDigi D-Gate","vendor":"Painter Engineering","class":"dstar"},"APND??":{"vendor":"PE1MEW","model":"DIGI_NED"},"APJY??":{"class":"software","vendor":"KA2DDO","model":"YAAC"},"APJE??":{"vendor":"Gregg Wonderly, W5GGW","model":"JeAPRS"},"APWA??":{"class":"software","os":"Android","vendor":"KJ4ERJ","model":"APRSISCE"},"APRG??":{"class":"software","os":"Linux/Unix","model":"aprsg","vendor":"OH2GVE"},"APDS??":{"model":"dsDIGI","vendor":"SP9UOB","os":"embedded"},"APJI??":{"model":"jAPRSIgate","vendor":"Peter Loveall, AE5PL","class":"software"},"APVR??":{"vendor":"unknown","model":"IRLP"},"APCWP8":{"model":"WinphoneAPRS","vendor":"GM7HHB","class":"app"},"APLM??":{"vendor":"WA0TQG","class":"software"},"APMI??":{"vendor":"Microsat","os":"embedded"}},"mice":{"_$":{"class":"ht","model":"FT1D","vendor":"Yaesu"},"_%":{"vendor":"Yaesu","model":"FTM-400DR","class":"rig"},"_)":{"class":"rig","model":"FTM-100D","vendor":"Yaesu"},"_\"":{"class":"rig","model":"FTM-350","vendor":"Yaesu"},"_ ":{"model":"VX-8","vendor":"Yaesu","class":"ht"},"_(":{"model":"FT2D","vendor":"Yaesu","class":"ht"},"|3":{"vendor":"Byonics","model":"TinyTrak3","class":"tracker"},"*v":{"class":"tracker","vendor":"KissOZ","model":"Tracker"},"_#":{"class":"ht","model":"VX-8G","vendor":"Yaesu"},"^v":{"vendor":"HinzTec","model":"anyfrog"},"|4":{"model":"TinyTrak4","vendor":"Byonics","class":"tracker"}},"micelegacy":{">":{"prefix":">","features":["messaging"],"model":"TH-D7A","vendor":"Kenwood","class":"ht"},">=":{"vendor":"Kenwood","model":"TH-D72","features":["messaging"],"suffix":"=","prefix":">","class":"ht"},"]=":{"suffix":"=","prefix":"]","vendor":"Kenwood","model":"TM-D710","features":["messaging"],"class":"rig"},"]":{"class":"rig","prefix":"]","features":["messaging"],"model":"TM-D700","vendor":"Kenwood"}}}`)

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
