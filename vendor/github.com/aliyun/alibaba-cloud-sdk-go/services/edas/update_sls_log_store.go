package edas

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

// UpdateSlsLogStore invokes the edas.UpdateSlsLogStore API synchronously
// api document: https://help.aliyun.com/api/edas/updateslslogstore.html
func (client *Client) UpdateSlsLogStore(request *UpdateSlsLogStoreRequest) (response *UpdateSlsLogStoreResponse, err error) {
	response = CreateUpdateSlsLogStoreResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSlsLogStoreWithChan invokes the edas.UpdateSlsLogStore API asynchronously
// api document: https://help.aliyun.com/api/edas/updateslslogstore.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateSlsLogStoreWithChan(request *UpdateSlsLogStoreRequest) (<-chan *UpdateSlsLogStoreResponse, <-chan error) {
	responseChan := make(chan *UpdateSlsLogStoreResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSlsLogStore(request)
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

// UpdateSlsLogStoreWithCallback invokes the edas.UpdateSlsLogStore API asynchronously
// api document: https://help.aliyun.com/api/edas/updateslslogstore.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateSlsLogStoreWithCallback(request *UpdateSlsLogStoreRequest, callback func(response *UpdateSlsLogStoreResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSlsLogStoreResponse
		var err error
		defer close(result)
		response, err = client.UpdateSlsLogStore(request)
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

// UpdateSlsLogStoreRequest is the request struct for api UpdateSlsLogStore
type UpdateSlsLogStoreRequest struct {
	*requests.RoaRequest
	Configs string `position:"Body" name:"Configs"`
	AppId   string `position:"Body" name:"AppId"`
}

// UpdateSlsLogStoreResponse is the response struct for api UpdateSlsLogStore
type UpdateSlsLogStoreResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Code      int    `json:"Code" xml:"Code"`
}

// CreateUpdateSlsLogStoreRequest creates a request to invoke UpdateSlsLogStore API
func CreateUpdateSlsLogStoreRequest() (request *UpdateSlsLogStoreRequest) {
	request = &UpdateSlsLogStoreRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UpdateSlsLogStore", "/pop/v5/k8s/sls/update_sls_log_store", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateSlsLogStoreResponse creates a response to parse from UpdateSlsLogStore response
func CreateUpdateSlsLogStoreResponse() (response *UpdateSlsLogStoreResponse) {
	response = &UpdateSlsLogStoreResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}