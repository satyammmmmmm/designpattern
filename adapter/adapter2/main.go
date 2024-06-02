// Let implement only for xml data to json and json data to xml .

// we have xml interface and json interface
// Now problems we need to solve is  .
// a->How xml data is printed in json ?
// b->How json data is printed in xml?

// Basic things to note ,other than that  comments on code help to grasp the program.

// xmlData and jsonData are incompatible interfaces representing XML and JSON data printing respectively.
// xmlToJsonAdapter and jsonToXmlAdapter act as adapters to make printing XML data in JSON format and vice versa compatible.
// ServiceXmltoJson and ServiceJsontoXml functions use these adapters to convert XML to JSON and JSON to XML respectively,
// thus providing compatibility between the two formats.

// The xmlInterface and jsonInterface interfaces define the common operations (printXml and printJson respectively) for the adapters to use.

package main

import "fmt"

type user struct{}

// used to print data in xml
type xmlData struct{}

// used to print data in json
type jsonData struct{}

// creating adapter to print xml data in json
type xmlToJsonAdapter struct {
	xmlAdapter jsonData
}

// we create adapter to print json data in xml
type jsonToXmlAdapter struct {
	jsonAdapter xmlData
}

// interface implementing json data printing function
type jsonInterface interface {
	printJson()
}

// interface implementing xml data printing function
type xmlInterface interface {
	printXml()
}

// adapter method to make printing xml data to json data compatible
func (adapter xmlToJsonAdapter) printXml() {
	adapter.xmlAdapter.printJson()

}

// adapter method to make printing json data to xml data compatible
func (adapter jsonToXmlAdapter) printJson() {
	adapter.jsonAdapter.printXml()

}

func (xml xmlData) printXml() {
	fmt.Println("printing xml Data")
}
func (xml jsonData) printJson() {
	fmt.Println("printing json Data")
}

// service to print xml data
func (u user) ServiceXml(xml xmlInterface) {
	fmt.Println("we want to print xml data ")
	xml.printXml()

}

// service to print json data
func (u user) ServiceJson(json jsonInterface) {
	fmt.Println("we want to print json data ")
	json.printJson()

}

// service for converting xml to json
func (u user) ServiceXmltoJson(j xmlInterface) {
	fmt.Println("we want to convert xml to json ")
	j.printXml()

}

// service for converting xml to json
func (u user) ServiceJsontoXml(x jsonInterface) {
	fmt.Println("we want to convert json to xml ")
	x.printJson()

}
func main() {
	u := user{}

	//printing xml data
	x := xmlData{}
	u.ServiceXml(x)
	fmt.Println("-----------------------")
	//printing json data
	j := jsonData{}
	u.ServiceJson(j)
	fmt.Println("-----------------------")
	// we have xml data and we want to print it in json
	xmlTojson := xmlToJsonAdapter{}
	u.ServiceXmltoJson(xmlTojson)
	fmt.Println("-----------------------")
	//we have json data and we want to print it in xml
	jsonToxml := jsonToXmlAdapter{}
	u.ServiceJsontoXml(jsonToxml)
}
