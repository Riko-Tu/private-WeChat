package controller

import (
	"encoding/json"
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

type Server struct {
	CastMsg chan string
}

var server = Server{CastMsg: make(chan string, 1024)}

const (
	CMDTYPE_ALL   = 20
	CMDTYPE_ONE   = 21
	CMDTYPE_HEART = 22
)

const ()

//message
type Msg struct {
	//消息id
	ID int64 `json:"id"`
	//发送者
	FromUser string `json:"from_user"`
	//消失类型:私聊消息 、 群聊消息 、 心跳消息
	CmdType int `json:"cmd_type"`
	//接收者
	ToRecver string `json:"to_recver"`
	// 消息内容类型：文本样式、图文消息、语言消息、图片消息、红包消息、emoj消息、超链接消息
	MdiaType int `json:"mdia_type"`
	//文本内容
	MsgContent string `json:"msg_content"`
	//图片内容
	PicContent string `json:"pic_content"`
	//URL链接
	Url string `json:"url"`
	//简单描述
	Memo string `json:"memo"`
	//附加消息
}

//接收当前客户端的信息
func recvPoce(conn *websocket.Conn, name string) {
	m := new(Msg)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			//用户下线
			delete(clientMap, name)
			onlineAdvice(name + "已下线")
			return
		}
		fmt.Println(string(msg))
		err = json.Unmarshal(msg, m)
		if err != nil {
			log.Println("msg解析错误：" + err.Error())
			break
		}
		switch m.CmdType {
		case CMDTYPE_ONE:
			OneSendMsg(m.ToRecver, msg)
		case CMDTYPE_ALL:

		case CMDTYPE_HEART:

		}

	}
}

//遍历该客户端的所有好友，上线通知
func onlineAdvice(msg string) {
	server.CastMsg <- msg
}

//下线通知
func offlineAdvice() {

}

//广播消息
func broadCast() {
	for {
		select {
		case msg := <-server.CastMsg:
			for _, conn := range clientMap {
				conn.DataQueue <- []byte(msg)
			}
		}
	}
}

//群聊发送
func AllSendMsg() {

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

func newClient(uid string, conn *websocket.Conn) *Conn {
	client := &Conn{
		Client:    conn,
		Name:      uid,
		Addr:      conn.RemoteAddr().String(),
		DataQueue: make(chan []byte, 50),
	}
	go client.lisentDate()
	return client

}

func (c *Conn) lisentDate() {
	for {
		select {
		case msg := <-c.DataQueue:
			err := c.Client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("写入错误:" + err.Error())
			}
		}
	}
}

//读写映射表锁
var once sync.Mutex

//用户连接映射表
var clientMap = make(map[string]*Conn)

func Chat(ctx *gin.Context) {
	uid := ctx.Query("uid")
	//协议升级
	fmt.Println(uid + "已上线")
	Upgrader := &websocket.Upgrader{}
	//每个用户为经过这个函数
	conn, err := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.String(http.StatusOK, err.Error())
	}
	client := newClient(uid, conn)

	//启动广播监听
	go broadCast()
	onlineAdvice(uid + "已上线")
	//把uid与uid的连接放到映射表里
	once.Lock()
	clientMap[fmt.Sprintf("%s", uid)] = client
	once.Unlock()

	go recvPoce(conn, uid)
	select {}
}
