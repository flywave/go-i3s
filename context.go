package i3s

// #include <stdlib.h>
// #include "i3s_api.h"
// #cgo CFLAGS: -I ./lib
// #cgo CXXFLAGS: -I ./lib
import "C"

type CtxProperties struct {
	m *C.struct__i3s_ctx_properties_t
}

func NewCtxProperties() *CtxProperties {
	if p := C.ctx_properties_create(); p != nil {
		return &CtxProperties{m: p}
	}
	return nil
}

func (n *CtxProperties) Free() {
	C.ctx_properties_free(n.m)
	n.m = nil
}

func (n *CtxProperties) SetMaxTextureSize(size uint16) {
	C.ctx_properties_set_max_texture_size(n.m, C.ushort(size))
}

func (n *CtxProperties) IsDropNormals(flag bool) {
	C.ctx_properties_is_drop_normals(n.m, C.bool(flag))
}

func (n *CtxProperties) SetGeomEncodingSupport(tp GeometryCompressionFlags) {
	C.ctx_properties_set_geom_encoding_support(n.m, C.uint(tp))
}

func (n *CtxProperties) SetGeomDecodingSupport(tp GeometryCompressionFlags) {
	C.ctx_properties_set_geom_decoding_support(n.m, C.uint(tp))
}

func (n *CtxProperties) SetGPUTexEncodingSupport(tp GeometryCompressionFlags) {
	C.ctx_properties_set_gpu_tex_encoding_support(n.m, C.uint(tp))
}

func (n *CtxProperties) SetGPUTexRenderingSupport(tp GeometryCompressionFlags) {
	C.ctx_properties_set_gpu_tex_rendering_support(n.m, C.uint(tp))
}
