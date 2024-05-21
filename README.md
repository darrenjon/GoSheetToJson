# GoSheetToJson

GoSheetToJson is a Go application that retrieves data from a specified Google Sheet and converts it into a JSON file. The application utilizes Google Sheets API and OAuth2 for authentication, and it reads configuration from a `.env` file.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (1.16 or higher) installed on your local machine. You can download it from [golang.org](https://golang.org/dl/).
- A Google Cloud project with Google Sheets API enabled.
- Service account credentials JSON file for your Google Cloud project.

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/your-username/GoSheetToJson.git
    cd GoSheetToJson
    ```

2. Install the required Go packages:

    ```sh
    go mod tidy
    ```

## Configuration

1. Place your service account credentials JSON file in the project root directory and name it `service_account.json`.

2. Create a `.env` file in the project root directory and add your Google Sheet ID:

    ```env
    SPREADSHEET_ID=your_google_sheet_id_here
    ```

## Usage

To run the application, simply use the `go run` command:

```sh
go run main.go
