package main

func main() {
	server := NewServer("127.0.0.1", 8888)

	server.Start()
}

//启动server命令（windows）
//"go build -o server.exe main.go server.go"
//"./server"
//"curl 127.0.0.1:8888"
