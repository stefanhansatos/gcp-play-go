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
export NAME=ws-echo
export GCP_PROJECT=$(gcloud config list --format="value(core.project)")
export GCP_REGION=$(gcloud config list --format="value(compute.region)")
export GCP_ZONE=$(gcloud config list --format="value(compute.zone)")

echo "NAME: $NAME
GCP_PROJECT: $GCP_PROJECT
GCP_REGION: $GCP_REGION
GCP_ZONE: $GCP_ZONE"
```

### Create custom image

Create compute engine with custom boot disk `play-go-debian` prepared for go deployment via git

```bash
gcloud beta compute --project=$GCP_PROJECT instances create ${NAME}-instance \
  --zone=europe-west1-b \
  --machine-type=f1-micro \
  --subnet=default \
  --network-tier=PREMIUM --maintenance-policy=MIGRATE \
  --service-account=53583873290-compute@developer.gserviceaccount.com \
  --scopes=https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring.write,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/trace.append \
  --tags=http-server \
  --image=play-go-debian \
  --image-project=aqueous-cargo-242610 \
  --boot-disk-size=10GB \
  --boot-disk-type=pd-standard \
  --boot-disk-device-name=ws-echo-instance \
  --reservation-affinity=any
```

Create firewall rule
```bash
gcloud compute --project=aqueous-cargo-242610 firewall-rules create allow-http-8080 \
  --direction=INGRESS \
  --priority=1000 \
  --network=default \
  --action=ALLOW \
  --rules=tcp:8080 \
  --source-ranges=0.0.0.0/0 \
  --target-tags=http-server
```


Connect via ssh to the instance
```bash
cd go/src/github.com
rm -rf *
git clone https://github.com/stefanhansatos/gcp-play-go.git
cd gcp-play-go/web-sockets-gorilla/ws-server

go mod init && go mod vendor
go build
./ws-server

```

Test the connection with the provided external IP of the instance
```bash
wscat -c ws://<provided ip>:8080

e.g.
wscat -c ws://35.195.255.145:8080
Connected (press CTRL+C to quit)
> hello
< hello
```

Stop the server and exit the instance.

Stop the instance
```bash
gcloud compute instances stop ${NAME}-instance \
  --zone=$GCP_ZONE
```


Create a custom image
```bash
gcloud compute images create ${NAME}-go-debian \
  --project=$GCP_PROJECT \
  --family=play-debian \
  --source-disk=${NAME}-instance \
  --source-disk-zone=$GCP_ZONE \
  --storage-location=$GCP_REGION
```


Add a startup script to the instance, i.e. add the custom metadata key `startup-script` 
with the following value:
```bash
#! /bin/bash

sudo ${HOME}/go/src/github.com/gcp-play-go/web-sockets-gorilla/ws-server/ws-server
```

Start the instance
```bash
gcloud compute instances start ${NAME}-instance \
  --zone=$GCP_ZONE
```

Test the connection with the provided external IP of the instance
```bash
wscat -c ws://<provided ip>:8080

e.g.
wscat -c ws://35.195.255.145:8080
Connected (press CTRL+C to quit)
> hello
< hello
```


Create an unmanaged instance group and add the newly created instance
```bash
gcloud compute instance-groups unmanaged create ${NAME}-instance-group \
  --project=$GCP_PROJECT \
  --zone=$GCP_ZONE

gcloud compute instance-groups unmanaged add-instances ${NAME}-instance-group \
  --project=$GCP_PROJECT \
  --zone=$GCP_ZONE \
  --instances=${NAME}-instance
```

gcloud compute instance-groups unmanaged describe ws-echo-instance-group
gcloud compute instance-groups unmanaged list-instances ws-echo-instance-group \
  --zone=$GCP_ZONE
gcloud compute instances list


Create HTTP(S) load balancer lb-ws-echo


wscat -c ws://34.107.171.105:8080	

gcloud compute health-checks describe ws-echo-healthcheck
gcloud compute backend-services describe backend-ws-echo --global
gcloud compute url-maps describe lb-ws-echo
gcloud compute target-http-proxies describe lb-ws-echo-target-proxy
gcloud compute forwarding-rules describe lb-ws-echo-forwarding-rule --global


gcloud compute forwarding-rules describe lb-ws-echo-forwarding-rule \
  --global --format="value(IPAddress)"
  
  
  
### Cleansing

Delete load balancer 
```bash
gcloud -q compute forwarding-rules delete lb-ws-echo-forwarding-rule --global
gcloud -q compute target-http-proxies delete lb-ws-echo-target-proxy
gcloud -q compute url-maps delete lb-ws-echo
gcloud -q compute backend-services delete backend-ws-echo --global
gcloud -q compute health-checks delete ws-echo-healthcheck
```

Delete instance group and instance
```bash
gcloud -q compute instance-groups unmanaged describe ws-echo-instance-group

gcloud -q compute instances stop ws-echo-instance
gcloud -q compute instances delete ws-echo-instance
```







