# Tarot Card Shuffle Draw

Tarot Card Shuffle Draw is a free and open-source project that shuffles and returns a selection of Tarot cards. Users can choose different decks, specify the number of cards to draw, and include reversed cards in the draw. Public domain illustrations of the cards are presented with the results. 

This port of the application is dockerised, with accompanying Helm chart that can be run in plain docker or installed on Kubernetes. There are other versions available - see [Alternative Deployment Ports](#alternative-deployment-ports) below


- [Tarot Card Shuffle Draw](#tarot-card-shuffle-draw)
  - [Features](#features)
  - [Getting Started](#getting-started)
    - [Running the Application Locally](#running-the-application-locally)
    - [Running the Application with Docker](#running-the-application-with-docker)
    - [Running the Application with Helm (Option 1)](#running-the-application-with-helm-option-1)
    - [Running the Application with ArgoCD](#running-the-application-with-argocd)
  - [Usage](#usage)
    - [Web Interface](#web-interface)
    - [API Endpoints](#api-endpoints)
  - [Development and Testing](#development-and-testing)
    - [Running Tests](#running-tests)
  - [Alternative Deployment Ports](#alternative-deployment-ports)

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

3. **Open your browser** and navigate to `http://localhost:8080`.

### Running the Application with Docker

1. **Build the Docker image**:

    ```sh
    docker build -t tarot_shuffle_draw .
    ```

2. **Run the Docker container**:

    ```sh
    docker run -p 8080:8080 tarot_shuffle_draw
    ```

3. **Open your browser** and navigate to `http://localhost:8080`.

### Running the Application with Helm (Option 1)

**Prerequisites**

- Kubernetes cluster
- Helm installed

1. **Install the Helm Chart**:

    ```sh
   helm install tarot-shuffle-draw helm/tarot-shuffle-draw --repo https://github.com/joshuamkite/tarot-card-shuffle-draw-web.git
    ```

2. **Access the Application**:

    ```sh
    kubectl get svc --namespace default
    ```

Look for the `tarot-shuffle-draw` service and note the `NodePort`. Access the application at `http://<node-ip>:<node-port>`.

### Running the Application with ArgoCD

**Prerequisites**

- Kubernetes cluster
- ArgoCD installed

1. **Login to ArgoCD**


2. **Add the Application to ArgoCD**:

    Create a new application in ArgoCD pointing to this repository.

    ```sh
   argocd app create tarot-shuffle-draw \
     --repo https://github.com/joshuamkite/tarot-card-shuffle-draw-web.git \
     --path helm/tarot-shuffle-draw \
     --dest-server https://kubernetes.default.svc \
     --dest-namespace default \
     --revision main
    ```

3. **Sync the Application**:

    ```sh
    argocd app sync tarot-shuffle-draw
    ```

4. **Access the Application**:

    ```sh
    kubectl get svc --namespace default
    ```

    Look for the `tarot-shuffle-draw` service and note the `NodePort`. Access the application at `http://<node-ip>:<node-port>`.

## Usage

### Web Interface

1. **Choose the deck type**: Full Deck/ Major Arcana only/ Minor Arcana only.
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

## Alternative Deployment Ports

There is a cross platform command line port (sorry, no illustrations!) available with binaries packaged for various operating systems - see [tarot-card-shuffle-draw](https://github.com/joshuamkite/tarot-card-shuffle-draw)

There is a serverless port deployed as AWS Lambda microservices orchestrated with API Gateway and backed with S3 and CloudFront - see [tarot-card-shuffle-draw-lambda](https://github.com/joshuamkite/tarot-card-shuffle-draw-lambda)