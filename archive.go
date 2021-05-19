package i3s

// #include <stdlib.h>
// #include <string.h>
// #include "i3s_api.h"
// #cgo CFLAGS: -I ./lib
// #cgo CXXFLAGS: -I ./lib
// #cgo linux LDFLAGS:  -L ./lib -L /usr/lib/x86_64-linux-gnu -Wl,--start-group  -lstdc++ -lm -pthread -ldl -lci3s -lzlib -ldraco -li3s -ljpeg -llepcc -lpng -lEtcLib -Wl,--end-group
// #cgo windows LDFLAGS: -L ./lib　-lci3s -lzlib -ldraco -li3s -ljpeg -llepcc -lpng -lEtcLib
// #cgo darwin LDFLAGS: -L　./lib -lci3s -lzlib -ldraco -li3s -ljpeg -llepcc -lpng -lEtcLib
import "C"
import (
	"image"
	"reflect"
	"unsafe"
)

const (
	IDS_I3S_VISITING_NODE                               = int32(7000)
	IDS_I3S_OUT_OF_RANGE_ID                             = int32(7001)
	IDS_I3S_TYPE_MISMATCH                               = int32(7002)
	IDS_I3S_MISSING_JPG_OR_PNG                          = int32(7003)
	IDS_I3S_INVALID_TREE_TOPOLOGY                       = int32(7004)
	IDS_I3S_EMPTY_LEAF_NODE                             = int32(7005)
	IDS_I3S_UNCONNECTED_NODE                            = int32(7006)
	IDS_I3S_MISSING_ATTRIBUTE_STATS                     = int32(7007)
	IDS_I3S_GEOMETRY_COMPRESSION_RATIO                  = int32(7008)
	IDS_I3S_JSON_PARSING_ERROR                          = int32(7009)
	IDS_I3S_INVALID_LOD_METRIC                          = int32(7010)
	IDS_I3S_EXPECTS                                     = int32(7011)
	IDS_I3S_UNSUPPORTED_VERSION                         = int32(7012)
	IDS_I3S_PATH_COMPATIBILITY_WARNING                  = int32(7013)
	IDS_I3S_INVALID_EXTENT                              = int32(7014)
	IDS_I3S_DUPLICATE_ATTRIBUTE_KEY                     = int32(7015)
	IDS_I3S_DUPLICATE_ATTRIBUTE_NAME                    = int32(7016)
	IDS_I3S_MISSING_ATTRIBUTE_STATS_DECL                = int32(7017)
	IDS_I3S_STATS_DECL_UNKNOWN_ATTRIBUTE                = int32(7018)
	IDS_I3S_MISSING_ATTRIBUTE_STORAGE_DECL              = int32(7019)
	IDS_I3S_INVALID_COMPRESSED_GEOMETRY_INDEX           = int32(7020)
	IDS_I3S_INVALID_UNCOMPRESSED_GEOMETRY_INDEX         = int32(7021)
	IDS_I3S_INVALID_VERTEX_COUNT_IN_BUFFER              = int32(7022)
	IDS_I3S_MISSING_ATTRIBUTE_SET_DECL                  = int32(7023)
	IDS_I3S_UNEXPECTED_ATTRIBUTE_IN_COMPRESSED_GEOMETRY = int32(7024)
	IDS_I3S_MISSING_ATTRIBUTE_IN_COMPRESSED_GEOMETRY    = int32(7025)
	IDS_I3S_INVALID_FEATURE_COUNT_IN_BUFFER             = int32(7026)
	IDS_I3S_MISSING_TEXEL_COUNT_ESTIMATE                = int32(7027)
	IDS_I3S_IMAGE_ENCODING_ERROR                        = int32(7028)
	IDS_I3S_IMAGE_DECODING_ERROR                        = int32(7029)
	IDS_I3S_UNEXPECTED_ALPHA_CHANNEL                    = int32(7030)
	IDS_I3S_MISSING_ALPHA_CHANNEL                       = int32(7031)
	IDS_I3S_MISSING_RESOURCE                            = int32(7032)
	IDS_I3S_INVALID_BYTE_ALIGNMENT                      = int32(7033)
	IDS_I3S_INVALID_BINARY_BUFFER_SIZE                  = int32(7034)
	IDS_I3S_CONNECTION_ERROR                            = int32(7035)
	IDS_I3S_MBS_OBB_CENTER_MISMATCH                     = int32(7036)
	IDS_I3S_INVALID_RADIUS_MBS                          = int32(7037)
	IDS_I3S_INVALID_STRING_LENGTH_IN_ATTRIBUTE_BUFFER   = int32(7038)
	IDS_I3S_INVALID_FEATURE_FACE_RANGE                  = int32(7039)
	IDS_I3S_BAD_UV                                      = int32(7040)
	IDS_I3S_MISSING_PROPERTY_COMPATIBILITY_WARNING      = int32(7041)
	IDS_I3S_COORDINATE_SYSTEM_COMPATIBILITY_WARNING     = int32(7042)
	IDS_I3S_IMAGE_TOO_LARGE                             = int32(7043)
	IDS_I3S_LEPCC_DECODE_ERROR                          = int32(7044)
	IDS_I3S_DRACO_DECODE_ERROR                          = int32(7045)
	IDS_I3S_ATTRIBUTE_VALUE_COUNT_MISMATCH              = int32(7046)
	IDS_I3S_WRONG_CHILD_LINK                            = int32(7047)
	IDS_I3S_TEX_ATLAS_VERTEX_REGION_MISMATCH            = int32(7048)
	IDS_I3S_LEGACY_TEX_ATLAS_FLAG_MISMATCH              = int32(7049)
	IDS_I3S_MISSING_VERSION                             = int32(7050)
	IDS_I3S_OK                                          = int32(8000)
	IDS_I3S_IO_OPEN_FAILED                              = int32(8004)
	IDS_I3S_IO_WRITE_FAILED                             = int32(8007)
	IDS_I3S_INTERNAL_ERROR                              = int32(8009)
	IDS_I3S_COMPRESSION_ERROR                           = int32(8032)
	IDS_I3S_DEGENERATED_MESH                            = int32(8090)
)

