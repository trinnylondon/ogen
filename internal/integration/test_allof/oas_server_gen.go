// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// NullableStrings implements nullableStrings operation.
	//
	// Nullable strings.
	//
	// POST /nullableStrings
	NullableStrings(ctx context.Context, req NilString) error
	// ObjectsWithConflictingArrayProperty implements objectsWithConflictingArrayProperty operation.
	//
	// Objects with conflicting array property.
	//
	// POST /objectsWithConflictingArrayProperty
	ObjectsWithConflictingArrayProperty(ctx context.Context, req *ObjectsWithConflictingArrayPropertyReq) error
	// ObjectsWithConflictingProperties implements objectsWithConflictingProperties operation.
	//
	// Objects with conflicting properties.
	//
	// POST /objectsWithConflictingProperties
	ObjectsWithConflictingProperties(ctx context.Context, req *ObjectsWithConflictingPropertiesReq) error
	// ReferencedAllOfNullable implements referencedAllOfNullable operation.
	//
	// Referenced allOf, but requestBody contains nullable refs.
	//
	// POST /referencedAllOfNullable
	ReferencedAllOfNullable(ctx context.Context, req ReferencedAllOfNullableReq) error
	// ReferencedAllof implements referencedAllof operation.
	//
	// Referenced allOf.
	//
	// POST /referencedAllof
	ReferencedAllof(ctx context.Context, req ReferencedAllofReq) error
	// ReferencedAllofOptional implements referencedAllofOptional operation.
	//
	// Referenced allOf, but requestBody is not required.
	//
	// POST /referencedAllofOptional
	ReferencedAllofOptional(ctx context.Context, req ReferencedAllofOptionalReq) error
	// SimpleInteger implements simpleInteger operation.
	//
	// Simple integers with validation.
	//
	// POST /simpleInteger
	SimpleInteger(ctx context.Context, req int) error
	// SimpleObjects implements simpleObjects operation.
	//
	// Simple objects.
	//
	// POST /simpleObjects
	SimpleObjects(ctx context.Context, req *SimpleObjectsReq) error
	// StringsNotype implements stringsNotype operation.
	//
	// POST /stringsNotype
	StringsNotype(ctx context.Context, req NilString) error
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
