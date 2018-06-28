// ************************************************************************
// * Copyright 2018 OSIsoft, LLC
// * Licensed under the Apache License, Version 2.0 (the "License");
// * you may not use this file except in compliance with the License.
// * You may obtain a copy of the License at
// * 
// *   <http://www.apache.org/licenses/LICENSE-2.0>
// * 
// * Unless required by applicable law or agreed to in writing, software
// * distributed under the License is distributed on an "AS IS" BASIS,
// * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// * See the License for the specific language governing permissions and
// * limitations under the License.
// ************************************************************************

package gowebapi

import (
	"time"
)

type StreamAnnotation struct {
	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Value *interface{} `json:"Value,omitempty"`

	Creator string `json:"Creator,omitempty"`

	CreationDate time.Time `json:"CreationDate,omitempty"`

	Modifier string `json:"Modifier,omitempty"`

	ModifyDate time.Time `json:"ModifyDate,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
