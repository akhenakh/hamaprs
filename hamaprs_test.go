package hamaprs

import "testing"

func TestParse(t *testing.T) {
	parser := NewParser()
	msg, err := parser.ParsePacket("JE6EET-9>S3SWY6,WIDE1-1,qAS,JH6ETS-10:`;\\ll} >/`\"3{}CM now GIGA No...5_$", false)
	if err != nil {
		t.Fatal(err.Error())
	}
	if msg.MicE != "in service" {
		t.Fatal("should be  in service is", msg.MicE)
	}

	msg, err = parser.ParsePacket("AF7PZ-7>TWSXYW,ERINB,WIDE1*,WIDE2-1,qAo,K7FZO:`2LHp@3[/`\"4(}_$", false)
	if err != nil {
		t.Fatal(err.Error())
	}
	if msg.MicE != "off duty" {
		t.Fatal("should be off duty is", msg.MicE)
	}

	msg, err = parser.ParsePacket("JE4MKV-9>S4SSY8,JM4WDK-2*,qAR,JA4YMC-10:`=D[m\\Tv\\`\"3z}_$", false)
	if err != nil {
		t.Fatal(err.Error())
	}
	if msg.MicE != "in service" {
		t.Fatal("should be in service is", msg.MicE)
	}

	msg, err = parser.ParsePacket("VK7QF-9>T2U1P4,WIDE1-1,WIDE2-1,qAR,VK7ZRO-2:`K1qm y>/'\"4/}|!$&<'V|!w4&!|3", false)
	if err != nil {
		t.Fatal(err.Error())
	}
	if msg.MicE != "in service" {
		t.Fatal("should be in service is", msg.MicE)
	}

	msg, err = parser.ParsePacket("DG1ABE-9>APOTC1,WIDE1-1,WIDE2-2,qAR,DG1ABE-10:/271651z5140.78N\\00938.27EP101/000!W01!/A=000660 12.2V 21C  Parken,Michael DOK H34/Uslar 145,575/439,100Mhz", false)
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(msg.Path) != 4 {
		t.Fatal("path should be 4 is", len(msg.Path))
	}

	msg, err = parser.ParsePacket("CW9786>APRS,TCPXX*,qAX,CWOP-3:@270308z3724.73N/12208.09W_225/000g000t060r000p000P000h65b10142eWUHU21623XXSP.", false)
	if err != nil {
		t.Fatal(err.Error())
	}
	if int(msg.Latitude) != 37 {
		t.Fatal("Invalid latitude")
	}
	if int(msg.Longitude) != -122 {
		t.Fatal("Invalid longitude")
	}

	if msg.PacketType != LocationPacketType {
		t.Fatal("should be a LocationPacketType is", msg.PacketType)
	}

	if msg.Weather == nil {
		t.Fatal("should have a weather report")
	}

}

func TestWeather(t *testing.T) {
	parser := NewParser()
	raw := "CYQB>APRS,TCPIP*,qAS,KK5WM-2:@221500z4648.00N/07123.00W_000/000g...t071h73b10210 Canada_Quebec, Que"
	msg, err := parser.ParsePacket(raw, false)
	if err != nil {
		t.Fatal(err.Error())
	}

	if msg.Weather == nil {
		t.Fatal("should find a weather report here")
	}
	if int(msg.Weather.Temperature) != 21 {
		t.Fatal("should find a temperature report here")
	}
}
