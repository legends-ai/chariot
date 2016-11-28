package runners

import (
	apb "github.com/asunaio/chariot/gen-go/asuna"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

func (r *Runners) VulgateChampions(ctx context.Context) proto.Message {
	champions, err := r.Vulgate.GetChampions(ctx, &apb.VulgateRpc_GetChampionsRequest{
		Context: &apb.VulgateRpc_Context{
			Locale: r.Flags.Locale,
			Region: r.Flags.Region,
			Release: &apb.VulgateRpc_Context_Version{
				Version: r.Flags.Version,
			},
		},
		Champions: r.Flags.ChampionId,
		Format:    apb.VulgateRpc_GetChampionsRequest_Format(apb.VulgateRpc_GetChampionsRequest_Format_value[r.Flags.VulgateFormat]),
	})
	if err != nil {
		r.Logger.Fatalf("Could not get champions from vulgate: %v", err)
	}
	return champions
}
