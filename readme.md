# Stock Price View Application

This is a stock price view application built using Golang, MongoDB, and a RESTful API.

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Setup](#setup)
- [API Routes](#api-routes)
  - [Get Top 10 Stocks](#get-top-10-stocks)
  - [Find Stocks by Name](#find-stocks-by-name)
  - [Get Stock Price History](#get-stock-price-history)
  - [Add Stock to Favorites](#add-stock-to-favorites)
  - [See Favorite Stocks](#see-favorite-stocks)
  - [Remove Stock from Favorites](#remove-stock-from-favorites)


## Overview

The Stock Price View Application is designed to access and manage stock data from the Bombay Stock Exchange (BSE). It utilizes Golang for scripting, MongoDB for data storage, and offers a RESTful API for seamless integration.

## Project Structure

The project is organized into several components:

- **Scripts:**
  - `downloader.go`: Downloads the Equity Bhavcopy ZIP from the BSE website and extracts the CSV file.
  - `extractor.go`: Handles the extraction of data from the CSV file.
  - `processor.go`: Manages data processing and storage in MongoDB.
  - `database.go`: Defines MongoDB operations.

- **API:**
  - `handler.go`: Implements API request handlers.
  - `routes.go`: Defines API routes.

- **Cache:**
  - `cache.go`: Implements a caching layer for improved performance.

- **Main:**
  - `main.go`: Orchestrates the execution of different components.

## Getting Started

### Prerequisites

Ensure you have the following installed:

- Golang
- MongoDB
- Additional Golang packages (dependencies) - Install them using `go get`.

### Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/atul-007/stock-price-view.git

## Endpoints

- `GET /health`: Endpoin to check if the service is up and running or not .
- `GET /top10stocks`: Endpoin to handle the request for getting the top 10 stocks.
- `GET /stockbyname?name:{name of the stock}`: Endpoin to handle the reuest for  getting the stock by name.
- `GET /stockpricehistory?code:{stock code}`: Endpoin to handle the reuest for  getting stock price history .
- `GET /favouritestocks`: Endpoin to handle the reuest for  getting  your favourite stock .


- `DELETE /removefromfavorites`: Endpoin to remove stocks from your favourites .
Payload:
  ```json
{
        "Code": "974274",
        "Name": "MSFL31022   ",
        "Group": "F ",
        "Type": "B",
        "Open": 1140002,
        "High": 1140002,
        "Low": 1140002,
        "Close": 1140002,
        "Last": 1140002,
        "PrevClose": 1139690,
        "NoTrades": 1,
        "NoOfShares": 4,
        "NetTurnover": 4560008
    

}
  ```


- `POST /addtofavorites`: Endpoint to add a stock as favourite. Expects a JSON payload with the form response.

Payload:
  ```json
   
    {
        "Code": "974274",
        "Name": "MSFL31022   ",
        "Group": "F ",
        "Type": "B",
        "Open": 1140002,
        "High": 1140002,
        "Low": 1140002,
        "Close": 1140002,
        "Last": 1140002,
        "PrevClose": 1139690,
        "NoTrades": 1,
        "NoOfShares": 4,
        "NetTurnover": 4560008
    

}

  ```