func GetMessage(stringID int32) string {
	cmsg := C.get_message_string(C.int(stringID))
	defer C.free(unsafe.Pointer(cmsg))
	return C.GoString(cmsg)
}

type NodeData struct {
	m *C.struct__i3s_node_data_t
}

func NewNodeData() *NodeData {
	if p := C.node_data_create(); p != nil {
		return &NodeData{m: p}
	}
	return nil
}

func (n *NodeData) Free() {
	C.node_data_free(n.m)
	n.m = nil
}

func (n *NodeData) SetLodThreshold(lodThreshold float64) {
	C.node_data_set_lod_threshold(n.m, C.double(lodThreshold))
}

func (n *NodeData) SetNodeDepth(nodeDepth int32) {
	C.node_data_set_node_depth(n.m, C.int(nodeDepth))
}

func (n *NodeData) ApppendChildren(nodeId NodeID) {
	C.node_data_append_children(n.m, C.size_t(nodeId))
}

func (n *NodeData) SetChildrens(nodeIds []NodeID) {
	for _, s := range nodeIds {
		n.ApppendChildren(s)
	}
}

func (n *NodeData) GetMeshData() *Mesh {
	return &Mesh{m: C.node_data_get_mesh_data(n.m)}
}

type RawMesh struct {
	m *C.struct__i3s_raw_mesh_t
}

func NewRawMesh() *RawMesh {
	if p := C.raw_mesh_create(); p != nil {
		return &RawMesh{m: p}
	}
	return nil
}

func (n *RawMesh) Free() {
	C.raw_mesh_free(n.m)
	n.m = nil
}

func (n *RawMesh) SetVertex(vertexs []float64, uvs []float64, indices []uint32) {
	C.raw_mesh_set_vertex(n.m, (*C.double)(unsafe.Pointer(&vertexs[0])), (*C.float)(unsafe.Pointer(&uvs[0])), C.size_t(len(vertexs)), (*C.uint)(unsafe.Pointer(&indices[0])), C.size_t(len(indices)))
}

func (n *RawMesh) SetTexture(size []int32, data []uint8) {
	C.raw_mesh_set_texture(n.m, C.int(size[0]), C.int(size[1]), C.int(size[2]), (*C.char)(unsafe.Pointer(&data[0])))
}

func (n *RawMesh) SetTextureImage(t interface{}) {
	switch img := t.(type) {
	case *image.NRGBA:
		n.SetTexture([]int32{int32(img.Rect.Dx()), int32(img.Rect.Dy()), int32(4)}, img.Pix)
	case *image.RGBA:
		n.SetTexture([]int32{int32(img.Rect.Dx()), int32(img.Rect.Dy()), int32(4)}, img.Pix)
	case *image.Gray:
		n.SetTexture([]int32{int32(img.Rect.Dx()), int32(img.Rect.Dy()), int32(1)}, img.Pix)
	}
}

