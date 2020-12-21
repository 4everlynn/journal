package cmd

import (
	"fmt"
	"github.com/4everlynn/journal/config"
	"github.com/4everlynn/journal/support"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"net"
	"strings"
)

func SysMsg(msg string) {
	fmt.Printf("\r\n\033[s\033[1A\033[K\033[u%s%s\n",
		color.FgGray.Render("Journal <- "),
		color.FgGray.Render(msg))
}

func ReceivedMsg(msg string) {
	fmt.Printf("\r\n\033[s\033[1A\033[K\033[u%s\n",
		color.FgGreen.Render(msg))
}

func UserMsg(msg string) {
	fmt.Printf("\033[s\033[1A\033[K\033[u%s\n", msg)
}

// exchangeCmd represents the exchange command
var exchangeCmd = &cobra.Command{
	Use:   "exchange",
	Short: "establish lan connection",
	Long: `If you want to initiate a communication, you need to start it as a server, 
and other people will establish a connection with you as a client. 
If you only need to connect, run it as a client.`,
	Run: func(cmd *cobra.Command, args []string) {
		isServer, _ := cmd.Flags().GetBool("server")
		isClient, _ := cmd.Flags().GetBool("client")
		port, _ := cmd.Flags().GetInt("port")
		host, _ := cmd.Flags().GetString("host")
		name, _ := cmd.Flags().GetString("name")

		if strings.Trim(name, " ") == "" {
			fmt.Println(color.FgRed.Render("journal error: server/client needs name"))
			return
		}

		properties := config.Exchange{
			Port:     port,
			Host:     host,
			Protocol: "tcp",
			Name:     name,
		}

		if isClient {

			support.Client(properties, support.ClientHooks{
				Started: func() {
					fmt.Printf("\033c")
				},
				Ready: func(ctx support.ClientContext) {
					for {
						print(color.FgWhite.Render(fmt.Sprintf("%s# ", ctx.Property.Name)))
						words := ctx.Scanner.ReadLine()
						ctx.Write([]byte(ctx.Property.Name + " <- " + words))
						UserMsg(color.FgYellow.Render(ctx.Property.Name+" -> ") + color.FgYellow.Render(words))
					}
				},
				Received: func(bytes []byte) {
					ReceivedMsg(fmt.Sprintf("%s", string(bytes)))
					print(color.FgWhite.Render(fmt.Sprintf("%s# ", properties.Name)))
				},
			})

		} else if isServer {
			support.Bootstrap(properties, support.ServerHooks{
				Started: func(ctx support.ServerContext) {
					//fmt.Printf("\033c")
				},
				Ready: func(ctx support.ServerContext) {
					for {
						print(color.FgWhite.Render(fmt.Sprintf("%s$ ", ctx.Property.Name)))
						words := ctx.Scanner.ReadLine()
						ctx.Write([]byte(ctx.Property.Name + " <- " + words))
						UserMsg(color.FgYellow.Render(ctx.Property.Name+" -> ") + color.FgYellow.Render(words))
					}
				},
				Deleted: func(client net.Conn, ctx support.ServerContext) {
					SysMsg("IP " + client.RemoteAddr().String() + " LEAVE")
				},
				Received: func(bytes []byte, ctx support.ServerContext, conn net.Conn) {
					ReceivedMsg(fmt.Sprintf("%s", string(bytes)))
					// dispatch msg to other clients
					for _, client := range ctx.Connections {
						if client.RemoteAddr().String() != conn.RemoteAddr().String() {
							_, _ = client.Write(bytes)
						}
					}
					print(color.FgWhite.Render(fmt.Sprintf("%s$ ", properties.Name)))
				},
				Accepted: func(client net.Conn, ctx support.ServerContext) {
					SysMsg("IP " + client.RemoteAddr().String() + " CONNECTED")
					print(color.FgWhite.Render(fmt.Sprintf("%s$ ", ctx.Property.Name)))
				},
			})
		}

	},
}

func init() {
	rootCmd.AddCommand(exchangeCmd)
	exchangeCmd.Flags().BoolP("server", "s", true, "run as server")
	exchangeCmd.Flags().BoolP("client", "c", false, "run as client")
	exchangeCmd.Flags().StringP("host", "H", "localhost", "specify server/client host address")
	exchangeCmd.Flags().IntP("port", "p", 9651, "specify server/client port")
	exchangeCmd.Flags().StringP("name", "n", "", "specify server/client temporary name")
	exchangeCmd.Flags().StringP("auth", "a", "", "specify connection authorization code")
}
