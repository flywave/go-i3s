package i3s

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	METADATA_NAME   = string("metadata.json")
	SCENE_LAYER     = string("3dSceneLayer.json")
	NODE_INDEX      = string("3dNodeIndexDocument.json")
	SHARED_RESOURCE = string("sharedResource.json")
	SLPK_HASH       = string("@specialIndexFileHASH128@")
	NODES           = string("nodes")
	SHARED          = string("shared")
	NODEPAGES       = string("nodepages")
	STATISTICS      = string("statistics")
	ATTRIBUTES      = string("attributes")
	GEOMETRIES      = string("geometries")
)

const (
	EXT_GZ     = string(".gz")
	EXT_JSON   = string(".json")
	EXT_BIN    = string(".bin")
	EXT_PCCXYZ = string(".bin.pccxyz")
	EXT_PCCRGB = string(".bin.pccrgb")
	EXT_PCCINT = string(".bin.pccint")
)

type SpecialFile struct {
	Name     string
	Required bool
}

type SpecialFiles []SpecialFile

var (
	DefaultLayout = SpecialFiles{
		SpecialFile{SCENE_LAYER, true},
		SpecialFile{SLPK_HASH, false},
		SpecialFile{NODE_INDEX, true},
		SpecialFile{SHARED_RESOURCE, true},
		SpecialFile{METADATA_NAME, false},
	}
)

type FolderPattern string

const (
	FOLDER_PATTERN_BASIC    = FolderPattern("basic")
	FOLDER_PATTERN_EXTENDED = FolderPattern("extended")
)

type ArchiveCompression string

const (
	ARCHIVE_COMPRESSION_STORE     = ArchiveCompression("store")
	ARCHIVE_COMPRESSION_DEFLATE64 = ArchiveCompression("deflate64")
	ARCHIVE_COMPRESSION_DEFLATE   = ArchiveCompression("deflate")
)

type ResourceCompression string

const (
	RESOURCE_COMPRESSION_NONE = ResourceCompression("none")
	RESOURCE_COMPRESSION_GZIP = ResourceCompression("gzip")
)

type VersionType string

type Metadata struct {
	FolderPattern       FolderPattern       `json:"folderPattern"`
	ArchiveCompression  ArchiveCompression  `json:"archiveCompression"`
	ResourceCompression ResourceCompression `json:"resourceCompression"`
	Version             VersionType         `json:"version"`
	NodeCount           int                 `json:"nodeCount"`
}

func (m *Metadata) ToJson() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func LoadMetadata(path string) *Metadata {
	bs, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	var cs *Metadata
	json.NewDecoder(bs).Decode(&cs)
	return cs
}

type HeightModelInfo struct {
}

type SpatialReference struct{}

type Cardinality struct{}

type IndexScheme struct{}

type HeaderAttribute struct{}

type GeometryAttribute struct{}

type GeometrySchema struct{}

type Encoding struct{}

type ExtInfo struct{}

type PcslIndex struct{}

type PreferredEncoding struct{}

type Store struct{}

type ElevatioInfo struct{}

type CompressionEncoding struct{}

type AttributeValue struct{}

type SceneLayerInfo struct{}

type MinimumBoundingSphere struct{}

type NodeReference struct{}

type FeatureRange struct{}

type Resource struct{}

type LodSelection struct{}

type Feature struct{}

type OrientedBoundingBoxes struct{}

type PcslNode struct{}

type NodePage struct{}

type Node struct{}

type Tree struct{}

type PcslTree struct{}

type Material struct{}

type Image struct{}

type Texture struct{}

type SharedResource struct{}

type GeometryReference struct{}

type arrayBufferView struct{}

type FeatureData struct{}

type NodeInfo struct{}
