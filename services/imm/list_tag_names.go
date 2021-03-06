package imm

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

// ListTagNames invokes the imm.ListTagNames API synchronously
// api document: https://help.aliyun.com/api/imm/listtagnames.html
func (client *Client) ListTagNames(request *ListTagNamesRequest) (response *ListTagNamesResponse, err error) {
	response = CreateListTagNamesResponse()
	err = client.DoAction(request, response)
	return
}

// ListTagNamesWithChan invokes the imm.ListTagNames API asynchronously
// api document: https://help.aliyun.com/api/imm/listtagnames.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTagNamesWithChan(request *ListTagNamesRequest) (<-chan *ListTagNamesResponse, <-chan error) {
	responseChan := make(chan *ListTagNamesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTagNames(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ListTagNamesWithCallback invokes the imm.ListTagNames API asynchronously
// api document: https://help.aliyun.com/api/imm/listtagnames.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTagNamesWithCallback(request *ListTagNamesRequest, callback func(response *ListTagNamesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTagNamesResponse
		var err error
		defer close(result)
		response, err = client.ListTagNames(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ListTagNamesRequest is the request struct for api ListTagNames
type ListTagNamesRequest struct {
	*requests.RpcRequest
	Marker  string `position:"Query" name:"Marker"`
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
}

// ListTagNamesResponse is the response struct for api ListTagNames
type ListTagNamesResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	NextMarker string     `json:"NextMarker" xml:"NextMarker"`
	Tags       []TagsItem `json:"Tags" xml:"Tags"`
}

// CreateListTagNamesRequest creates a request to invoke ListTagNames API
func CreateListTagNamesRequest() (request *ListTagNamesRequest) {
	request = &ListTagNamesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "ListTagNames", "imm", "openAPI")
	return
}

// CreateListTagNamesResponse creates a response to parse from ListTagNames response
func CreateListTagNamesResponse() (response *ListTagNamesResponse) {
	response = &ListTagNamesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
