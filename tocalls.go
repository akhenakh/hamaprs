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
var ToCallsJSONData = []byte(`{"classes":{"satellite":{"description":"Satellite-based station","shown":"Satellite"},"ht":{"description":"Hand-held radio","shown":"HT"},"digi":{"shown":"Digipeater","description":"Digipeater software"},"wx":{"description":"Dedicated weather station","shown":"Weather station"},"igate":{"description":"iGate software","shown":"iGate"},"dstar":{"shown":"D-Star","description":"D-Star radio"},"software":{"description":"Desktop software","shown":"Software"},"app":{"shown":"Mobile app","description":"Mobile phone or tablet app"},"tracker":{"description":"Tracker device","shown":"Tracker"},"rig":{"description":"Mobile or desktop radio","shown":"Rig"}},"tocalls":{"APAF??":{"model":"AFilter"},"APHH?":{"class":"tracker","vendor":"Steven D. Bragg, KA9MVA","model":"HamHud"},"APR8??":{"model":"APRSdos","class":"software","vendor":"Bob Bruninga, WB4APR"},"APECAN":{"model":"Pecan Pico APRS Balloon Tracker","vendor":"KT5TK/DL7AD","class":"tracker"},"APZG??":{"vendor":"OH2GVE","class":"software","os":"Linux/Unix","model":"aprsg"},"APWW??":{"vendor":"KJ4ERJ","os":"Windows","class":"software","model":"APRSIS32","features":["messaging","item-in-msg"]},"APRRT?":{"model":"RTrak","class":"tracker","vendor":"RPC Electronics"},"APJS??":{"vendor":"Peter Loveall, AE5PL","model":"javAPRSSrvr"},"APLIG?":{"model":"LightAPRS Tracker","class":"tracker","vendor":"TA2MUN/TA9OHC"},"APAVT5":{"vendor":"SainSonic","class":"tracker","model":"AP510"},"APCSS":{"vendor":"AMSAT","model":"CubeSatSim CubeSat Simulator"},"APBL??":{"vendor":"BigRedBee","class":"tracker","model":"BeeLine GPS"},"APY05D":{"class":"ht","vendor":"Yaesu","model":"FT5D"},"APHK??":{"vendor":"LA1BR","model":"Digipeater/tracker"},"APAT81":{"model":"AT-D878","vendor":"Anytone","class":"ht"},"APRHH?":{"model":"HamHud","class":"tracker","vendor":"Steven D. Bragg, KA9MVA"},"AP1WWX":{"class":"wx","vendor":"TAPR","model":"T-238+"},"APSFLG":{"model":"LoRa/APRS Gateway","vendor":"F5OPV, SFCP_LABS","os":"embedded","class":"digi"},"APRARX":{"model":"radiosonde_auto_rx","os":"Linux/Unix","class":"software","vendor":"Open Source"},"APDG??":{"model":"ircDDB Gateway","vendor":"Jonathan, G4KLX","class":"dstar"},"APNIC4":{"model":"BidaTrak","os":"embedded","class":"tracker","vendor":"SQ5EKU"},"APCSMS":{"model":"Cosmos","vendor":"USNA"},"APDnnn":{"model":"aprsd","vendor":"Open Source","os":"Linux/Unix","class":"software"},"APTNG?":{"vendor":"Filip YU1TTN","class":"tracker","model":"Tango Tracker"},"APTPN?":{"model":"TARPN Packet Node Tracker","vendor":"KN4ORB","class":"tracker"},"APZTKP":{"model":"TrackPoint","vendor":"Nick Hanks, N0LP","class":"tracker","os":"embedded"},"APRFGL":{"model":"Lora APRS Digipeater","os":"embedded","class":"digi","contact":"info@rf.guru","vendor":"RF.Guru"},"APRFGM":{"contact":"info@rf.guru","vendor":"RF.Guru","os":"embedded","class":"rig","model":"Mobile Radio"},"API970":{"model":"IC-9700","vendor":"Icom","class":"dstar"},"APRFGT":{"model":"LoRa APRS Tracker","class":"tracker","os":"embedded","vendor":"RF.Guru","contact":"info@rf.guru"},"APUDR?":{"vendor":"NW Digital Radio","model":"UDR"},"APFI??":{"class":"app","vendor":"aprs.fi"},"AP4R??":{"model":"APRS4R","class":"software","vendor":"Open Source"},"APY300":{"model":"FTM-300D","vendor":"Yaesu","class":"rig"},"APNV1?":{"model":"VP-Node","os":"embedded","vendor":"SQ8L"},"APND??":{"model":"DIGI_NED","vendor":"PE1MEW"},"APMI05":{"os":"embedded","vendor":"Microsat","model":"PLXTracker"},"APAT??":{"vendor":"Anytone"},"APB2MF":{"model":"MF2APRS Radiosonde tracking tool","class":"software","os":"Windows","vendor":"Mike, DL2MF"},"APHPIB":{"model":"Python APRS Beacon","vendor":"HP3ICC"},"API410":{"model":"IC-4100","class":"dstar","vendor":"Icom"},"APHAX?":{"vendor":"PY2UEP","class":"software","os":"Windows","model":"SM2APRS SondeMonitor"},"APVE??":{"vendor":"unknown","model":"EchoLink"},"APTR??":{"model":"MotoTRBO","vendor":"Motorola"},"APWA??":{"vendor":"KJ4ERJ","os":"Android","class":"software","model":"APRSISCE"},"APCLWX":{"class":"wx","vendor":"ZS6EY","model":"EYWeather"},"APERXQ":{"model":"PE1RXQ APRS Tracker","class":"tracker","vendor":"PE1RXQ"},"APS???":{"vendor":"Brent Hildebrand, KH2Z","class":"software","model":"APRS+SA"},"APBK??":{"vendor":"PY5BK","class":"tracker","model":"Bravo Tracker"},"APK0??":{"class":"ht","vendor":"Kenwood","model":"TH-D7"},"APPM??":{"class":"software","vendor":"DL1MX","model":"rtl-sdr Python iGate"},"APMI04":{"model":"WX3in1 Mini","vendor":"Microsat","os":"embedded"},"APNK80":{"model":"KAM","vendor":"Kantronics"},"APNM??":{"vendor":"MFJ","model":"TNC"},"APRRF?":{"class":"tracker","os":"embedded","contact":"f1evm@f1evm.fr","vendor":"Jean-Francois Huet F1EVM","features":["messaging"],"model":"Tracker for RRF"},"APFII?":{"model":"iPhone/iPad app","vendor":"aprs.fi","os":"ios","class":"app"},"APE2A?":{"model":"Email-2-APRS gateway","os":"Linux/Unix","class":"software","vendor":"NoseyNick, VA3NNW"},"APDR??":{"model":"APRSdroid","class":"app","os":"Android","vendor":"Open Source"},"APBSD?":{"vendor":"hambsd.org","model":"HamBSD"},"APCDS0":{"model":"cell tracker","vendor":"ZS6LMG","class":"tracker"},"APRX??":{"model":"Aprx","os":"Linux/Unix","class":"igate","vendor":"Kenneth W. Finnegan, W6KWF"},"APVM??":{"class":"igate","vendor":"Digital Radio China Club","model":"DRCC-DVM"},"APN3??":{"model":"KPC-3","vendor":"Kantronics"},"APBPQ?":{"model":"BPQ32","os":"Windows","class":"software","vendor":"John Wiseman, G8BPQ"},"APRFGR":{"contact":"info@rf.guru","vendor":"RF.Guru","os":"embedded","class":"rig","model":"Repeater"},"API710":{"class":"dstar","vendor":"Icom","model":"IC-7100"},"APNK01":{"vendor":"Kenwood","class":"rig","model":"TM-D700","features":["messaging"]},"APTCMA":{"vendor":"Cleber, PU1CMA","class":"tracker","model":"CAPI Tracker"},"APIN??":{"model":"PinPoint","vendor":"AB0WV"},"APKHTW":{"vendor":"Kip, W3SN","contact":"w3sn@moxracing.33mail.com","class":"wx","os":"embedded","model":"Tempest Weather Bridge"},"APNV??":{"vendor":"SQ8L"},"APTB??":{"model":"TinyAPRS","vendor":"BG5HHP"},"APNX??":{"model":"TNC-X","vendor":"K6DBG"},"APRPR?":{"contact":"dm4rw@skywaves.de","vendor":"Robert DM4RW, Peter DL6MAA","os":"embedded","class":"tracker","model":"Teensy RPR TNC"},"APDW??":{"vendor":"WB2OSZ","model":"DireWolf"},"APZMDR":{"class":"tracker","os":"embedded","vendor":"Open Source","model":"HaMDR"},"APR2MF":{"model":"MF2wxAPRS Tinkerforge gateway","vendor":"Mike, DL2MF","class":"wx","os":"Windows"},"APNCM":{"contact":"wa0tjt@gmail.com","vendor":"Keith Kaiser, WA0TJT","class":"software","os":"browser","model":"Net Control Manager"},"APAX??":{"model":"AFilterX"},"APMI03":{"os":"embedded","vendor":"Microsat","model":"PLXDigi"},"APDPRS":{"model":"D-Star APDPRS","class":"dstar","vendor":"unknown"},"APTW??":{"model":"WXTrak","class":"wx","vendor":"Byonics"},"APFG??":{"class":"software","vendor":"KP4DJT","model":"Flood Gage"},"APNKMP":{"model":"KAM+","vendor":"Kantronics"},"APLP1?":{"model":"LORA/FSK/AFSK fajny tracker","os":"embedded","class":"tracker","contact":"sq9p.peter@gmail.com","vendor":"SQ9P"},"APAND?":{"class":"app","os":"Android","vendor":"Open Source","model":"APRSdroid"},"APU1??":{"model":"UI-View16","os":"Windows","class":"software","vendor":"Roger Barker, G4IDE"},"APRRDZ":{"model":"rdzTTGOsonde","vendor":"DL9RDZ","class":"tracker"},"APDF??":{"model":"Automatic DF units"},"APN2??":{"vendor":"VE4KLM","model":"NOSaprs for JNOS 2.0"},"APIC??":{"vendor":"HA9MCQ","model":"PICiGATE"},"APNU??":{"model":"UIdigi","vendor":"IW3FQG","class":"digi"},"APERS?":{"class":"tracker","vendor":"Jason, KG7YKZ","model":"Runner tracking"},"APTKJ?":{"os":"embedded","vendor":"W9JAJ","model":"ATTiny APRS Tracker"},"APMG??":{"vendor":"Alex, AB0TJ","class":"software","model":"PiCrumbs and MiniGate"},"APLG??":{"vendor":"OE5BPA","class":"digi","model":"LoRa Gateway/Digipeater"},"APJY??":{"model":"YAAC","class":"software","vendor":"KA2DDO"},"APDU??":{"model":"U2APRS","os":"Android","class":"app","vendor":"JA7UDE"},"APIE??":{"vendor":"W7KMV","model":"PiAPRS"},"APLS??":{"model":"SARIMESH","vendor":"SARIMESH","class":"software"},"APY400":{"model":"FTM-400","class":"rig","vendor":"Yaesu"},"APOT??":{"model":"OpenTracker","class":"tracker","vendor":"Argent Data Systems"},"APNKMX":{"model":"KAM-XL","vendor":"Kantronics"},"APZ19":{"class":"digi","vendor":"IW3FQG","model":"UIdigi"},"PSKAPR":{"class":"software","vendor":"Open Source","model":"PSKmail"},"APYS??":{"model":"Python APRS","class":"software","vendor":"W2GMD"},"APT3??":{"model":"TinyTrak3","class":"tracker","vendor":"Byonics"},"APN9??":{"model":"KPC-9612","vendor":"Kantronics"},"APDI??":{"model":"DIXPRS","vendor":"Bela, HA5DI","class":"software"},"APDS??":{"model":"dsDIGI","os":"embedded","vendor":"SP9UOB"},"API80":{"vendor":"Icom","class":"dstar","model":"IC-80"},"APNT??":{"class":"digi","vendor":"SV2AGW","model":"TNT TNC as a digipeater"},"APAGW":{"model":"AGWtracker","vendor":"SV2AGW","os":"Windows","class":"software"},"APP6??":{"model":"APRSlib"},"APHPIA":{"vendor":"HP3ICC","model":"Arduino APRS"},"APESP?":{"os":"embedded","vendor":"LY3PH","model":"APRS-ESP"},"APK003":{"model":"TH-D72","vendor":"Kenwood","class":"ht"},"APHW??":{"vendor":"HamWAN"},"APMI02":{"os":"embedded","vendor":"Microsat","model":"WXEth"},"APK004":{"model":"TH-D74","class":"ht","vendor":"Kenwood"},"APZ18":{"vendor":"IW3FQG","class":"digi","model":"UIdigi"},"APNP??":{"vendor":"PacComm","model":"TNC"},"APSAR":{"model":"SARTrack","vendor":"ZL4FOX","os":"Windows","class":"software"},"APMPAD":{"model":"WXBot clone and extension","vendor":"DF1JSL"},"APHBL?":{"vendor":"KF7EEL","class":"software","model":"HBLink D-APRS Gateway"},"APDT??":{"vendor":"unknown","model":"APRStouch Tone (DTMF)"},"APATAR":{"vendor":"TA7W/OH2UDS Baris Dinc and TA6AEU","class":"digi","model":"ATA-R APRS Digipeater"},"APCN??":{"vendor":"DG5OAW","model":"carNET"},"APRFG?":{"contact":"info@rf.guru","vendor":"RF.Guru"},"APZWKR":{"model":"NetSked","vendor":"GM1WKR","class":"software"},"APZMAJ":{"vendor":"M1MAJ","model":"DeLorme inReach Tracker"},"APSFWX":{"model":"embedded Weather Station","vendor":"F5OPV, SFCP_LABS","os":"embedded","class":"wx"},"APN102":{"model":"APRSNow","class":"app","os":"ipad","vendor":"Gregg Wonderly, W5GGW"},"APCLEZ":{"vendor":"ZS6EY","class":"tracker","model":"Telit EZ10 GSM application"},"APDV??":{"class":"software","vendor":"OE6PLD","model":"SSTV with APRS"},"APLRG?":{"contact":"richonguzman@gmail.com","vendor":"Ricardo, CD2RXU","class":"igate","os":"embedded","model":"ESP32 LoRa iGate"},"APCTLK":{"class":"app","vendor":"Open Source","model":"Codec2Talkie"},"APOA??":{"class":"app","os":"ios","vendor":"OpenAPRS","model":"app"},"APJA??":{"vendor":"K4HG & AE5PL","model":"JavAPRS"},"APSC??":{"class":"software","vendor":"OH2MQK, OH7LZB","model":"aprsc"},"APQTH?":{"vendor":"Weston Bustraan, W8WJB","os":"macOS","class":"software","model":"QTH.app","features":["messaging"]},"APCLUB":{"model":"Brazil APRS network"},"APSFRP":{"os":"embedded","vendor":"F5OPV, SFCP_LABS","model":"VHF/UHF Repeater"},"APAM??":{"class":"tracker","vendor":"Altus Metrum","model":"AltOS"},"APY01D":{"class":"ht","vendor":"Yaesu","model":"FT1D"},"APWnnn":{"os":"Windows","class":"software","vendor":"Sproul Brothers","model":"WinAPRS"},"APAT51":{"model":"AT-D578","vendor":"Anytone","class":"rig"},"APAGW?":{"os":"Windows","class":"software","vendor":"SV2AGW","model":"AGWtracker"},"APPIC?":{"model":"PicoAPRS","vendor":"DB1NTO","class":"tracker"},"APLP0?":{"os":"embedded","class":"digi","vendor":"SQ9P","contact":"sq9p.peter@gmail.com","model":"fajne digi"},"APSTM?":{"vendor":"W7QO","class":"tracker","model":"Balloon tracker"},"APLRT?":{"class":"tracker","os":"embedded","contact":"richonguzman@gmail.com","vendor":"Ricardo, CD2RXU","model":"ESP32 LoRa Tracker"},"APK005":{"vendor":"Kenwood","class":"ht","model":"TH-D75"},"APW9??":{"vendor":"Mile Strk, 9A9Y","class":"wx","os":"embedded","model":"WX Katarina","features":["messaging"]},"API31":{"model":"IC-31","class":"dstar","vendor":"Icom"},"APSK63":{"vendor":"Chris Moulding, G4HYG","os":"Windows","class":"software","model":"APRS Messenger"},"APMI01":{"model":"WX3in1","os":"embedded","vendor":"Microsat"},"APPT??":{"vendor":"JF6LZE","class":"tracker","model":"KetaiTracker"},"APCWP8":{"model":"WinphoneAPRS","class":"app","vendor":"GM7HHB"},"APMI06":{"vendor":"Microsat","os":"embedded","model":"WX3in1 Plus 2.0"},"APE???":{"model":"Telemetry devices"},"APRNOW":{"os":"ipad","class":"app","vendor":"Gregg Wonderly, W5GGW","model":"APRSNow"},"API282":{"model":"IC-2820","class":"dstar","vendor":"Icom"},"APKRAM":{"vendor":"kramstuff.com","os":"ios","class":"app","model":"Ham Tracker"},"APLM??":{"vendor":"WA0TQG","class":"software"},"APWEE?":{"model":"WeeWX Weather Software","os":"Linux/Unix","class":"software","vendor":"Tom Keffer and Matthew Wall"},"APTT*":{"model":"TinyTrak","class":"tracker","vendor":"Byonics"},"APHPIW":{"model":"Python APRS WX","vendor":"HP3ICC"},"API???":{"class":"dstar","vendor":"Icom","model":"unknown"},"APXR??":{"vendor":"G8PZT","model":"Xrouter"},"APBT62":{"vendor":"BTech","model":"DMR 6x2"},"APWM??":{"vendor":"KJ4ERJ","os":"Windows Mobile","class":"software","model":"APRSISCE","features":["messaging","item-in-msg"]},"APHMEY":{"model":"APRS-IS Client for Athom Homey","vendor":"Tapio Heiskanen, OH2TH","contact":"oh2th@iki.fi"},"APU2*":{"os":"Windows","class":"software","vendor":"Roger Barker, G4IDE","model":"UI-View32"},"APMI??":{"vendor":"Microsat","os":"embedded"},"APVR??":{"model":"IRLP","vendor":"unknown"},"APRG??":{"model":"aprsg","vendor":"OH2GVE","class":"software","os":"Linux/Unix"},"APX???":{"model":"Xastir","os":"Linux/Unix","class":"software","vendor":"Open Source"},"APRFGH":{"model":"Hotspot","contact":"info@rf.guru","vendor":"RF.Guru","class":"rig","os":"embedded"},"API51":{"model":"IC-51","class":"dstar","vendor":"Icom"},"APnnnD":{"model":"uSmartDigi D-Gate","vendor":"Painter Engineering","class":"dstar"},"APC???":{"class":"app","vendor":"Rob Wittner, KZ5RW","model":"APRS/CE"},"APAH??":{"model":"AHub"},"APSMS?":{"model":"SMS gateway","vendor":"Paul Dufresne","class":"software"},"APRFGP":{"contact":"info@rf.guru","vendor":"RF.Guru","class":"ht","os":"embedded","model":"Portable Radio"},"APZ186":{"class":"digi","vendor":"IW3FQG","model":"UIdigi"},"API510":{"vendor":"Icom","class":"dstar","model":"IC-5100"},"APTCHE":{"vendor":"PU3IKE","class":"tracker","model":"TcheTracker, Tcheduino"},"APNW??":{"vendor":"SQ3FYK","os":"embedded","model":"WX3in1"},"API880":{"model":"IC-880","vendor":"Icom","class":"dstar"},"APK1??":{"model":"TM-D700","class":"rig","vendor":"Kenwood"},"APAW??":{"os":"Windows","class":"software","vendor":"SV2AGW","model":"AGWPE"},"APRFGW":{"model":"LoRa APRS Weather Station","class":"wx","os":"embedded","contact":"info@rf.guru","vendor":"RF.Guru"},"APSFTL":{"model":"LoRa/APRS Telemetry Reporter","vendor":"F5OPV, SFCP_LABS","os":"embedded"},"APDNO?":{"model":"APRSduino","os":"embedded","class":"tracker","vendor":"DO3SWW"},"APELK?":{"class":"tracker","vendor":"WB8ELK","model":"Balloon tracker"},"APRS":{"model":"Unknown","vendor":"Unknown"},"APT4??":{"model":"TinyTrak4","vendor":"Byonics","class":"tracker"},"APNV0?":{"vendor":"SQ8L","os":"embedded","model":"VP-Digi"},"APnnnU":{"vendor":"Painter Engineering","class":"digi","model":"uSmartDigi Digipeater"},"APMT??":{"vendor":"LZ1PPL","class":"tracker","model":"Micro APRS Tracker"},"APGBLN":{"class":"tracker","vendor":"NW5W","model":"GoBalloon"},"APT2??":{"vendor":"Byonics","class":"tracker","model":"TinyTrak2"},"APAG??":{"model":"AGate"},"APOZ??":{"class":"tracker","vendor":"OZ1EKD, OZ7HVO","model":"KissOZ"},"API910":{"class":"dstar","vendor":"Icom","model":"IC-9100"},"APJI??":{"model":"jAPRSIgate","class":"software","vendor":"Peter Loveall, AE5PL"},"API92":{"model":"IC-92","vendor":"Icom","class":"dstar"},"APCLEY":{"vendor":"ZS6EY","class":"tracker","model":"EYTraker"},"APLC??":{"model":"APRScube","vendor":"DL3DCW"},"APBM??":{"model":"BrandMeister DMR","vendor":"R3ABM"},"APJID2":{"vendor":"Peter Loveall, AE5PL","class":"dstar","model":"D-Star APJID2"},"APOVU?":{"model":"BeliefSat","vendor":"K J Somaiya Institute"},"APSTPO":{"class":"software","vendor":"N0AGI","model":"Satellite Tracking and Operations"},"APY02D":{"vendor":"Yaesu","class":"ht","model":"FT2D"},"APRFGI":{"model":"LoRa APRS iGate","vendor":"RF.Guru","contact":"info@rf.guru","os":"embedded","class":"igate"},"APMQ??":{"model":"Ham Radio of Things","vendor":"WB2OSZ"},"APRFGD":{"vendor":"RF.Guru","contact":"info@rf.guru","os":"embedded","class":"digi","model":"APRS Digipeater"},"APLO??":{"model":"LoRa KISS TNC/Tracker","class":"tracker","vendor":"SQ9MDD"},"APPCO?":{"os":"embedded","class":"tracker","contact":"ab4mw@radcommsoft.com","vendor":"RadCommSoft, LLC","model":"PicoAPRSTracker"},"APOCSG":{"model":"POCSAG","vendor":"N0AGI"},"APLT??":{"model":"LoRa Tracker","class":"tracker","vendor":"OE5BPA"},"APZ247":{"model":"UPRS","vendor":"NR0Q"},"APSF??":{"vendor":"F5OPV, SFCP_LABS","os":"embedded","model":"embedded APRS devices"},"APJE??":{"model":"JeAPRS","vendor":"Gregg Wonderly, W5GGW"},"APIZCI":{"model":"hymTR IZCI Tracker","vendor":"TA7W/OH2UDS and TA6AEU","os":"embedded","class":"tracker"},"APGO??":{"model":"APRS-Go","class":"app","vendor":"AA3NJ"},"APHT??":{"class":"tracker","vendor":"IU0AAC","model":"HMTracker"},"APOLU?":{"model":"Oscar","vendor":"AMSAT-LU","class":"satellite"},"APDST?":{"model":"dsTracker","vendor":"SP9UOB","os":"embedded"}},"micelegacy":{">=":{"suffix":"=","prefix":">","features":["messaging"],"model":"TH-D72","class":"ht","vendor":"Kenwood"},"]=":{"model":"TM-D710","suffix":"=","features":["messaging"],"prefix":"]","vendor":"Kenwood","class":"rig"},">":{"features":["messaging"],"prefix":">","model":"TH-D7A","class":"ht","vendor":"Kenwood"},">^":{"vendor":"Kenwood","class":"ht","model":"TH-D74","features":["messaging"],"prefix":">","suffix":"^"},">&":{"vendor":"Kenwood","class":"ht","model":"TH-D75","features":["messaging"],"prefix":">","suffix":"&"},"]":{"class":"rig","vendor":"Kenwood","features":["messaging"],"prefix":"]","model":"TM-D700"}},"mice":{"_$":{"model":"FT1D","class":"ht","vendor":"Yaesu"},"_%":{"class":"rig","vendor":"Yaesu","model":"FTM-400DR"},"(8":{"class":"ht","vendor":"Anytone","model":"D878UV"},"|3":{"model":"TinyTrak3","class":"tracker","vendor":"Byonics"},"*v":{"class":"tracker","vendor":"KissOZ","model":"Tracker"},"_ ":{"model":"VX-8","class":"ht","vendor":"Yaesu"},"|4":{"model":"TinyTrak4","vendor":"Byonics","class":"tracker"},"_(":{"model":"FT2D","class":"ht","vendor":"Yaesu"},"_#":{"class":"ht","vendor":"Yaesu","model":"VX-8G"},"_0":{"model":"FT3D","vendor":"Yaesu","class":"ht"},"_\"":{"vendor":"Yaesu","class":"rig","model":"FTM-350"},"_3":{"class":"ht","vendor":"Yaesu","model":"FT5D"},"_1":{"class":"rig","vendor":"Yaesu","model":"FTM-300D"},"(5":{"class":"ht","vendor":"Anytone","model":"D578UV"},"^v":{"model":"anyfrog","vendor":"HinzTec"},"_)":{"model":"FTM-100D","class":"rig","vendor":"Yaesu"}}}`)

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
