package runners

import (
	apb "github.com/asunaio/chariot/gen-go/asuna"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

// CharonMatch is Charon::GetMatch
func (r *Runners) CharonMatch(ctx context.Context) proto.Message {
	match, err := r.Charon.GetMatch(ctx, &apb.CharonRpc_MatchRequest{
		Match: &apb.MatchId{
			Region: r.Flags.Region,
			Id:     r.Flags.MatchId,
		},
	})
	if err != nil {
		r.Logger.Fatalf("Could not get match: %v", err)
	}
	return match
}

// CharonDominionMatch is Charon::GetDominionMatch
func (r *Runners) CharonDominionMatch(ctx context.Context) proto.Message {
	match, err := r.Charon.GetMatch(ctx, &apb.CharonRpc_MatchRequest{
		Match: &apb.MatchId{
			Region: r.Flags.Region,
			Id:     r.Flags.MatchId,
		},
	})
	if err != nil {
		r.Logger.Fatalf("Could not get dominion match: %v", err)
	}
	return match
}

// CharonMatchList is Charon::GetMatchList
func (r *Runners) CharonMatchList(ctx context.Context) proto.Message {
	match, err := r.Charon.GetMatchList(ctx, &apb.CharonRpc_MatchListRequest{
		Summoner: &apb.SummonerId{
			Region: r.Flags.Region,
			Id:     r.Flags.SummonerId[0],
		},
		Seasons: []string{
			"PRESEASON2015",
			"SEASON2015",
			"PRESEASON2016",
			"SEASON2016",
		},
	})
	if err != nil {
		r.Logger.Fatalf("Could not get match list: %v", err)
	}
	return match
}

// CharonRankings is Charon::GetRankings
func (r *Runners) CharonRankings(ctx context.Context) proto.Message {
	rankings, err := r.Charon.GetRankings(ctx, &apb.CharonRpc_RankingsRequest{
		Region:      r.Flags.Region,
		SummonerIds: r.Flags.SummonerId,
	})
	if err != nil {
		r.Logger.Fatalf("Could not get rankings: %v", err)
	}
	return rankings
}

// CharonStaticChampions is Charon::GetStatic
func (r *Runners) CharonStatic(ctx context.Context) proto.Message {
	sc, err := r.Charon.GetStatic(ctx, &apb.CharonRpc_StaticRequest{
		Region:  r.Flags.Region,
		Locale:  r.Flags.Locale,
		Version: r.Flags.Version,
	})
	if err != nil {
		r.Logger.Fatalf("Could not get static: %v", err)
	}
	return sc
}

// CharonStaticVersions is Charon::GetVersions
func (r *Runners) CharonStaticVersions(ctx context.Context) proto.Message {
	sv, err := r.Charon.GetStaticVersions(ctx, &apb.CharonRpc_StaticVersionsRequest{
		Region: r.Flags.Region,
	})
	if err != nil {
		r.Logger.Fatalf("Could not get static versions: %v", err)
	}
	return sv
}
