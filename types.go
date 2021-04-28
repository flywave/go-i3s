package i3s

type NodeId C.size_t
type LayerId C.size_t

type AttribIndex C.int
type AttribSchemaId C.int
type TextureId C.int
type GeometryId C.int

type Rgba8 C.uchar[4]
type UVRegion C.uchar[4]

type GeometryFormat C.int

const (
	GF_NOT_SET      = GeometryFormat(0)
	GF_LEGACY      = GeometryFormat(1)
	GF_POINT_LEGACY           = GeometryFormat(2)
	GF_DRACO         = GeometryFormat(3)
	GF_LEPCC = GeometryFormat(4)
)

type ImageFormat C.uint

const (
	IF_NOT_SET      = ImageFormat(0)
	IF_JPG      = ImageFormat(1)
	IF_PNG      = ImageFormat(2)
	IF_DDS      = ImageFormat(4)
	IF_KTX      = ImageFormat(8)
	IF_RAW_RGBA8      = ImageFormat(16)
	IF_RAW_RGB8      = ImageFormat(32)
	IF_DEFAULT = IF_JPG | IF_PNG
	IF_UNCOMPRESSED = IF_RAW_RGBA8 | IF_RAW_RGB8
	IF_DESKTOP = IF_DEFAULT | IF_DDS
	IF_ALL_COMPRESSED = IF_DESKTOP | IF_KTX
	IF_NOT_GPU_COMPRESSED = IF_UNCOMPRESSED | IF_JPG | IF_PNG
)

type AlphaMode C.int
const (
	AM_OPAQUE      = AlphaMode(0)
	AM_MASK      = AlphaMode(1)
	AM_BLEND      = AlphaMode(2)
	AM_NOT_SET      = AlphaMode(3)
)

type FaceCullingMode C.int
const (
	AM_NONE      = FaceCullingMode(0)
	AM_FRONT      = FaceCullingMode(1)
	AM_BACK      = FaceCullingMode(2)
)

type DataType C.int
const (
	DT_INT8      = DataType(0)
	DT_UINT8      = DataType(1)
	DT_INT16      = DataType(2)
	DT_UINT16      = DataType(3)
	DT_INT32      = DataType(4)
	DT_UINT32      = DataType(5)
	DT_OID32      = DataType(6)
	DT_INT64      = DataType(7)
	DT_UINT64      = DataType(8)
	DT_OID64      = DataType(9)
	DT_FLOAT32      = DataType(10)
	DT_FLOAT64      = DataType(11)
	DT_STRING      = DataType(12)
	DT_DATE      = DataType(13)
	DT_GLOBALID      = DataType(14)
	DT_GUID      = DataType(15)
	DT_NOT_SET      = DataType(16)
)

type EsriFieldType C.int
const (
	FT_DATA      = EsriFieldType(0)
	FT_SINGLE      = EsriFieldType(1)
	FT_DOUBLE      = EsriFieldType(2)
	FT_GUID      = EsriFieldType(3)
	FT_GLOBALID      = EsriFieldType(4)
	FT_INTEGER      = EsriFieldType(5)
	FT_OID      = EsriFieldType(6)
	FT_SMALLINTEGER      = EsriFieldType(7)
	FT_STRING      = EsriFieldType(8)
	FT_XML      = EsriFieldType(9)
	FT_NOT_SET      = EsriFieldType(10)
)

type LegacyTopology C.int
const (
	LT_PERATTRIBUTEARRAY = LegacyTopology(0)
)

type VertexAttribOrdering C.int
const (
	VA_POSITION = VertexAttribOrdering(0)
	VA_NORMAL = VertexAttribOrdering(1)
	VA_UV0 = VertexAttribOrdering(2)
	VA_COLOR = VertexAttribOrdering(3)
	VA_REGION = VertexAttribOrdering(4)
)

type FeatureAttribOrdering C.int
const (
	FA_FID = FeatureAttribOrdering(0)
	FA_FACERANGE = FeatureAttribOrdering(1)
)

type GeometryHeaderProperty C.int
const (
	GH_VERTEXCOUNT = GeometryHeaderProperty(0)
	GH_FEATURECOUNT = GeometryHeaderProperty(1)
)

type MeshTopology C.int
const (
	MT_TRIANGLES = MeshTopology(0)
	MT_LINES = MeshTopology(1)
	MT_POINTS = MeshTopology(2)
)

