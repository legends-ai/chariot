package runners

import (
	apb "github.com/asunaio/chariot/gen-go/asuna"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

// LunaGetSeasonRankings is Luna::GetSeasonRankings
func (r *Runners) LunaGetSeasonRankings(ctx context.Context) proto.Message {
	var summoners []*apb.SummonerId
	for _, sid := range r.Flags.SummonerId {
		summoners = append(summoners, &apb.SummonerId{
			Region: r.Flags.Region,
			Id:     sid,
		})
	}

	rankings, err := r.Luna.GetSeasonRankings(ctx, &apb.LunaRpc_GetSeasonRankingsRequest{
		Summoners: summoners,
		Season:    r.Flags.Season,
	})
	if err != nil {
		r.Logger.Fatalf("Could not get match: %v", err)
	}
	return rankings
}
