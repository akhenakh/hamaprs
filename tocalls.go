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
var ToCallsJSONData = []byte(`{"micelegacy":{">":{"prefix":">","vendor":"Kenwood","features":["messaging"],"model":"TH-D7A","class":"ht"},"]":{"features":["messaging"],"model":"TM-D700","class":"rig","vendor":"Kenwood","prefix":"]"},"]=":{"suffix":"=","class":"rig","features":["messaging"],"model":"TM-D710","prefix":"]","vendor":"Kenwood"},">=":{"class":"ht","suffix":"=","features":["messaging"],"model":"TH-D72","vendor":"Kenwood","prefix":">"}},"mice":{"|4":{"model":"TinyTrak4","class":"tracker","vendor":"Byonics"},"_(":{"class":"ht","model":"FT2D","vendor":"Yaesu"},"_%":{"model":"FTM-400DR","class":"rig","vendor":"Yaesu"},"_#":{"vendor":"Yaesu","model":"VX-8G","class":"ht"},"|3":{"vendor":"Byonics","model":"TinyTrak3","class":"tracker"},"_\"":{"model":"FTM-350","class":"rig","vendor":"Yaesu"},"_ ":{"model":"VX-8","class":"ht","vendor":"Yaesu"},"_)":{"model":"FTM-100D","class":"rig","vendor":"Yaesu"},"_$":{"model":"FT1D","class":"ht","vendor":"Yaesu"},"^v":{"model":"anyfrog","vendor":"HinzTec"},"*v":{"vendor":"KissOZ","model":"Tracker","class":"tracker"}},"tocalls":{"APAG??":{"model":"AGate"},"APNK80":{"vendor":"Kantronics","model":"KAM"},"APU1??":{"vendor":"Roger Barker, G4IDE","os":"Windows","model":"UI-View16","class":"software"},"APZ186":{"vendor":"IW3FQG","class":"digi","model":"UIdigi"},"APAX??":{"model":"AFilterX"},"APZ18":{"vendor":"IW3FQG","class":"digi","model":"UIdigi"},"APDST?":{"model":"dsTracker","vendor":"SP9UOB","os":"embedded"},"APLM??":{"vendor":"WA0TQG","class":"software"},"APKRAM":{"class":"app","model":"Ham Tracker","os":"ios","vendor":"kramstuff.com"},"APGO??":{"class":"app","model":"APRS-Go","vendor":"AA3NJ"},"APCL??":{"vendor":"maprs.org","model":"maprs","class":"app"},"APND??":{"model":"DIGI_NED","vendor":"PE1MEW"},"APXR??":{"vendor":"G8PZT","model":"Xrouter"},"APZ247":{"model":"UPRS","vendor":"NR0Q"},"APX???":{"class":"software","model":"Xastir","os":"Linux/Unix","vendor":"Open Source"},"APRNOW":{"class":"app","model":"APRSNow","os":"ipad","vendor":"Gregg Wonderly, W5GGW"},"APSC??":{"vendor":"OH2MQK, OH7LZB","model":"aprsc","class":"software"},"APMI04":{"model":"WX3in1 Mini","vendor":"Microsat","os":"embedded"},"APRS":{"model":"Unknown","vendor":"Unknown"},"APMI05":{"model":"PLXTracker","os":"embedded","vendor":"Microsat"},"APC???":{"model":"APRS/CE","class":"app","vendor":"Rob Wittner, KZ5RW"},"APDI??":{"model":"DIXPRS","class":"software","vendor":"Bela, HA5DI"},"APMI02":{"model":"WXEth","vendor":"Microsat","os":"embedded"},"APJI??":{"model":"jAPRSIgate","class":"software","vendor":"Peter Loveall, AE5PL"},"APCLWX":{"class":"wx","model":"EYWeather","vendor":"ZS6EY"},"APDT??":{"model":"APRStouch Tone (DTMF)","vendor":"unknown"},"APNT??":{"vendor":"SV2AGW","class":"digi","model":"TNT TNC as a digipeater"},"APOZ??":{"class":"tracker","model":"KissOZ","vendor":"OZ1EKD, OZ7HVO"},"APY01D":{"model":"FT1D","class":"ht","vendor":"Yaesu"},"APnnnU":{"vendor":"Painter Engineering","model":"uSmartDigi Digipeater","class":"digi"},"APRRT?":{"model":"RTrak","class":"tracker","vendor":"RPC Electronics"},"APNP??":{"model":"TNC","vendor":"PacComm"},"APNK01":{"vendor":"Kenwood","class":"rig","model":"TM-D700","features":["messaging"]},"APZTKP":{"vendor":"Nick Hanks, N0LP","os":"embedded","model":"TrackPoint","class":"tracker"},"APS???":{"vendor":"Brent Hildebrand, KH2Z","class":"software","model":"APRS+SA"},"APTR??":{"model":"MotoTRBO","vendor":"Motorola"},"APSTM?":{"vendor":"W7QO","class":"tracker","model":"Balloon tracker"},"APR2MF":{"model":"MF2wxAPRS Tinkerforge gateway","class":"wx","vendor":"Mike, DL2MF","os":"Windows"},"APECAN":{"vendor":"KT5TK/DL7AD","model":"Pecan Pico APRS Balloon Tracker","class":"tracker"},"APE???":{"model":"Telemetry devices"},"APVR??":{"vendor":"unknown","model":"IRLP"},"APCLEY":{"class":"tracker","model":"EYTraker","vendor":"ZS6EY"},"APSK63":{"os":"Windows","vendor":"Chris Moulding, G4HYG","class":"software","model":"APRS Messenger"},"APK003":{"class":"ht","model":"TH-D72","vendor":"Kenwood"},"APFII?":{"class":"app","os":"ios","vendor":"aprs.fi"},"APMG??":{"vendor":"Alex, AB0TJ","os":"Netduino","model":"MiniGate","class":"software"},"APTT*":{"class":"tracker","model":"TinyTrak","vendor":"Byonics"},"APY02D":{"vendor":"Yaesu","model":"FT2D","class":"ht"},"APNU??":{"vendor":"IW3FQG","class":"digi","model":"UIdigi"},"APDR??":{"os":"Android","vendor":"Open Source","class":"app","model":"APRSdroid"},"APDF??":{"model":"Automatic DF units"},"APSAR":{"os":"Windows","vendor":"ZL4FOX","class":"software","model":"SARTrack"},"APMI01":{"model":"WX3in1","os":"embedded","vendor":"Microsat"},"APUDR?":{"model":"UDR","vendor":"NW Digital Radio"},"APJA??":{"vendor":"K4HG & AE5PL","model":"JavAPRS"},"APSTPO":{"model":"Satellite Tracking and Operations","class":"software","vendor":"N0AGI"},"APAGW?":{"model":"AGWtracker","class":"software","vendor":"SV2AGW","os":"Windows"},"APMI??":{"os":"embedded","vendor":"Microsat"},"APDU??":{"vendor":"JA7UDE","os":"Android","model":"U2APRS","class":"app"},"APJY??":{"class":"software","model":"YAAC","vendor":"KA2DDO"},"APAF??":{"model":"AFilter"},"APHH?":{"model":"HamHud","class":"tracker","vendor":"Steven D. Bragg, KA9MVA"},"APJE??":{"vendor":"Gregg Wonderly, W5GGW","model":"JeAPRS"},"APTW??":{"vendor":"Byonics","class":"wx","model":"WXTrak"},"APR8??":{"class":"software","model":"APRSdos","vendor":"Bob Bruninga, WB4APR"},"APZ19":{"vendor":"IW3FQG","class":"digi","model":"UIdigi"},"APDG??":{"vendor":"Jonathan, G4KLX","class":"dstar","model":"ircDDB Gateway"},"APRG??":{"class":"software","model":"aprsg","os":"Linux/Unix","vendor":"OH2GVE"},"APZG??":{"class":"software","model":"aprsg","os":"Linux/Unix","vendor":"OH2GVE"},"APAW??":{"model":"AGWPE","class":"software","vendor":"SV2AGW","os":"Windows"},"APCLEZ":{"model":"Telit EZ10 GSM application","class":"tracker","vendor":"ZS6EY"},"APBPQ?":{"os":"Windows","vendor":"John Wiseman, G8BPQ","class":"software","model":"BPQ32"},"APHT??":{"class":"tracker","model":"HMTracker","vendor":"IU0AAC"},"APAVT5":{"vendor":"SainSonic","class":"tracker","model":"AP510"},"API???":{"vendor":"Icom","class":"dstar","model":"unknown"},"APMI03":{"vendor":"Microsat","os":"embedded","model":"PLXDigi"},"APW???":{"vendor":"Sproul Brothers","os":"Windows","model":"WinAPRS","class":"software"},"APWM??":{"vendor":"KJ4ERJ","os":"Windows Mobile","model":"APRSISCE","features":["messaging","item-in-msg"],"class":"software"},"APNX??":{"model":"TNC-X","vendor":"K6DBG"},"APDW??":{"model":"DireWolf","vendor":"WB2OSZ"},"APWW??":{"os":"Windows","vendor":"KJ4ERJ","class":"software","features":["messaging","item-in-msg"],"model":"APRSIS32"},"APAH??":{"model":"AHub"},"APT4??":{"model":"TinyTrak4","class":"tracker","vendor":"Byonics"},"APVE??":{"model":"EchoLink","vendor":"unknown"},"APN9??":{"model":"KPC-9612","vendor":"Kantronics"},"APIC??":{"model":"PICiGATE","vendor":"HA9MCQ"},"APRHH?":{"model":"HamHud","class":"tracker","vendor":"Steven D. Bragg, KA9MVA"},"APK1??":{"model":"TM-D700","class":"rig","vendor":"Kenwood"},"APNW??":{"os":"embedded","vendor":"SQ3FYK","model":"WX3in1"},"APOLU?":{"vendor":"AMSAT-LU","class":"satellite","model":"Oscar"},"APNKMP":{"model":"KAM+","vendor":"Kantronics"},"APAGW":{"model":"AGWtracker","class":"software","vendor":"SV2AGW","os":"Windows"},"APN3??":{"vendor":"Kantronics","model":"KPC-3"},"APB2MF":{"vendor":"Mike, DL2MF","os":"Windows","model":"MF2APRS Radiosonde tracking tool","class":"software"},"APAM??":{"model":"AltOS","class":"tracker","vendor":"Altus Metrum"},"APFI??":{"vendor":"aprs.fi","class":"app"},"APRX??":{"class":"software","model":"aprx","vendor":"OH2MQK"},"AP1WWX":{"class":"wx","model":"T-238+","vendor":"TAPR"},"APK0??":{"vendor":"Kenwood","model":"TH-D7","class":"ht"},"APT3??":{"vendor":"Byonics","class":"tracker","model":"TinyTrak3"},"APT2??":{"class":"tracker","model":"TinyTrak2","vendor":"Byonics"},"APnnnD":{"class":"dstar","model":"uSmartDigi D-Gate","vendor":"Painter Engineering"},"PSKAPR":{"model":"PSKmail","class":"software","vendor":"Open Source"},"AP4R??":{"model":"APRS4R","class":"software","vendor":"Open Source"},"APPT??":{"class":"tracker","model":"KetaiTracker","vendor":"JF6LZE"},"APNM??":{"vendor":"MFJ","model":"TNC"},"APWA??":{"model":"APRSISCE","class":"software","vendor":"KJ4ERJ","os":"Android"},"APOA??":{"os":"ios","vendor":"OpenAPRS","class":"app","model":"app"},"APJS??":{"vendor":"Peter Loveall, AE5PL","model":"javAPRSSrvr"},"APZMDR":{"os":"embedded","vendor":"Open Source","class":"tracker","model":"HaMDR"},"APSMS?":{"vendor":"Paul Defrusne","model":"SMS gateway","class":"software"},"APU2*":{"vendor":"Roger Barker, G4IDE","os":"Windows","model":"UI-View32","class":"software"},"APERXQ":{"class":"tracker","model":"PE1RXQ APRS Tracker","vendor":"PE1RXQ"},"APHK??":{"vendor":"LA1BR","model":"Digipeater/tracker"},"APDS??":{"model":"dsDIGI","vendor":"SP9UOB","os":"embedded"},"APFG??":{"vendor":"KP4DJT","class":"software","model":"Flood Gage"},"APZMAJ":{"model":"DeLorme inReach Tracker","vendor":"M1MAJ"},"APBL??":{"class":"tracker","model":"BeeLine GPS","vendor":"BigRedBee"},"APDPRS":{"model":"D-Star APDPRS","class":"dstar","vendor":"unknown"},"APY400":{"model":"FTM-400","class":"ht","vendor":"Yaesu"},"APMT??":{"vendor":"LZ1PPL","class":"tracker","model":"Micro APRS Tracker"},"APJID2":{"class":"dstar","model":"D-Star APJID2","vendor":"Peter Loveall, AE5PL"},"APN102":{"model":"APRSNow","class":"app","vendor":"Gregg Wonderly, W5GGW","os":"ipad"},"APDnnn":{"os":"Linux/Unix","vendor":"Open Source","class":"software","model":"aprsd"},"APOT??":{"class":"tracker","model":"OpenTracker","vendor":"Argent Data Systems"},"APHAX?":{"os":"Windows","vendor":"PY2UEP","class":"software","model":"SM2APRS SondeMonitor"},"APAND?":{"vendor":"Open Source","os":"Android","model":"APRSdroid","class":"app"},"APCWP8":{"vendor":"GM7HHB","model":"WinphoneAPRS","class":"app"},"APZWKR":{"class":"software","model":"NetSked","vendor":"GM1WKR"}},"classes":{"app":{"shown":"Mobile app","description":"Mobile phone or tablet app"},"dstar":{"description":"D-Star radio","shown":"D-Star"},"ht":{"description":"Hand-held radio","shown":"HT"},"satellite":{"shown":"Satellite","description":"Satellite-based station"},"digi":{"shown":"Digipeater","description":"Digipeater firmware"},"rig":{"description":"Mobile or desktop radio","shown":"Rig"},"tracker":{"shown":"Tracker","description":"Tracker device"},"wx":{"description":"Dedicated weather station","shown":"Weather station"},"software":{"description":"Desktop software","shown":"Software"}}}`)

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
