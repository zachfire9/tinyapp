# Tiny App

## Run the app locally

go run main.go

## Docker

docker build -t tinyapp .

docker run -p 8080:8080 tinyapp

## Endpoints

http://localhost:8080/

http://localhost:8080/health

## Run the app in AWS

Setup terraform locally by running the commands [here](./terraform/)

Add an env var for the ecr image pointing to the image in your AWS instance (replace `<aws link>`)

```
export TF_VAR_ecr_image="<aws link>/tinyapp:latest"
```

Create the cluster and pod: `terraform apply`

When finished delete everything: `terraform destroy`