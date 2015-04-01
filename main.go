package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type TwiML struct {
	XMLName xml.Name `xml:"Response"`

	Say     string `xml:",omitempty"`
	Play    string `xml:",omitempty"`
	Message string `xml:",omitempty"`
}

func main() {

	populate_database()
	fmt.Println("Database populated")
	http.HandleFunc("/", hello)
	http.HandleFunc("/twiml", twiml)
	http.HandleFunc("/sms", sms)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}

type Request struct {
	Ok      bool
	Members []User
}

func sms(w http.ResponseWriter, r *http.Request) {
	// Set initial variables
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	r.ParseForm()
	fmt.Println(r.PostForm)
	slack_channel := r.PostForm["channel_name"][0]
	text := r.PostForm["text"][0]
	bodyArray := strings.Fields(text)
	to_slack_name := bodyArray[0]
	slack_msg := strings.Join(bodyArray[1:], " ")
	fmt.Println(to_slack_name, slack_msg)
	for key, value := range user_info {
		if value["slack_name"] == to_slack_name {
			// Build out the data for our message
			v := url.Values{}
			if key[:2] != "+1" {
				key = "+1" + key
			}
			v.Set("To", key)
			v.Set("From", os.Getenv("TWILIO_NUMBER"))
			v.Set("Body", "#"+slack_channel+" "+slack_msg)
			rb := *strings.NewReader(v.Encode())

			// Create client
			client := &http.Client{}

			req, _ := http.NewRequest("POST", urlStr, &rb)
			req.SetBasicAuth(accountSid, authToken)
			req.Header.Add("Accept", "application/json")
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			// Make request
			resp, _ := client.Do(req)
			fmt.Println(resp.Status)
		}
	}

}

var user_info = make(map[string]map[string]string)

func populate_database() {

	data := Request{}
	token := "xoxp-2757127568-2813046014-4209073904-35634c"
	resp, _ := http.Get("https://slack.com/api/users.list?token=" + token)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &data)
	for _, j := range data.Members {
		if j.Profile.Phone != "" {
			temp := make(map[string]string)
			temp["slack_id"] = j.Id
			temp["slack_name"] = j.Name
			temp["slack_email"] = j.Profile.Email
			temp["slack_profilepic"] = j.Profile.Image_48
			user_info[j.Profile.Phone] = temp
		}
	}
	fmt.Println(user_info)
}

func twiml(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body := r.PostForm["Body"][0]
	from := r.PostForm["From"][0]
	if from[:2] == "+1" {
		from = from[2:]
	}
	if val, ok := user_info[from]; ok {
		bodyArray := strings.Fields(body)
		slack_channel := bodyArray[0]
		slack_name := val["slack_name"]
		slack_profilepic := val["slack_profilepic"]
		slack_msg := strings.Join(bodyArray[1:], " ")
		resp, _ := http.Post("https://hooks.slack.com/services/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", "text/json", strings.NewReader("{\"text\": \""+slack_msg+"\", \"channel\" : \""+slack_channel+"\", \"username\" : \""+slack_name+"\", \"icon_url\":\""+slack_profilepic+"\"}"))
		fmt.Println(resp.Status)
		msg := "Responding..."
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
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
	} else {
		fmt.Println("Unrecognized number", from)
	}

}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Welcome to SlackMS!!!")
}

type User struct {
	Id        string
	Name      string
	Deleted   bool
	Color     string
	Profile   profile
	Is_Admin  bool
	Is_Owner  bool
	Has_2fa   bool
	Has_Files bool
}

type profile struct {
	First_Name string
	Last_Name  string
	Real_Name  string
	Email      string
	Skype      string
	Phone      string
	Image_24   string
	Image_32   string
	Image_48   string
	Image_72   string
	Image_192  string
}
