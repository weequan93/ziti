// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package raft

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RaftTranferLeadershipHandlerFunc turns a function with the right signature into a raft tranfer leadership handler
type RaftTranferLeadershipHandlerFunc func(RaftTranferLeadershipParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RaftTranferLeadershipHandlerFunc) Handle(params RaftTranferLeadershipParams) middleware.Responder {
	return fn(params)
}

// RaftTranferLeadershipHandler interface for that can handle valid raft tranfer leadership params
type RaftTranferLeadershipHandler interface {
	Handle(RaftTranferLeadershipParams) middleware.Responder
}

// NewRaftTranferLeadership creates a new http.Handler for the raft tranfer leadership operation
func NewRaftTranferLeadership(ctx *middleware.Context, handler RaftTranferLeadershipHandler) *RaftTranferLeadership {
	return &RaftTranferLeadership{Context: ctx, Handler: handler}
}

/* RaftTranferLeadership swagger:route POST /raft/transfer-leadership Raft raftTranferLeadership

Attempts to transfer leadership to a different member of the cluster

Attempts to transfer leadership to a different member of the cluster. Requires admin access.

*/
type RaftTranferLeadership struct {
	Context *middleware.Context
	Handler RaftTranferLeadershipHandler
}

func (o *RaftTranferLeadership) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRaftTranferLeadershipParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
