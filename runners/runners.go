package runners

import (
	"github.com/Sirupsen/logrus"
	apb "github.com/asunaio/chariot/gen-go/asuna"
	"golang.org/x/net/context"
)

// Runners runs some shit
type Runners struct {
	Logger *logrus.Logger

	Apollo apb.ApolloClient
	Charon apb.CharonClient
}

func (r *Runners) Run(ctx context.Context, runner string) {
	switch runner {

	case "Apollo::GetChampion":
		r.ApolloChampion(ctx)
		break

	case "Apollo::GetMatchup":
		r.ApolloMatchup(ctx)
		break

	case "Charon::GetMatch":
		r.CharonMatch(ctx)
		break

	case "Charon::GetMatchList":
		r.CharonMatchList(ctx)
		break

	case "Charon::GetRankings":
		r.CharonRankings(ctx)
		break

	default:
		r.Logger.Fatalf("Unknown runner %q", runner)
		break

	}
}
