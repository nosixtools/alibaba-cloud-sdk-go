package iot

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)


// DisableThing invokes the iot.DisableThing API synchronously
// api document: https://help.aliyun.com/document_detail/69602.html
func (client *Client) DisableThing(request *DisableThingRequest) (response *DisableThingResponse, err error) {
	response = CreateDisableThingResponse()
	err = client.DoAction(request, response)
	return
}

// DisableThingRequest is the request struct for api DisableThing
type DisableThingRequest struct {
	*requests.RpcRequest
	IotId      string `position:"Query" name:"IotId"`
	ProductKey string `position:"Query" name:"ProductKey"`
	DeviceName string `position:"Query" name:"DeviceName"`
}

// DisableThingResponse is the response struct for api DisableThing
type DisableThingResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableThingResponse creates a response to parse from DisableThing response
func CreateDisableThingResponse() (response *DisableThingResponse) {
	response = &DisableThingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

// CreateDisableThingRequest creates a request to invoke DisableThing request
func CreateDisableThingRequest() (request *DisableThingRequest) {
	request = &DisableThingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DisableThing", "iot", "openAPI")
	return
}