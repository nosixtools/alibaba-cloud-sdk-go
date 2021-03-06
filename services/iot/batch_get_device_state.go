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
	"strconv"
)


// BatchGetDeviceState invokes the iot.BatchGetDeviceState API synchronously
// api document: https://help.aliyun.com/document_detail/69906.html
func (client *Client) BatchGetDeviceState(request *BatchGetDeviceStateRequest) (response *BatchGetDeviceStateResponse, err error) {
	response = CreateBatchGetDeviceStateResponse()
	err = client.DoAction(request, response)
	return
}

// BatchGetDeviceStateRequest is the request struct for api BatchGetDeviceState
type BatchGetDeviceStateRequest struct {
	*requests.RpcRequest
	IotId      string `position:"Query" name:"IotId"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

// SetDeviceName is the request struct for set device name
func (request *BatchGetDeviceStateRequest) SetDeviceName(deviceName []string) {
	if deviceName == nil || len(deviceName) == 0 {
		return
	}
	paramsMap := request.GetQueryParams()
	for i := 1; i <= len(deviceName); i++ {
		paramsMap["DeviceName."+strconv.Itoa(i)] = deviceName[i-1]
	}
}


// BatchGetDeviceStateResponse is the response struct for api BatchGetDeviceState
type BatchGetDeviceStateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBatchGetDeviceStateResponse creates a response to parse from BatchGetDeviceState response
func CreateBatchGetDeviceStateResponse() (response *BatchGetDeviceStateResponse) {
	response = &BatchGetDeviceStateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

// CreateBatchGetDeviceStateRequest creates a request to invoke BatchGetDeviceState request
func CreateBatchGetDeviceStateRequest() (request *BatchGetDeviceStateRequest) {
	request = &BatchGetDeviceStateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BatchGetDeviceState", "iot", "openAPI")
	return
}
