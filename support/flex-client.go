package support

import (
	"fmt"
	"github.com/4everlynn/journal/config"
	"io"
	"net"
)

type ClientHooks struct {
	Started   func()
	Connected func(server net.Conn, ctx ClientContext)
	Ready     func(ctx ClientContext)
	Received  func([]byte)
}

type ClientContext struct {
	socket   net.Conn
	Scanner  config.Scanner
	Property config.Exchange
}

func (ctx ClientContext) Write(buffer []byte) {
	_, _ = ctx.socket.Write(buffer)
}

func (ctx ClientContext) Shutdown() {
	Catch(ctx.socket.Close(), func() {
		fmt.Printf("\nINFO | client exited \n ")
	})
}

func Client(client config.Exchange, hooks ClientHooks) ClientContext {
	socket, err := net.Dial(client.Protocol, fmt.Sprintf("%s:%d", client.Host, client.Port))
	ctx := ClientContext{
		socket:   socket,
		Scanner:  config.ScannerInstance(),
		Property: client,
	}
	// return when the connection fails
	if err != nil {
		fmt.Printf("\nERROR | Fail to start client, please check whether the parameters are normal\n")
		return ctx
	}

	defer func() {
		Catch(socket.Close(), func() {
			fmt.Printf("\nINFO | client exited \n ")
		})
	}()

	// no error
	Catch(err, func() {
		if hooks.Started != nil {
			hooks.Started()
		}

		if hooks.Ready != nil {
			go hooks.Ready(ctx)
		}

		buffer := make([]byte, 1024)
		for {
			size, e := socket.Read(buffer)

			if e == io.EOF {
				break
			} else {
				if hooks.Received != nil {
					hooks.Received(buffer[0:size])
				}
			}
		}
	})

	return ctx
}
