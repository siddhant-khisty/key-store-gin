This is a simple key-value store made using the Gin framework, and simulates the data a computer retail shop would need 

## Avaliable API endpoints 
 
The data is stored directly to memory, which means there is no data persistance. Once you close the server, the data is lost for good.

Below are the endpoints that exist for the current version of the app. Use [curl](https://curl.se/), or a tool such as [Postman](https://www.postman.com/).

Assuming that you are compiling the binary and running locally, the app will run on `localhost:9000`.

- `/getAll`: GET call that retrives all the values stored in-memory when the call is made.

  `curl localhost:9000/getAll`

- `/get/{id}`: GET call that retrives a particular entry based on the ID value of the entry.
               For example, if the element has a ID of `21`, and you do a `curl localhost:9000/get/21`,
               you will get all the attributes for that element.

  `curl localhost:9000/get/21`
  
- `/search`: GET method that searches for a certain ID string as either a prefix or suffix which is provided using a query(?) in the URL

    Assuming we have the following ID's stored `AAC4326`,`AA926`, `BLO004` and `GVP021404`

--- `/search?prefix=AA`: This will return the data associated with IDs `AAC4326` and `AA926`
  
 `curl localhost:9000/search?prefix=AA`
  
 --- `/search?suffix=02`: This will return the data associated with IDs `BLO004` and `GVP021404`
  
  `curl localhost:9000/search?suffix=02`
  
- `/key/sey`: POST call that adds data to the memory. Please make sure that the body is formatted correctly and is a JSON array.
              If not, you will get a 400 error due to an invalid body. You can make sure of the `seed_data.json` file and enter data in bulk with it.

`curl -X POST localhost:9000/key/set -d '@seed_data.json'` - Make sure you're in the project's root directory when running this command to avoid path mismatch errors

# Running the Project
## Building from source and running locally

In order to build this project and run it locally, you will need the [Go binary](https://go.dev/doc/install) installed.

Once you have the binary installed, clone this project locally using `git clone https://github.com/siddhant-khisty/key-store-gin.git` or create a fork and clone the fork

`cd` into the project's root directory.

Use `go mod download` to download all the dependencies of the project

Run `go build -o server` to build a single binary file for starting the server
This will create a `server` executable file. Simply run `./server` from your terminal to start the server. Make sure you're path is correct.
    
OR

Alternatively, you can simply run `go run main.go` to start a live server. This is best used for development purposes and is not recommended for production.

## Running as a Docker container

If you'd like to run this app as a Docker container, there are a few different ways you can do this.

1. Pull the image from DockerHub and get the container running locally
    `docker run siddhantkhisty/gin-kv:latest`
2. Build the image from the Dockerfile and then run locally
    `docker build ./ -t [custom-image-name]:[custom-tag-number]`
        AND
    `docker start [image-id]`

## Running inside a Kubernetes cluster

If you're feeling extra daring, and want to run this inside a K8S cluster, we've got you covered.

All you need, is a running k8s cluster, and this repository itself.

Assuming that you already have a cluster, all you need to do is run this single command. Make sure you're in the project's root directory
``` kubectl apply -f ./Kubernetes/deply.yaml ./Kubernetes/ingress.yaml ./Kubernetes/svc.yaml```





