```bash
gcloud config list
```

```bash
export GCP_PROJECT=$(gcloud config list --format="value(core.project)")
export GCP_REGION=$(gcloud config list --format="value(compute.region)")
```

echo $GCP_REGION $GCP_PROJECT

```bash
gcloud builds submit --tag gcr.io/$GCP_PROJECT/ws-server
```

export GCP_REGION=us-central1

docker run gcr.io/$GCP_PROJECT/ws-server
docker run gcr.io/aqueous-cargo-242610/ws-server


gcloud run deploy http-service \
  --image gcr.io/$GCP_PROJECT/http-server \
  --platform managed \
  --region $GCP_REGION \
  --allow-unauthenticated

```bash
gcloud run deploy ws-echo-service \
  --image gcr.io/$GCP_PROJECT/ws-server \
  --platform managed \
  --region $GCP_REGION \
  --allow-unauthenticated
```

```bash
export WS_ECHO_SERVICE_URL=$(gcloud run services describe ws-echo-service \
  --platform managed \
  --region us-central1 \
  --format="value(status.address.url)")
  
echo "WS_ECHO_SERVICE_URL: $WS_ECHO_SERVICE_URL"
```  