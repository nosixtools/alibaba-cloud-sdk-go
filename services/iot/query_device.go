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

// QueryDevice invokes the iot.QueryDevice API synchronously
// api document: https://help.aliyun.com/document_detail/69905.html
func (client *Client) QueryDevice(request *QueryDeviceRequest) (response *QueryDeviceResponse, err error) {
	response = CreateQueryDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceRequest is the request struct for api QueryDevice
type QueryDeviceRequest struct {
	*requests.RpcRequest
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	ProductKey  string           `position:"Query" name:"ProductKey"`
}

// QueryDeviceResponse is the response struct for api QueryDevice
type QueryDeviceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateQueryDeviceResponse creates a response to parse from QueryDevice response
func CreateQueryDeviceResponse() (response *QueryDeviceResponse) {
	response = &QueryDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

// CreateQueryDeviceRequest creates a request to invoke QueryDevice request
func CreateQueryDeviceRequest() (request *QueryDeviceRequest) {
	request = &QueryDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDevice", "iot", "openAPI")
	return
}
