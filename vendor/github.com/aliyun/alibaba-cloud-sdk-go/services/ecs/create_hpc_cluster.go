package ecs

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

// CreateHpcCluster invokes the ecs.CreateHpcCluster API synchronously
// api document: https://help.aliyun.com/api/ecs/createhpccluster.html
func (client *Client) CreateHpcCluster(request *CreateHpcClusterRequest) (response *CreateHpcClusterResponse, err error) {
	response = CreateCreateHpcClusterResponse()
	err = client.DoAction(request, response)
	return
}

// CreateHpcClusterWithChan invokes the ecs.CreateHpcCluster API asynchronously
// api document: https://help.aliyun.com/api/ecs/createhpccluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateHpcClusterWithChan(request *CreateHpcClusterRequest) (<-chan *CreateHpcClusterResponse, <-chan error) {
	responseChan := make(chan *CreateHpcClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateHpcCluster(request)
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

// CreateHpcClusterWithCallback invokes the ecs.CreateHpcCluster API asynchronously
// api document: https://help.aliyun.com/api/ecs/createhpccluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateHpcClusterWithCallback(request *CreateHpcClusterRequest, callback func(response *CreateHpcClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateHpcClusterResponse
		var err error
		defer close(result)
		response, err = client.CreateHpcCluster(request)
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

// CreateHpcClusterRequest is the request struct for api CreateHpcCluster
type CreateHpcClusterRequest struct {
	*requests.RpcRequest
}

// CreateHpcClusterResponse is the response struct for api CreateHpcCluster
type CreateHpcClusterResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	HpcClusterId string `json:"HpcClusterId" xml:"HpcClusterId"`
}

// CreateCreateHpcClusterRequest creates a request to invoke CreateHpcCluster API
func CreateCreateHpcClusterRequest() (request *CreateHpcClusterRequest) {
	request = &CreateHpcClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateHpcCluster", "ecs", "openAPI")
	return
}

// CreateCreateHpcClusterResponse creates a response to parse from CreateHpcCluster response
func CreateCreateHpcClusterResponse() (response *CreateHpcClusterResponse) {
	response = &CreateHpcClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}