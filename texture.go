package i3s

// #include <stdlib.h>
// #include "i3s_api.h"
// #cgo CFLAGS: -I ./lib
// #cgo CXXFLAGS: -I ./lib
import "C"

type TextureBuffer struct {
	m *C.struct__i3s_texture_buffer_t
}

type TextureMeta struct {
	m *C.struct__i3s_texture_meta_t
}
