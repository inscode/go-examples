package znet

import (
	"x-zinx/ziface"
	"fmt"
	"net"
	"x-zinx/utils"
)

//IServer的接口实现，定义一个Server的服务器模块
type Server struct {
	//服务器名称
	Name string

	//服务器绑定的ip版本
	IPVersion string

	//监听的IP
	IP string

	//服务器监听的端口
	Port int

	//当前的server 添加一个router, server注册的连接对应的处理业务
	//Router ziface.IRouter
	MsgHandler ziface.IMsgHandler
}

//开启服务器
func (s *Server) Start() {
	fmt.Printf("[zinx]server name : %s; listenner at ip : %s, port : %d\n",
		utils.GlobalObject.Name,
		utils.GlobalObject.Host,
		utils.GlobalObject.TCPPort,
	)

	//异步处理
	go func() {
		//1：获取tcp的addr---解析TCP
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != err {
			fmt.Println("ResolveTCPAddr err: ", err)
			return
		}

		//2：监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("ListenTCP err: ", err)
			return
		}
		fmt.Println("start zinx server success", listener.Addr(), " server name ", s.Name)

		//3：阻塞的（等待客户端链接，处理客户端连接（读写））
		var connID uint32
		connID = 0
		for {
			//如果有客户端连接请求过来，阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("AcceptTCP err：", err)
				continue
			}

			//客户端已经与服务端建立连接，处理业务

			//将处理新连接的业务方法 和 conn 进行绑定，得到我们的连接模块
			dealConn := NewConnection(conn, connID, s.MsgHandler)
			connID++

			go dealConn.Start()
		}
	}()
}

//停止服务器
func (s *Server) Stop() {
	//todo 将一些服务器的资源，状态或者一些开辟的连接信息进行停止处理（回收）
}

//运行服务器
func (s *Server) Serve() {

	//启动server的服务功能
	s.Start()

	//todo 处理一些启动服务器之后的额外业务

	//阻塞状态
	select {}
}

//运行服务器
func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Println("add router success")
}

//初始化Server模块的方法
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TCPPort,
		MsgHandler: NewMsgHandler(),
	}
	return s
}
