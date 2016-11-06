package runners

import (
	"github.com/Sirupsen/logrus"
	apb "github.com/asunaio/chariot/gen-go/asuna"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

// Runners runs some shit
type Runners struct {
	Logger *logrus.Logger

	Apollo apb.ApolloClient
	Charon apb.CharonClient
}

func (r *Runners) Run(ctx context.Context, runner string) proto.Message {
	switch runner {

	case "Apollo::GetChampion":
		return r.ApolloChampion(ctx)

	case "Apollo::GetMatchup":
		return r.ApolloMatchup(ctx)

	case "Apollo::GetMatchSum":
		return r.ApolloMatchSum(ctx)

	case "Charon::GetMatch":
		return r.CharonMatch(ctx)

	case "Charon::GetDominionMatch":
		return r.CharonDominionMatch(ctx)

	case "Charon::GetMatchList":
		return r.CharonMatchList(ctx)

	case "Charon::GetRankings":
		return r.CharonRankings(ctx)

	case "Charon::GetStaticChampions":
		return r.CharonStaticChampions(ctx)

	default:
		r.Logger.Fatalf("Unknown runner %q", runner)
	}
	return nil
}
