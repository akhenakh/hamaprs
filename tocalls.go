package hamaprs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// map for exact matches callsign
var toCalls map[string]Device

// miceLegacy a small list of device to test for MicE
var miceLegacy []Device

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

// MiceLegacyJSON struct handle the tocalls Mice Legacy
type MiceLegacyJSON struct {
	Mice map[string]Device `json:"micelegacy"`
}

// Device stores the model information for a radio device
type Device struct {
	Model  string `json:"model"`
	Vendor string `json:"vendor"`
	Class  string `json:"class"`
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

// source 	https://raw.githubusercontent.com/hessu/aprs-deviceid/master/generated/tocalls.pretty.json
var ToCallsJSONData = []byte(`{"mice":{"*v":{"model":"Tracker","class":"tracker","vendor":"KissOZ"},"|4":{"model":"TinyTrak4","class":"tracker","vendor":"Byonics"},"^v":{"model":"anyfrog","vendor":"HinzTec"},"_\"":{"model":"FTM-350","class":"rig","vendor":"Yaesu"},"|3":{"vendor":"Byonics","class":"tracker","model":"TinyTrak3"},"_%":{"class":"rig","vendor":"Yaesu","model":"FTM-400DR"},"_)":{"model":"FTM-100D","vendor":"Yaesu","class":"rig"},"_$":{"model":"FT1D","class":"ht","vendor":"Yaesu"},"_(":{"class":"ht","vendor":"Yaesu","model":"FT2D"},"_ ":{"vendor":"Yaesu","class":"ht","model":"VX-8"},"_#":{"model":"VX-8G","vendor":"Yaesu","class":"ht"}},"micelegacy":{">":{"model":"TH-D7A","prefix":">","features":["messaging"],"vendor":"Kenwood","class":"ht"},"]=":{"model":"TM-D710","vendor":"Kenwood","prefix":"]","features":["messaging"],"class":"rig","suffix":"="},">=":{"model":"TH-D72","suffix":"=","class":"ht","features":["messaging"],"prefix":">","vendor":"Kenwood"},"]":{"features":["messaging"],"prefix":"]","vendor":"Kenwood","class":"rig","model":"TM-D700"}},"tocalls":{"API???":{"model":"unknown","vendor":"Icom","class":"dstar"},"APJE??":{"vendor":"Gregg Wonderly, W5GGW","model":"JeAPRS"},"APAM??":{"model":"AltOS","vendor":"Altus Metrum","class":"tracker"},"APUDR?":{"model":"UDR","vendor":"NW Digital Radio"},"APGO??":{"class":"app","vendor":"AA3NJ","model":"APRS-Go"},"APDG??":{"class":"dstar","vendor":"Jonathan, G4KLX","model":"ircDDB Gateway"},"APMI05":{"vendor":"Microsat","os":"embedded","model":"PLXTracker"},"APMI01":{"model":"WX3in1","os":"embedded","vendor":"Microsat"},"APSK63":{"model":"APRS Messenger","os":"Windows","vendor":"Chris Moulding, G4HYG","class":"software"},"APZ18":{"model":"UIdigi","vendor":"IW3FQG","class":"digi"},"APSTPO":{"model":"Satellite Tracking and Operations","vendor":"N0AGI","class":"software"},"APHH?":{"vendor":"Steven D. Bragg, KA9MVA","class":"tracker","model":"HamHud"},"APT3??":{"model":"TinyTrak3","vendor":"Byonics","class":"tracker"},"APHK??":{"model":"Digipeater/tracker","vendor":"LA1BR"},"APRX??":{"model":"aprx","class":"software","vendor":"OH2MQK"},"APMG??":{"model":"MiniGate","vendor":"Alex, AB0TJ","os":"Netduino","class":"software"},"APND??":{"model":"DIGI_NED","vendor":"PE1MEW"},"APK003":{"vendor":"Kenwood","class":"ht","model":"TH-D72"},"APFI??":{"class":"app","vendor":"aprs.fi"},"APSMS?":{"model":"SMS gateway","class":"software","vendor":"Paul Defrusne"},"APJS??":{"vendor":"Peter Loveall, AE5PL","model":"javAPRSSrvr"},"APOZ??":{"vendor":"OZ1EKD, OZ7HVO","class":"tracker","model":"KissOZ"},"APB2MF":{"class":"software","vendor":"Mike, DL2MF","os":"Windows","model":"MF2APRS Radiosonde tracking tool"},"APECAN":{"model":"Pecan Pico APRS Balloon Tracker","vendor":"KT5TK/DL7AD","class":"tracker"},"APJA??":{"model":"JavAPRS","vendor":"K4HG & AE5PL"},"APFG??":{"model":"Flood Gage","vendor":"KP4DJT","class":"software"},"APBL??":{"model":"BeeLine GPS","class":"tracker","vendor":"BigRedBee"},"APAGW":{"model":"AGWtracker","class":"software","os":"Windows","vendor":"SV2AGW"},"APERXQ":{"model":"PE1RXQ APRS Tracker","class":"tracker","vendor":"PE1RXQ"},"APY02D":{"model":"FT2D","class":"ht","vendor":"Yaesu"},"APOLU?":{"class":"satellite","vendor":"AMSAT-LU","model":"Oscar"},"APMI03":{"os":"embedded","vendor":"Microsat","model":"PLXDigi"},"APMI02":{"model":"WXEth","vendor":"Microsat","os":"embedded"},"APSTM?":{"class":"tracker","vendor":"W7QO","model":"Balloon tracker"},"APMI??":{"os":"embedded","vendor":"Microsat"},"APY01D":{"vendor":"Yaesu","class":"ht","model":"FT1D"},"APCL??":{"class":"app","vendor":"maprs.org","model":"maprs"},"APAH??":{"model":"AHub"},"APSAR":{"os":"Windows","vendor":"ZL4FOX","class":"software","model":"SARTrack"},"APNK80":{"model":"KAM","vendor":"Kantronics"},"APWW??":{"features":["messaging","item-in-msg"],"os":"Windows","vendor":"KJ4ERJ","class":"software","model":"APRSIS32"},"APK1??":{"vendor":"Kenwood","class":"rig","model":"TM-D700"},"APN102":{"vendor":"Gregg Wonderly, W5GGW","os":"ipad","class":"app","model":"APRSNow"},"APNM??":{"vendor":"MFJ","model":"TNC"},"APMI04":{"model":"WX3in1 Mini","os":"embedded","vendor":"Microsat"},"APAG??":{"model":"AGate"},"APBPQ?":{"class":"software","os":"Windows","vendor":"John Wiseman, G8BPQ","model":"BPQ32"},"APAX??":{"model":"AFilterX"},"APDnnn":{"class":"software","os":"Linux/Unix","vendor":"Open Source","model":"aprsd"},"APT2??":{"model":"TinyTrak2","class":"tracker","vendor":"Byonics"},"APZTKP":{"class":"tracker","os":"embedded","vendor":"Nick Hanks, N0LP","model":"TrackPoint"},"APWM??":{"model":"APRSISCE","class":"software","vendor":"KJ4ERJ","os":"Windows Mobile","features":["messaging","item-in-msg"]},"APU1??":{"model":"UI-View16","vendor":"Roger Barker, G4IDE","os":"Windows","class":"software"},"APCLEY":{"vendor":"ZS6EY","class":"tracker","model":"EYTraker"},"APZWKR":{"vendor":"GM1WKR","class":"software","model":"NetSked"},"APRS":{"vendor":"Unknown","model":"Unknown"},"APAGW?":{"model":"AGWtracker","os":"Windows","vendor":"SV2AGW","class":"software"},"APKRAM":{"os":"ios","vendor":"kramstuff.com","class":"app","model":"Ham Tracker"},"APAVT5":{"model":"AP510","class":"tracker","vendor":"SainSonic"},"APNP??":{"vendor":"PacComm","model":"TNC"},"APNK01":{"model":"TM-D700","features":["messaging"],"vendor":"Kenwood","class":"rig"},"APNT??":{"vendor":"SV2AGW","class":"digi","model":"TNT TNC as a digipeater"},"APZ247":{"model":"UPRS","vendor":"NR0Q"},"APIC??":{"model":"PICiGATE","vendor":"HA9MCQ"},"APWA??":{"model":"APRSISCE","vendor":"KJ4ERJ","os":"Android","class":"software"},"APFII?":{"vendor":"aprs.fi","os":"ios","class":"app"},"APN3??":{"vendor":"Kantronics","model":"KPC-3"},"APZ186":{"model":"UIdigi","class":"digi","vendor":"IW3FQG"},"APDST?":{"vendor":"SP9UOB","os":"embedded","model":"dsTracker"},"APnnnD":{"class":"dstar","vendor":"Painter Engineering","model":"uSmartDigi D-Gate"},"APDPRS":{"model":"D-Star APDPRS","class":"dstar","vendor":"unknown"},"APNKMP":{"model":"KAM+","vendor":"Kantronics"},"APRRT?":{"model":"RTrak","class":"tracker","vendor":"RPC Electronics"},"APTT*":{"model":"TinyTrak","vendor":"Byonics","class":"tracker"},"APRG??":{"model":"aprsg","os":"Linux/Unix","vendor":"OH2GVE","class":"software"},"APSC??":{"class":"software","vendor":"OH2MQK, OH7LZB","model":"aprsc"},"APE???":{"model":"Telemetry devices"},"APZMDR":{"class":"tracker","vendor":"Open Source","os":"embedded","model":"HaMDR"},"APXR??":{"model":"Xrouter","vendor":"G8PZT"},"APDW??":{"vendor":"WB2OSZ","model":"DireWolf"},"APDR??":{"os":"Android","vendor":"Open Source","class":"app","model":"APRSdroid"},"APDI??":{"vendor":"Bela, HA5DI","class":"software","model":"DIXPRS"},"APRHH?":{"vendor":"Steven D. Bragg, KA9MVA","class":"tracker","model":"HamHud"},"APAF??":{"model":"AFilter"},"APJI??":{"vendor":"Peter Loveall, AE5PL","class":"software","model":"jAPRSIgate"},"APX???":{"model":"Xastir","os":"Linux/Unix","vendor":"Open Source","class":"software"},"APZG??":{"model":"aprsg","class":"software","os":"Linux/Unix","vendor":"OH2GVE"},"APVE??":{"model":"EchoLink","vendor":"unknown"},"APCWP8":{"model":"WinphoneAPRS","vendor":"GM7HHB","class":"app"},"APC???":{"model":"APRS/CE","class":"app","vendor":"Rob Wittner, KZ5RW"},"APS???":{"vendor":"Brent Hildebrand, KH2Z","class":"software","model":"APRS+SA"},"APAW??":{"os":"Windows","vendor":"SV2AGW","class":"software","model":"AGWPE"},"APNU??":{"model":"UIdigi","class":"digi","vendor":"IW3FQG"},"APW???":{"model":"WinAPRS","os":"Windows","vendor":"Sproul Brothers","class":"software"},"APMT??":{"class":"tracker","vendor":"LZ1PPL","model":"Micro APRS Tracker"},"APY400":{"model":"FTM-400","vendor":"Yaesu","class":"ht"},"APJY??":{"class":"software","vendor":"KA2DDO","model":"YAAC"},"APJID2":{"class":"dstar","vendor":"Peter Loveall, AE5PL","model":"D-Star APJID2"},"APPT??":{"vendor":"JF6LZE","class":"tracker","model":"KetaiTracker"},"AP4R??":{"model":"APRS4R","vendor":"Open Source","class":"software"},"APDU??":{"class":"app","os":"Android","vendor":"JA7UDE","model":"U2APRS"},"APVR??":{"vendor":"unknown","model":"IRLP"},"APZMAJ":{"model":"DeLorme inReach Tracker","vendor":"M1MAJ"},"APZ19":{"class":"digi","vendor":"IW3FQG","model":"UIdigi"},"APN9??":{"model":"KPC-9612","vendor":"Kantronics"},"APOT??":{"model":"OpenTracker","vendor":"Argent Data Systems","class":"tracker"},"APDNO?":{"class":"tracker","vendor":"DO3SWW","os":"embedded","model":"APRSduino"},"APTW??":{"vendor":"Byonics","class":"wx","model":"WXTrak"},"APHAX?":{"model":"SM2APRS SondeMonitor","class":"software","os":"Windows","vendor":"PY2UEP"},"APRNOW":{"class":"app","os":"ipad","vendor":"Gregg Wonderly, W5GGW","model":"APRSNow"},"APOA??":{"model":"app","class":"app","vendor":"OpenAPRS","os":"ios"},"APCLEZ":{"model":"Telit EZ10 GSM application","vendor":"ZS6EY","class":"tracker"},"APDF??":{"model":"Automatic DF units"},"PSKAPR":{"model":"PSKmail","vendor":"Open Source","class":"software"},"APDT??":{"model":"APRStouch Tone (DTMF)","vendor":"unknown"},"APNX??":{"model":"TNC-X","vendor":"K6DBG"},"APR2MF":{"vendor":"Mike, DL2MF","os":"Windows","class":"wx","model":"MF2wxAPRS Tinkerforge gateway"},"APAND?":{"model":"APRSdroid","os":"Android","vendor":"Open Source","class":"app"},"APHT??":{"model":"HMTracker","vendor":"IU0AAC","class":"tracker"},"APnnnU":{"vendor":"Painter Engineering","class":"digi","model":"uSmartDigi Digipeater"},"APDS??":{"os":"embedded","vendor":"SP9UOB","model":"dsDIGI"},"APK0??":{"vendor":"Kenwood","class":"ht","model":"TH-D7"},"APT4??":{"model":"TinyTrak4","class":"tracker","vendor":"Byonics"},"APR8??":{"model":"APRSdos","class":"software","vendor":"Bob Bruninga, WB4APR"},"AP1WWX":{"vendor":"TAPR","class":"wx","model":"T-238+"},"APNW??":{"os":"embedded","vendor":"SQ3FYK","model":"WX3in1"},"APTR??":{"model":"MotoTRBO","vendor":"Motorola"},"APLM??":{"vendor":"WA0TQG","class":"software"},"APCLWX":{"model":"EYWeather","class":"wx","vendor":"ZS6EY"},"APU2*":{"class":"software","vendor":"Roger Barker, G4IDE","os":"Windows","model":"UI-View32"}},"classes":{"digi":{"description":"Digipeater firmware","shown":"Digipeater"},"wx":{"description":"Dedicated weather station","shown":"Weather station"},"rig":{"description":"Mobile or desktop radio","shown":"Rig"},"app":{"shown":"Mobile app","description":"Mobile phone or tablet app"},"dstar":{"description":"D-Star radio","shown":"D-Star"},"tracker":{"shown":"Tracker","description":"Tracker device"},"satellite":{"description":"Satellite-based station","shown":"Satellite"},"ht":{"shown":"HT","description":"Hand-held radio"},"software":{"shown":"Software","description":"Desktop software"}}}`)

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

	var miceLegacyJSON MiceLegacyJSON
	json.Unmarshal(ToCallsJSONData, &miceLegacyJSON)
	for _, d := range miceLegacyJSON.Mice {
		miceLegacy = append(miceLegacy, d)
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

// MiceLegacyMatch test for a matching device using Prefix & Suffix againt msg.Comment
func (d *Device) MiceLegacyMatch(p *Packet) bool {
	if len(p.Comment) < 2 {
		return false
	}

	if (d.Prefix != "" && p.Comment[:1] == d.Prefix) || (d.Suffix != "" && p.Comment[len(p.Comment)-1:] == d.Suffix) {
		return true
	}

	return false
}

// return true if the rune is between 0 - 9
func isNumeric(r rune) bool {
	// convert a rune to it's numerica value
	if int(r-'0') >= 0 && int(r-'0') <= 9 {
		return true
	}
	return false
}
