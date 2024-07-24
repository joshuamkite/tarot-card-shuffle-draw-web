# Tarot Card Shuffle Draw

Tarot Card Shuffle Draw is a web application that simulates a tarot card reading. Users can choose different decks, specify the number of cards to draw, and include reversed cards in the draw.

- [Tarot Card Shuffle Draw](#tarot-card-shuffle-draw)
  - [Features](#features)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Running the Application Locally](#running-the-application-locally)
    - [Running the Application with Docker](#running-the-application-with-docker)
  - [Usage](#usage)
    - [Web Interface](#web-interface)
    - [API Endpoints](#api-endpoints)
  - [Development and Testing](#development-and-testing)
    - [Running Tests](#running-tests)

## Features

- **Deck Options**: Full Deck, Major Arcana only, Minor Arcana only
- **Reversed Cards**: Option to include reversed cards
- **Random Draw**: Secure randomness using `crypto/rand`
- **Web Interface**: User-friendly web interface built with Gin
- **Dockerized**: Easy deployment with Docker
- **Continuous Integration**: Automated testing and deployment with GitHub Actions

## Getting Started

### Prerequisites

- Go 1.22 or later
- Docker (for containerized deployment)

### Running the Application Locally

1. **Clone the repository**:

    ```sh
    git clone https://github.com/yourusername/tarot_shuffle_draw.git
    cd tarot_shuffle_draw
    ```

2. **Install dependencies and run the application**:

    ```sh
    go mod tidy
    go run main.go
    ```

3. **Open your browser** and navigate to `http://localhost:8080`.

### Running the Application with Docker

1. **Build the Docker image**:

    ```sh
    docker build -t tarot_shuffle_draw .
    ```

2. **Run the Docker container**:

    ```sh
    docker run -p 80:80 tarot_shuffle_draw
    ```

3. **Open your browser** and navigate to `http://localhost`.

## Usage

### Web Interface

1. **Choose the deck type**: Full Deck, Major Arcana only, Minor Arcana only.
2. **Select reversed cards option**: Include or exclude reversed cards.
3. **Specify the number of cards to draw**.
4. **Click "Draw Cards"** to see the results.

### API Endpoints

- `GET /`: Displays the options page.
- `POST /draw`: Handles drawing cards based on user input.
- `GET /license`: Displays the license page.

## Development and Testing

### Running Tests

Run the tests using the following command:

```sh
go test -v -cover ./...
```

This project is licensed under the GNU Affero General Public License v3.0. See the [LICENSE](LICENSE