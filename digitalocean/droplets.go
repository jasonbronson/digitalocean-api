package digitalocean

import (
	"context"
	"docontroller/utils"

	"github.com/davecgh/go-spew/spew"
	"github.com/digitalocean/godo"
)

func CreateDroplet(ctx context.Context, client *godo.Client, dropletRequest *godo.DropletCreateRequest) (*godo.Droplet, error) {

	droplet, response, err := client.Droplets.Create(ctx, dropletRequest)
	if err != nil {
		utils.LogError(err)
		return nil, err
	}
	//debug
	spew.Dump(response)
	spew.Dump(droplet)
	return droplet, nil
}

func DeleteDroplet(ctx context.Context, client *godo.Client, id int) (*godo.Response, error) {

	response, err := client.Droplets.Delete(ctx, id)
	if err != nil {
		utils.LogError(err)
		return nil, err
	}
	spew.Dump(response)
	return response, nil
}

func GetDroplets(ctx context.Context, client *godo.Client) ([]godo.Droplet, error) {
	// create a list to hold our droplets
	list := []godo.Droplet{}

	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		droplets, resp, err := client.Droplets.List(ctx, opt)
		if err != nil {
			return nil, err
		}

		// append the current page's droplets to our list
		list = append(list, droplets...)

		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		// set the page we want for the next request
		opt.Page = page + 1
	}

	return list, nil
}
