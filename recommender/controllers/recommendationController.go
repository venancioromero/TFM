package controllers

import (
	"fmt"
	"strconv"

	recommender "../recommenderSystem"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

var (
	recommenderSystem recommender.RecommenderSystem
)

func init() {
	recommenderSystem = recommender.New()
}

func getIDs(c *routing.Context) (int, int, error) {
	idClient, err := strconv.Atoi(c.Param("idClient"))

	if err != nil {
		return 0, 0, err
	}

	idProduct, err := strconv.Atoi(c.Param("idProduct"))

	if err != nil {
		return 0, 0, err
	}
	return idClient, idProduct, nil
}

//GetRecommentation method is a controller that return a list of recommendations
func GetRecommentation(c *routing.Context) error {
	c.Response.Header.Set("Content-type", "application/json")
	c.Response.Header.Set("Server", "Recommender system")

	clientID, productID, err := getIDs(c)

	if err != nil {
		c.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(c, "{\"error\":\"Bad request\"}")
		return nil
	}

	if clientID != 0 {
		fmt.Fprintf(c, recommenderSystem.GetRecommendationByClientIDAndProductID(clientID, productID))
	} else {
		fmt.Fprintf(c, recommenderSystem.GetRecommendationByProductID(productID))
	}

	return nil
}
