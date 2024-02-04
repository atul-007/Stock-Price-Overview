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
   git clone https://github.com/yourusername/stock-price-view.git

### API Routes

#### Get Top 10 Stocks

**Request:**

```bash
curl http://localhost:8080/api/top10


Response:

[
  {
    "Code": "ABC123",
    "Name": "Company ABC"
    // Other fields...
  },
  // ... (up to 10 stocks)
]


#### Find Stocks by Name
**Request:**

curl http://localhost:8080/api/stocks?name=Company


**Response:**

[
  {
    "Code": "ABC123",
    "Name": "Company ABC"
    // Other fields...
  },
  // ... (matching stocks)
]


Get Stock Price History
**Request:**

curl http://localhost:8080/api/history?code=ABC123


Response:

[
  {
    "Date": "2024-01-31",
    "Price": 123.45
    // Other fields...
  },
  // ... (historical data)
]


Add Stock to Favorites
**Request:**

curl -X POST -H "Content-Type: application/json" -d '{"code": "ABC123", "name": "Company ABC"}' http://localhost:8080/api/favorites

Response:

{
  "message": "Stock added to favorites successfully."
}

See Favorite Stocks
**Request:**

curl http://localhost:8080/api/favorites

Response:

[
  {
    "Code": "ABC123",
    "Name": "Company ABC"
    // Other fields...
  },
  // ... (favorite stocks)
]


Remove Stock from Favorites
**Request:**

curl -X DELETE http://localhost:8080/api/favorites?code=ABC123

Response:

{
  "message": "Stock removed from favorites successfully."
}
