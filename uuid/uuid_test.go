package uuid

import (
	"encoding/json"
	"testing"
)

func TestNewUUID(t *testing.T) {
	u := NewUUID()
	if u.String() == "" || !Validate(u.String()) {
		t.Errorf("NewUUID generated invalid UUID: %s", u.String())
	}
}

func TestParse(t *testing.T) {
	u1 := NewUUID()
	u2, err := Parse(u1.String())
	if err != nil {
		t.Fatalf("Parse() error: %v", err)
	}
	if !u1.Equal(u2) {
		t.Errorf("Parsed UUID does not match original")
	}
}

func TestValidate(t *testing.T) {
	u := NewUUID().String()
	if !Validate(u) {
		t.Errorf("Validate() failed for valid UUID: %s", u)
	}
	if Validate("not-a-uuid") {
		t.Errorf("Validate() passed for invalid UUID")
	}
}

func TestEqual(t *testing.T) {
	u1 := NewUUID()
	u2, _ := Parse(u1.String())
	if !u1.Equal(u2) {
		t.Errorf("Equal() failed for identical UUIDs")
	}
	u3 := NewUUID()
	if u1.Equal(u3) {
		t.Errorf("Equal() passed for different UUIDs")
	}
}

func TestNil(t *testing.T) {
	nilUUID := Nil()
	if nilUUID.String() != "00000000-0000-0000-0000-000000000000" {
		t.Errorf("Nil() did not return zero UUID")
	}
}

func TestMarshalUnmarshalJSON(t *testing.T) {
	u := NewUUID()
	data, err := json.Marshal(u)
	if err != nil {
		t.Fatalf("MarshalJSON() error: %v", err)
	}

	var u2 UUID
	if err := json.Unmarshal(data, &u2); err != nil {
		t.Fatalf("UnmarshalJSON() error: %v", err)
	}
	if !u.Equal(u2) {
		t.Errorf("UnmarshalJSON() did not produce equal UUID")
	}
}
