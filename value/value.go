package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

/*
# Types
JSON <-> Go
string <-> string
null <-> nil
true/false <-> bool
number <-> float64, float32, int, int8 ... int64, uint, uint8 ... uint64
array <-> []T, []any
object <-> struct, map[string]any

MIA:
- time.Time: string in RFC3339 format, seconds since epoch
- []byte: string in base64

# Design
API       type User struct { ... }
Business  type User struct { ... }
Data      type User struct { ... }

# encoding/json API
Go -> io.Wrier -> JSON: json.Encoder
Go -> []byte -> JSON: json.Marshal
JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte -> Go: json.Unmarshal

[]byte -> io.Writer/i.Reader: bytes.Buffer
*/

func main() {
	data, err := json.Marshal(1)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	var i any
	if err := json.Unmarshal(data, &i); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("%v (%T)\n", i, i)

	buf := bytes.NewBuffer(data)
	dec := json.NewDecoder(buf)
	var n int
	if err := dec.Decode(&n); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("n:", n)

	v := Value{
		Unit:   "cm",
		Amount: 20.3,
	}

	// "23.3cm"

	data, err = json.Marshal(v)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(string(data))

	var v2 Value
	if err := json.Unmarshal(data, &v2); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("%#v\n", v2)

	fmt.Println(v)
}

// Different message on same endpoint
/*
{
	"type": "login",
	"payload": {
		"user": "elliot"
	}
}

{
	"type": "access",
	"payload": {
		"uri": "file:///etc/passwd"
	}
}
*/

type Message struct {
	Type    string
	Payload json.RawMessage // json won't decode this
}

type Login struct{}
type Access struct{}

// implement fmt.Stringer
func (v Value) String() string {
	return "I'm a value"

}

// https://github.com/353solutions/para

// Exercise: Implement json.Unmarshaler for Value
// Hint: fmt.Sscanf

func (v *Value) UnmarshalJSON(data []byte) error {
	// Trim surrounding "" in "20.300000cm"
	s := string(data[1 : len(data)-1])
	var a float64
	var u Unit
	if _, err := fmt.Sscanf(s, "%f%s", &a, &u); err != nil {
		return err
	}

	v.Amount = a
	v.Unit = u
	return nil
}

// Use value receiver, it'll work both for values & pointers
func (v Value) MarshalJSON() ([]byte, error) {
	// Step 1: Convert to type that encoding/json knows
	s := fmt.Sprintf("%f%s", v.Amount, v.Unit)

	// Step 2: Use json.Marshal
	return json.Marshal(s)

	// Step 3: There is not step 3
}

type Value struct {
	Unit   Unit    `json:"unit,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}

const (
	Centimeter Unit = "cm"
	Inch       Unit = "in"
)

type Unit string
