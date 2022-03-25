//
//Copyright The Kubernetes Authors.
//
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

// This file was autogenerated by go-to-protobuf. Do not edit it manually!

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: k8s.io/api/authentication/v1/generated.proto

package v1

import (
	"github.com/sourcegraph/k8s/apis/meta/v1"
	_ "github.com/sourcegraph/k8s/runtime"
	_ "github.com/sourcegraph/k8s/runtime/schema"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BoundObjectReference is a reference to an object that a token is bound to.
type BoundObjectReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kind of the referent. Valid kinds are 'Pod' and 'Secret'.
	// +optional
	Kind *string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	// API version of the referent.
	// +optional
	APIVersion *string `protobuf:"bytes,2,opt,name=aPIVersion" json:"aPIVersion,omitempty"`
	// Name of the referent.
	// +optional
	Name *string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// UID of the referent.
	// +optional
	UID *string `protobuf:"bytes,4,opt,name=uID" json:"uID,omitempty"`
}

func (x *BoundObjectReference) Reset() {
	*x = BoundObjectReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoundObjectReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoundObjectReference) ProtoMessage() {}

func (x *BoundObjectReference) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoundObjectReference.ProtoReflect.Descriptor instead.
func (*BoundObjectReference) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_authentication_v1_generated_proto_rawDescGZIP(), []int{0}
}

func (x *BoundObjectReference) GetKind() string {
	if x != nil && x.Kind != nil {
		return *x.Kind
	}
	return ""
}

func (x *BoundObjectReference) GetAPIVersion() string {
	if x != nil && x.APIVersion != nil {
		return *x.APIVersion
	}
	return ""
}

func (x *BoundObjectReference) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *BoundObjectReference) GetUID() string {
	if x != nil && x.UID != nil {
		return *x.UID
	}
	return ""
}

// ExtraValue masks the value so protobuf can generate
// +protobuf.nullable=true
// +protobuf.options.(gogoproto.goproto_stringer)=false
type ExtraValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []string `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (x *ExtraValue) Reset() {
	*x = ExtraValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtraValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtraValue) ProtoMessage() {}

func (x *ExtraValue) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtraValue.ProtoReflect.Descriptor instead.
func (*ExtraValue) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_authentication_v1_generated_proto_rawDescGZIP(), []int{1}
}

func (x *ExtraValue) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

// TokenRequest requests a token for a given service account.
type TokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Metadata *v1.ObjectMeta    `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *TokenRequestSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	// +optional
	Status *TokenRequestStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (x *TokenRequest) Reset() {
	*x = TokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRequest) ProtoMessage() {}

func (x *TokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRequest.ProtoReflect.Descriptor instead.
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_authentication_v1_generated_proto_rawDescGZIP(), []int{2}
}

func (x *TokenRequest) GetMetadata() *v1.ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *TokenRequest) GetSpec() *TokenRequestSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *TokenRequest) GetStatus() *TokenRequestStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// TokenRequestSpec contains client provided parameters of a token request.
type TokenRequestSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Audiences are the intendend audiences of the token. A recipient of a
	// token must identitfy themself with an identifier in the list of
	// audiences of the token, and otherwise should reject the token. A
	// token issued for multiple audiences may be used to authenticate
	// against any of the audiences listed but implies a high degree of
	// trust between the target audiences.
	Audiences []string `protobuf:"bytes,1,rep,name=audiences" json:"audiences,omitempty"`
	// ExpirationSeconds is the requested duration of validity of the request. The
	// token issuer may return a token with a different validity duration so a
	// client needs to check the 'expiration' field in a response.
	// +optional
	ExpirationSeconds *int64 `protobuf:"varint,4,opt,name=expirationSeconds" json:"expirationSeconds,omitempty"`
	// BoundObjectRef is a reference to an object that the token will be bound to.
	// The token will only be valid for as long as the bound object exists.
	// NOTE: The API server's TokenReview endpoint will validate the
	// BoundObjectRef, but other audiences may not. Keep ExpirationSeconds
	// small if you want prompt revocation.
	// +optional
	BoundObjectRef *BoundObjectReference `protobuf:"bytes,3,opt,name=boundObjectRef" json:"boundObjectRef,omitempty"`
}

