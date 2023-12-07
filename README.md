# Backoff Retry Mechanisms Feature

## Overview

This repository includes a feature for implementing HTTP request retries with multiple retry mechanisms, including incremental, exponential, and constant backoff intervals. The implementation follows the strategy design pattern, SOLID principles, and leverages various techniques for efficient handling of concurrent requests.

## Table of Contents

- [Integration](#integration)
  - [Example Usage](#example-usage)
- [API Endpoints](#api-endpoints)
  - [/read-file](#read-file)
  - [/operation](#operation)
- [Retry Mechanisms](#retry-mechanisms)

## Integration

To integrate the backoff retry mechanisms feature into your repository, follow these steps:

### 1. Clone the Repository

```bash
git clone https://github.com/Ishhyoboytarun/Backoff-Retry-Mechanisms.git
```

### 2. Navigate to the Repository

```bash
cd Backoff-Retry-Mechanisms
```

### 3. Run the Feature

```bash
go run main.go
```

This will start the server and make the retry mechanisms feature available for use.

### Example Usage

Assuming the feature is running locally, you can integrate it into your existing Go application as follows:

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Make a GET request to the /operation endpoint
	response, err := http.Get("http://localhost:8080/operation")
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	// Process the response
	fmt.Println("Response:", response)
}
```

## API Endpoints

### /read-file

Assumed to be an external API which reads a file. No additional retry mechanisms are applied here.

### /operation

Calls the /read-file API and applies the retry mechanism based on the specified strategy.

## Retry Mechanisms

```bash
curl http://localhost:8080/operation
```

Waits for an incremental/exponential/constant period of time for every retry.

Feel free to customize this template based on your specific project structure, requirements, and any additional information you'd like to provide to users integrating this feature into their repositories.

<img width="1512" alt="Screenshot 2023-12-07 at 6 18 56 PM" src="https://github.com/Ishhyoboytarun/Backoff-Retry-Mechanisms/assets/36428256/9e0bb75d-ece4-4d05-8456-aae981574f7a">
<img width="1512" alt="Screenshot 2023-12-07 at 6 17 10 PM" src="https://github.com/Ishhyoboytarun/Backoff-Retry-Mechanisms/assets/36428256/957cd1ae-d49e-46c9-be45-29d2d461a643">
<img width="1512" alt="Screenshot 2023-12-07 at 6 17 17 PM" src="https://github.com/Ishhyoboytarun/Backoff-Retry-Mechanisms/assets/36428256/4a7b61fd-dee0-402c-8938-082c222a26a9">
<img width="1512" alt="Screenshot 2023-12-07 at 6 19 11 PM" src="https://github.com/Ishhyoboytarun/Backoff-Retry-Mechanisms/assets/36428256/c696eb2e-b07b-4987-8e9c-f60fa305a37a">
<img width="1512" alt="Screenshot 2023-12-07 at 6 24 06 PM" src="https://github.com/Ishhyoboytarun/Backoff-Retry-Mechanisms/assets/36428256/37f7a24a-3913-4b3c-af54-9e759faa7287">
<img width="1512" alt="Screenshot 2023-12-07 at 6 24 59 PM" src="https://github.com/Ishhyoboytarun/Backoff-Retry-Mechanisms/assets/36428256/9397b6df-aca8-46da-997f-0b1dc78a97f3">


