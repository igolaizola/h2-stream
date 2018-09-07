package cli

import (
	"github.com/igarciaolaizola/h2-stream/internal/client"
	"github.com/igarciaolaizola/h2-stream/internal/server"
	"github.com/spf13/cobra"
)

// NewCommand create and returns the root cli command
func NewCommand() *cobra.Command {
	tunnelCmd := &cobra.Command{
		Use:   "tunnel",
		Short: "CLI tool aimed to handle http2 streams",
		Long:  `CLI tool aimed to handle http2 streams: client or server side.`,
	}

	tunnelCmd.AddCommand(newServerCommand())
	tunnelCmd.AddCommand(newClientCommand())
	return tunnelCmd
}

func newServerCommand() *cobra.Command {
	var addr, cert, key string
	var tls bool

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Launches the http2 stream echo server",
		Long:  `Launches the http2 stream server that echoes any incoming data.`,
		RunE: func(c *cobra.Command, args []string) error {
			return server.Run(addr, tls, cert, key)
		},
	}

	cmd.Flags().StringVar(&addr, "addr", "localhost:8080", "http listening address")
	cmd.Flags().BoolVar(&tls, "tls", true, "enable tls (https)")
	cmd.Flags().StringVar(&cert, "cert", "certs/cert.pem", "tls certificate file")
	cmd.Flags().StringVar(&key, "key", "certs/key.pem", "tls key file")
	return cmd
}

func newClientCommand() *cobra.Command {
	var cfg client.Config

	cmd := &cobra.Command{
		Use:   "cli",
		Short: "Launches the http2 stream client",
		Long:  `Launches the http2 stream client that sends any writen to stdin.`,
		RunE: func(c *cobra.Command, args []string) error {
			client, err := client.New(cfg)
			if err != nil {
				return err
			}
			return client.Run()
		},
	}

	cmd.Flags().StringVar(&cfg.Addr, "addr", "https://localhost:8080", "http server address")
	cmd.Flags().StringVar(&cfg.Method, "method", "POST", "http method (GET, POST, PUT, DELETE...)")
	cmd.Flags().BoolVar(&cfg.Insecure, "insecure", false, "skip tls validation")
	cmd.Flags().StringVar(&cfg.Data, "data", "", "initial body data to be sent")
	cmd.Flags().StringArrayVar(&cfg.Headers, "header", nil, "custom headers, ex.: \"Content-Type: application/json\"")
	return cmd
}
