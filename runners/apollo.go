package runners

import (
	"bytes"
	"fmt"

	apb "github.com/asunaio/chariot/gen-go/asuna"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

// ApolloChampion is Apollo::Champion
func (r *Runners) ApolloChampion(ctx context.Context) {
	champion, err := r.Apollo.GetChampion(ctx, &apb.GetChampionRequest{
		ChampionId: 64, // this is Lee Sin
		Patch: &apb.PatchRange{
			Min: "6.16",
			Max: "6.18",
		},
		// match everything
		Tier: &apb.TierRange{
			Min: 0x0000,
			Max: 0x1000,
		},
		Region: apb.Region_NA,
		Role:   apb.Role_JUNGLE,
	})
	if err != nil {
		r.Logger.Fatalf("Could not get champion: %v", err)
	}

	var out bytes.Buffer
	if err := proto.MarshalText(&out, champion); err != nil {
		r.Logger.Fatalf("Could not marshal champion: %v", err)
	}

	fmt.Println(out.String())
}

// ApolloMatchup is Apollo::Matchup
func (r *Runners) ApolloMatchup(ctx context.Context) {
	matchup, err := r.Apollo.GetMatchup(ctx, &apb.GetMatchupRequest{
		FocusChampionId: 51, // this is Caitlyn
		EnemyChampionId: 81, // this is Ezreal
		Patch: &apb.PatchRange{
			Min: "6.16",
			Max: "6.18",
		},
		// match everything
		Tier: &apb.TierRange{
			Min: 0x0000,
			Max: 0x1000,
		},
		Region: apb.Region_NA,
		Role:   apb.Role_BOT,
	})
	if err != nil {
		r.Logger.Fatalf("Could not get matchup: %v", err)
	}

	var out bytes.Buffer
	if err := proto.MarshalText(&out, matchup); err != nil {
		r.Logger.Fatalf("Could not marshal matchup: %v", err)
	}

	fmt.Println(out.String())
}