type RawPoints struct {
	m *C.struct__i3s_raw_points_t
}

func NewRawPoints() *RawPoints {
	if p := C.raw_points_create(); p != nil {
		return &RawPoints{m: p}
	}
	return nil
}

func (n *RawPoints) Free() {
	C.raw_points_free(n.m)
	n.m = nil
}

func (n *RawPoints) SetPoints(vertexs []float64, fids []uint32) {
	if len(fids) > 0 {
		C.raw_points_set_points(n.m, (*C.double)(unsafe.Pointer(&vertexs[0])), (*C.uint)(unsafe.Pointer(&fids[0])), C.size_t(len(vertexs)))
	} else {
		C.raw_points_set_points(n.m, (*C.double)(unsafe.Pointer(&vertexs[0])), nil, C.size_t(len(vertexs)))
	}
}

type WriterContext struct {
	m *C.struct__i3s_writer_context_t
}

func NewWriterContext(prop *CtxProperties) *WriterContext {
	if p := C.writer_context_create(prop.m, nil); p != nil {
		return &WriterContext{m: p}
	}
	return nil
}

func NewWriterContextWithCartesianTransformation(prop *CtxProperties, tran *CartesianTransformation) *WriterContext {
	if p := C.writer_context_create(prop.m, tran.m); p != nil {
		return &WriterContext{m: p}
	}
	return nil
}

func (n *WriterContext) Free() {
	C.writer_context_free(n.m)
	n.m = nil
}

type LayerMeta struct {
	m *C.struct__i3s_layer_meta_t
}

func NewLayerMeta() *LayerMeta {
	if p := C.layer_meta_create(); p != nil {
		return &LayerMeta{m: p}
	}
	return nil
}

func (n *LayerMeta) Free() {
	C.layer_meta_free(n.m)
	n.m = nil
}

func (n *LayerMeta) SetType(tp LayerType) {
	C.layer_meta_set_type(n.m, C.int(tp))
}

func (n *LayerMeta) SetName(name string) {
	cname := C.CString(name)
	C.layer_meta_set_name(n.m, cname)
	C.free(unsafe.Pointer(cname))
}

func (n *LayerMeta) SetDesc(desc string) {
	cdesc := C.CString(desc)
	C.layer_meta_set_desc(n.m, cdesc)
	C.free(unsafe.Pointer(cdesc))
}

func (n *LayerMeta) SetCopyright(copyright string) {
	ccopyright := C.CString(copyright)
	C.layer_meta_set_copyright(n.m, ccopyright)
	C.free(unsafe.Pointer(ccopyright))
}

func (n *LayerMeta) SetSpatialReferenceWkid(wkid int32) {
	C.layer_meta_set_spatial_reference_wkid(n.m, C.int(wkid))
}

func (n *LayerMeta) SetSpatialReferenceLatestWkid(wkid int32) {
	C.layer_meta_set_spatial_reference_latest_wkid(n.m, C.int(wkid))
}

func (n *LayerMeta) SetSpatialReferenceVcsid(vcsid int32) {
	C.layer_meta_set_spatial_reference_vcs_id(n.m, C.int(vcsid))
}

func (n *LayerMeta) SetSpatialReferenceLatestVcsid(vcsid int32) {
	C.layer_meta_set_spatial_reference_latest_vcs_id(n.m, C.int(vcsid))
}

func (n *LayerMeta) SetSpatialReferenceWkt(wkt string) {
	cwkt := C.CString(wkt)
	C.layer_meta_set_spatial_reference_wkt(n.m, cwkt)
	defer C.free(unsafe.Pointer(cwkt))
}

func (n *LayerMeta) SetUID(uid string) {
	cuid := C.CString(uid)
	C.layer_meta_set_uid(n.m, cuid)
	defer C.free(unsafe.Pointer(cuid))
}

func (n *LayerMeta) SetDrawingInfo(drawingInfo string) {
	cdrawingInfo := C.CString(drawingInfo)
	C.layer_meta_set_drawing_info(n.m, cdrawingInfo)
	defer C.free(unsafe.Pointer(cdrawingInfo))
}

func (n *LayerMeta) SetElevationInfo(elevationInfo string) {
	celevationInfo := C.CString(elevationInfo)
	C.layer_meta_set_elevation_info(n.m, celevationInfo)
	defer C.free(unsafe.Pointer(celevationInfo))
}

