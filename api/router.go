package api

import "github.com/gorilla/mux"

// SetupRoutes sets up the API routes
func Router() *mux.Router {
	router := mux.NewRouter()

	// Define API routes
	router.HandleFunc("/", Health).Methods("GET")
	router.HandleFunc("/top10stocks", GetTop10StocksHandler).Methods("GET")
	router.HandleFunc("/stockbyname", GetStockByNameHandler).Methods("GET")
	router.HandleFunc("/stockpricehistory", GetStockPriceHistoryHandler).Methods("GET")
	router.HandleFunc("/addtofavorites", AddToFavoritesHandler).Methods("POST")
	router.HandleFunc("/favouritestocks", GetFavoriteStocksHandler).Methods("GET")
	router.HandleFunc("/removefromfavorites", RemoveFromFavoritesHandler).Methods("DELETE")

	return router
}
