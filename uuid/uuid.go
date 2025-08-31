// Package uuid provides utilities for UUID generation and manipulation using Go's standard library only.
//
// Features:
//   - Generate new UUIDs (v4)
//   - Parse UUID from string
//   - Convert UUID to string
//   - Validate a UUID string
//   - Compare two UUIDs
//   - Marshal/Unmarshal UUID to/from JSON
//   - Generate nil (zero value) UUID
//
// Usage Example:
//
//	import "yourmodule/uuid"
//	id := uuid.NewV4()
//	valid := uuid.Validate(id.String())
//	...
package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// UUID represents a 128-bit universally unique identifier (UUID).
type UUID [16]byte

// NewV4 generates a new random (version 4) UUID.
func NewUUID() UUID {
	var u UUID
	_, err := rand.Read(u[:])
	if err != nil {
		panic("uuid: cannot generate random UUID: " + err.Error())
	}
	// Set version (4) and variant bits as per RFC 4122
	u[6] = (u[6] & 0x0f) | 0x40 // Version 4
	u[8] = (u[8] & 0x3f) | 0x80 // Variant is 10
	return u
}

// Parse parses a UUID from string (accepts canonical form only).
func Parse(s string) (UUID, error) {
	var u UUID
	s = strings.ToLower(s)
	if len(s) != 36 {
		return Nil(), errors.New("uuid: invalid length")
	}
	// Remove dashes
	hexStr := strings.ReplaceAll(s, "-", "")
	if len(hexStr) != 32 {
		return Nil(), errors.New("uuid: invalid format")
	}
	b, err := hex.DecodeString(hexStr)
	if err != nil || len(b) != 16 {
		return Nil(), errors.New("uuid: invalid hex")
	}
	copy(u[:], b)
	return u, nil
}

// String returns the canonical string representation of the UUID.
func (u UUID) String() string {
	b := u[:]
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16],
	)
}

// Validate checks if a string is a valid UUID.
func Validate(s string) bool {
	_, err := Parse(s)
	return err == nil
}

// Equal compares two UUIDs for equality.
func (u UUID) Equal(other UUID) bool {
	return u == other
}

// Nil returns a nil (zero value) UUID.
func Nil() UUID {
	return UUID{}
}

// MarshalJSON implements the json.Marshaler interface.
func (u UUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (u *UUID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := Parse(s)
	if err != nil {
		return err
	}
	*u = parsed
	return nil
}

// Format implements fmt.Formatter for UUID.
func (u UUID) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, "%s", u.String())
}
