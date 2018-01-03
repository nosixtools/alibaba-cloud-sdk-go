package rds

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

func (client *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
	response = CreateDescribeAccountsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeAccountsWithChan(request *DescribeAccountsRequest) (<-chan *DescribeAccountsResponse, <-chan error) {
	responseChan := make(chan *DescribeAccountsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAccounts(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeAccountsWithCallback(request *DescribeAccountsRequest, callback func(response *DescribeAccountsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAccountsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAccounts(request)
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

type DescribeAccountsRequest struct {
	*requests.RpcRequest
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	AccountName          string           `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeAccountsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Accounts  struct {
		DBInstanceAccount []struct {
			DBInstanceId       string `json:"DBInstanceId" xml:"DBInstanceId"`
			AccountName        string `json:"AccountName" xml:"AccountName"`
			AccountStatus      string `json:"AccountStatus" xml:"AccountStatus"`
			AccountType        string `json:"AccountType" xml:"AccountType"`
			AccountDescription string `json:"AccountDescription" xml:"AccountDescription"`
			DatabasePrivileges struct {
				DatabasePrivilege []struct {
					DBName           string `json:"DBName" xml:"DBName"`
					AccountPrivilege string `json:"AccountPrivilege" xml:"AccountPrivilege"`
				} `json:"DatabasePrivilege" xml:"DatabasePrivilege"`
			} `json:"DatabasePrivileges" xml:"DatabasePrivileges"`
		} `json:"DBInstanceAccount" xml:"DBInstanceAccount"`
	} `json:"Accounts" xml:"Accounts"`
}

func CreateDescribeAccountsRequest() (request *DescribeAccountsRequest) {
	request = &DescribeAccountsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeAccounts", "", "")
	return
}

func CreateDescribeAccountsResponse() (response *DescribeAccountsResponse) {
	response = &DescribeAccountsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}