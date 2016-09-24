package main

import (
	"flag"

	"github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	apb "github.com/asunaio/chariot/gen-go/asuna"
	"github.com/asunaio/chariot/runners"
)

var (
	apolloHost = flag.String("apollo_host", "", "Host of the Apollo server.")
	charonHost = flag.String("charon_host", "", "Host of the Charon server.")

	runner = flag.String("runner", "", "Runner")
)

func main() {
	logger := logrus.New()

	// set up runners
	r := &runners.Runners{
		Logger: logger,
	}

	if *apolloHost != "" {
		conn, err := grpc.Dial(*apolloHost, grpc.WithInsecure())
		if err != nil {
			logger.Fatalf("Could not connect to Apollo: %v", err)
		}
		r.Apollo = apb.NewApolloClient(conn)
	}

	if *charonHost != "" {
		conn, err := grpc.Dial(*charonHost, grpc.WithInsecure())
		if err != nil {
			logger.Fatalf("Could not connect to Charon: %v", err)
		}
		r.Charon = apb.NewCharonClient(conn)
	}

	ctx := context.Background()
	r.Run(ctx, *runner)
}
