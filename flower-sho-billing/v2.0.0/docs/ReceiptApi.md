# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/Charnley-Test/Billing-Flower-Shop/2.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReceiptIdGet**](ReceiptApi.md#ReceiptIdGet) | **Get** /receipt/{id} | Get a receipt.
[**ReceiptPost**](ReceiptApi.md#ReceiptPost) | **Post** /receipt | Create a receipt.

# **ReceiptIdGet**
> Receipt ReceiptIdGet(ctx, id)
Get a receipt.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Receipt id. | 

### Return type

[**Receipt**](Receipt.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReceiptPost**
> Receipt ReceiptPost(ctx, body)
Create a receipt.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Receipt**](Receipt.md)|  | 

### Return type

[**Receipt**](Receipt.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

