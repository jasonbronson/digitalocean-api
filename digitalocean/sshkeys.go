package digitalocean

import (
	"context"
	"docontroller/utils"

	"github.com/digitalocean/godo"
)

func GetSSHKeys(ctx context.Context, client *godo.Client) *[]godo.Key {
	opt := &godo.ListOptions{
		Page:    1,
		PerPage: 200,
	}

	keys, _, err := client.Keys.List(ctx, opt)
	if err != nil {
		utils.LogError(err)
		return nil
	}
	return &keys
}