func (x *TokenRequestSpec) Reset() {
	*x = TokenRequestSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRequestSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRequestSpec) ProtoMessage() {}

func (x *TokenRequestSpec) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRequestSpec.ProtoReflect.Descriptor instead.
func (*TokenRequestSpec) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_authentication_v1_generated_proto_rawDescGZIP(), []int{3}
}

func (x *TokenRequestSpec) GetAudiences() []string {
	if x != nil {
		return x.Audiences
	}
	return nil
}

func (x *TokenRequestSpec) GetExpirationSeconds() int64 {
	if x != nil && x.ExpirationSeconds != nil {
		return *x.ExpirationSeconds
	}
	return 0
}

func (x *TokenRequestSpec) GetBoundObjectRef() *BoundObjectReference {
	if x != nil {
		return x.BoundObjectRef
	}
	return nil
}

// TokenRequestStatus is the result of a token request.
type TokenRequestStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token is the opaque bearer token.
	Token *string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	// ExpirationTimestamp is the time of expiration of the returned token.
	ExpirationTimestamp *v1.Time `protobuf:"bytes,2,opt,name=expirationTimestamp" json:"expirationTimestamp,omitempty"`
}

func (x *TokenRequestStatus) Reset() {
	*x = TokenRequestStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRequestStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRequestStatus) ProtoMessage() {}

func (x *TokenRequestStatus) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRequestStatus.ProtoReflect.Descriptor instead.
func (*TokenRequestStatus) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_authentication_v1_generated_proto_rawDescGZIP(), []int{4}
}

func (x *TokenRequestStatus) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *TokenRequestStatus) GetExpirationTimestamp() *v1.Time {
	if x != nil {
		return x.ExpirationTimestamp
	}
	return nil
}

// TokenReview attempts to authenticate a token to a known user.
// Note: TokenReview requests may be cached by the webhook token authenticator
// plugin in the kube-apiserver.
type TokenReview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Spec holds information about the request being evaluated
	Spec *TokenReviewSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	// Status is filled in by the server and indicates whether the request can be authenticated.
	// +optional
	Status *TokenReviewStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (x *TokenReview) Reset() {
	*x = TokenReview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenReview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenReview) ProtoMessage() {}

func (x *TokenReview) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenReview.ProtoReflect.Descriptor instead.
func (*TokenReview) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_authentication_v1_generated_proto_rawDescGZIP(), []int{5}
}

func (x *TokenReview) GetMetadata() *v1.ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *TokenReview) GetSpec() *TokenReviewSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *TokenReview) GetStatus() *TokenReviewStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// TokenReviewSpec is a description of the token authentication request.
type TokenReviewSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token is the opaque bearer token.
	// +optional
	Token *string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	// Audiences is a list of the identifiers that the resource server presented
	// with the token identifies as. Audience-aware token authenticators will
	// verify that the token was intended for at least one of the audiences in
	// this list. If no audiences are provided, the audience will default to the
	// audience of the Kubernetes apiserver.
	// +optional
	Audiences []string `protobuf:"bytes,2,rep,name=audiences" json:"audiences,omitempty"`
}

func (x *TokenReviewSpec) Reset() {
	*x = TokenReviewSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenReviewSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenReviewSpec) ProtoMessage() {}

func (x *TokenReviewSpec) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenReviewSpec.ProtoReflect.Descriptor instead.
func (*TokenReviewSpec) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_authentication_v1_generated_proto_rawDescGZIP(), []int{6}
}

func (x *TokenReviewSpec) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *TokenReviewSpec) GetAudiences() []string {
	if x != nil {
		return x.Audiences
	}
	return nil
}

