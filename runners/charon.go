package runners

import (
	apb "github.com/asunaio/chariot/gen-go/asuna"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

// CharonMatch is Charon::GetMatch
func (r *Runners) CharonMatch(ctx context.Context) proto.Message {
	match, err := r.Charon.GetMatch(ctx, &apb.CharonMatchRequest{
		Match: &apb.MatchId{
			Region: apb.Region_NA,
			Id:     1365660506,
		},
	})
	if err != nil {
		r.Logger.Fatalf("Could not get match: %v", err)
	}
	return match
}

// CharonMatchList is Charon::GetMatchList
func (r *Runners) CharonMatchList(ctx context.Context) proto.Message {
	match, err := r.Charon.GetMatchList(ctx, &apb.CharonMatchListRequest{
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
	rankings, err := r.Charon.GetRankings(ctx, &apb.CharonRankingsRequest{
		Region:      apb.Region_NA,
		SummonerIds: []uint64{29236065, 24575247},
	})
	if err != nil {
		r.Logger.Fatalf("Could not get rankings: %v", err)
	}
	return rankings
}
