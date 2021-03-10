package cmd

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/7cav/api/proto"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/grpclog"
	"strconv"
)

var exampleCmd = &cobra.Command{
	Use:   "getProfile",
	Short: "example of a golang client to use the API method: getProfile",
	Run: func(cmd *cobra.Command, args []string) {
		token := "NNCCz5jHXWCPpd07Bw0l92Hn8VYdeLEvkyIprbpjpAUHFSrRlG"
		rpcCreds := oauth.NewOauthAccess(&oauth2.Token{AccessToken: token})
		creds, err:= credentials.NewClientTLSFromFile("out/localhost.crt", "")
		if err != nil {
			panic(err)
		}

		fmt.Println("gathering creds grpc")
		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(creds),
			grpc.WithPerRPCCredentials(rpcCreds),
			grpc.WithBlock(),
		}

		fmt.Println("dialing grpc")
		conn, err := grpc.Dial("127.0.0.1:1443", opts...)

		fmt.Println("dialed...")
		if err != nil {
			grpclog.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()

		fmt.Println("creating client")
		client := proto.NewMilpacsClient(conn)

		if len(args) != 1{
			grpclog.Fatalln("must supply id to request as argument")
		}

		id, _ := strconv.ParseUint(args[0], 10, 64)

		fmt.Println("Searching for client with ID:", id)
		msg, err := client.GetProfile(context.Background(), &proto.ProfileRequest{UserId: id})
		if err != nil {
			grpclog.Fatalf("fail to get profile: %v", err)
		}
		fmt.Println(msg)
	},
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("certs/server-cert.pem", "certs/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

func init() {
	rootCmd.AddCommand(exampleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exampleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exampleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
