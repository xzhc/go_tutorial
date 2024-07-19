package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	go user.ListenMessage()
	return user
}

// 用户上线功能
func (this *User) Online() {
	//用户上线后加入到onlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//广播当前用户上线消息
	this.server.BroadCast(this, "go online")
}

// 用户下线功能
func (this *User) Offline() {
	//用户下线将用户从OnlineMap删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	//广播当前用户下线消息
	this.server.BroadCast(this, "offline")
}

// 用户处理消息
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//查询当前在线用户有哪些
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "online...\n"
			this.sendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		newName := msg[7:]
		//判断name是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.sendMsg("the username is used \n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.sendMsg("You have updated username:" + this.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		//消息格式：to|zhangsan|content

		//1.获取用户的客户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.sendMsg("message formatation is uncorrect, please use 'to | zhangsan | content'. \n")
			return
		}
		//2.根据用户名得到对方的User对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.sendMsg("the username isn't exist \n")
			return
		}

		//3.获取消息内容
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.sendMsg("no content,please resend!")
			return
		}

		remoteUser.sendMsg(this.Name + "to you:" + content)

	} else {
		this.server.BroadCast(this, msg)
	}

}

// 给当前user对应的客户端发送信息
func (this *User) sendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 监听当前User channel的方法，一旦有消息就发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
