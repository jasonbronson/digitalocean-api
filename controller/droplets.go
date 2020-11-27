package controller

import (
	"context"
	"docontroller/config"
	"docontroller/digitalocean"
	"docontroller/middleware"
	"docontroller/repositories"
	"docontroller/utils"
	"strconv"

	"github.com/digitalocean/godo"
	"github.com/gorilla/mux"

	"net/http"
)

func GetDropletsHandler(w http.ResponseWriter, r *http.Request, config *config.Config) {

	ctx := context.Background()
	droplets, err := digitalocean.GetDroplets(ctx, config.GodoClient)
	if err != nil {
		utils.LogError(err)
		return
	}

	//record to the db all current droplets
	repositories.CreateDroplets(droplets)
	middleware.Response(w, droplets)

}

func CreateDropletHandler(w http.ResponseWriter, r *http.Request, config *config.Config) {

	ctx := context.Background()

	//get ssh keys
	keysData := digitalocean.GetSSHKeys(ctx, config.GodoClient)
	keys := []godo.DropletCreateSSHKey{}
	for _, key := range *keysData {
		k := godo.DropletCreateSSHKey{
			ID: key.ID,
		}
		keys = append(keys, k)
	}
	image := godo.DropletCreateImage{
		Slug: "debian-10-x64",
	}
	dropletRequest := godo.DropletCreateRequest{
		Name:              "TestingScript",
		Region:            "nyc3",
		Size:              "s-1vcpu-1gb",
		Image:             image,
		SSHKeys:           keys,
		Backups:           false,
		IPv6:              false,
		UserData:          "",
		PrivateNetworking: true,
		Volumes:           nil,
	}
	droplet, err := digitalocean.CreateDroplet(ctx, config.GodoClient, &dropletRequest)
	if err != nil {
		utils.LogError(err)
		return
	}

	middleware.Response(w, droplet)

}

func DeleteDropletHandler(w http.ResponseWriter, r *http.Request, config *config.Config) {

	ctx := context.Background()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		middleware.ResponseError(w, "Invalid droplet ID")
		return
	}
	response, err := digitalocean.DeleteDroplet(ctx, config.GodoClient, id)
	if err != nil {
		utils.LogError(err)
		return
	}

	middleware.Response(w, response)
}