// TokenReviewStatus is the result of the token authentication request.
type TokenReviewStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Authenticated indicates that the token was associated with a known user.
	// +optional
	Authenticated *bool `protobuf:"varint,1,opt,name=authenticated" json:"authenticated,omitempty"`
	// User is the UserInfo associated with the provided token.
	// +optional
	User *UserInfo `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	// Audiences are audience identifiers chosen by the authenticator that are
	// compatible with both the TokenReview and token. An identifier is any
	// identifier in the intersection of the TokenReviewSpec audiences and the
	// token's audiences. A client of the TokenReview API that sets the
	// spec.audiences field should validate that a compatible audience identifier
	// is returned in the status.audiences field to ensure that the TokenReview
	// server is audience aware. If a TokenReview returns an empty
	// status.audience field where status.authenticated is "true", the token is
	// valid against the audience of the Kubernetes API server.
	// +optional
	Audiences []string `protobuf:"bytes,4,rep,name=audiences" json:"audiences,omitempty"`
	// Error indicates that the token couldn't be checked
	// +optional
	Error *string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (x *TokenReviewStatus) Reset() {
	*x = TokenReviewStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenReviewStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenReviewStatus) ProtoMessage() {}

func (x *TokenReviewStatus) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenReviewStatus.ProtoReflect.Descriptor instead.
func (*TokenReviewStatus) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_authentication_v1_generated_proto_rawDescGZIP(), []int{7}
}

func (x *TokenReviewStatus) GetAuthenticated() bool {
	if x != nil && x.Authenticated != nil {
		return *x.Authenticated
	}
	return false
}

func (x *TokenReviewStatus) GetUser() *UserInfo {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *TokenReviewStatus) GetAudiences() []string {
	if x != nil {
		return x.Audiences
	}
	return nil
}

func (x *TokenReviewStatus) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

// UserInfo holds the information about the user needed to implement the
// user.Info interface.
type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name that uniquely identifies this user among all active users.
	// +optional
	Username *string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	// A unique value that identifies this user across time. If this user is
	// deleted and another user by the same name is added, they will have
	// different UIDs.
	// +optional
	Uid *string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
	// The names of groups this user is a part of.
	// +optional
	Groups []string `protobuf:"bytes,3,rep,name=groups" json:"groups,omitempty"`
	// Any additional information provided by the authenticator.
	// +optional
	Extra map[string]*ExtraValue `protobuf:"bytes,4,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_io_api_authentication_v1_generated_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_k8s_io_api_authentication_v1_generated_proto_rawDescGZIP(), []int{8}
}

func (x *UserInfo) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *UserInfo) GetUid() string {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return ""
}

func (x *UserInfo) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *UserInfo) GetExtra() map[string]*ExtraValue {
	if x != nil {
		return x.Extra
	}
	return nil
}

var File_k8s_io_api_authentication_v1_generated_proto protoreflect.FileDescriptor

var file_k8s_io_api_authentication_v1_generated_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c,
	0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x34, 0x6b, 0x38,
	0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72,
	0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2f, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x14, 0x42,
	0x6f, 0x75, 0x6e, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x50, 0x49, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x50, 0x49,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x49, 0x44, 0x22, 0x22, 0x0a,
	0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0xea, 0x01, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x4c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x42, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x12, 0x48, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xba,
	0x01, 0x0a, 0x10, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12,
	0x5a, 0x0a, 0x0e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0e, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x22, 0x88, 0x01, 0x0a, 0x12,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x5c, 0x0a, 0x13, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x13, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xe7, 0x01, 0x0a, 0x0b, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x4c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x47, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x45, 0x0a, 0x0f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x64,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x11, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a,
	0x0d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0xfd, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x47, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a,
	0x62, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x3e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x1e, 0x5a, 0x1c, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31,
}

var (
	file_k8s_io_api_authentication_v1_generated_proto_rawDescOnce sync.Once
	file_k8s_io_api_authentication_v1_generated_proto_rawDescData = file_k8s_io_api_authentication_v1_generated_proto_rawDesc
)

func file_k8s_io_api_authentication_v1_generated_proto_rawDescGZIP() []byte {
	file_k8s_io_api_authentication_v1_generated_proto_rawDescOnce.Do(func() {
		file_k8s_io_api_authentication_v1_generated_proto_rawDescData = protoimpl.X.CompressGZIP(file_k8s_io_api_authentication_v1_generated_proto_rawDescData)
	})
	return file_k8s_io_api_authentication_v1_generated_proto_rawDescData
}

