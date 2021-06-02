package i3s

// #include <stdlib.h>
// #include "i3s_api.h"
// #cgo CFLAGS: -I ./lib
// #cgo CXXFLAGS: -I ./lib
import "C"
import (
	"runtime"
	"unsafe"
)

type Mesh struct {
	m *C.struct__i3s_mesh_data_t
}

func NewMesh(m *C.struct__i3s_mesh_data_t) *Mesh {
	mh := &Mesh{m: m}
	runtime.SetFinalizer(mh, (*Mesh).free)
	return mh
}

func (n *Mesh) free() {
	C.mesh_data_free(n.m)
	n.m = nil
}

func (n *Mesh) GetMaterial() *MaterialData {
	return &NewMaterialData(C.mesh_data_get_material(n.m))
}

func (n *Mesh) SetColor(c []byte) {
	C.mesh_data_set_color(n.m, (*C.char)(unsafe.Pointer(&c[0])), C.ulong(len(c)))
}
