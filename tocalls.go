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

// ToCallsJSONData raw JSON source 	https://raw.githubusercontent.com/hessu/aprs-deviceid/master/generated/tocalls.pretty.json
var ToCallsJSONData = []byte(`{"micelegacy":{">^":{"features":["messaging"],"class":"ht","model":"TH-D74","vendor":"Kenwood","suffix":"^","prefix":">"},">":{"class":"ht","features":["messaging"],"prefix":">","model":"TH-D7A","vendor":"Kenwood"},">=":{"vendor":"Kenwood","model":"TH-D72","prefix":">","suffix":"=","class":"ht","features":["messaging"]},"]":{"features":["messaging"],"class":"rig","model":"TM-D700","vendor":"Kenwood","prefix":"]"},"]=":{"model":"TM-D710","vendor":"Kenwood","suffix":"=","prefix":"]","class":"rig","features":["messaging"]}},"mice":{"_#":{"vendor":"Yaesu","model":"VX-8G","class":"ht"},"_(":{"class":"ht","vendor":"Yaesu","model":"FT2D"},"_\"":{"model":"FTM-350","vendor":"Yaesu","class":"rig"},"|4":{"class":"tracker","model":"TinyTrak4","vendor":"Byonics"},"*v":{"model":"Tracker","vendor":"KissOZ","class":"tracker"},"_$":{"class":"ht","model":"FT1D","vendor":"Yaesu"},"_ ":{"model":"VX-8","vendor":"Yaesu","class":"ht"},"|3":{"class":"tracker","model":"TinyTrak3","vendor":"Byonics"},"^v":{"vendor":"HinzTec","model":"anyfrog"},"_)":{"class":"rig","vendor":"Yaesu","model":"FTM-100D"},"_%":{"model":"FTM-400DR","vendor":"Yaesu","class":"rig"}},"tocalls":{"APNW??":{"model":"WX3in1","os":"embedded","vendor":"SQ3FYK"},"APJI??":{"class":"software","model":"jAPRSIgate","vendor":"Peter Loveall, AE5PL"},"APU1??":{"os":"Windows","model":"UI-View16","vendor":"Roger Barker, G4IDE","class":"software"},"APAVT5":{"class":"tracker","vendor":"SainSonic","model":"AP510"},"APOA??":{"class":"app","vendor":"OpenAPRS","os":"ios","model":"app"},"APAG??":{"model":"AGate"},"APAX??":{"model":"AFilterX"},"APNKMP":{"model":"KAM+","vendor":"Kantronics"},"APIN??":{"model":"PinPoint","vendor":"AB0WV"},"APMI05":{"model":"PLXTracker","os":"embedded","vendor":"Microsat"},"APDF??":{"model":"Automatic DF units"},"APCLWX":{"model":"EYWeather","vendor":"ZS6EY","class":"wx"},"APHT??":{"vendor":"IU0AAC","model":"HMTracker","class":"tracker"},"APAM??":{"model":"AltOS","vendor":"Altus Metrum","class":"tracker"},"APZMAJ":{"vendor":"M1MAJ","model":"DeLorme inReach Tracker"},"AP1WWX":{"vendor":"TAPR","model":"T-238+","class":"wx"},"APCLEY":{"class":"tracker","model":"EYTraker","vendor":"ZS6EY"},"APYS??":{"class":"software","model":"Python APRS","vendor":"W2GMD"},"APDnnn":{"class":"software","os":"Linux/Unix","model":"aprsd","vendor":"Open Source"},"APDT??":{"model":"APRStouch Tone (DTMF)","vendor":"unknown"},"PSKAPR":{"vendor":"Open Source","model":"PSKmail","class":"software"},"APMI01":{"model":"WX3in1","os":"embedded","vendor":"Microsat"},"APOZ??":{"class":"tracker","model":"KissOZ","vendor":"OZ1EKD, OZ7HVO"},"APPT??":{"model":"KetaiTracker","vendor":"JF6LZE","class":"tracker"},"APMT??":{"vendor":"LZ1PPL","model":"Micro APRS Tracker","class":"tracker"},"APAW??":{"class":"software","os":"Windows","model":"AGWPE","vendor":"SV2AGW"},"APCL??":{"model":"maprs","vendor":"maprs.org","class":"app"},"APK0??":{"vendor":"Kenwood","model":"TH-D7","class":"ht"},"APDU??":{"os":"Android","model":"U2APRS","vendor":"JA7UDE","class":"app"},"APGO??":{"class":"app","vendor":"AA3NJ","model":"APRS-Go"},"APK1??":{"vendor":"Kenwood","model":"TM-D700","class":"rig"},"APRNOW":{"class":"app","vendor":"Gregg Wonderly, W5GGW","os":"ipad","model":"APRSNow"},"APMI02":{"os":"embedded","model":"WXEth","vendor":"Microsat"},"APY01D":{"class":"ht","model":"FT1D","vendor":"Yaesu"},"APAGW":{"os":"Windows","model":"AGWtracker","vendor":"SV2AGW","class":"software"},"APC???":{"class":"app","model":"APRS/CE","vendor":"Rob Wittner, KZ5RW"},"APZMDR":{"class":"tracker","vendor":"Open Source","os":"embedded","model":"HaMDR"},"APBPQ?":{"vendor":"John Wiseman, G8BPQ","os":"Windows","model":"BPQ32","class":"software"},"APZG??":{"class":"software","vendor":"OH2GVE","os":"Linux/Unix","model":"aprsg"},"APDS??":{"model":"dsDIGI","os":"embedded","vendor":"SP9UOB"},"APVE??":{"model":"EchoLink","vendor":"unknown"},"APK004":{"vendor":"Kenwood","model":"TH-D74","class":"ht"},"APECAN":{"model":"Pecan Pico APRS Balloon Tracker","vendor":"KT5TK/DL7AD","class":"tracker"},"APSTM?":{"model":"Balloon tracker","vendor":"W7QO","class":"tracker"},"APWW??":{"features":["messaging","item-in-msg"],"class":"software","vendor":"KJ4ERJ","os":"Windows","model":"APRSIS32"},"APN3??":{"vendor":"Kantronics","model":"KPC-3"},"APN9??":{"vendor":"Kantronics","model":"KPC-9612"},"APHK??":{"model":"Digipeater/tracker","vendor":"LA1BR"},"APIC??":{"vendor":"HA9MCQ","model":"PICiGATE"},"APXR??":{"model":"Xrouter","vendor":"G8PZT"},"APSAR":{"class":"software","model":"SARTrack","os":"Windows","vendor":"ZL4FOX"},"APS???":{"class":"software","model":"APRS+SA","vendor":"Brent Hildebrand, KH2Z"},"APSTPO":{"class":"software","vendor":"N0AGI","model":"Satellite Tracking and Operations"},"APZTKP":{"class":"tracker","vendor":"Nick Hanks, N0LP","os":"embedded","model":"TrackPoint"},"APR8??":{"class":"software","model":"APRSdos","vendor":"Bob Bruninga, WB4APR"},"APDG??":{"class":"dstar","model":"ircDDB Gateway","vendor":"Jonathan, G4KLX"},"APOT??":{"class":"tracker","model":"OpenTracker","vendor":"Argent Data Systems"},"APTR??":{"vendor":"Motorola","model":"MotoTRBO"},"APFI??":{"vendor":"aprs.fi","class":"app"},"APMI03":{"os":"embedded","model":"PLXDigi","vendor":"Microsat"},"APRX??":{"os":"Linux/Unix","vendor":"Kenneth W. Finnegan, W6KWF","model":"Aprx","class":"software"},"APK003":{"vendor":"Kenwood","model":"TH-D72","class":"ht"},"APWA??":{"class":"software","os":"Android","model":"APRSISCE","vendor":"KJ4ERJ"},"APBL??":{"model":"BeeLine GPS","vendor":"BigRedBee","class":"tracker"},"APB2MF":{"class":"software","vendor":"Mike, DL2MF","os":"Windows","model":"MF2APRS Radiosonde tracking tool"},"APAGW?":{"class":"software","os":"Windows","vendor":"SV2AGW","model":"AGWtracker"},"APY02D":{"class":"ht","vendor":"Yaesu","model":"FT2D"},"APFG??":{"class":"software","vendor":"KP4DJT","model":"Flood Gage"},"APWM??":{"class":"software","features":["messaging","item-in-msg"],"os":"Windows Mobile","model":"APRSISCE","vendor":"KJ4ERJ"},"APX???":{"class":"software","vendor":"Open Source","os":"Linux/Unix","model":"Xastir"},"APTW??":{"class":"wx","vendor":"Byonics","model":"WXTrak"},"APNT??":{"vendor":"SV2AGW","model":"TNT TNC as a digipeater","class":"digi"},"APZ19":{"class":"digi","vendor":"IW3FQG","model":"UIdigi"},"APNM??":{"model":"TNC","vendor":"MFJ"},"APERXQ":{"model":"PE1RXQ APRS Tracker","vendor":"PE1RXQ","class":"tracker"},"APNK01":{"class":"rig","features":["messaging"],"model":"TM-D700","vendor":"Kenwood"},"APZ247":{"vendor":"NR0Q","model":"UPRS"},"APRS":{"model":"Unknown","vendor":"Unknown"},"APN102":{"class":"app","model":"APRSNow","os":"ipad","vendor":"Gregg Wonderly, W5GGW"},"APMG??":{"class":"software","vendor":"Alex, AB0TJ","os":"Netduino","model":"MiniGate"},"APAF??":{"model":"AFilter"},"APRG??":{"class":"software","os":"Linux/Unix","vendor":"OH2GVE","model":"aprsg"},"APMI??":{"os":"embedded","vendor":"Microsat"},"APZ18":{"model":"UIdigi","vendor":"IW3FQG","class":"digi"},"AP4R??":{"class":"software","vendor":"Open Source","model":"APRS4R"},"APJID2":{"model":"D-Star APJID2","vendor":"Peter Loveall, AE5PL","class":"dstar"},"APVR??":{"model":"IRLP","vendor":"unknown"},"APAND?":{"class":"app","os":"Android","model":"APRSdroid","vendor":"Open Source"},"APJS??":{"model":"javAPRSSrvr","vendor":"Peter Loveall, AE5PL"},"APnnnU":{"model":"uSmartDigi Digipeater","vendor":"Painter Engineering","class":"digi"},"APSMS?":{"class":"software","model":"SMS gateway","vendor":"Paul Defrusne"},"APT2??":{"class":"tracker","model":"TinyTrak2","vendor":"Byonics"},"APCLEZ":{"class":"tracker","model":"Telit EZ10 GSM application","vendor":"ZS6EY"},"APFII?":{"class":"app","vendor":"aprs.fi","os":"ios"},"APDI??":{"model":"DIXPRS","vendor":"Bela, HA5DI","class":"software"},"APT4??":{"vendor":"Byonics","model":"TinyTrak4","class":"tracker"},"APRRT?":{"model":"RTrak","vendor":"RPC Electronics","class":"tracker"},"APW???":{"model":"WinAPRS","os":"Windows","vendor":"Sproul Brothers","class":"software"},"APZWKR":{"vendor":"GM1WKR","model":"NetSked","class":"software"},"APND??":{"model":"DIGI_NED","vendor":"PE1MEW"},"APNP??":{"model":"TNC","vendor":"PacComm"},"API???":{"vendor":"Icom","model":"unknown","class":"dstar"},"APAH??":{"model":"AHub"},"APHH?":{"vendor":"Steven D. Bragg, KA9MVA","model":"HamHud","class":"tracker"},"APDPRS":{"class":"dstar","vendor":"unknown","model":"D-Star APDPRS"},"APUDR?":{"vendor":"NW Digital Radio","model":"UDR"},"APY400":{"class":"ht","vendor":"Yaesu","model":"FTM-400"},"APNK80":{"model":"KAM","vendor":"Kantronics"},"APNU??":{"class":"digi","model":"UIdigi","vendor":"IW3FQG"},"APSK63":{"vendor":"Chris Moulding, G4HYG","os":"Windows","model":"APRS Messenger","class":"software"},"APJA??":{"vendor":"K4HG & AE5PL","model":"JavAPRS"},"APIE??":{"model":"PiAPRS","vendor":"W7KMV"},"APDST?":{"os":"embedded","model":"dsTracker","vendor":"SP9UOB"},"APCWP8":{"vendor":"GM7HHB","model":"WinphoneAPRS","class":"app"},"APDW??":{"vendor":"WB2OSZ","model":"DireWolf"},"APOLU?":{"class":"satellite","model":"Oscar","vendor":"AMSAT-LU"},"APT3??":{"class":"tracker","model":"TinyTrak3","vendor":"Byonics"},"APLM??":{"vendor":"WA0TQG","class":"software"},"APRHH?":{"model":"HamHud","vendor":"Steven D. Bragg, KA9MVA","class":"tracker"},"APDNO?":{"os":"embedded","vendor":"DO3SWW","model":"APRSduino","class":"tracker"},"APKRAM":{"class":"app","os":"ios","vendor":"kramstuff.com","model":"Ham Tracker"},"APNX??":{"vendor":"K6DBG","model":"TNC-X"},"APnnnD":{"class":"dstar","model":"uSmartDigi D-Gate","vendor":"Painter Engineering"},"APJY??":{"vendor":"KA2DDO","model":"YAAC","class":"software"},"APE???":{"model":"Telemetry devices"},"APHAX?":{"model":"SM2APRS SondeMonitor","os":"Windows","vendor":"PY2UEP","class":"software"},"APR2MF":{"class":"wx","os":"Windows","model":"MF2wxAPRS Tinkerforge gateway","vendor":"Mike, DL2MF"},"APU2*":{"vendor":"Roger Barker, G4IDE","os":"Windows","model":"UI-View32","class":"software"},"APDR??":{"class":"app","model":"APRSdroid","os":"Android","vendor":"Open Source"},"APJE??":{"vendor":"Gregg Wonderly, W5GGW","model":"JeAPRS"},"APTT*":{"class":"tracker","model":"TinyTrak","vendor":"Byonics"},"APMI04":{"model":"WX3in1 Mini","vendor":"Microsat","os":"embedded"},"APZ186":{"class":"digi","model":"UIdigi","vendor":"IW3FQG"},"APSC??":{"class":"software","vendor":"OH2MQK, OH7LZB","model":"aprsc"}},"classes":{"rig":{"description":"Mobile or desktop radio","shown":"Rig"},"app":{"shown":"Mobile app","description":"Mobile phone or tablet app"},"tracker":{"description":"Tracker device","shown":"Tracker"},"satellite":{"shown":"Satellite","description":"Satellite-based station"},"dstar":{"description":"D-Star radio","shown":"D-Star"},"software":{"shown":"Software","description":"Desktop software"},"digi":{"description":"Digipeater firmware","shown":"Digipeater"},"wx":{"shown":"Weather station","description":"Dedicated weather station"},"ht":{"shown":"HT","description":"Hand-held radio"}}}`)

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

	// it needs to be inserted in the right order
	// those with prefix & suffix 1st or it will match the wrong one
	for _, d := range miceLegacyJSON.Mice {
		if d.Prefix != "" && d.Suffix != "" {
			miceLegacy = append(miceLegacy[:0], append([]Device{d}, miceLegacy[0:]...)...)
		} else {
			miceLegacy = append(miceLegacy, d)
		}
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

	if d.Prefix != "" && d.Suffix != "" {
		return p.Comment[:1] == d.Prefix && p.Comment[len(p.Comment)-1:] == d.Suffix
	}

	if d.Prefix != "" && p.Comment[:1] == d.Prefix {
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
