// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
