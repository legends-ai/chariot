package runners

import (
	apb "github.com/asunaio/chariot/gen-go/asuna"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

// LucindaStatistics
func (r *Runners) LucindaStatistics(ctx context.Context) proto.Message {
	statistics, err := r.Lucinda.GetStatistics(ctx, &apb.LucindaRpc_GetStatisticsRequest{
		EnemyChampionId: -1,
		Patch: &apb.PatchRange{
			Min: r.Flags.Version,
			Max: r.Flags.Version,
		},
		// match everything
		Tier: &apb.TierRange{
			Min: apb.Tier_BRONZE,
			Max: apb.Tier_CHALLENGER,
		},
		Region:       r.Flags.Region,
		ForceRefresh: true,
	})
	if err != nil {
		r.Logger.Fatalf("Could not get statistics: %v", err)
	}
	return statistics
}

// LucindaChampion is Lucinda::Champion
func (r *Runners) LucindaChampion(ctx context.Context) proto.Message {
	champion, err := r.Lucinda.GetChampion(ctx, &apb.LucindaRpc_GetChampionRequest{
		ChampionId: r.Flags.ChampionId[0],
		Patch: &apb.PatchRange{
			Min: r.Flags.Version,
			Max: r.Flags.Version,
		},
		// match everything
		Tier: &apb.TierRange{
			Min: apb.Tier_BRONZE,
			Max: apb.Tier_CHALLENGER,
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
			Min: r.Flags.Version,
			Max: r.Flags.Version,
		},
		// match everything
		Tier: &apb.TierRange{
			Min: apb.Tier_BRONZE,
			Max: apb.Tier_CHALLENGER,
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
				ChampionId: &apb.ChampionId{Value: r.Flags.ChampionId[0]},
				Patch:      r.Flags.Version,
				Tier:       0x40,
				Region:     r.Flags.Region,
				Role:       r.Flags.Role,
			},
			{
				ChampionId: &apb.ChampionId{Value: r.Flags.ChampionId[0]},
				Patch:      r.Flags.Version,
				Tier:       0x50,
				Region:     r.Flags.Region,
				Role:       r.Flags.Role,
			},
			{
				ChampionId: &apb.ChampionId{Value: r.Flags.ChampionId[0]},
				Patch:      r.Flags.Version,
				Tier:       0x60,
				Region:     r.Flags.Region,
				Role:       r.Flags.Role,
			},
			{
				ChampionId: &apb.ChampionId{Value: r.Flags.ChampionId[0]},
				Patch:      r.Flags.Version,
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
