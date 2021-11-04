# HTTP Service - Starter

Based on **Golang Gin** > https://github.com/gin-gonic/gin

> Provide a Controller implementation.


## Google Cloud - Google Cloud Run 

Google Cloud Project configuration

```
export PROJECT_NUMBER="$(gcloud projects describe $(gcloud config get-value project) --format='value(projectNumber)')" 
```

```
export PROJECT=[YOUR-PROJECT-ID]
```

``` 
export SERVICE=[SERVICE-NAME]
```

``` 
export REGION=[REGION]
```

``` 
export IMAGE_URL="gcr.io/$(gcloud config get-value project)/editor-service" 
``` 

* IAM

``` 
gcloud projects add-iam-policy-binding $(gcloud config get-value project) \
    --member="serviceAccount:service-${PROJECT_NUMBER}@gcp-sa-pubsub.iam.gserviceaccount.com"\
    --role='roles/iam.serviceAccountTokenCreator' 
```

* Enabling Cloud Build
  
```
gcloud services enable cloudbuild.googleapis.com
```

* Enabling Cloud Run
  
```
gcloud services enable run.googleapis.com
```     

* Cloud Run configuration
  
```
gcloud config set run/platform managed
```

```
gcloud config set run/region ${REGION}
```

* BUILD & DEPLOY
  
```
gcloud builds submit --tag ${IMAGE_URL} 
```  

``` 
gcloud beta run deploy ${SERVICE} --image ${IMAGE_URL} --platform managed --no-allow-unauthenticated
``` 

