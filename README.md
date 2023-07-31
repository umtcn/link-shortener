# Link Shortener Service

Link Shortener Service is a web application that provides a URL shortening service to transform long and complex URLs into shorter and more convenient links. This application is built using Golang and follows the Domain-Driven Design (DDD) principles. It provides an in-memory service to handle URL shortening and redirection using HTTP redirection.

## Getting Started

These instructions will help you set up and run the Link Shortener Service on your local machine for development and testing purposes.

### Prerequisites

- Golang should be installed on your machine.
- Docker and Kubernetes should be set up if you wish to deploy the service on Kubernetes.

### Installing

1. Clone the repository to your local machine:

git clone https://github.com/umtcn/link-shortener.git
cd link-shortener

2. Build the Docker image:


docker build -t your-docker-username/link-shortener:latest .

### Running the Service

To run the Link Shortener Service locally:

docker run -p 8080:8080 your-docker-username/link-shortener

The service will be accessible at `http://localhost:8080`.

### Usage

- To create a short link, make a POST request to the following endpoint:

POST http://localhost:8080/shorten
Request Body: { "url": "https://example.com/long-and-complex-url" }

- To access the original URL using the short link, make a GET request to the short link URL:

GET http://localhost:8080/{short-link}

### Kubernetes Deployment

To deploy the Link Shortener Service on Kubernetes, you can use the provided `link-shortener-deployment.yaml` file. Ensure you have Kubernetes set up and running. Then, apply the deployment:

kubectl apply -f link-shortener-deployment.yaml

The service will be accessible using Kubernetes service.

## Contributing

Pull requests and contributions are welcome. Please feel free to open an issue for any bug fixes, improvements, or new features.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.