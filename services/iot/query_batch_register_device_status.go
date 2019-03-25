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

// QueryBatchRegisterDeviceStatus invokes the iot.QueryBatchRegisterDeviceStatus API synchronously
// api document: https://help.aliyun.com/document_detail/69483.html
func (client *Client) QueryBatchRegisterDeviceStatus(request *QueryBatchRegisterDeviceStatusRequest) (response *QueryBatchRegisterDeviceStatusResponse, err error) {
	response = CreateQueryBatchRegisterDeviceStatusResponse()
	err = client.DoAction(request, response)
	return
}

// QueryBatchRegisterDeviceStatusRequest is the request struct for api QueryBatchRegisterDeviceStatus
type QueryBatchRegisterDeviceStatusRequest struct {
	*requests.RpcRequest
	ProductKey string           `position:"Query" name:"ProductKey"`
	ApplyId    requests.Integer `position:"Query" name:"ApplyId"`
}

// QueryBatchRegisterDeviceStatusResponse is the response struct for api QueryBatchRegisterDeviceStatus
type QueryBatchRegisterDeviceStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateQueryBatchRegisterDeviceStatusResponse creates a response to parse from QueryBatchRegisterDeviceStatus response
func CreateQueryBatchRegisterDeviceStatusResponse() (response *QueryBatchRegisterDeviceStatusResponse) {
	response = &QueryBatchRegisterDeviceStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

// CreateQueryBatchRegisterDeviceStatusRequest creates a request to invoke QueryBatchRegisterDeviceStatus request
func CreateQueryBatchRegisterDeviceStatusRequest() (request *QueryBatchRegisterDeviceStatusRequest) {
	request = &QueryBatchRegisterDeviceStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryBatchRegisterDeviceStatus", "iot", "openAPI")
	return
}
