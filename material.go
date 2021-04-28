package i3s

type MetallicRoughness struct {
	BaseColorFactor      [4]float32
	MetallicFactor       float32
	RoughnessFactor      float32
	BaseColorTex         []TextureBuffer
	MetallicRoughnessTex []TextureBuffer
}

type MaterialProperties struct {
	AMode          AlphaMode
	ACutOff        int32
	DoubleSided    bool
	EmissiveFactor [3]float32
	CullFace       FaceCullingMode
}

type MaterialData struct {
	m *C.struct__i3s_material_data_t
}
