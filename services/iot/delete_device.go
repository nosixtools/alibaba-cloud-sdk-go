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


// DeleteDevice invokes the iot.DeleteDevice API synchronously
// api document: https://help.aliyun.com/document_detail/69281.html
func (client *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
	response = CreateDeleteDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDeviceRequest is the request struct for api DeleteDevice
type DeleteDeviceRequest struct {
	*requests.RpcRequest
	IotId      string `position:"Query" name:"IotId"`
	ProductKey string `position:"Query" name:"ProductKey"`
	DeviceName string `position:"Query" name:"DeviceName"`
}

// DeleteDeviceResponse is the response struct for api DeleteDevice
type DeleteDeviceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDeviceResponse creates a response to parse from DeleteDevice response
func CreateDeleteDeviceResponse() (response *DeleteDeviceResponse) {
	response = &DeleteDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

// CreateDeleteDeviceRequest creates a request to invoke DeleteDevice request
func CreateDeleteDeviceRequest() (request *DeleteDeviceRequest) {
	request = &DeleteDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DeleteDevice", "iot", "openAPI")
	return
}
