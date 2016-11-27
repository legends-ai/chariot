package main

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
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
	lucindaHost = flag.String("lucinda_host", "localhost:45045", "Host of the Lucinda server.")
	charonHost  = flag.String("charon_host", "localhost:5609", "Host of the Charon server.")
	printJSON   = flag.Bool("json", false, "Prints output in JSON")
	runner      = flag.String("runner", "", "Runner")

	locale  = flag.String("locale", "en_US", "Locale")
	matchId = flag.Uint64("matchId", 2300639987, "Match ID to use with Charon")
	version = flag.String("version", "", "Version")
)

func main() {
	flag.Parse()
	logger := logrus.New()

	// set up runners
	r := &runners.Runners{
		Logger: logger,
	}

	if strings.Contains(*runner, "Lucinda") {
		logger.Infof("Connecting to Lucinda at %q", *lucindaHost)
		conn, err := grpc.Dial(*lucindaHost, grpc.WithInsecure())
		if err != nil {
			logger.Fatalf("Could not connect to Lucinda: %v", err)
		}
		r.Lucinda = apb.NewLucindaClient(conn)
	}

	if strings.Contains(*runner, "Charon") {
		logger.Infof("Connecting to Charon at %q", *charonHost)
		conn, err := grpc.Dial(*charonHost, grpc.WithInsecure())
		if err != nil {
			logger.Fatalf("Could not connect to Charon: %v", err)
		}
		r.Charon = apb.NewCharonClient(conn)
	}

	ctx := context.WithValue(context.Background(), "locale", *locale)
	ctx = context.WithValue(ctx, "matchId", *matchId)
	ctx = context.WithValue(ctx, "version", *version)
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
