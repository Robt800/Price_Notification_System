package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetItemPriceHandler(w http.ResponseWriter, r *http.Request) {
	//Create key-value pairs (map) from the URL
	vars := mux.Vars(r)

	//Obtain from the url which item we are interested in reporting
	itemsToReport := vars["id"]

}
