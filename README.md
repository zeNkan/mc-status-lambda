# mc-status-lambda
AWS Lambda Function Fetching MC Status

## Building instructions
**NOTE** the following instructions require that you first provide docker with
credentials to ECR

### 1. Build the container locally
```shell
$ ls
$ docker build -t minecraft-status:latest . 
...
```

### 2. Tag the newly built image
```shell
$ docker tag minecraft-status:latest {AWS_ACCOUNT_ID}.dkr.ecr.{AWS_REGION}.amazonaws.com/minecraft-status:latest
```

### 3. Push the image to ECR
```shell
$ docker push {AWS_ACCOUNT_ID}.dkr.ecr.{AWS_REGION}.amazonaws.com/minecraft-status:latest
```
