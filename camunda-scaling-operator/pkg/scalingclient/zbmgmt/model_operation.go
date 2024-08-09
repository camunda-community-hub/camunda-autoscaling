/*
Cluster Topology Management API

API for managing cluster membership and partition distribution.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zbmgmt

import (
	"encoding/json"
)

// checks if the Operation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Operation{}

// Operation struct for Operation
type Operation struct {
	Operation *string `json:"operation,omitempty"`
	// The ID of a broker, starting from 0
	BrokerId *int32 `json:"brokerId,omitempty"`
	// The ID of a partition, starting from 1
	PartitionId *int32 `json:"partitionId,omitempty"`
	// The priority of the partition
	Priority *int32  `json:"priority,omitempty"`
	Brokers  []int32 `json:"brokers,omitempty"`
	// The ID of an exporter
	ExporterId *string `json:"exporterId,omitempty"`
}

// NewOperation instantiates a new Operation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperation() *Operation {
	this := Operation{}
	return &this
}

// NewOperationWithDefaults instantiates a new Operation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationWithDefaults() *Operation {
	this := Operation{}
	return &this
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *Operation) GetOperation() string {
	if o == nil || IsNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetOperationOk() (*string, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *Operation) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *Operation) SetOperation(v string) {
	o.Operation = &v
}

// GetBrokerId returns the BrokerId field value if set, zero value otherwise.
func (o *Operation) GetBrokerId() int32 {
	if o == nil || IsNil(o.BrokerId) {
		var ret int32
		return ret
	}
	return *o.BrokerId
}

// GetBrokerIdOk returns a tuple with the BrokerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetBrokerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.BrokerId) {
		return nil, false
	}
	return o.BrokerId, true
}

// HasBrokerId returns a boolean if a field has been set.
func (o *Operation) HasBrokerId() bool {
	if o != nil && !IsNil(o.BrokerId) {
		return true
	}

	return false
}

// SetBrokerId gets a reference to the given int32 and assigns it to the BrokerId field.
func (o *Operation) SetBrokerId(v int32) {
	o.BrokerId = &v
}

// GetPartitionId returns the PartitionId field value if set, zero value otherwise.
func (o *Operation) GetPartitionId() int32 {
	if o == nil || IsNil(o.PartitionId) {
		var ret int32
		return ret
	}
	return *o.PartitionId
}

// GetPartitionIdOk returns a tuple with the PartitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetPartitionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PartitionId) {
		return nil, false
	}
	return o.PartitionId, true
}

// HasPartitionId returns a boolean if a field has been set.
func (o *Operation) HasPartitionId() bool {
	if o != nil && !IsNil(o.PartitionId) {
		return true
	}

	return false
}

// SetPartitionId gets a reference to the given int32 and assigns it to the PartitionId field.
func (o *Operation) SetPartitionId(v int32) {
	o.PartitionId = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *Operation) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *Operation) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *Operation) SetPriority(v int32) {
	o.Priority = &v
}

// GetBrokers returns the Brokers field value if set, zero value otherwise.
func (o *Operation) GetBrokers() []int32 {
	if o == nil || IsNil(o.Brokers) {
		var ret []int32
		return ret
	}
	return o.Brokers
}

// GetBrokersOk returns a tuple with the Brokers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetBrokersOk() ([]int32, bool) {
	if o == nil || IsNil(o.Brokers) {
		return nil, false
	}
	return o.Brokers, true
}

// HasBrokers returns a boolean if a field has been set.
func (o *Operation) HasBrokers() bool {
	if o != nil && !IsNil(o.Brokers) {
		return true
	}

	return false
}

// SetBrokers gets a reference to the given []int32 and assigns it to the Brokers field.
func (o *Operation) SetBrokers(v []int32) {
	o.Brokers = v
}

// GetExporterId returns the ExporterId field value if set, zero value otherwise.
func (o *Operation) GetExporterId() string {
	if o == nil || IsNil(o.ExporterId) {
		var ret string
		return ret
	}
	return *o.ExporterId
}

// GetExporterIdOk returns a tuple with the ExporterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetExporterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExporterId) {
		return nil, false
	}
	return o.ExporterId, true
}

// HasExporterId returns a boolean if a field has been set.
func (o *Operation) HasExporterId() bool {
	if o != nil && !IsNil(o.ExporterId) {
		return true
	}

	return false
}

// SetExporterId gets a reference to the given string and assigns it to the ExporterId field.
func (o *Operation) SetExporterId(v string) {
	o.ExporterId = &v
}

func (o Operation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Operation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.BrokerId) {
		toSerialize["brokerId"] = o.BrokerId
	}
	if !IsNil(o.PartitionId) {
		toSerialize["partitionId"] = o.PartitionId
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Brokers) {
		toSerialize["brokers"] = o.Brokers
	}
	if !IsNil(o.ExporterId) {
		toSerialize["exporterId"] = o.ExporterId
	}
	return toSerialize, nil
}

type NullableOperation struct {
	value *Operation
	isSet bool
}

func (v NullableOperation) Get() *Operation {
	return v.value
}

func (v *NullableOperation) Set(val *Operation) {
	v.value = val
	v.isSet = true
}

func (v NullableOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperation(val *Operation) *NullableOperation {
	return &NullableOperation{value: val, isSet: true}
}

func (v NullableOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}