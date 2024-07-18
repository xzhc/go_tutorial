package main

import "net"

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

//创建一个用户的API
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

//用户上线功能
func (this *User) Online() {
	//用户上线后加入到onlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//广播当前用户上线消息
	this.server.BroadCast(this, "go online")
}

//用户下线功能
func (this *User) Offline() {
	//用户下线将用户从OnlineMap删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	//广播当前用户下线消息
	this.server.BroadCast(this, "offline")
}

//用户处理消息
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//查询当前在线用户有哪些
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "online...\n"
			this.sendMsg(onlineMsg)
		}

		this.server.mapLock.Unlock()
	} else {
		this.server.BroadCast(this, msg)
	}

}

//给当前user对应的客户端发送信息
func (this *User) sendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

//监听当前User channel的方法，一旦有消息就发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
