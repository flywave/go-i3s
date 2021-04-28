package i3s

// #include <stdlib.h>
// #include "i3s_api.h"
// #cgo CFLAGS: -I ./lib
// #cgo CXXFLAGS: -I ./lib
import "C"

type MaterialData struct {
	m *C.struct__i3s_material_data_t
}
