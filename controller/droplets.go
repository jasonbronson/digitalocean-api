package controller

import (
	"context"
	"docontroller/config"
	"docontroller/digitalocean"
	"docontroller/middleware"
	logger "docontroller/middleware"
	"net/http"
)

func GetDropletsHandler(w http.ResponseWriter, r *http.Request, config *config.Config) {

	ctx := context.TODO()
	droplets, err := digitalocean.GetDroplets(ctx, config.GodoClient)
	if err != nil {
		logger.LogError(err)
		return
	}

	middleware.Response(w, droplets)

}
