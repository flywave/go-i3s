package i3s

// #include <stdlib.h>
// #include "i3s_api.h"
// #cgo CFLAGS: -I ./lib
// #cgo CXXFLAGS: -I ./lib
import "C"

type CtxProperties struct {
	m *C.struct__i3s_ctx_properties_t
}

type Context struct {
	m *C.struct__i3s_context_t
}
