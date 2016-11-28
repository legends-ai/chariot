package runners

import (
	apb "github.com/asunaio/chariot/gen-go/asuna"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

// LucindaChampion is Lucinda::Champion
func (r *Runners) LucindaChampion(ctx context.Context) proto.Message {
	champion, err := r.Lucinda.GetChampion(ctx, &apb.LucindaRpc_GetChampionRequest{
		ChampionId: r.Flags.ChampionId[0],
		Patch: &apb.PatchRange{
			Min: "6.18",
			Max: "6.18",
		},
		// match everything
		Tier: &apb.TierRange{
			Min: 0x0000,
			Max: 0x1000,
		},
		Region: r.Flags.Region,
		Role:   r.Flags.Role,
	})
	if err != nil {
		r.Logger.Fatalf("Could not get champion: %v", err)
	}
	return champion
}

// LucindaMatchup is Lucinda::Matchup
func (r *Runners) LucindaMatchup(ctx context.Context) proto.Message {
	matchup, err := r.Lucinda.GetMatchup(ctx, &apb.LucindaRpc_GetMatchupRequest{
		FocusChampionId: r.Flags.ChampionId[0],
		EnemyChampionId: r.Flags.ChampionId[1],
		Patch: &apb.PatchRange{
			Min: "6.16",
			Max: "6.18",
		},
		// match everything
		Tier: &apb.TierRange{
			Min: 0x0000,
			Max: 0x1000,
		},
		Region: r.Flags.Region,
		Role:   r.Flags.Role,
	})
	if err != nil {
		r.Logger.Fatalf("Could not get matchup: %v", err)
	}
	return matchup
}

// LucindaMatchSum is Lucinda::GetMatchSum
func (r *Runners) LucindaMatchSum(ctx context.Context) proto.Message {
	sum, err := r.Lucinda.GetMatchSum(ctx, &apb.LucindaRpc_GetMatchSumRequest{
		Filters: []*apb.MatchFilters{
			{
				ChampionId: int32(r.Flags.ChampionId[0]),
				EnemyId:    -1,
				Patch:      "6.18",
				Tier:       0x40,
				Region:     r.Flags.Region,
				Role:       r.Flags.Role,
			},
			{
				ChampionId: int32(r.Flags.ChampionId[0]),
				EnemyId:    -1,
				Patch:      "6.18",
				Tier:       0x50,
				Region:     r.Flags.Region,
				Role:       r.Flags.Role,
			},
			{
				ChampionId: int32(r.Flags.ChampionId[0]),
				EnemyId:    -1,
				Patch:      "6.18",
				Tier:       0x60,
				Region:     r.Flags.Region,
				Role:       r.Flags.Role,
			},
			{
				ChampionId: int32(r.Flags.ChampionId[0]),
				EnemyId:    -1,
				Patch:      "6.18",
				Tier:       0x70,
				Region:     r.Flags.Region,
				Role:       r.Flags.Role,
			},
		},
	})
	if err != nil {
		r.Logger.Fatalf("Could not get match sum: %v", err)
	}
	return sum
}
