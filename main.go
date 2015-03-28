package main
 
import (
  "encoding/xml"
  "net/http"
  "fmt"
)
 
type TwiML struct {
  XMLName xml.Name `xml:"Response"`
 
  Say    string `xml:",omitempty"`
  Play   string `xml:",omitempty"`
  Message   string `xml:",omitempty"`
}
 
func main() {
  http.HandleFunc("/twiml", twiml)
  http.ListenAndServe(":3000", nil)
}
 
func twiml(w http.ResponseWriter, r *http.Request) {
  twiml := TwiML{Message: "Hello World!!!"}
  x, err := xml.Marshal(twiml)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
 
  w.Header().Set("Content-Type", "application/xml")
  w.Write(x)
}