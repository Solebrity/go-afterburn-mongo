# !!! WARNING THIS IS EXPERIMENTAL !!!
Not an official [OpenFaaS](https://ghithub.com/openfaas/faas) project

# go-afterburn-mongo
An attempt to use OpenFaaS [go-afterburn](https://github.com/openfaas-incubator/go-afterburn) with a MongoDB connection, and the [of-watchdog](https://github.com/openfaas-incubator/of-watchdog).

# How it works
Because afterburn-based functions keep the target fprocess running, it is possible to have long-lived connections to other services such as databases, kafka streams, websockets, etc.  Taking advantage of this feature, a function can create a connection to an external service when its container is first started and maintain it for the lifetime of the function.

The template for the deployed function is uses the same template from [go-afterburn](https://github.com/openfaas-incubator/go-afterburn) which relies on the [of-watchdog](https://github.com/openfaas-incubator/of-watchdog)

For a detailed explanation of how afterburn works see [this blog post by Alex Ellis](https://blog.alexellis.io/openfaas-serverless-acceleration/)

# Setup
To run this example, you will need an accessible OpenFaaS deployment.  In your container scheduling platform, you'll also need a "secret" named `mongo-creds` with a key named `mongo-url` accessible from the container.  The `mongo-url` should contain the mongo connection string to an reachable MongoDB deployment.

Build, Push, Deploy, Invoke, Invoke, Invoke!!!  Use the [sample.json](mongo-function/sample.json) as your input.  The function should return the number of inserted documents on eacy successful call.