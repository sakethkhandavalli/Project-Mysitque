package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	//"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	//"github.com/rs/cors"
	data "gitlab.com/ashwini/ProjectMystique/src/databases"
	handler "gitlab.com/ashwini/ProjectMystique/src/handler"
	wallet "gitlab.com/ashwini/ProjectMystique/src/wallet"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "index.html")
}
func PackageHandler(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "packages.html")
}
func AllBlocksHandler(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "allblocks.html")
}
func TrackingHandler(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "tracking.html")
}

func main() {
	data.InitBlockList()
	data.InitPackageList()
	wallet.InitWallets()
	r := mux.NewRouter()

	cssHandler := http.FileServer(http.Dir("./public/css/"))
	imagesHandler := http.FileServer(http.Dir("./public/img/"))
	jshandler := http.FileServer(http.Dir("./public/js"))
	vendorhandler := http.FileServer(http.Dir("./public/vendor"))

	http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
	http.Handle("/img/", http.StripPrefix("/img/", imagesHandler))
	http.Handle("/js/", http.StripPrefix("/js/", jshandler))
	http.Handle("/vendor/", http.StripPrefix("/vendor/", vendorhandler))
	r.PathPrefix("/home").Handler(http.FileServer(http.Dir("./public/")))

	r.HandleFunc("/packagePage", PackageHandler)
	r.HandleFunc("/trackingPage", TrackingHandler)
	r.HandleFunc("/allBlocksPage", AllBlocksHandler)
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	r.Handle("/", handler.NotImplemented).Methods("POST")
	//Blocks
	r.Handle("/getRecentBlocks", handler.GetRecentBlocks).Methods("GET")
	r.Handle("/getRecentBlocks/{number}", handler.GetRecentBlocksByNumber).Methods("GET")
	r.Handle("/getAllBlocks", handler.GetAllBlocks).Methods("GET")
	r.Handle("/settings", handler.GetSettings).Methods("GET")
	r.Handle("/getPackageTable/{packageId}", handler.GetPackageTable).Methods("GET")
	r.Handle("/getAllPackages", handler.GetAllPackages).Methods("GET")
	r.Handle("/search/{searchstring}", handler.Search).Methods("GET")

	//Second page
	r.Handle("/trackPackage/{packageId}", handler.TrackPackageWithIdLtd).Methods("GET")
	r.Handle("/trackPackageEndToEnd/{packageId}", handler.TrackPackageWithId).Methods("GET")
	r.Handle("/getBlockInfo/{blockId}", handler.GetBlockInfoWithId).Methods("GET")

	//Wallet handler
	r.Handle("/wallet", wallet.FetchWalletPage).Methods("GET")
	r.Handle("/wallet/{userid}", wallet.FetchWalletPageWithUserId).Methods("GET")
	r.Handle("/wallet/getBalance/{userid}", wallet.GetBalanceByUserId).Methods("GET")
	r.Handle("/wallet/listAssociates/{userid}", wallet.ListAssociatesWithUserId).Methods("GET")
	r.Handle("/wallet/addAssociates/{userid}", wallet.AddAssociateWithUserId).Methods("POST")
	r.Handle("/wallet/latesttransactions/{userid}", wallet.GetLatestTransactionsByUserId).Methods("GET")
	r.Handle("/wallet/getAllTransactions/{userid}", wallet.GetAllTransactionsByUserId).Methods("GET")
	r.Handle("/wallet/addTransaction/{userid}", wallet.AddTransactionByUserId).Methods("POST")

	//handle server interrupts
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Stopping Server....")
		os.Exit(1)
	}()

	fmt.Println("Listening on 9040")
	//COmpile with this for CSS
	http.ListenAndServe(":9040", nil)

	//Compile wit this for AJAX
	//handler := cors.Default().Handler(r)
	//http.ListenAndServe(":9040", handlers.LoggingHandler(os.Stdout, handler))

}
