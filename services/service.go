/*
 * Copyright 2018-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package services

import (
	"fmt"
)

// Service represents a service bound to the application.
type Service struct {
	// BindingName is the binding name of this service.
	BindingName string `json:"binding_name"`

	// Credentials is the collection of credential keys.
	Credentials Credentials `json:"credentials"`

	// InstanceName is the instance name of this service.
	InstanceName string `json:"instance_name"`

	// Label is the type of service.
	Label string `json:"label"`

	// Plan is the plan type of this service.
	Plan string `json:"plan"`

	// Tags is the collection of tags of the service.
	Tags []string `json:"tags"`
}

// String makes Service satisfy the Stringer interface.
func (s Service) String() string {
	return fmt.Sprintf("Service{ BindingName: %s, Credentials: %s, InstanceName: %s, Label: %s, Plan: %s, Tags: %s }",
		s.BindingName, s.Credentials, s.InstanceName, s.Label, s.Plan, s.Tags)
}
