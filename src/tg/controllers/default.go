package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
)

// SendController 集成beego控制器
type SendController struct {
	beego.Controller
}

// Message 结构体
type Message struct {
	Message string
	Chatid  string
	Token   string
}

func launchawx(m string) {
	var msg map[string]interface{}
	err := json.Unmarshal([]byte(m), &msg)
	if err != nil {
		fmt.Println(err)
	}
	extt := make(map[string]interface{})
	extt["IP"] = msg["ip"]
	extt["state"] = "present"
	msg["inventory"] = 38
	msg["limit"] = "ADMIN"
	msg["extra_vars"] = extt
	delete(msg, "ip")
	delete(msg, "request")

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("执行AWX job", string(data))
	payload := strings.NewReader(string(data))

	req, err := http.NewRequest("POST", beego.AppConfig.String("awxurl"), payload)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Basic YWRtaW46STAxSWdzVGpxVzRrbHVh")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "c6de34b2-0697-b41f-9658-88abb96462f3")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

}

func (msg Message) httpPost() {
	url := "https://api.telegram.org/bot" + msg.Token + "/sendMessage?" + "chat_id=" + msg.Chatid + "&text=" + msg.Message + "&parse_mode=markdown"
	req, _ := http.NewRequest("POST", url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	launchawx(msg.Message)
}

// Post 请求
func (c *SendController) Post() {
	jsondata := c.Ctx.Input.RequestBody
	var msg map[string]interface{}
	err := json.Unmarshal([]byte(jsondata), &msg)
	if err != nil {
		fmt.Println(err)
	}

	var message Message
	message.Token = fmt.Sprint(msg["token"])
	message.Chatid = fmt.Sprint(msg["chat_id"])
	message.Message = fmt.Sprint(msg["message"])
	go message.httpPost()
}
