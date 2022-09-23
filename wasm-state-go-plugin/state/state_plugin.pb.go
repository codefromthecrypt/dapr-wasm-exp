//go:build tinygo.wasm

// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v3.21.6
// source: state.proto

//import "google/protobuf/any.proto";

package state

import (
	context "context"
	wasm "github.com/knqyf263/go-plugin/wasm"
)

const StorePluginAPIVersion = 1

//export store_api_version
func _store_api_version() uint64 {
	return StorePluginAPIVersion
}

var store Store

func RegisterStore(p Store) {
	store = p
}

//export store_init
func _store_init(ptr, size uint32) uint64 {
	b := wasm.PtrToByte(ptr, size)
	var req InitRequest
	if err := req.UnmarshalVT(b); err != nil {
		return 0
	}
	response, err := store.Init(context.Background(), req)
	if err != nil {
		return 0
	}

	b, err = response.MarshalVT()
	if err != nil {
		return 0
	}
	ptr, size = wasm.ByteToPtr(b)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}

//export store_features
func _store_features(ptr, size uint32) uint64 {
	b := wasm.PtrToByte(ptr, size)
	var req FeaturesRequest
	if err := req.UnmarshalVT(b); err != nil {
		return 0
	}
	response, err := store.Features(context.Background(), req)
	if err != nil {
		return 0
	}

	b, err = response.MarshalVT()
	if err != nil {
		return 0
	}
	ptr, size = wasm.ByteToPtr(b)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}

//export store_get
func _store_get(ptr, size uint32) uint64 {
	b := wasm.PtrToByte(ptr, size)
	var req GetRequest
	if err := req.UnmarshalVT(b); err != nil {
		return 0
	}
	response, err := store.Get(context.Background(), req)
	if err != nil {
		return 0
	}

	b, err = response.MarshalVT()
	if err != nil {
		return 0
	}
	ptr, size = wasm.ByteToPtr(b)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}

//export store_set
func _store_set(ptr, size uint32) uint64 {
	b := wasm.PtrToByte(ptr, size)
	var req SetRequest
	if err := req.UnmarshalVT(b); err != nil {
		return 0
	}
	response, err := store.Set(context.Background(), req)
	if err != nil {
		return 0
	}

	b, err = response.MarshalVT()
	if err != nil {
		return 0
	}
	ptr, size = wasm.ByteToPtr(b)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}

//export store_delete
func _store_delete(ptr, size uint32) uint64 {
	b := wasm.PtrToByte(ptr, size)
	var req DeleteRequest
	if err := req.UnmarshalVT(b); err != nil {
		return 0
	}
	response, err := store.Delete(context.Background(), req)
	if err != nil {
		return 0
	}

	b, err = response.MarshalVT()
	if err != nil {
		return 0
	}
	ptr, size = wasm.ByteToPtr(b)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}
