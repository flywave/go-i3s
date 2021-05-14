package i3s

// #include <stdlib.h>
// #include "i3s_api.h"
// #cgo CFLAGS: -I ./lib
// #cgo CXXFLAGS: -I ./lib
import "C"
import (
	"reflect"
	"unsafe"
)

type SpatialReference struct {
	m *C.struct__i3s_spatial_reference_t
}

func (n *SpatialReference) Free() {
	C.spatial_reference_free(n.m)
	n.m = nil
}

func (n *SpatialReference) GetWKID() int32 {
	return int32(C.spatial_reference_get_wkid(n.m))
}

func (n *SpatialReference) GetLatestWKID() int32 {
	return int32(C.spatial_reference_get_latest_wkid(n.m))
}

func (n *SpatialReference) GetVCSID() int32 {
	return int32(C.spatial_reference_get_vcs_id(n.m))
}

func (n *SpatialReference) GetLatestVCSID() int32 {
	return int32(C.spatial_reference_get_latest_vcs_id(n.m))
}

func (n *SpatialReference) GetWKT() string {
	return C.GoString(C.spatial_reference_get_wkt(n.m))
}

type CartesianTransformation struct {
	m    *C.struct__i3s_cartesian_transformation_t
	to   func(sr *SpatialReference, xyz []C.float) bool
	from func(sr *SpatialReference, xyz []C.float) bool
}

func NewCartesianTransformation(to func(sr *SpatialReference, xyz []C.float) bool, from func(sr *SpatialReference, xyz []C.float) bool) *CartesianTransformation {
	ret := &CartesianTransformation{m: nil, to: to, from: from}
	inptr := uintptr(unsafe.Pointer(ret))
	if p := C.cartesian_transformation_create((unsafe.Pointer)(&inptr)); p != nil {
		ret.m = p
		return ret
	}
	return nil
}

func (c *CartesianTransformation) toCartesian(sr *SpatialReference, xyz []C.float) bool {
	return c.toCartesian(sr, xyz)
}

func (c *CartesianTransformation) fromCartesian(sr *SpatialReference, xyz []C.float) bool {
	return c.fromCartesian(sr, xyz)
}

//export toCartesian
func toCartesian(ctx unsafe.Pointer, sr *C.struct__i3s_spatial_reference_t, xyz *C.float, count C.int) C.bool {
	var cpointsSlice []C.float
	cpointsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&cpointsSlice)))
	cpointsHeader.Cap = int(count)
	cpointsHeader.Len = int(count)
	cpointsHeader.Data = uintptr(unsafe.Pointer(xyz))

	return C.bool(*(**CartesianTransformation)(ctx).toCartesian(&SpatialReference{m: sr}, cpointsSlice))
}

//export fromCartesian
func fromCartesian(ctx unsafe.Pointer, sr *C.struct__i3s_spatial_reference_t, xyz *C.float, count C.int) C.bool {
	var cpointsSlice []C.float
	cpointsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&cpointsSlice)))
	cpointsHeader.Cap = int(count)
	cpointsHeader.Len = int(count)
	cpointsHeader.Data = uintptr(unsafe.Pointer(xyz))

	return C.bool(*(**CartesianTransformation)(ctx).fromCartesian(&SpatialReference{m: sr}, cpointsSlice))
}