type Encoding C.int
const (
	ENC_NONE = Encoding(0)
	ENC_STRING = Encoding(1)
)

type LodMetricType C.int
const (
	LMT_MAX_SCREEN_AREA = LodMetricType(0)
	LMT_MAX_SCREEN_SIZE = LodMetricType(1)
	LMT_DISTANCE = LodMetricType(2)
	LMT_EFFECTIVE_DENSITY = LodMetricType(3)
	LMT_SCREEN_SPACE_RELATIVE = LodMetricType(4)
	LMT_NOT_SET = LodMetricType(5)
)

type LayerType C.int
const (
	LT_MESH3D = LayerType(0)
	LT_MESHIM = LayerType(1)
	LT_POINT = LayerType(2)
	LT_POINTCLOUD = LayerType(3)
	LT_BUILDING = LayerType(4)
	LT_VOXEL = LayerType(5)
	LT_GROUP = LayerType(6)
)

type VBBinding C.int
const (
	VBB_PREVERTEX = VBBinding(0)
	VBB_UVREGION = VBBinding(1)
	VBB_FEATURE = VBBinding(2)
)

type NormalReferenceFrame C.int
const (
	NRF_EAST_NORTH_UP = NormalReferenceFrame(0)
	NRF_EAST_CENTERED = NormalReferenceFrame(1)
	NRF_VERTEXREFERENCEFRAME = NormalReferenceFrame(2)
	NRF_NOT_SET = NormalReferenceFrame(3)
)

type CompressedMeshAttribute C.int
const (
	CMA_POSITION = CompressedMeshAttribute(0)
	CMA_NORMAL = CompressedMeshAttribute(1)
	CMA_UV0 = CompressedMeshAttribute(2)
	CMA_COLOR = CompressedMeshAttribute(3)
	CMA_UVREGION = CompressedMeshAttribute(4)
	CMA_FEATUREINDEX = CompressedMeshAttribute(5)
	CMA_FLAG = CompressedMeshAttribute(6)
)

type AttribHeaderProperty C.int
const (
	AH_COUNT = AttribHeaderProperty(0)
	AH_ATTRIBHEADERPROPERTY = AttribHeaderProperty(1)
)

type AttribOrdering C.int
const (
	AO_ATTRIBVALUES = AttribOrdering(0)
	AO_BYTECOUNTS = AttribOrdering(1)
	AO_OBJECTIDS = AttribOrdering(1)
)

type CompressedGeometryFormat C.int
const (
	CGF_NOT_INIT = CompressedGeometryFormat(0)
	CGF_DRACO = CompressedGeometryFormat(1)
	CGF_LEPCC = CompressedGeometryFormat(2)
)

type LegacyImageChannel C.int
const (
	LIC_RGBA = LegacyImageChannel(0)
	LIC_RGB = LegacyImageChannel(1)
	LIC_NOT_SET = LegacyImageChannel(2)
)

type LegacyWrapMode C.int
const (
	LWM_NONE = LegacyWrapMode(0)
	LWM_REPEAT = LegacyWrapMode(1)
	LWM_MIRROR = LegacyWrapMode(2)
)

type MimeImageFormat C.int
const (
	MIF_JPG = MimeImageFormat(0)
	MIF_PNG = MimeImageFormat(1)
	MIF_DDS = MimeImageFormat(2)
	MIF_KTX = MimeImageFormat(3)
	MIF_NOT_SET = MimeImageFormat(4)
)

type LegacyUVSet C.int
const (
	LUS_UV0 = LegacyUVSet(0)
)

type ValueEncoding C.int
const (
	VENC_UTF8 = ValueEncoding(0)
	VENC_NOT_SET = ValueEncoding(1)
)

type AttributeStorageInfoEncoding C.int
const (
	ASE_EMBEDDEDELEVATION = AttributeStorageInfoEncoding(0)
	ASE_LEPCC_XYZ = AttributeStorageInfoEncoding(1)
	ASE_LEPCC_RGB = AttributeStorageInfoEncoding(2)
	ASE_LEPCC_INTENSITY = AttributeStorageInfoEncoding(3)
	ASE_NOT_SET = AttributeStorageInfoEncoding(4)
)

type BslFilterMode C.int
const (
	BFM_SOLID = BslFilterMode(0)
	BFM_WIREFRAME = BslFilterMode(1)
	BFM_NOT_SET = BslFilterMode(2)
)

