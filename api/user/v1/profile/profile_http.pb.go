// Code generated by protoc-gen-go-http. DO NOT EDIT.

package profile

import (
	context "context"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	mux "github.com/gorilla/mux"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(http.Request)
var _ = new(context.Context)
var _ = binding.MapProto
var _ = mux.NewRouter

const _ = http1.SupportPackageIsVersion1

type ProfileHandler interface {
	CreateProfile(context.Context, *CreateProfileRequest) (*CreateProfileReply, error)

	DeleteProfile(context.Context, *DeleteProfileRequest) (*DeleteProfileReply, error)

	GetProfile(context.Context, *GetProfileRequest) (*GetProfileReply, error)

	ListProfile(context.Context, *ListProfileRequest) (*ListProfileReply, error)

	UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileReply, error)
}

func NewProfileHandler(srv ProfileHandler, opts ...http1.HandleOption) http.Handler {
	h := http1.DefaultHandleOptions()
	for _, o := range opts {
		o(&h)
	}
	r := mux.NewRouter()

	r.HandleFunc("/v1/user", func(w http.ResponseWriter, r *http.Request) {
		var in CreateProfileRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateProfile(ctx, req.(*CreateProfileRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*CreateProfileReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	r.HandleFunc("/v1/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		var in UpdateProfileRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.MapProto(&in, mux.Vars(r)); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateProfile(ctx, req.(*UpdateProfileRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*UpdateProfileReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("PUT")

	r.HandleFunc("/v1/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		var in DeleteProfileRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.MapProto(&in, mux.Vars(r)); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteProfile(ctx, req.(*DeleteProfileRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*DeleteProfileReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("DELETE")

	r.HandleFunc("/v1/user/{id}/{name}", func(w http.ResponseWriter, r *http.Request) {
		var in GetProfileRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.MapProto(&in, mux.Vars(r)); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProfile(ctx, req.(*GetProfileRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*GetProfileReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/api.user.v1.profile.Profile/ListProfile", func(w http.ResponseWriter, r *http.Request) {
		var in ListProfileRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListProfile(ctx, req.(*ListProfileRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*ListProfileReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	return r
}