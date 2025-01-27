# Go API client for qase

You can use our API to access Qase.io API endpoints, which allows to retrieve information about entities stored in database and perform actions with them. The API is organized around REST. 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [https://qase.io](https://qase.io)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./qase"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.qase.io/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ProjectsApi* | [**ProjectList**](docs/ProjectsApi.md#projectlist) | **Get** /project | Get All Projects.

## Documentation For Models

 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse200Result](docs/InlineResponse200Result.md)
 - [Project](docs/Project.md)
 - [ProjectCounts](docs/ProjectCounts.md)
 - [ProjectCountsDefects](docs/ProjectCountsDefects.md)
 - [ProjectCountsRuns](docs/ProjectCountsRuns.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author

support@qase.io
