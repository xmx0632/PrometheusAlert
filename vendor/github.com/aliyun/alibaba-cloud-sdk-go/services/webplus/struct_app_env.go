package webplus

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

// AppEnv is a nested struct in webplus response
type AppEnv struct {
	EnvId                string `json:"EnvId" xml:"EnvId"`
	EnvName              string `json:"EnvName" xml:"EnvName"`
	EnvDescription       string `json:"EnvDescription" xml:"EnvDescription"`
	CreateUsername       string `json:"CreateUsername" xml:"CreateUsername"`
	UpdateUsername       string `json:"UpdateUsername" xml:"UpdateUsername"`
	CreateTime           int64  `json:"CreateTime" xml:"CreateTime"`
	UpdateTime           int64  `json:"UpdateTime" xml:"UpdateTime"`
	StackId              string `json:"StackId" xml:"StackId"`
	StackName            string `json:"StackName" xml:"StackName"`
	AppName              string `json:"AppName" xml:"AppName"`
	AppId                string `json:"AppId" xml:"AppId"`
	ApplyingChange       bool   `json:"ApplyingChange" xml:"ApplyingChange"`
	AbortingChange       bool   `json:"AbortingChange" xml:"AbortingChange"`
	EnvType              string `json:"EnvType" xml:"EnvType"`
	PkgVersionId         string `json:"PkgVersionId" xml:"PkgVersionId"`
	PkgVersionLabel      string `json:"PkgVersionLabel" xml:"PkgVersionLabel"`
	EnvStatus            string `json:"EnvStatus" xml:"EnvStatus"`
	LastEnvStatus        string `json:"LastEnvStatus" xml:"LastEnvStatus"`
	StorageBase          string `json:"StorageBase" xml:"StorageBase"`
	DataRoot             string `json:"DataRoot" xml:"DataRoot"`
	LatestChangeId       string `json:"LatestChangeId" xml:"LatestChangeId"`
	ChangeBanner         string `json:"ChangeBanner" xml:"ChangeBanner"`
	CategoryName         string `json:"CategoryName" xml:"CategoryName"`
	TotalInstances       int64  `json:"TotalInstances" xml:"TotalInstances"`
	LogBase              string `json:"LogBase" xml:"LogBase"`
	PkgVersionStorageKey string `json:"PkgVersionStorageKey" xml:"PkgVersionStorageKey"`
}