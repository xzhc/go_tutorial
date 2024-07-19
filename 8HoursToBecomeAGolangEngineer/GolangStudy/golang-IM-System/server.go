package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

// 定义server结构体
type Server struct {
	Ip   string
	Port int

	//在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//消息广播的channel
	Message chan string
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

func (this *Server) Handler(conn net.Conn) {
	//...当前链接的业务
	//fmt.Println("链接建立成功")

	user := NewUser(conn, this)
	user.Online()

	//监听用户是否活跃的channel
	isLive := make(chan bool)

	//接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read Err:", err)
			}

			//提取用户信息
			msg := string(buf[:n-1])
			//得到的消息进行广播
			user.DoMessage(msg)
			//用户的任意消息代表当前用户是活跃的
			isLive <- true
		}
	}()
	//当前handler堵塞
	select {
	case <-isLive:
	//当前用户是活跃的，应该重置定时器
	//不做任何事情，为了激活select，更新下面的定时器

	case <-time.After(time.Second * 10):
		//超时将当前user强制关闭
		user.sendMsg("you are offline because time exceed!")
		//销毁用的资源
		close(user.C)
		//关闭链接
		conn.Close()
		//退出当前handler
		return
	}
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

// 监听Message广播消息的goroutine,一旦有消息就发送给全部在线的user
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		//将msg发送给全部在线的User
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 启动服务器的接口
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
	}
	//close socket
	defer listener.Close()

	//启动监听Messenager的goroutine
	go this.ListenMessager()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error:", err)
			continue
		}

		//do handler
		go this.Handler(conn)
	}

}
