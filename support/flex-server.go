package support

import (
	"diswares.com.journal/config"
	"fmt"
	"io"
	"net"
)

type ServerHooks struct {
	Started func(ctx ServerContext)
	// client connection is successful
	Accepted func(client net.Conn, ctx ServerContext)
	Received func(buffer []byte, ctx ServerContext, conn net.Conn)
	Ready    func(ctx ServerContext)
	Deleted  func(client net.Conn, ctx ServerContext)
}

type ServerContext struct {
	socket      net.Listener
	Scanner     config.Scanner
	Connections map[string]net.Conn
	Property    config.Exchange
}

func (ctx ServerContext) Write(buffer []byte) {
	if ctx.Connections != nil {
		conn := ctx.Connections
		for _, client := range conn {
			_, _ = client.Write(buffer)
		}
	}
}

func (ctx ServerContext) Shutdown() {
	Catch(ctx.socket.Close(), func() {
		fmt.Printf("\nINFO | server exited \n ")
	})
}

func Bootstrap(server config.Exchange, hooks ServerHooks) ServerContext {

	if server.Host == "localhost" {
		ip, err := LocalIP()
		Catch(err, func() {
			server.Host = ip
		})
	}

	socket, err := net.Listen(server.Protocol, fmt.Sprintf("%s:%d", server.Host, server.Port))
	context := ServerContext{
		socket:   socket,
		Scanner:  config.ScannerInstance(),
		Property: server,
	}
	// return when the connection fails
	if nil == socket {
		fmt.Printf("\nERROR | Fail to start server, please check whether the parameters are normal\n")
		return context
	}

	defer func() {
		Catch(socket.Close(), func() {
			fmt.Printf("\nINFO | server exited \n ")
		})
	}()

	// no error
	Catch(err, func() {
		if hooks.Started != nil {
			context.Connections = map[string]net.Conn{}
			hooks.Started(context)
		}

		if hooks.Ready != nil {
			go hooks.Ready(context)
		}

		for {
			conn, e := socket.Accept()

			buffer := make([]byte, 1024)

			Catch(e, func() {
				context.Connections[conn.RemoteAddr().String()] = conn
				if hooks.Accepted != nil {
					go hooks.Accepted(conn, context)

					go func() {
						// release
						defer func() {
							_ = conn.Close()
						}()

						for {
							size, err := conn.Read(buffer)
							if err == io.EOF {
								delete(context.Connections, conn.RemoteAddr().String())
								if hooks.Deleted != nil {
									hooks.Deleted(conn, context)
								}
								break
							} else {
								if hooks.Received != nil && size > 0 {
									hooks.Received(buffer[0:size], context, conn)
								}
							}
						}
					}()

				}
			})
		}
	})

	return context
}
