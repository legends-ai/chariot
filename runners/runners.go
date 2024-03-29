package runners

import (
	"github.com/Sirupsen/logrus"
	apb "github.com/asunaio/chariot/gen-go/asuna"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

type Flags struct {
	Locale        apb.Locale
	Region        apb.Region
	Role          apb.Role
	ChampionId    []uint32
	MatchId       uint64
	Season        apb.Season
	SummonerId    []uint64
	Version       string
	VulgateFormat string
}

// Runners runs some shit
type Runners struct {
	Logger *logrus.Logger
	Flags  Flags

	Charon  apb.CharonClient
	Lucinda apb.LucindaClient
	Luna    apb.LunaClient
	Vulgate apb.VulgateClient
}

func (r *Runners) Run(ctx context.Context, runner string) proto.Message {
	switch runner {

	case "Charon::GetMatch":
		return r.CharonMatch(ctx)

	case "Charon::GetVilemawMatch":
		return r.CharonVilemawMatch(ctx)

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

	case "Lucinda::GetStatistics":
		return r.LucindaStatistics(ctx)

	case "Lucinda::GetChampion":
		return r.LucindaChampion(ctx)

	case "Lucinda::GetMatchup":
		return r.LucindaMatchup(ctx)

	case "Lucinda::GetMatchSum":
		return r.LucindaMatchSum(ctx)

	case "Luna::GetSeasonRankings":
		return r.LunaGetSeasonRankings(ctx)

	case "Vulgate::GetChampions":
		return r.VulgateChampions(ctx)

	case "Vulgate::GetEntry":
		return r.VulgateEntry(ctx)

	default:
		r.Logger.Fatalf("Unknown runner %q", runner)
	}
	return nil
}
