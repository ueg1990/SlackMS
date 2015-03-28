package main
 
import (
  "encoding/xml"
  //"encoding/json"
  "net/http"
  "fmt"
  //"bytes"
  // "net/url"
  //"io/ioutil"
  "strings"
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
	// data := map[string]interface{}{}
	// token := "xoxp-2757127568-2813046014-4209073904-35634c"
	// resp, _ := http.Get("https://slack.com/api/users.list?token="+token)
	// fmt.Println(resp)
	// defer resp.Body.Close()
 //    body, _ := ioutil.ReadAll(resp.Body)
 //    json.Unmarshal(body, &data)
 //    fmt.Println(data)
  

 //  defer resp.Body.Close()
 //  body, err := ioutil.ReadAll(resp.Body)

 //  if nil != err {
 //    fmt.Println("errorination happened reading the body", err)
 //    return
 //  }
    
}
 
func twiml(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  body := r.PostForm["Body"][0]
  fmt.Println(body)
  bodyArray := strings.Fields(body)
  fmt.Println(bodyArray)
  slack_name := bodyArray[0]
  slack_channel := bodyArray[1]
  slack_msg := strings.Join(bodyArray[2:], " ")
  //resp, _ := http.Post("https://hooks.slack.com/services/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", "text/json",strings.NewReader("{\"text\":\"Random message from usman using Go\", \"channel\" : \"#general\", \"username\" : \"uehtesham90\", \"icon_url\":\"https://secure.gravatar.com/avatar/595b1952765efa4ff448f55a0e71b49a.jpg?s=72&d=https%3A%2F%2Fslack.global.ssl.fastly.net%2F3654%2Fimg%2Favatars%2Fava_0006-72.png\"}"))
  //resp, _ := http.Post("https://hooks.slack.com/services/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", "text/json",strings.NewReader("{\"text\": \""+slack_msg+"\", \"channel\" : \""+channel+"\", \"username\" : \"uehtesham90\", \"icon_url\":\"https://secure.gravatar.com/avatar/595b1952765efa4ff448f55a0e71b49a.jpg?s=72&d=https%3A%2F%2Fslack.global.ssl.fastly.net%2F3654%2Fimg%2Favatars%2Fava_0006-72.png\"}"))
  resp, _ := http.Post("https://hooks.slack.com/services/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", "text/json",strings.NewReader("{\"text\": \""+slack_msg+"\", \"channel\" : \""+slack_channel+"\", \"username\" : \""+slack_name+"\"}"))
  fmt.Println(resp.Status)
  msg := "Responding..."
  if( resp.StatusCode >= 200 && resp.StatusCode < 300 ) {
  	msg = "Message successfully sent!!!"
  
} else {
  msg = "Message NOT sent!!!"
}
  twiml := TwiML{Message: msg}
  x, err := xml.Marshal(twiml)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
 
  w.Header().Set("Content-Type", "application/xml")
  w.Write(x)
}