type HeightModel C.int
const (
	HM_GRAVITY_RELATED = HeightModel(0)
	HM_ELLIPSOIDAL = HeightModel(1)
	HM_ORTHOMETRIC = HeightModel(2)
	HM_NOT_SET = HeightModel(3)
)

type HeightUnit C.int
const (
	HU_METER = HeightUnit(0)
	HU_USFOOT = HeightUnit(1)
	HU_FOOT = HeightUnit(2)
	HU_CLARKEFOOT = HeightUnit(3)
	HU_CLARKEYARD = HeightUnit(4)
	HU_CLARKELINK = HeightUnit(5)
	HU_SEARSYARD = HeightUnit(6)
	HU_SEARSFOOT = HeightUnit(7)
	HU_SEARSCHAIN = HeightUnit(8)
	HU_BENOITCHAIN = HeightUnit(9)
	HU_INDIANYARD = HeightUnit(10)
	HU_INDIAN1937YARD = HeightUnit(11)
	HU_GOLDCOASTFOOT = HeightUnit(12)
	HU_SEARSTRUNCCHAIN = HeightUnit(13)
	HU_USINCH = HeightUnit(14)
	HU_USMILE = HeightUnit(15)
	HU_USYARD = HeightUnit(16)
	HU_MILLIMETER = HeightUnit(17)
	HU_DECIMETER = HeightUnit(18)
	HU_CENTIMETER = HeightUnit(19)
	HU_KILOMETER = HeightUnit(20)
	HU_NOT_SET = HeightUnit(21)
)

type Continuity C.int
const (
	C_CONTINUITY = Continuity(0)
	C_DISCRETE = Continuity(1)
	C_NOT_SET = Continuity(2)
)

type BaseQuantity C.int
const (
	BQ_HORIZONTAL_COORDINATE = BaseQuantity(0)
	BQ_VERTICAL_COORDINATE = BaseQuantity(1)
	BQ_TIME = BaseQuantity(2)
	BQ_NONE = BaseQuantity(3)
	BQ_NOT_SET = BaseQuantity(4)
)

type StatisticsStatus C.int
const (
	SS_FINAL = StatisticsStatus(0)
	SS_PARTIAL = StatisticsStatus(1)
	SS_NOT_SET = StatisticsStatus(2)
)

type PcslAttributeBufferType C.int
const (
	PABT_ELEVATION = PcslAttributeBufferType(1)
	PABT_INTENSITY = PcslAttributeBufferType(2)
	PABT_RGB = PcslAttributeBufferType(4)
	PABT_CLASS_CODE = PcslAttributeBufferType(8)
	PABT_FLAGS = PcslAttributeBufferType(16)
	PABT_RETURNS = PcslAttributeBufferType(32)
	PABT_POINTID = PcslAttributeBufferType(64)
	PABT_USER_DATA = PcslAttributeBufferType(128)
	PABT_POINT_SRC_ID = PcslAttributeBufferType(256)
	PABT_GPS_TIME = PcslAttributeBufferType(512)
	PABT_SCAN_ANGLE = PcslAttributeBufferType(1024)
	PABT_NEAR_INFRARED = PcslAttributeBufferType(2048)
	PABT_NOT_SET = PcslAttributeBufferType(0)
)

type BoundingVolumeType C.int
const (
	BVT_OBB = BoundingVolumeType(0)
	BVT_MBS = BoundingVolumeType(1)
	BVT_NOT_SET = BoundingVolumeType(2)
)

type VxlVariableSemantic C.int
const (
	VVS_STC_HOT_SPOT_RESULTS = VxlVariableSemantic(0)
	VVS_STC_CLUSTER_OUTLIER_RESULTS = VxlVariableSemantic(1)
	VVS_STC_ESTIMATED_BIN = VxlVariableSemantic(2)
	VVS_GENERIC_NEAREST_INTEPOLATED = VxlVariableSemantic(3)
	VVS_NOT_SET = VxlVariableSemantic(4)
)

type VerticalExagMode C.int
const (
	VEM_SCALE_POSITION = VerticalExagMode(0)
	VEM_SCALE_HEIGHT = VerticalExagMode(1)
	VEM_NOT_SET = VerticalExagMode(2)
)

type AttributeMeta struct {
	Key string
	Name string
	Alias string
}

type AttributeDefinition struct {
	Meta AttributeMeta
}