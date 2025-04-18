# Go API client for swagger

This is a sample Flower Shop server.   This includes the Billing REST APIs. You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/). 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://virtserver.swaggerhub.com/Charnley-Test/Billing-Flower-Shop/2.0.0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BillApi* | [**BillIdGet**](docs/BillApi.md#billidget) | **Get** /bill/{id} | Get a bill.
*BillApi* | [**BillPost**](docs/BillApi.md#billpost) | **Post** /bill | Submit a bill.
*ProcessApi* | [**ProcessPost**](docs/ProcessApi.md#processpost) | **Post** /process | Process a bill.
*ReceiptApi* | [**ReceiptIdGet**](docs/ReceiptApi.md#receiptidget) | **Get** /receipt/{id} | Get a receipt.
*ReceiptApi* | [**ReceiptPost**](docs/ReceiptApi.md#receiptpost) | **Post** /receipt | Create a receipt.

## Documentation For Models

 - [Address](docs/Address.md)
 - [Bill](docs/Bill.md)
 - [Person](docs/Person.md)
 - [Receipt](docs/Receipt.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author

rosemary.charnley@smartbear.com
