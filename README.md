# Hello World Go Lambda - Post Body

## Abstract

Couldn't find many examples on how to post to a Lambda from Go (golang) with serverless and api-gateway.  Worked through a few of the currently less documented features and created a working example.

### Install
```
make build
sls deploy
```

### Utilization

Via [Postman](https://www.getpostman.com) or some other client of your choice, POST to the aws endpoint serverless returns (ie: https://xxx.execute-api.us-east-1.amazonaws.com/dev/hello) with a body as follows:

```
{"name":"Shawn"}
```

Response will resemble below

```
{
    "message": "Hello Shawn thank you for calling me!"
}
```
