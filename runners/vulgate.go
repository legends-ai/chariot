package runners

import (
	apb "github.com/asunaio/chariot/gen-go/asuna"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

func (r *Runners) VulgateChampions(ctx context.Context) proto.Message {
	champions, err := r.Vulgate.GetChampions(ctx, &apb.VulgateRpc_GetChampionsRequest{
		Context: &apb.VulgateRpc_Context{
			Locale: apb.Locale(apb.Locale_value[ctx.Value("locale").(string)]),
			Region: apb.Region(apb.Region_value[ctx.Value("region").(string)]),
		},
	})
	if err != nil {
		r.Logger.Fatalf("Could not get champions from vulgate: %v", err)
	}
	return champions
}
