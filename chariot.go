package main

import (
	"bytes"
	"flag"
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	apb "github.com/asunaio/chariot/gen-go/asuna"
	"github.com/asunaio/chariot/runners"
)

var (
	apolloHost = flag.String("apollo_host", "", "Host of the Apollo server.")
	charonHost = flag.String("charon_host", "", "Host of the Charon server.")

	printJSON = flag.Bool("json", false, "Prints output in JSON")
	runner    = flag.String("runner", "", "Runner")
)

func main() {
	flag.Parse()
	logger := logrus.New()

	// set up runners
	r := &runners.Runners{
		Logger: logger,
	}

	if *apolloHost != "" {
		logger.Infof("Connecting to Apollo at %q", *apolloHost)
		conn, err := grpc.Dial(*apolloHost, grpc.WithInsecure())
		if err != nil {
			logger.Fatalf("Could not connect to Apollo: %v", err)
		}
		r.Apollo = apb.NewApolloClient(conn)
	}

	if *charonHost != "" {
		logger.Infof("Connecting to Charon at %q", *charonHost)
		conn, err := grpc.Dial(*charonHost, grpc.WithInsecure())
		if err != nil {
			logger.Fatalf("Could not connect to Charon: %v", err)
		}
		r.Charon = apb.NewCharonServiceClient(conn)
	}

	ctx := context.Background()
	logger.Infof("Running runner %q", *runner)
	start := time.Now()

	var out bytes.Buffer
	msg := r.Run(ctx, *runner)
	if *printJSON {
		if err := (&jsonpb.Marshaler{
			EnumsAsInts:  false,
			EmitDefaults: true,
			OrigName:     false,
		}).Marshal(&out, msg); err != nil {
			r.Logger.Fatalf("Could not marshal msg: %v", err)
		}
	} else {
		if err := proto.MarshalText(&out, msg); err != nil {
			r.Logger.Fatalf("Could not marshal msg: %v", err)
		}
	}
	fmt.Println(out.String())

	logger.Infof("Completed; took %s", time.Now().Sub(start))
}
