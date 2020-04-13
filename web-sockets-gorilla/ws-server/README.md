### Configuration

Check current GCP configuration
```bash
gcloud config list
```

Set variables accordingly
```bash
export NAME=ws-echo
export GCP_PROJECT=$(gcloud config list --format="value(core.project)")
export GCP_REGION=$(gcloud config list --format="value(compute.region)")

echo "GCP_PROJECT: $GCP_PROJECT\nGCP_REGION: $GCP_REGION"
```
### Container

Create container image
```bash
gcloud builds submit --tag gcr.io/${GCP_PROJECT}/${NAME}-server
```


### Troubleshooting

Run go code 
```bash
docker run gcr.io/${GCP_PROJECT}/${NAME}-server
```