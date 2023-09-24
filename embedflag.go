package gofaiss

// #cgo CPPFLAGS: -I../../../vendor/github.com/arjunsk/c-faiss/internal/c_api
// #cgo CPPFLAGS: -I../../../vendor/github.com/arjunsk/c-faiss/internal/faiss
// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -lfaiss -lfaiss_c
import "C"

import (
	_ "github.com/arjunsk/c-faiss"
)
