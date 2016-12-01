package runners

import (
	apb "github.com/asunaio/chariot/gen-go/asuna"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

func (r *Runners) VulgateChampions(ctx context.Context) proto.Message {
	context := &apb.VulgateData_Context{
		Locale: r.Flags.Locale,
		Region: r.Flags.Region,
	}

	if r.Flags.Version != "" {
		context.Release = &apb.VulgateData_Context_Version{
			Version: r.Flags.Version,
		}
	}

	champions, err := r.Vulgate.GetChampions(ctx, &apb.VulgateRpc_GetChampionsRequest{
		Context:   context,
		Champions: r.Flags.ChampionId,
		Format:    apb.VulgateData_Format(apb.VulgateData_Format_value[r.Flags.VulgateFormat]),
	})

	if err != nil {
		r.Logger.Fatalf("Could not get champions from vulgate: %v", err)
	}

	return champions
}

func (r *Runners) VulgateEntry(ctx context.Context) proto.Message {
	context := &apb.VulgateData_Context{
		Locale: r.Flags.Locale,
		Region: r.Flags.Region,
	}

	if r.Flags.Version != "" {
		context.Release = &apb.VulgateData_Context_Version{
			Version: r.Flags.Version,
		}
	}

	entry, err := r.Vulgate.GetEntry(ctx, &apb.VulgateRpc_GetEntryRequest{
		Context: context,
		Format:  apb.VulgateData_Format(apb.VulgateData_Format_value[r.Flags.VulgateFormat]),
	})

	if err != nil {
		r.Logger.Fatalf("Could not get entry from vulgate: %v", err)
	}

	return entry
}
