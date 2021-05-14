package i3s

// #include <stdlib.h>
// #include "i3s_api.h"
// #cgo CFLAGS: -I ./lib
// #cgo CXXFLAGS: -I ./lib
import "C"
import (
	"image"
	"unsafe"
)

type TextureBuffer struct {
	m *C.struct__i3s_texture_buffer_t
}

func NewTextureBuffer(size []int32, data []uint8) *TextureBuffer {
	if p := C.texture_buffer_create(C.int(size[0]), C.int(size[1]), C.int(size[2]), (*C.char)(unsafe.Pointer(&data[0]))); p != nil {
		return &TextureBuffer{m: p}
	}
	return nil
}

func NewTextureBufferWithImage(t interface{}) *TextureBuffer {
	switch img := t.(type) {
	case *image.NRGBA:
		return NewTextureBuffer([]int32{int32(img.Rect.Dx()), int32(img.Rect.Dy()), int32(4)}, img.Pix)
	case *image.RGBA:
		return NewTextureBuffer([]int32{int32(img.Rect.Dx()), int32(img.Rect.Dy()), int32(4)}, img.Pix)
	case *image.Gray:
		return NewTextureBuffer([]int32{int32(img.Rect.Dx()), int32(img.Rect.Dy()), int32(1)}, img.Pix)
	}
	return nil
}

func (n *TextureBuffer) Free() {
	C.texture_buffer_free(n.m)
	n.m = nil
}

func (n *TextureBuffer) GetMeta() *TextureMeta {
	return &TextureMeta{m: C.texture_buffer_get_meta(n.m)}
}

type TextureMeta struct {
	m *C.struct__i3s_texture_meta_t
}

func (n *TextureMeta) Free() {
	C.texture_meta_free(n.m)
	n.m = nil
}

func (n *TextureMeta) SetMip0Width(mip0Width int32) {
	C.texture_meta_set_mip0_width(n.m, C.int(mip0Width))
}

func (n *TextureMeta) SetMip0Height(mip0Height int32) {
	C.texture_meta_set_mip0_height(n.m, C.int(mip0Height))
}

func (n *TextureMeta) SetMipCount(mipCount int32) {
	C.texture_meta_set_mip_count(n.m, C.int(mipCount))
}

func (n *TextureMeta) SetUVSet(uvSet int32) {
	C.texture_meta_set_uv_set(n.m, C.int(uvSet))
}

func (n *TextureMeta) SetAlphaStatus(alphaStatus AlphaStatus) {
	C.texture_meta_set_alpha_status(n.m, C.int(alphaStatus))
}

func (n *TextureMeta) SetWrapMode(wrapMode WrapMode) {
	C.texture_meta_set_wrap_mode(n.m, C.int(wrapMode))
}

func (n *TextureMeta) SetFormat(format ImageFormat) {
	C.texture_meta_set_format(n.m, C.int(format))
}

func (n *TextureMeta) SetIsAtlas(isAtlas bool) {
	C.texture_meta_set_is_atlas(n.m, C.bool(isAtlas))
}
