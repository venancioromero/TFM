package recommendersystem

import (
	"encoding/json"
	"fmt"
	"log"

	"../config"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// Recommendation struct
type Recommendation struct {
	URL   string `json:"url"`
	Image string `json:"image"`
	//Ref   string
}

// RecommenderSystem struct
type RecommenderSystem struct {
	conf    config.Config
	neoConn bolt.Conn
}

// GetRecommendationByClientIDAndProductID method will return a recommendation depending idClient and idProduct
func (r *RecommenderSystem) GetRecommendationByClientIDAndProductID(clientID int, productID int) string {
	rec := &Recommendation{URL: "http://CLientID,PRODUCTid", Image: "image.jpg"}
	json, err := json.Marshal(rec)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(json)

}

// GetRecommendationByProductID method will returns a recommendation dependign idProduct
func (r *RecommenderSystem) GetRecommendationByProductID(productID int) string {
	rec := &Recommendation{URL: "http://ProductID", Image: "image.jpg"}
	json, err := json.Marshal(rec)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(json)

}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}

// New method will return recommenderSystem object
func New() RecommenderSystem {
	cfg := config.New()
	neo, err := bolt.NewDriver().OpenNeo(cfg.Neo4JURL)
	checkErr(err)
	recom := &RecommenderSystem{conf: cfg, neoConn: neo}
	return *recom
}

//Destroy method will close connection to graph database
func (r *RecommenderSystem) Destroy() {
	r.neoConn.Close()
}
