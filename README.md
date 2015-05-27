hamAPRS
=======

HamAPRS is a Go library to parse and decode APRS packet, mainly based on libfap but provides additional features

* Recognize the APRS devices

Usage
=====
You may have to set CGO_LDFLAGS and CGO_CFLAGS environment according to your path:
```
export CGO_CFLAGS=-I/opt/local/include
export CGO_LDFLAGS=-L/opt/local/lib
```

Create the parser and then parse your packet.
```
parser := NewParser()
packet, _ = parser.parsePacket("KK6NXK-7>T2U1P4,WIDE1-1,WIDE2-1,qAR,VK7ZRO-2:`K1qm y>/'\"4/}|!$&<'V|!w4&!|3", false)
fmt.Println(packet.Latitude, packet.Longitude)
```

Get the model transceiver used to send a packet.
```
packet, _ := parser.ParsePacket("KK6NXK-7>S3SWY6,WIDE1-1,qAS,JH6ETS-10:`;\\ll} >/`\"3{}CM now GIGA No...5_$", false)
fmt.Println(DeviceForCallsign(packet).Model)
```