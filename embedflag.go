//go:build embed
// +build embed

package gofaiss

// #cgo CXXFLAGS: -std=c++11
// #cgo CPPFLAGS: -I${SRCDIR}/../../arjunsk/c-faiss/internal/c_api
// #cgo LDFLAGS: -lstdc++
// #cgo darwin LDFLAGS: -Wl,-undefined -Wl,dynamic_lookup
// #cgo !darwin LDFLAGS: -Wl,-unresolved-symbols=ignore-all -lrt
import "C"

import (
	_ "github.com/arjunsk/c-faiss"
)
