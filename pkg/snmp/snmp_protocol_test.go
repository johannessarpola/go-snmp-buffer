package snmp

import "testing"

func TestTargetToString(t *testing.T) {

	s := `
	{ 
		"field" : "value",
		"protocol" : "V1"
	}
	`

	p, _ := DetermineProtocol(s)
	if p != V1 {
		t.Errorf("DetermineProtocol was invalid, expected %s got %s", "V3", p)
	}

}

func TestTargetToString2(t *testing.T) {

	s := `
	{ 
		"field" : "value",
		"protocol" : "V3"
	}
	`

	p, _ := DetermineProtocol(s)
	if p != V3 {
		t.Errorf("DetermineProtocol was invalid, expected %s got %s", "V3", p)
	}

}

func TestTargetToString3(t *testing.T) {

	s := `
	{ 
		"field" : "value",
		"protocol" : "V2"
	}
	`

	p, _ := DetermineProtocol(s)
	if p != V2 {
		t.Errorf("DetermineProtocol was invalid, expected %s got %s", "V3", p)
	}

}

func TestTargetToString4(t *testing.T) {

	s := `
	{ 
		"field" : "value",
		"protocol" : "nnnnn"
	}
	`

	p, e := DetermineProtocol(s)

	if e == nil {
		t.Error("expected error, was nil")
	}

	if p != invalid {
		t.Errorf("DetermineProtocol was invalid, expected %s got %s", "V3", p)
	}

}
