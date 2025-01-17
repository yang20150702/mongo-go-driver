// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package options

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// QueryType describes the type of query the result of Encrypt is used for.
type QueryType int

// These constants specify valid values for QueryType
const (
	QueryTypeEquality QueryType = 1
)

// EncryptOptions represents options to explicitly encrypt a value.
type EncryptOptions struct {
	KeyID            *primitive.Binary
	KeyAltName       *string
	Algorithm        string
	QueryType        *QueryType
	ContentionFactor *int64
}

// Encrypt creates a new EncryptOptions instance.
func Encrypt() *EncryptOptions {
	return &EncryptOptions{}
}

// SetKeyID specifies an _id of a data key. This should be a UUID (a primitive.Binary with subtype 4).
func (e *EncryptOptions) SetKeyID(keyID primitive.Binary) *EncryptOptions {
	e.KeyID = &keyID
	return e
}

// SetKeyAltName identifies a key vault document by 'keyAltName'.
func (e *EncryptOptions) SetKeyAltName(keyAltName string) *EncryptOptions {
	e.KeyAltName = &keyAltName
	return e
}

// SetAlgorithm specifies an algorithm to use for encryption. This should be one of the following:
// - AEAD_AES_256_CBC_HMAC_SHA_512-Deterministic
// - AEAD_AES_256_CBC_HMAC_SHA_512-Random
// - Indexed
// - Unindexed
// This is required.
func (e *EncryptOptions) SetAlgorithm(algorithm string) *EncryptOptions {
	e.Algorithm = algorithm
	return e
}

// SetQueryType specifies the intended query type. It is only valid to set if algorithm is "Indexed".
func (e *EncryptOptions) SetQueryType(queryType QueryType) *EncryptOptions {
	e.QueryType = &queryType
	return e
}

// SetContentionFactor specifies the contention factor. It is only valid to set if algorithm is "Indexed".
func (e *EncryptOptions) SetContentionFactor(contentionFactor int64) *EncryptOptions {
	e.ContentionFactor = &contentionFactor
	return e
}

// MergeEncryptOptions combines the argued EncryptOptions in a last-one wins fashion.
func MergeEncryptOptions(opts ...*EncryptOptions) *EncryptOptions {
	eo := Encrypt()
	for _, opt := range opts {
		if opt == nil {
			continue
		}

		if opt.KeyID != nil {
			eo.KeyID = opt.KeyID
		}
		if opt.KeyAltName != nil {
			eo.KeyAltName = opt.KeyAltName
		}
		if opt.Algorithm != "" {
			eo.Algorithm = opt.Algorithm
		}
		if opt.QueryType != nil {
			eo.QueryType = opt.QueryType
		}
		if opt.ContentionFactor != nil {
			eo.ContentionFactor = opt.ContentionFactor
		}
	}

	return eo
}
