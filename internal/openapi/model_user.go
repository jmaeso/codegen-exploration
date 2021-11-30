/*
Users service API

A REST API for creating, listing and deleting users.

API version: 0.0.1
Contact: jmaeso@enti.cat
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// User Main entity of the service
type User struct {
	// Unique identifier of the user.
	Id string `json:"id"`
	// nick name of the user.
	NickName string `json:"nickName"`
	// password for the user to login.
	Password string `json:"password"`
	// email of the user.
	Email string `json:"email"`
	// country of the user.
	Country *string `json:"country,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(id string, nickName string, password string, email string) *User {
	this := User{}
	this.Id = id
	this.NickName = nickName
	this.Password = password
	this.Email = email
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value
func (o *User) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *User) SetId(v string) {
	o.Id = v
}

// GetNickName returns the NickName field value
func (o *User) GetNickName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NickName
}

// GetNickNameOk returns a tuple with the NickName field value
// and a boolean to check if the value has been set.
func (o *User) GetNickNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NickName, true
}

// SetNickName sets field value
func (o *User) SetNickName(v string) {
	o.NickName = v
}

// GetPassword returns the Password field value
func (o *User) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *User) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *User) SetPassword(v string) {
	o.Password = v
}

// GetEmail returns the Email field value
func (o *User) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *User) SetEmail(v string) {
	o.Email = v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *User) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *User) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *User) SetCountry(v string) {
	o.Country = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["nickName"] = o.NickName
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

