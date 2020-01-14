package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/astaxie/beego"
)

type SendController struct {
	beego.Controller
}

type Message struct {
	Message string
	Chat_id string
	Token   string
}

func (msg Message) httpPost() {

	url := "https://api.telegram.org/bot" + msg.Token + "/sendMessage?" + "chat_id=" + msg.Chat_id + "&text=" + msg.Message + "&parse_mode=markdown"
	req, _ := http.NewRequest("POST", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func reflectTest(b interface{}) string {
	var text string
	rVal := reflect.ValueOf(b).Interface()
	msg, ok := rVal.(map[string]interface{})
	if ok {
		for k, v := range msg {
			text += fmt.Sprint("`", k, ":` ", "*", v, "*", "%0a")
		}
	}
	return text
}

func (c *SendController) Post() {
	json_data := c.Ctx.Input.RequestBody
	var msg map[string]interface{}
	err := json.Unmarshal([]byte(json_data), &msg)
	if err != nil {
		fmt.Println(err)
	}

	var message Message
	message.Token = fmt.Sprint(msg["token"])
	message.Chat_id = fmt.Sprint(msg["chat_id"])
	message.Message = reflectTest(msg["message"])
	go message.httpPost()
}
