package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

func main() {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8080", Path: "/user/ws"}
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
	message := make(chan int, 1024)
	interrupt := make(chan bool, 1)
	go func() {
		for i := 0; i < 20; i++ {
			message <- i
			time.Sleep(2 * time.Second)
		}
		interrupt <- true
	}()

	for {
		select {
		case msg := <-message:
			err := dial.WriteMessage(websocket.TextMessage, []byte("我是客户端2"+fmt.Sprintf("%d", msg)))
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
