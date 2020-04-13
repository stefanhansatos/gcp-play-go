### Intro

We establish a "Hello World" HTTP server implemented in Go running 
in a container in a managed Cloud Run service.

### Configuration

Check current GCP configuration
```bash
gcloud config list
```

Set variables accordingly
```bash
export NAME=hello-world
export GCP_PROJECT=$(gcloud config list --format="value(core.project)")
export GCP_REGION=$(gcloud config list --format="value(compute.region)")

echo "NAME: $NAME
GCP_PROJECT: $GCP_PROJECT
GCP_REGION: $GCP_REGION"
```
### Container

Create container image
```bash
gcloud builds submit --tag gcr.io/${GCP_PROJECT}/${NAME}-server
```

Deploy container as Cloud Run service
```bash
gcloud run deploy ${NAME}-service \
  --image gcr.io/${GCP_PROJECT}/${NAME}-server \
  --platform managed \
  --region $GCP_REGION \
  --allow-unauthenticated
```

### Service

Check service
```bash
export SERVICE_URL=$(gcloud run services describe ${NAME}-service \
  --platform managed \
  --region $GCP_REGION \
  --format="value(status.address.url)")
  
echo "SERVICE_URL: $SERVICE_URL"

curl $SERVICE_URL
``` 

### Troubleshooting

Run go code
```bash
go run main.go
```

Test from another terminal
```bash
curl <provided address>
```
<br>
---

Run container locally 
```bash
docker run gcr.io/${GCP_PROJECT}/${NAME}-server 
```

Test from another terminal
```bash
curl <local IP address>

e.g. 
curl 192.168.1.132:8080
```

### Cleansing


