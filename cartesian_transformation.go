package i3s

// #include <stdlib.h>
// #include "i3s_api.h"
// #cgo CFLAGS: -I ./lib
// #cgo CXXFLAGS: -I ./lib
import "C"

type CartesianTransformation struct {
	m *C.struct__i3s_cartesian_transformation_t
}