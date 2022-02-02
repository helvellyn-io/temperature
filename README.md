# Demo App

# Configure
    - set environment variables
      - export REGISTRY= "docker.index.io" # example
      - export REPO= "helvellyn/demo" # example
      - export VERSION=eg. "v1.0" # example

# Build 
    - make build-app 
      - builds the demo application
    - make push-app
      - push image to registry 

# Deployment 
    - active triggers
      - on new image is **TRUE
      - on Git commit(main) is **TRUE
    - make deploy
      - manual deploy to development environment 
    - Orchestration
      - shorturl.at/rvxJ4

# Observe 
    - APM
      -  shorturl.at/hrxLW
    
# Troubleshooting
    - HTTP 404
      - Check URL
      - Check main.go route is not misconfigured
    - HTTP 500
      - Check application logs.
      - Check kubernetes svc is exposed and healthy
    - Latency
      - Check CPU/MEM load for workloads.
      - Check Replica count is adequate.
    - Timeout
    - DNS 



