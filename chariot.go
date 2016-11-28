package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"

	apb "github.com/asunaio/chariot/gen-go/asuna"
	"github.com/asunaio/chariot/runners"
)

var (
	charonHost  = flag.String("charon_host", "localhost:5609", "Host of the Charon server.")
	lucindaHost = flag.String("lucinda_host", "localhost:45045", "Host of the Lucinda server.")
	vulgateHost = flag.String("vulgate_host", "localhost:6205", "Host of the Vulgate server.")

	runner    = flag.String("runner", "", "Runner")
	printJSON = flag.Bool("json", false, "Prints output in JSON")

	locale        = flag.String("locale", "en_US", "Locale")
	region        = flag.String("region", "NA", "Region")
	role          = flag.String("role", "MID", "Role")
	version       = flag.String("version", "6.23.1", "Version")
	vulgateFormat = flag.String("vulgate_format", "BASIC", "Vulgate response format")

	// Default is Ahri,Annie
	championId = flag.String("championId", "103,1", "Champion ID. Pass multiple champions delimited by a ,")
	matchId    = flag.Uint64("matchId", 2300639987, "Match ID to use with Charon")
	summonerId = flag.String("summonerId", "29236065,24575247", "Summoner ID to use with Charon. Pass multiple summoners delimited by a ,")
)

func main() {
	flag.Parse()
	logger := logrus.New()

	r := setupRunners(logger)
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

func setupRunners(logger *logrus.Logger) *runners.Runners {
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

	if strings.Contains(*runner, "Vulgate") {
		logger.Infof("Connecting to Vulgate at %q", *vulgateHost)
		conn, err := grpc.Dial(*vulgateHost, grpc.WithInsecure())
		if err != nil {
			logger.Fatalf("Could not connect to Vulgate: %v", err)
		}
		r.Vulgate = apb.NewVulgateClient(conn)
	}

	r.Flags = runners.Flags{
		Locale:        apb.Locale(apb.Locale_value[*locale]),
		Region:        apb.Region(apb.Region_value[*region]),
		Role:          apb.Role(apb.Role_value[*role]),
		ChampionId:    stringToUint32Slice(*championId),
		MatchId:       *matchId,
		SummonerId:    stringToUint64Slice(*summonerId),
		Version:       *version,
		VulgateFormat: *vulgateFormat,
	}

	return r
}

func stringToUint32Slice(s string) []uint32 {
	out := []uint32{}
	for _, token := range strings.Split(s, ",") {
		// TODO return error on failure
		if uint, err := strconv.ParseUint(token, 10, 32); err == nil {
			out = append(out, uint32(uint))
		}
	}
	return out
}

func stringToUint64Slice(s string) []uint64 {
	out := []uint64{}
	for _, token := range strings.Split(s, ",") {
		// TODO return error on failure
		if uint, err := strconv.ParseUint(token, 10, 64); err == nil {
			out = append(out, uint)
		}
	}
	return out
}
