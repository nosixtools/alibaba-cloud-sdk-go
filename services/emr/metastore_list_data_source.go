package emr

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

// MetastoreListDataSource invokes the emr.MetastoreListDataSource API synchronously
// api document: https://help.aliyun.com/api/emr/metastorelistdatasource.html
func (client *Client) MetastoreListDataSource(request *MetastoreListDataSourceRequest) (response *MetastoreListDataSourceResponse, err error) {
	response = CreateMetastoreListDataSourceResponse()
	err = client.DoAction(request, response)
	return
}

// MetastoreListDataSourceWithChan invokes the emr.MetastoreListDataSource API asynchronously
// api document: https://help.aliyun.com/api/emr/metastorelistdatasource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetastoreListDataSourceWithChan(request *MetastoreListDataSourceRequest) (<-chan *MetastoreListDataSourceResponse, <-chan error) {
	responseChan := make(chan *MetastoreListDataSourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MetastoreListDataSource(request)
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

// MetastoreListDataSourceWithCallback invokes the emr.MetastoreListDataSource API asynchronously
// api document: https://help.aliyun.com/api/emr/metastorelistdatasource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetastoreListDataSourceWithCallback(request *MetastoreListDataSourceRequest, callback func(response *MetastoreListDataSourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MetastoreListDataSourceResponse
		var err error
		defer close(result)
		response, err = client.MetastoreListDataSource(request)
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

// MetastoreListDataSourceRequest is the request struct for api MetastoreListDataSource
type MetastoreListDataSourceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClusterReleased requests.Boolean `position:"Query" name:"ClusterReleased"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	SourceType      string           `position:"Query" name:"SourceType"`
	DataSourceName  string           `position:"Query" name:"DataSourceName"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
}

// MetastoreListDataSourceResponse is the response struct for api MetastoreListDataSource
type MetastoreListDataSourceResponse struct {
	*responses.BaseResponse
	RequestId      string                                  `json:"RequestId" xml:"RequestId"`
	TotalCount     int                                     `json:"TotalCount" xml:"TotalCount"`
	PageNumber     int                                     `json:"PageNumber" xml:"PageNumber"`
	PageSize       int                                     `json:"PageSize" xml:"PageSize"`
	DataSourceList DataSourceListInMetastoreListDataSource `json:"DataSourceList" xml:"DataSourceList"`
}

// CreateMetastoreListDataSourceRequest creates a request to invoke MetastoreListDataSource API
func CreateMetastoreListDataSourceRequest() (request *MetastoreListDataSourceRequest) {
	request = &MetastoreListDataSourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "MetastoreListDataSource", "emr", "openAPI")
	return
}

// CreateMetastoreListDataSourceResponse creates a response to parse from MetastoreListDataSource response
func CreateMetastoreListDataSourceResponse() (response *MetastoreListDataSourceResponse) {
	response = &MetastoreListDataSourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
