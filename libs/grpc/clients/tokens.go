package clients

import (
	"fmt"
	"github.com/satont/twir/libs/grpc/generated/tokens"
	"log"

	"github.com/satont/twir/libs/grpc/servers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
)

func NewTokens(env string) tokens.TokensClient {
	serverAddress := createClientAddr(env, "tokens", servers.TOKENS_SERVER_PORT)

	conn, err := grpc.Dial(
		serverAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingConfig": [{"%s":{}}]}`, roundrobin.Name)),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := tokens.NewTokensClient(conn)

	return c
}
