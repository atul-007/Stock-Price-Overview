package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atul-007/stockPriceView/api"
	"github.com/atul-007/stockPriceView/data"
)

func main() {
	// URL for the Equity Bhavcopy ZIP file
	zipURL := "https://www.bseindia.com/download/BhavCopy/Equity/EQ250124_CSV.ZIP"

	// Local path to save the downloaded ZIP file
	zipFilePath := "./equity_bhavcopy.zip"

	// Local path to save the extracted CSV file
	csvFilePath := "./equity_bhavcopy.csv"

	// Download the Equity Bhavcopy ZIP file
	err := data.DownloadFile(zipFilePath, zipURL)
	if err != nil {
		log.Fatal("Error downloading file:", err)
	}

	// Extract the CSV file from the ZIP
	err = data.ExtractCSVFromZIP(zipFilePath, csvFilePath)
	if err != nil {
		log.Fatal("Error extracting CSV file:", err)
	}

	// err = data.SaveToMongoDB(csvFilePath)
	// if err != nil {
	// 	log.Fatal("Error saving data to MongoDB:", err)
	// }

	data.GetMongoDBCollection()
	data.InitializeMongoDBConnection()
	//data.GetTop10Stocks()

	fmt.Println("Data processing completed successfully.")

	fmt.Println("Hello World!")
	r := api.Router()
	fmt.Println("server is getting started")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000...")
}
