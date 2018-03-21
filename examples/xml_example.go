package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type address struct {
	City, State string
}

type person struct {
	XMLName   xml.Name `xml:"person"`
	ID        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Age       int      `xml:"age"`
	Height    float32  `xml:"height,omitempty"`
	Married   bool
	address
	Comment string `xml:",comment"`
}

// Example of XML Marshalling in Golang
func exampleMarshalIndent() {
	v := &person{ID: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.address = address{"Hanga Roa", "Easter Island"}

	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
}

func exampleEncoder() {
	v := &person{ID: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = "Need more details."
	v.address = address{"Hanga Roa", "Easter Islang"}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func main() {
	fmt.Println("Example Marshal Indent:")
	exampleMarshalIndent()
	fmt.Println("")

	fmt.Println("Example Encoder:")
	exampleEncoder()
	fmt.Println("")
}
