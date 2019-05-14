package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

// tcp server 服务端代码

//服务端 tcpListener.AcceptTCP() 接受一个客户端连接请求，通过go tcpPipe(tcpConn) 开启一个新协程来管理这对连接。
// 在func tcpPipe(conn *net.TCPConn) 中，处理服务端和客户端数据的交换，在这段代码for中，通过 bufio.NewReader 读取客户端发送过来的数据。
func main() {

	var tcpAddr *net.TCPAddr

	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:999")

	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	defer tcpListener.Close()

	fmt.Println("Server ready to read ...")
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}

}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()

	defer func() {
		fmt.Println(" Disconnected : " + ipStr)
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	i := 0

	for {
		message, err := reader.ReadString('\n') //将数据按照换行符进行读取。
		if err != nil || err == io.EOF {
			break
		}

		fmt.Println(string(message))

		time.Sleep(time.Second * 3)

		msg := time.Now().String() + conn.RemoteAddr().String() + " Server Say hello! \n"
		b := []byte(msg)

		conn.Write(b)
		i++

		if i > 10 {
			break
		}
	}
}

//客户端net.DialTCP("tcp", nil, tcpAddr) 向服务端发起一个连接请求，调用onMessageRecived(conn)，处理客户端和服务端数据的发送与接收。
// 在func onMessageRecived(conn *net.TCPConn) 中，通过 bufio.NewReader 读取客户端发送过来的数据。
//客户端代码

//func main() {
//	var tcpAddr *net.TCPAddr
//	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:999")
//
//	conn, err := net.DialTCP("tcp", nil, tcpAddr)
//	if err != nil {
//		fmt.Println("Client connect error ! " + err.Error())
//		return
//	}
//
//	defer conn.Close()
//
//	fmt.Println(conn.LocalAddr().String() + " : Client connected!")
//
//	onMessageRecived(conn)
//}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	b := []byte(conn.LocalAddr().String() + " Say hello to Server... \n")
	conn.Write(b)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println("ReadString")
		fmt.Println(msg)

		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Second * 2)

		fmt.Println("writing...")

		b := []byte(conn.LocalAddr().String() + " write data to Server... \n")
		_, err = conn.Write(b)

		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
