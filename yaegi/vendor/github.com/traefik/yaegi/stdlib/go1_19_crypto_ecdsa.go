// Code generated by 'yaegi extract crypto/ecdsa'. DO NOT EDIT.

//go:build go1.19
// +build go1.19

package stdlib

import (
	"crypto/ecdsa"
	"reflect"
)

func init() {
	Symbols["crypto/ecdsa/ecdsa"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"GenerateKey": reflect.ValueOf(ecdsa.GenerateKey),
		"Sign":        reflect.ValueOf(ecdsa.Sign),
		"SignASN1":    reflect.ValueOf(ecdsa.SignASN1),
		"Verify":      reflect.ValueOf(ecdsa.Verify),
		"VerifyASN1":  reflect.ValueOf(ecdsa.VerifyASN1),

		// type definitions
		"PrivateKey": reflect.ValueOf((*ecdsa.PrivateKey)(nil)),
		"PublicKey":  reflect.ValueOf((*ecdsa.PublicKey)(nil)),
	}
}
