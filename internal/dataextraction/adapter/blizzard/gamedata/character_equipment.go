package pvpseason

import (
	"context"
	"fmt"
	"github.com/alessandro54/stats/internal/dataextraction/adapter/blizzard"
)

func FetchCharacterEquipment(ctx context.Context, realm string, name string, opts map[string]string) ([]byte, error) {
	client, _ := blizzard.GetClient(ctx, opts["region"], opts["locale"])

	return client.Get(
		ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/equipment", realm, name),
		map[string]string{
			"namespace": "dynamic-" + client.Region,
			"locale":    client.Locale,
		},
	)
}
