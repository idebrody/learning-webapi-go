# learning-webapi-go
Learning how to write a webapi in golang

Using this exercise to learn about golang and how to write a web api in it. 

Based on 
https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj


It is a very simple API, exposes add, substract, division and random at /api/v1 
It also contains:

* /metrics endpoint for prometheus, 
* /liveness  and /readiness for kubernetes healthchecks
* /docs/ endpoint for documentation using Swagger (swaggo/swag)


# Unit testing
```
go test -v 
```
or with a docker image
```
docker run -v ${PWD}:/testme  golang:1.17-buster bash -c "cd /testme && go test -v "
```

# Building

```
docker build -t idebrody/go-web-api .
```

# Running with Docker

```
docker run -p 8080:8080 idebrody/go-web-api
```

# Running it in Kubernetes

There is a combined manifest at kubernetes/go-web-api-nginx.yaml
* Deployment with healthchecks
* Service
* HorizontalPodAutoscaler
* Ingress, nginx ingressClassName, the route is set to go-web-api.dev.engineering.somecompany.cloud, which you might need to change according to your hosted domain

Tested on Kubernetes version 1.21.3

**Using your default namespace**
```
kubectl apply -f kubernetes/go-web-api-nginx.yaml
```

**Using idebrody as a namespace**
```
kubectl apply -f kubernetes/namespace.yaml
kubectl apply -n idebrody -f kubernetes/go-web-api-nginx.yaml
```

# Availabe Endpoints and Examples

Assuming you are running the service on your localhost:8080

**Documentation**

http://localhost:8080/docs/

**ADD**

Add together 2 numbers

http://localhost:8080/api/v1/add?num1=13&num2=3.1415926

**Division**

Divide 2 numbers

http://localhost:8080/api/v1/division?num1=9.27&num2=3

**Random**

Generate 23 random numbers, default is 10

http://localhost:8080/api/v1/random?Count=23

**Substract**

Substract 2 numbers

http://localhost:8080/api/v1/substract?num1=9&num2=1.414

**Metrics**

Prometheus metrics

http://localhost:8080/metrics

**Liveliness**

It has no output it only returns StatusOK or StatusInternalServerError

http://localhost:8080/liveness

**Readiness**

It has no output it only returns StatusOK or StatusInternalServerError

http://localhost:8080/readiness