func (n *LayerMeta) SetPopupInfo(popupInfo string) {
	cpopupInfo := C.CString(popupInfo)
	C.layer_meta_set_popup_info(n.m, cpopupInfo)
	defer C.free(unsafe.Pointer(cpopupInfo))
}

func (n *LayerMeta) SetTimestamp(timestamp uint64) {
	C.layer_meta_set_timestamp(n.m, C.uint64_t(timestamp))
}

func (n *LayerMeta) SetNormalReferenceFrame(normal_reference_frame NormalReferenceFrame) {
	C.layer_meta_set_normal_reference_frame(n.m, C.int(normal_reference_frame))
}

func (n *LayerMeta) SetHeightModel(hmode HeightModel) {
	C.layer_meta_set_height_model_info_height_model(n.m, C.int(hmode))
}

func (n *LayerMeta) SetHeightModelInfoVertCrs(crs string) {
	ccrs := C.CString(crs)
	C.layer_meta_set_height_model_info_vert_crs(n.m, ccrs)
	defer C.free(unsafe.Pointer(ccrs))
}

func (n *LayerMeta) SetHeightModelInfoHeightUnit(hu HeightUnit) {
	C.layer_meta_set_height_model_info_height_unit(n.m, C.int(hu))
}

type AttributeMeta struct {
	m *C.struct__i3s_attribute_meta_t
}

func (n *AttributeMeta) Free() {
	C.attribute_meta_free(n.m)
	n.m = nil
}

func (n *AttributeMeta) SetKey(key string) {
	ckey := C.CString(key)
	C.attribute_meta_set_key(n.m, ckey)
	defer C.free(unsafe.Pointer(ckey))
}

func (n *AttributeMeta) SetName(name string) {
	cname := C.CString(name)
	C.attribute_meta_set_name(n.m, cname)
	defer C.free(unsafe.Pointer(cname))
}

func (n *AttributeMeta) SetAlias(alias string) {
	calias := C.CString(alias)
	C.attribute_meta_set_alias(n.m, calias)
	defer C.free(unsafe.Pointer(calias))
}

type LayerWriter struct {
	m *C.struct__i3s_layer_writer_t
}

func (n *LayerWriter) Free() {
	C.layer_writer_free(n.m)
	n.m = nil
}

func NewLayerWriter(ctx *WriterContext, path string) *LayerWriter {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	if p := C.layer_writer_create(ctx.m, cpath); p != nil {
		return &LayerWriter{m: p}
	}
	return nil
}

func (n *LayerWriter) SetLayerMeta(meta *LayerMeta) {
	C.layer_writer_set_layer_meta(n.m, meta.m)
}

func (n *LayerWriter) SetAttributeMeta(idx int32, meta *AttributeMeta, schemaId int32) {
	C.layer_writer_set_attribute_meta(n.m, C.int(idx), meta.m, C.int(schemaId))
}

func (n *LayerWriter) CreateOutputNode(node *NodeData, nodeId NodeID) {
	C.layer_writer_create_output_node(n.m, node.m, C.size_t(nodeId))
}

func (n *LayerWriter) CreateNode(node *NodeData, nodeId NodeID) {
	C.layer_writer_create_node(n.m, node.m, C.size_t(nodeId))
}

func (n *LayerWriter) ProcessChildren(node *NodeData, nodeId NodeID) {
	C.layer_writer_process_children(n.m, node.m, C.size_t(nodeId))
}

func (n *LayerWriter) CreateMesh(mesh *RawMesh, node *NodeData) {
	C.layer_writer_create_mesh_from_raw(n.m, mesh.m, node.m)
}

func (n *LayerWriter) CreatePoints(pts *RawPoints, node *NodeData) {
	C.layer_writer_create_points_from_raw(n.m, pts.m, node.m)
}

func (n *LayerWriter) Save() {
	C.layer_writer_save(n.m)
}

func CreateDefaultWriter(lt LayerType, name string, slpkPath string) *LayerWriter {
	ctx := NewCtxProperties()
	writer_context := NewWriterContext(ctx)
	meta := NewLayerMeta()
	meta.SetType(lt)
	meta.SetName(name)
	meta.SetDesc("Generated with flywave")
	meta.SetSpatialReferenceWkid(3857)
	meta.SetUID(name)
	meta.SetNormalReferenceFrame(NRF_NOT_SET)

	writer := NewLayerWriter(writer_context, slpkPath)

	writer.SetLayerMeta(meta)
	return writer
}

