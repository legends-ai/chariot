package main

import (
	"flag"

	"github.com/Sirupsen/logrus"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	apb "github.com/asunaio/chariot/gen-go/asuna"
)

var (
	apolloHost = flag.String("apollo_host", "127.0.0.1:4834", "Host of the Apollo server.")
)

func main() {
	logger := logrus.New()

	conn, err := grpc.Dial(*apolloHost)
	if err != nil {
		logger.Fatalf("Could not connect to Apollo: %v", err)
	}
	client := apb.NewApolloClient(conn)

	ctx := context.Background()
	champion, err := client.GetChampion(ctx, &apb.GetChampionRequest{
		ChampionId: 64, // this is Lee Sin
		Patch: &apb.PatchRange{
			Min: "6.16",
			Max: "6.18",
		},
		// match everything
		Tier: &apb.TierRange{
			Min: 0x0000,
			Max: 0x1000,
		},
		Region: apb.Region_NA,
		Role:   apb.Role_JUNGLE,
	})
	if err != nil {
		logger.Fatalf("Could not get champion: %v", err)
	}

	spew.Dump(champion)
}
