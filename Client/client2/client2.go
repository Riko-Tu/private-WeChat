package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
	"turan.com/WeChat-Private/controller"
)

func main() {
	message := make(chan []byte, 1024)
	go func() {
		for true {
			var msgConent string
			fmt.Scan(&msgConent)
			msgC := controller.Msg{
				ID:         1,
				FromUser:   "23",
				CmdType:    21,
				ToRecver:   "31",
				MdiaType:   0,
				MsgContent: msgConent,
				PicContent: "",
				Url:        "",
				Memo:       "",
			}
			marshal, err := json.Marshal(msgC)
			if err != nil {
				panic("序列化失败:" + err.Error())
			}
			message <- marshal
		}
	}()

	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8080", Path: "/user/ws", RawQuery: "uid=23"}
	dial, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err.Error())
	}

	go func() {
		for {
			_, message, err := dial.ReadMessage()
			if err != nil {
				log.Println(err.Error())
				return
			}
			fmt.Println(string(message))
		}
	}()

	interrupt := make(chan bool, 1)

	for {
		select {
		case msg := <-message:
			err := dial.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				println(err.Error())
				break
			}
			time.Sleep(1 * time.Second)
		case <-interrupt:
			err := dial.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println(err.Error())
			}
			break
			dial.Close()
		default:
			time.Sleep(time.Second)
		}
	}
}
