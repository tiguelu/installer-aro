package ddoscoo

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

// DescribeWebInstanceRelations invokes the ddoscoo.DescribeWebInstanceRelations API synchronously
func (client *Client) DescribeWebInstanceRelations(request *DescribeWebInstanceRelationsRequest) (response *DescribeWebInstanceRelationsResponse, err error) {
	response = CreateDescribeWebInstanceRelationsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWebInstanceRelationsWithChan invokes the ddoscoo.DescribeWebInstanceRelations API asynchronously
func (client *Client) DescribeWebInstanceRelationsWithChan(request *DescribeWebInstanceRelationsRequest) (<-chan *DescribeWebInstanceRelationsResponse, <-chan error) {
	responseChan := make(chan *DescribeWebInstanceRelationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWebInstanceRelations(request)
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

// DescribeWebInstanceRelationsWithCallback invokes the ddoscoo.DescribeWebInstanceRelations API asynchronously
func (client *Client) DescribeWebInstanceRelationsWithCallback(request *DescribeWebInstanceRelationsRequest, callback func(response *DescribeWebInstanceRelationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWebInstanceRelationsResponse
		var err error
		defer close(result)
		response, err = client.DescribeWebInstanceRelations(request)
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

// DescribeWebInstanceRelationsRequest is the request struct for api DescribeWebInstanceRelations
type DescribeWebInstanceRelationsRequest struct {
	*requests.RpcRequest
	Domains         *[]string `position:"Query" name:"Domains"  type:"Repeated"`
	ResourceGroupId string    `position:"Query" name:"ResourceGroupId"`
	SourceIp        string    `position:"Query" name:"SourceIp"`
}

// DescribeWebInstanceRelationsResponse is the response struct for api DescribeWebInstanceRelations
type DescribeWebInstanceRelationsResponse struct {
	*responses.BaseResponse
	RequestId            string                `json:"RequestId" xml:"RequestId"`
	WebInstanceRelations []WebInstanceRelation `json:"WebInstanceRelations" xml:"WebInstanceRelations"`
}

// CreateDescribeWebInstanceRelationsRequest creates a request to invoke DescribeWebInstanceRelations API
func CreateDescribeWebInstanceRelationsRequest() (request *DescribeWebInstanceRelationsRequest) {
	request = &DescribeWebInstanceRelationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeWebInstanceRelations", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeWebInstanceRelationsResponse creates a response to parse from DescribeWebInstanceRelations response
func CreateDescribeWebInstanceRelationsResponse() (response *DescribeWebInstanceRelationsResponse) {
	response = &DescribeWebInstanceRelationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}