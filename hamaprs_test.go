package hamaprs

import "testing"

func TestParse(t *testing.T) {
	parser := NewParser()
	msg, _ := parser.ParsePacket("JE6EET-9>S3SWY6,WIDE1-1,qAS,JH6ETS-10:`;\\ll} >/`\"3{}CM now GIGA No...5_$", false)
	if msg.MicE != "in service" {
		t.Error("should be  in service is", msg.MicE)
	}

	msg, _ = parser.ParsePacket("AF7PZ-7>TWSXYW,ERINB,WIDE1*,WIDE2-1,qAo,K7FZO:`2LHp@3[/`\"4(}_$", false)
	if msg.MicE != "off duty" {
		t.Error("should be off duty is", msg.MicE)
	}

	msg, _ = parser.ParsePacket("JE4MKV-9>S4SSY8,JM4WDK-2*,qAR,JA4YMC-10:`=D[m\\Tv\\`\"3z}_$", false)
	if msg.MicE != "in service" {
		t.Error("should be in service is", msg.MicE)
	}

	msg, _ = parser.ParsePacket("VK7QF-9>T2U1P4,WIDE1-1,WIDE2-1,qAR,VK7ZRO-2:`K1qm y>/'\"4/}|!$&<'V|!w4&!|3", false)
	if msg.MicE != "in service" {
		t.Error("should be in service is", msg.MicE)
	}

}
