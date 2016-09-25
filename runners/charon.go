package runners

import (
	"bytes"
	"fmt"

	apb "github.com/asunaio/chariot/gen-go/asuna"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

// CharonMatch is Charon::GetMatch
func (r *Runners) CharonMatch(ctx context.Context) {
	match, err := r.Charon.GetMatch(ctx, &apb.CharonMatchRequest{
		Match: &apb.MatchId{
			Region: apb.Region_NA,
			Id:     1365660506,
		},
	})
	if err != nil {
		r.Logger.Fatalf("Could not get match: %v", err)
	}

	var out bytes.Buffer
	if err := proto.MarshalText(&out, match); err != nil {
		r.Logger.Fatalf("Could not marshal match: %v", err)
	}

	fmt.Println(out.String())
}
