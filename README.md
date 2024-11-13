# Jumia Laptop Scraper

This Go application scrapes laptop data from Jumia Kenya's website every two minutes and saves the data in a CSV file. The scraped data includes the name and price of each laptop, and it is stored in a `Laptops.csv` file for analysis and future reference.

## Features
- Scrapes laptop names and prices from Jumia Kenya.
- Runs as a scheduled task every two minutes using `gocron`.
- Saves the data to a CSV file (`Laptops.csv`).

## Requirements
- Go 1.15+
- Required Go packages:
  - `colly`: For web scraping.
  - `gocron`: For scheduling the scraper to run periodically.

## Installation
1. Clone this repository:
    ```bash
    git clone https://github.com/your-username/jumia-laptop-scraper.git
    cd jumia-laptop-scraper
    ```

2. Install dependencies:
    ```bash
    go get -u github.com/gocolly/colly
    go get -u github.com/go-co-op/gocron
    ```

3. Run the application:
    ```bash
    go run main.go
    ```

## Usage
After running the program, it will start scraping the Jumia Kenya laptops page every two minutes, saving the results in `Laptops.csv` located in the same directory as the application. Each entry in the CSV file will contain:
- **Laptop Name**
- **Laptop Price**

## About the Code
- **Scheduler**: The `gocron` library schedules the scraper to run every two minutes.
- **Web Scraper**: `colly` is used to parse the HTML elements of the Jumia Kenya laptops page, selecting the laptop name and price.
- **CSV Writer**: The scraped data is written to `Laptops.csv` using the built-in `csv` package.

## About
This project is a simple web scraper written in Go, designed to gather laptop data periodically from Jumia Kenya. It demonstrates the use of scheduled tasks, web scraping with Colly, and file handling with CSV in Go.

## Disclaimer
This project is intended for educational purposes. Please review Jumia Kenya's terms of service before using this scraper in production environments.
