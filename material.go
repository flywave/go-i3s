package i3s

// #include <stdlib.h>
// #include "i3s_api.h"
// #cgo CFLAGS: -I ./lib
// #cgo CXXFLAGS: -I ./lib
import "C"
import "unsafe"

type MaterialData struct {
	m *C.struct__i3s_material_data_t
}

func (n *MaterialData) Free() {
	C.material_data_free(n.m)
	n.m = nil
}

func (n *MaterialData) SetMaterialPropertiesAlphaMode(alpha_mode AlphaMode) {
	C.material_data_set_material_properties_alpha_mode(n.m, C.int(alpha_mode))
}

func (n *MaterialData) SetMaterialPropertiesAlphaCutOff(alpha_cut_off int) {
	C.material_data_set_material_properties_alpha_cut_off(n.m, C.int(alpha_cut_off))
}

func (n *MaterialData) SetMaterialPropertiesDoubleSided(double_sided bool) {
	C.material_data_set_material_properties_double_sided(n.m, C.bool(double_sided))
}

func (n *MaterialData) SetMaterialPropertiesEmissiveFactor(emissive_factor []float32) {
	C.material_data_set_material_properties_emissive_factor(n.m, (*C.float)(unsafe.Pointer(&emissive_factor[0])))
}

func (n *MaterialData) SetMaterialPropertiesCullFace(cull_face FaceCullingMode) {
	C.material_data_set_material_properties_cull_face(n.m, C.int(cull_face))
}

func (n *MaterialData) SetMetallicRoughnessBaseColorFactor(base_color_factor []float32) {
	C.material_data_set_metallic_roughness_base_color_factor(n.m, (*C.float)(unsafe.Pointer(&base_color_factor[0])))
}

func (n *MaterialData) SetMetallicRoughnessMetallicFactor(metallic_factor float32) {
	C.material_data_set_metallic_roughness_metallic_factor(n.m, C.float(metallic_factor))
}

func (n *MaterialData) SetMetallicRoughnessRoughnessFactor(roughness_factor float32) {
	C.material_data_set_metallic_roughness_roughness_factor(n.m, C.float(roughness_factor))
}

func (n *MaterialData) AppendMetallicRoughnessBaseColorTex(tex *TextureBuffer) {
	C.material_data_append_metallic_roughness_base_color_tex(n.m, tex.m)
}

func (n *MaterialData) SetMetallicRoughnessBaseColorTex(texs []*TextureBuffer) {
	for _, s := range texs {
		n.AppendMetallicRoughnessBaseColorTex(s)
	}
}

func (n *MaterialData) AppendMetallicRoughnessMetallicRoughnessTex(tex *TextureBuffer) {
	C.material_data_append_metallic_roughness_metallic_roughness_tex(n.m, tex.m)
}

func (n *MaterialData) SetMetallicRoughnessMetallicRoughnessTex(texs []*TextureBuffer) {
	for _, s := range texs {
		n.AppendMetallicRoughnessMetallicRoughnessTex(s)
	}
}

func (n *MaterialData) AppendNormalMapTex(tex *TextureBuffer) {
	C.material_data_append_normal_map_tex(n.m, tex.m)
}

func (n *MaterialData) SetNormalMapTex(texs []*TextureBuffer) {
	for _, s := range texs {
		n.AppendNormalMapTex(s)
	}
}
