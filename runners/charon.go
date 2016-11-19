package runners

import (
	apb "github.com/asunaio/chariot/gen-go/asuna"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

// CharonRpc_Match is Charon::GetMatch
func (r *Runners) CharonMatch(ctx context.Context) proto.Message {
	match, err := r.Charon.GetMatch(ctx, &apb.CharonRpc_MatchRequest{
		Match: &apb.MatchId{
			Region: apb.Region_NA,
			Id:     ctx.Value("matchId").(uint64),
		},
	})
	if err != nil {
		r.Logger.Fatalf("Could not get match: %v", err)
	}
	return match
}

// CharonRpc_Match is Charon::GetDominionMatch
func (r *Runners) CharonDominionMatch(ctx context.Context) proto.Message {
	match, err := r.Charon.GetMatch(ctx, &apb.CharonRpc_MatchRequest{
		Match: &apb.MatchId{
			Region: apb.Region_NA,
			Id:     2315462640,
		},
	})
	if err != nil {
		r.Logger.Fatalf("Could not get dominion match: %v", err)
	}
	return match
}

// CharonRpc_MatchList is Charon::GetMatchList
func (r *Runners) CharonMatchList(ctx context.Context) proto.Message {
	match, err := r.Charon.GetMatchList(ctx, &apb.CharonRpc_MatchListRequest{
		Summoner: &apb.SummonerId{
			Region: apb.Region_NA,
			Id:     29236065,
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
		Region:      apb.Region_NA,
		SummonerIds: []uint64{29236065, 24575247},
	})
	if err != nil {
		r.Logger.Fatalf("Could not get rankings: %v", err)
	}
	return rankings
}

// CharonStaticChampions is Charon::GetStatic
func (r *Runners) CharonStatic(ctx context.Context) proto.Message {
	sc, err := r.Charon.GetStatic(ctx, &apb.CharonRpc_StaticRequest{
		Region:  apb.Region_NA,
		Locale:  apb.Locale(apb.Locale_value[ctx.Value("locale").(string)]),
		Version: ctx.Value("version").(string),
	})
	if err != nil {
		r.Logger.Fatalf("Could not get static: %v", err)
	}
	return sc
}
