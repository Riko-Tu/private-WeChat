package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

//客户端
type Conn struct {
	Client *websocket.Conn
	//并行转串行数据
	Name      string
	Addr      string
	DataQueue chan []byte
}

//message
type Msg struct {
	//发送者
	FromUser string
	//消失类型
	MsgType string
	//接收者
	ToRecver string
}


//上线通知
func onlineAdvice() {

}
//下线通知
func offlineAdvice()  {

}


//广播消息
func () {
	
}

//群聊发送
func AllSendMsg()  {
	
}

//私聊发送
func OneSendMsg(userId string, msg []byte) {
	once.Lock()
	conn, ok := clientMap[userId]
	once.Unlock()
	if ok {
		conn.DataQueue <- msg
	}
}

//读写映射表锁
var once sync.Mutex

//用户连接映射表
var clientMap = make(map[string]*Conn)

func Chat(ctx *gin.Context) {
	uid := ctx.Request.RemoteAddr
	//协议升级
	Upgrader := &websocket.Upgrader{}
	//每个用户为经过这个函数
	conn, err := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.String(http.StatusOK, err.Error())
	}
	UserClient := &Conn{
		Client:    conn,
		DataQueue: make(chan []byte, 50),
	}
	//把uid与uid的连接放到映射表里
	once.Lock()
	clientMap[fmt.Sprintf("%s", uid)] = UserClient
	once.Unlock()

	go func() {
		for {
			//读取消息
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			//写入消息
			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}
		}
	}()

}