type CreateFlags C.uint

const (
	CF_Unfinalized_only                      = CreateFlags(1)
	CF_Disable_auto_finalize                 = CreateFlags(2)
	CF_Overwrite_if_exists_and_auto_finalize = CreateFlags(0)
)

type MimeType C.uint

const (
	MT_Not_set     = MimeType(1)
	MT_Json        = MimeType(2)
	MT_Binary      = MimeType(4)
	MT_Jpeg        = MimeType(8)
	MT_Png         = MimeType(16)
	MT_Dds_proper  = MimeType(32)
	MT_Dds_legacy  = MimeType(64)
	MT_Ktx         = MimeType(128)
	MT__next_      = MimeType(256)
	MT_Default     = MT_Json | MT_Binary | MT_Not_set
	MT_Dds_any     = MT_Dds_proper | MT_Dds_legacy
	MT_Any_texture = MT_Jpeg | MT_Png | MT_Dds_any | MT_Ktx | MT_Binary
)

type MimeEncoding C.uint

const (
	ME_Not_set             = MimeEncoding(1)
	ME_Gzip                = MimeEncoding(2)
	ME_Lepcc_xyz           = MimeEncoding(4)
	ME_Lepcc_rgb           = MimeEncoding(8)
	ME_Lepcc_int           = MimeEncoding(16)
	ME__next_              = MimeEncoding(32)
	ME_Default             = ME_Not_set | ME_Gzip
	ME_Any_pcsl_attributes = ME_Gzip | ME_Lepcc_rgb | ME_Lepcc_int | ME_Not_set
	ME_Any_geometry        = ME_Not_set | ME_Gzip | ME_Lepcc_xyz
)

type SlpkWriter interface {
	CreateArchive(path string, flags CreateFlags) bool
	AppendFile(path string, buf []byte, tp MimeType, pack MimeEncoding) bool
	GetFile(path string) []byte
	Finalize() bool
}

type SlpkWriterAdapt struct {
	m *C.struct__i3s_slpk_writer_t
	w SlpkWriter
}

func (n *SlpkWriterAdapt) Free() {
	C.slpk_writer_free(n.m)
	n.m = nil
}

//export createArchive
func createArchive(ctx unsafe.Pointer, path *C.char, flags C.uint) C.bool {
	return C.bool((*(**SlpkWriterAdapt)(ctx)).w.CreateArchive(C.GoString(path), CreateFlags(flags)))
}

//export appendFile
func appendFile(ctx unsafe.Pointer, path *C.char, buf *C.uchar, count C.int, mimeType C.uint, pack C.uint) C.bool {
	var cpointsSlice []byte
	cpointsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&cpointsSlice)))
	cpointsHeader.Cap = int(count)
	cpointsHeader.Len = int(count)
	cpointsHeader.Data = uintptr(unsafe.Pointer(buf))

	return C.bool((*(**SlpkWriterAdapt)(ctx)).w.AppendFile(C.GoString(path), cpointsSlice, MimeType(mimeType), MimeEncoding(pack)))
}

//export getFile
func getFile(ctx unsafe.Pointer, path *C.char, size *C.size_t) *C.uchar {
	buf := (*(**SlpkWriterAdapt)(ctx)).w.GetFile(C.GoString(path))
	*size = C.size_t(len(buf))
	cbuf := C.malloc(*size)

	C.memcpy(unsafe.Pointer(&buf[0]), unsafe.Pointer(cbuf), C.size_t(*size))
	return (*C.uchar)(cbuf)
}

//export finalize
func finalize(ctx unsafe.Pointer) C.bool {
	return C.bool((*(**SlpkWriterAdapt)(ctx)).w.Finalize())
}

func NewSlpkWriterAdapt(ctx SlpkWriter) *SlpkWriterAdapt {
	ret := &SlpkWriterAdapt{m: nil, w: ctx}
	inptr := uintptr(unsafe.Pointer(ret))

	if p := C.slpk_writer_create((unsafe.Pointer)(&inptr)); p != nil {
		ret.m = p
		return ret
	}
	return nil
}

func NewLayerWriterWithSlpkWriter(ctx *WriterContext, adapt *SlpkWriterAdapt) *LayerWriter {
	if p := C.layer_writer_create_with_slpk_writer(ctx.m, adapt.m); p != nil {
		return &LayerWriter{m: p}
	}
	return nil
}
