# Protocol Documentation
<a name="top"/>

## Table of Contents

- [api-gateway.proto](#api-gateway.proto)
    - [EchoRequest](#grpc.gomeetexamples.apigateway.EchoRequest)
    - [EchoResponse](#grpc.gomeetexamples.apigateway.EchoResponse)
    - [EmptyMessage](#grpc.gomeetexamples.apigateway.EmptyMessage)
    - [ProfileCreationRequest](#grpc.gomeetexamples.apigateway.ProfileCreationRequest)
    - [ProfileInfo](#grpc.gomeetexamples.apigateway.ProfileInfo)
    - [ProfileList](#grpc.gomeetexamples.apigateway.ProfileList)
    - [ProfileListRequest](#grpc.gomeetexamples.apigateway.ProfileListRequest)
    - [ProfileRequest](#grpc.gomeetexamples.apigateway.ProfileRequest)
    - [ProfileResponse](#grpc.gomeetexamples.apigateway.ProfileResponse)
    - [ServiceStatus](#grpc.gomeetexamples.apigateway.ServiceStatus)
    - [ServicesStatusList](#grpc.gomeetexamples.apigateway.ServicesStatusList)
    - [VersionResponse](#grpc.gomeetexamples.apigateway.VersionResponse)
  
    - [Genders](#grpc.gomeetexamples.apigateway.Genders)
    - [ServiceStatus.Status](#grpc.gomeetexamples.apigateway.ServiceStatus.Status)
  
  
    - [ApiGateway](#grpc.gomeetexamples.apigateway.ApiGateway)
  

- [Scalar Value Types](#scalar-value-types)



<a name="api-gateway.proto"/>
<p align="right"><a href="#top">Top</a></p>

## api-gateway.proto



<a name="grpc.gomeetexamples.apigateway.EchoRequest"/>

### EchoRequest
EchoRequest represents a simple message sent to the Echo service.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | Uuid represents the message identifier. |
| content | [string](#string) |  | some content |






<a name="grpc.gomeetexamples.apigateway.EchoResponse"/>

### EchoResponse
EchoResponse represents a simple message that the Echo service return.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | Uuid represents the message identifier. |
| content | [string](#string) |  | some content |






<a name="grpc.gomeetexamples.apigateway.EmptyMessage"/>

### EmptyMessage







<a name="grpc.gomeetexamples.apigateway.ProfileCreationRequest"/>

### ProfileCreationRequest
ProfileCreationRequest encodes a profile creation request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gender | [Genders](#grpc.gomeetexamples.apigateway.Genders) |  | profile role |
| email | [string](#string) |  | profile email |
| name | [string](#string) |  | profile name |
| birthday | [string](#string) |  | profile birthday |






<a name="grpc.gomeetexamples.apigateway.ProfileInfo"/>

### ProfileInfo
ProfileInfo encodes information about a profile.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | internal profile ID |
| gender | [Genders](#grpc.gomeetexamples.apigateway.Genders) |  | profile role |
| email | [string](#string) |  | profile email |
| name | [string](#string) |  | profile name |
| birthday | [string](#string) |  | profile birthday |
| created_at | [string](#string) |  | creation time (UTC - RFC 3339 format) |
| updated_at | [string](#string) |  | modification time (UTC - RFC 3339 format) |
| deleted_at | [string](#string) |  | deletion time (UTC - RFC 3339 format if the profile was logically deleted, empty otherwise) |






<a name="grpc.gomeetexamples.apigateway.ProfileList"/>

### ProfileList
ProfileList encodes the result of a ProfileListRequest.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result_set_size | [uint32](#uint32) |  | total number of results |
| has_more | [bool](#bool) |  | true if there are more results for the ProfileListRequest |
| profiles | [ProfileInfo](#grpc.gomeetexamples.apigateway.ProfileInfo) | repeated | list of ProfileInfo messages |






<a name="grpc.gomeetexamples.apigateway.ProfileListRequest"/>

### ProfileListRequest
ProfileListRequest encodes a set of criteria for the retrieval of a list of profiles.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_number | [uint32](#uint32) |  | page number (starting from 1) |
| page_size | [uint32](#uint32) |  | number of results in a page |
| order | [string](#string) |  | result ordering specification (default &#34;created_at asc&#34;) |
| exclude_soft_deleted | [bool](#bool) |  | if true, excludes logically-deleted profiles from the result set |
| soft_deleted_only | [bool](#bool) |  | if true, restricts the result set to logically-deleted profiles |
| gender | [Genders](#grpc.gomeetexamples.apigateway.Genders) |  | role to search for |






<a name="grpc.gomeetexamples.apigateway.ProfileRequest"/>

### ProfileRequest
ProfileRequest encodes a profile identifier.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | profile ID |






<a name="grpc.gomeetexamples.apigateway.ProfileResponse"/>

### ProfileResponse
ProfileResponse encodes the result of a profile operation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | indicates whether the operation (authentication, creation, update or delete) was successful |
| info | [ProfileInfo](#grpc.gomeetexamples.apigateway.ProfileInfo) |  | profile information (unreliable if the operation failed) |






<a name="grpc.gomeetexamples.apigateway.ServiceStatus"/>

### ServiceStatus
SeviceStatus represents a sub services status message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | name of service |
| version | [string](#string) |  | version of service |
| status | [ServiceStatus.Status](#grpc.gomeetexamples.apigateway.ServiceStatus.Status) |  | status of service see enum Status |
| e_msg | [string](#string) |  |  |






<a name="grpc.gomeetexamples.apigateway.ServicesStatusList"/>

### ServicesStatusList
ServicesStatusList is the sub services status list


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| services | [ServiceStatus](#grpc.gomeetexamples.apigateway.ServiceStatus) | repeated |  |






<a name="grpc.gomeetexamples.apigateway.VersionResponse"/>

### VersionResponse
VersionMessage represents a version message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Id represents the message identifier. |
| version | [string](#string) |  |  |





 


<a name="grpc.gomeetexamples.apigateway.Genders"/>

### Genders


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOW | 0 | normaly never |
| MALE | 1 | male gender |
| FEMALE | 2 | female gender |



<a name="grpc.gomeetexamples.apigateway.ServiceStatus.Status"/>

### ServiceStatus.Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 |  |
| UNAVAILABLE | 1 |  |


 

 


<a name="grpc.gomeetexamples.apigateway.ApiGateway"/>

### ApiGateway


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [EmptyMessage](#grpc.gomeetexamples.apigateway.EmptyMessage) | [VersionResponse](#grpc.gomeetexamples.apigateway.EmptyMessage) | Version method receives no paramaters and returns a version message. |
| ServicesStatus | [EmptyMessage](#grpc.gomeetexamples.apigateway.EmptyMessage) | [ServicesStatusList](#grpc.gomeetexamples.apigateway.EmptyMessage) | ServicesStatus method receives no paramaters and returns all services status message |
| Echo | [EchoRequest](#grpc.gomeetexamples.apigateway.EchoRequest) | [EchoResponse](#grpc.gomeetexamples.apigateway.EchoRequest) | Echo method receives a simple message and returns it from svc-echo. |
| CreateProfile | [ProfileCreationRequest](#grpc.gomeetexamples.apigateway.ProfileCreationRequest) | [ProfileResponse](#grpc.gomeetexamples.apigateway.ProfileCreationRequest) | CreateProfile attempts to create a new profile via svc-profile. |
| ReadProfile | [ProfileRequest](#grpc.gomeetexamples.apigateway.ProfileRequest) | [ProfileInfo](#grpc.gomeetexamples.apigateway.ProfileRequest) | ReadProfile returns information about an existing profile via svc-profile. |
| ListProfile | [ProfileListRequest](#grpc.gomeetexamples.apigateway.ProfileListRequest) | [ProfileList](#grpc.gomeetexamples.apigateway.ProfileListRequest) | ListProfile returns a list of profiles matching a set of criteria via svc-profile. |
| UpdateProfile | [ProfileInfo](#grpc.gomeetexamples.apigateway.ProfileInfo) | [ProfileResponse](#grpc.gomeetexamples.apigateway.ProfileInfo) | UpdateProfile attempts to update an existing profile via svc-profile. |
| DeleteProfile | [ProfileRequest](#grpc.gomeetexamples.apigateway.ProfileRequest) | [ProfileResponse](#grpc.gomeetexamples.apigateway.ProfileRequest) | DeleteProfile attempts to delete (logically) an existing profile via svc-profile. |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