var file_k8s_io_api_authentication_v1_generated_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_k8s_io_api_authentication_v1_generated_proto_goTypes = []interface{}{
	(*BoundObjectReference)(nil), // 0: k8s.io.api.authentication.v1.BoundObjectReference
	(*ExtraValue)(nil),           // 1: k8s.io.api.authentication.v1.ExtraValue
	(*TokenRequest)(nil),         // 2: k8s.io.api.authentication.v1.TokenRequest
	(*TokenRequestSpec)(nil),     // 3: k8s.io.api.authentication.v1.TokenRequestSpec
	(*TokenRequestStatus)(nil),   // 4: k8s.io.api.authentication.v1.TokenRequestStatus
	(*TokenReview)(nil),          // 5: k8s.io.api.authentication.v1.TokenReview
	(*TokenReviewSpec)(nil),      // 6: k8s.io.api.authentication.v1.TokenReviewSpec
	(*TokenReviewStatus)(nil),    // 7: k8s.io.api.authentication.v1.TokenReviewStatus
	(*UserInfo)(nil),             // 8: k8s.io.api.authentication.v1.UserInfo
	nil,                          // 9: k8s.io.api.authentication.v1.UserInfo.ExtraEntry
	(*v1.ObjectMeta)(nil),        // 10: k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	(*v1.Time)(nil),              // 11: k8s.io.apimachinery.pkg.apis.meta.v1.Time
}
var file_k8s_io_api_authentication_v1_generated_proto_depIdxs = []int32{
	10, // 0: k8s.io.api.authentication.v1.TokenRequest.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	3,  // 1: k8s.io.api.authentication.v1.TokenRequest.spec:type_name -> k8s.io.api.authentication.v1.TokenRequestSpec
	4,  // 2: k8s.io.api.authentication.v1.TokenRequest.status:type_name -> k8s.io.api.authentication.v1.TokenRequestStatus
	0,  // 3: k8s.io.api.authentication.v1.TokenRequestSpec.boundObjectRef:type_name -> k8s.io.api.authentication.v1.BoundObjectReference
	11, // 4: k8s.io.api.authentication.v1.TokenRequestStatus.expirationTimestamp:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.Time
	10, // 5: k8s.io.api.authentication.v1.TokenReview.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	6,  // 6: k8s.io.api.authentication.v1.TokenReview.spec:type_name -> k8s.io.api.authentication.v1.TokenReviewSpec
	7,  // 7: k8s.io.api.authentication.v1.TokenReview.status:type_name -> k8s.io.api.authentication.v1.TokenReviewStatus
	8,  // 8: k8s.io.api.authentication.v1.TokenReviewStatus.user:type_name -> k8s.io.api.authentication.v1.UserInfo
	9,  // 9: k8s.io.api.authentication.v1.UserInfo.extra:type_name -> k8s.io.api.authentication.v1.UserInfo.ExtraEntry
	1,  // 10: k8s.io.api.authentication.v1.UserInfo.ExtraEntry.value:type_name -> k8s.io.api.authentication.v1.ExtraValue
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_k8s_io_api_authentication_v1_generated_proto_init() }
func file_k8s_io_api_authentication_v1_generated_proto_init() {
	if File_k8s_io_api_authentication_v1_generated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_k8s_io_api_authentication_v1_generated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoundObjectReference); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_k8s_io_api_authentication_v1_generated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtraValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_k8s_io_api_authentication_v1_generated_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_k8s_io_api_authentication_v1_generated_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRequestSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_k8s_io_api_authentication_v1_generated_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRequestStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_k8s_io_api_authentication_v1_generated_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenReview); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_k8s_io_api_authentication_v1_generated_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenReviewSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_k8s_io_api_authentication_v1_generated_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenReviewStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_k8s_io_api_authentication_v1_generated_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_k8s_io_api_authentication_v1_generated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_k8s_io_api_authentication_v1_generated_proto_goTypes,
		DependencyIndexes: file_k8s_io_api_authentication_v1_generated_proto_depIdxs,
		MessageInfos:      file_k8s_io_api_authentication_v1_generated_proto_msgTypes,
	}.Build()
	File_k8s_io_api_authentication_v1_generated_proto = out.File
	file_k8s_io_api_authentication_v1_generated_proto_rawDesc = nil
	file_k8s_io_api_authentication_v1_generated_proto_goTypes = nil
	file_k8s_io_api_authentication_v1_generated_proto_depIdxs = nil
}
