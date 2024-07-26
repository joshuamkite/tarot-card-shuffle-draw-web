# Tarot Card Shuffle Draw

Tarot Card Shuffle Draw is a web application that simulates a tarot card reading. Users can choose different decks, specify the number of cards to draw, and include reversed cards in the draw.

- [Tarot Card Shuffle Draw](#tarot-card-shuffle-draw)
  - [Features](#features)
  - [Getting Started](#getting-started)
    - [Running the Application Locally](#running-the-application-locally)
    - [Running the Application with Docker](#running-the-application-with-docker)
    - [Running the Application with Helm](#running-the-application-with-helm)
    - [Running the Application with ArgoCD](#running-the-application-with-argocd)
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

The included GitHub actions workflow covers test, build, package, publish.

### Running the Application Locally

**Prerequisites**

- Go 1.22 or later
- Docker (for containerized deployment)

There is a helper script for downloading the (included) images from Wikipedia/Wikimedia commons in `image_downloader`

1. **Clone the repository**:

    ```sh
    git clone https://github.com/joshuamkite/tarot-card-shuffle-draw-web.git
    cd tarot-card-shuffle-draw-web
    ```

2. **Install dependencies and run the application**:

    ```sh
    go mod tidy
    go run main.go
    ```

3. **Open your browser** and navigate to `http://localhost:80`.

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

### Running the Application with Helm

**Prerequisites**

- Kubernetes cluster
- Helm installed

1. **Add the Helm Repository**:

    ```sh
    helm repo add tarot-card-shuffle-draw-web https://github.com/joshuamkite/tarot-card-shuffle-draw-web
    ```

2. **Update the Helm Repository**:

    ```sh
    helm repo update
    ```

3. **Install the Helm Chart**:

    ```sh
    helm install tarot-shuffle-draw tarot-card-shuffle-draw-web/helm/tarot-shuffle-draw
    ```

4. **Access the Application**:

    ```sh
    kubectl get svc --namespace default
    ```

    Look for the `tarot-shuffle-draw` service and note the `NodePort`. Access the application at `http://<node-ip>:<node-port>`.

### Running the Application with ArgoCD

**Prerequisites**

- Kubernetes cluster
- ArgoCD installed

1. **Add the Application to ArgoCD**:

    Create a new application in ArgoCD pointing to this repository.

    ```sh
    argocd app create tarot-shuffle-draw \
      --repo https://github.com/joshuamkite/tarot-card-shuffle-draw-web.git \
      --path helm/tarot-shuffle-draw \
      --dest-server https://kubernetes.default.svc \
      --dest-namespace default
    ```

2. **Sync the Application**:

    ```sh
    argocd app sync tarot-shuffle-draw
    ```

3. **Access the Application**:

    ```sh
    kubectl get svc --namespace default
    ```

    Look for the `tarot-shuffle-draw` service and note the `NodePort`. Access the application at `http://<node-ip>:<node-port>`.

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
