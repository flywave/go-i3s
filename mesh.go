package i3s

// #include <stdlib.h>
// #include "i3s_api.h"
// #cgo CFLAGS: -I ./lib
// #cgo CXXFLAGS: -I ./lib
import "C"

type Mesh struct {
	m *C.struct__i3s_mesh_data_t
}

func (n *Mesh) Free() {
	C.mesh_data_free(n.m)
	n.m = nil
}

func (n *Mesh) GetMaterial() *MaterialData {
	return &MaterialData{m: C.mesh_data_get_material(n.m)}
}
