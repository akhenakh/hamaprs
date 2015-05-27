package hamaprs

import "testing"

func TestFile(t *testing.T) {
	p := &Packet{DestinationCallsign: "APU25N"}
	if p.Device().Model != "UI-View32" {
		t.Error("Should be UI-View32")
	}
	if p.Device().Model != "UI-View32" {
		t.Error("Should be UI-View32")
	}
	/* browse the  Tries */
	// cnode := trieRoot.root
	// var f func()
	// f = func() {
	// 	for _, n := range cnode.children {
	// 		log.Println(n, n.children)
	// 		cnode = n
	// 		f()
	// 	}
	// }
	// f()
}

func TestTrie(t *testing.T) {
	// initiliazing the trie
	trieRoot := &Trie{root: &TrieNode{children: make([]*TrieNode, 0)}}

	trieRoot.Add("TEST", Device{Model: "test"})

	if trieRoot.match("T") != nil {
		t.Error("Should not match")
	}

	if trieRoot.match("TEST") == nil {
		t.Error("Should match")
	}

	if trieRoot.match("TESTE") != nil {
		t.Error("Should not match")
	}

	if trieRoot.match("TESTE") != nil {
		t.Error("Should not match")
	}

	trieRoot.Add("TESTE", Device{Model: "teste"})

	if trieRoot.match("TESTE") == nil {
		t.Error("Should match")
	}

	if trieRoot.match("TEST4") != nil {
		t.Error("Should not match")
	}

	trieRoot.Add("TESTn", Device{Model: "testn"})

	if trieRoot.match("TEST4") == nil {
		t.Error("Should match")
	}

	if trieRoot.match("TEST4T") != nil {
		t.Error("Should not match")
	}

	trieRoot.Add("TESTnT", Device{Model: "testn"})

	if trieRoot.match("TEST4T") == nil {
		t.Error("Should match")
	}

	trieRoot.Add("TES??K", Device{Model: "test??k"})
	if trieRoot.match("TESBBK") == nil {
		t.Error("Should match")
	}

}

func TestTrieWildcard(t *testing.T) {
	// initiliazing the trie
	trieRoot := &Trie{root: &TrieNode{children: make([]*TrieNode, 0)}}
	trieRoot.Add("AEST*", Device{Model: "test"})
	trieRoot.Add("AESTnB", Device{Model: "testnB"})
	trieRoot.Add("AESTA*", Device{Model: "testa*"})

	cnode := trieRoot.root

	for {
		if len(cnode.children) > 0 {
			cnode = cnode.children[0]
		} else {
			break
		}

		if cnode.value == 'T' {
			if cnode.children[len(cnode.children)-1].value != '*' {
				t.Error("Last should be n *")
			}

			if cnode.children[len(cnode.children)-2].value != 'n' {
				t.Error("should be n right before *")
			}
		}
	}

	if trieRoot.match("AEST8B") == nil {
		t.Error("Should match")
	}

	if trieRoot.match("AEST8B").Model != "testnB" {
		t.Error("Should find model testnb")
	}

	if trieRoot.match("AESTZZ").Model != "test" {
		t.Error("Should find model test")
	}

	if trieRoot.match("AESTAAA") == nil {
		t.Error("Should match")
	}
}

func TestTrieSortedWildcard(t *testing.T) {
	trieRoot := &Trie{root: &TrieNode{children: make([]*TrieNode, 0)}}
	trieRoot.Add("AESTA*", Device{Model: "testa*"})
	trieRoot.Add("AEST*", Device{Model: "test"})
	trieRoot.Add("AESTnB", Device{Model: "testnB"})
	cnode := trieRoot.root

	if trieRoot.match("AEST8B") == nil {
		t.Error("Should match")
	}

	if trieRoot.match("AEST8B").Model != "testnB" {
		t.Error("Should find model testnb")
	}

	if trieRoot.match("AESTZZ").Model != "test" {
		t.Error("Should find model test")
	}

	if trieRoot.match("AESTAAA").Model != "testa*" {
		t.Error("Should match")
	}

	for {
		if len(cnode.children) > 0 {
			cnode = cnode.children[0]
		} else {
			break
		}

		if cnode.value == 'T' {
			if cnode.children[1].value != 'n' {
				t.Error("Should be n first then * got", cnode.children)
			}
			if cnode.children[2].value != '*' {
				t.Error("Should be n first then *")
			}
		}
	}

}

func TestMice(t *testing.T) {
	parser := NewParser()
	msg, _ := parser.ParsePacket("JE6EET-9>S3SWY6,WIDE1-1,qAS,JH6ETS-10:`;\\ll} >/`\"3{}CM now GIGA No...5_$", false)
	if msg.Device().Model != "FT1D" {
		t.Error("Should be FT1D")
	}
	msg, _ = parser.ParsePacket("AF7PZ-7>TWSXYW,ERINB,WIDE1*,WIDE2-1,qAo,K7FZO:`2LHp@3[/`\"4(}_$", false)
	if msg.Device().Model != "FT1D" {
		t.Error("Should be FT1D")
	}
	msg, _ = parser.ParsePacket("JE4MKV-9>S4SSY8,JM4WDK-2*,qAR,JA4YMC-10:`=D[m\\Tv\\`\"3z}_$", false)
	if msg.Device().Model != "FT1D" {
		t.Error("Should be FT1D")
	}

	msg, _ = parser.ParsePacket("VK7QF-9>T2U1P4,WIDE1-1,WIDE2-1,qAR,VK7ZRO-2:`K1qm y>/'\"4/}|!$&<'V|!w4&!|3", false)
	if msg.Device().Model != "TinyTrak3" {
		t.Error("Should be TinyTrak")
	}

}
