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

	Charon  apb.CharonClient
	Lucinda apb.LucindaClient
	Vulgate apb.VulgateClient
}

func (r *Runners) Run(ctx context.Context, runner string) proto.Message {
	switch runner {

	case "Charon::GetMatch":
		return r.CharonMatch(ctx)

	case "Charon::GetDominionMatch":
		return r.CharonDominionMatch(ctx)

	case "Charon::GetMatchList":
		return r.CharonMatchList(ctx)

	case "Charon::GetRankings":
		return r.CharonRankings(ctx)

	case "Charon::GetStatic":
		return r.CharonStatic(ctx)

	case "Charon::GetStaticVersions":
		return r.CharonStaticVersions(ctx)

	case "Lucinda::GetChampion":
		return r.LucindaChampion(ctx)

	case "Lucinda::GetMatchup":
		return r.LucindaMatchup(ctx)

	case "Lucinda::GetMatchSum":
		return r.LucindaMatchSum(ctx)

	default:
		r.Logger.Fatalf("Unknown runner %q", runner)
	}
	return nil
}
