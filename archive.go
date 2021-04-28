package i3s

// #include <stdlib.h>
// #include "i3s_api.h"
// #cgo CFLAGS: -I ./lib
// #cgo CXXFLAGS: -I ./lib
// #cgo linux LDFLAGS:  -L ./lib -L /usr/lib/x86_64-linux-gnu -Wl,--start-group  -lstdc++ -lm -pthread -ldl -lci3s -lzlib -ldraco -li3s -ljpeg -llepcc -lpng -Wl,--end-group
// #cgo windows LDFLAGS: -L ./lib　-lci3s -lzlib -ldraco -li3s -ljpeg -llepcc -lpng
// #cgo darwin LDFLAGS: -L　./lib -lci3s -lzlib -ldraco -li3s -ljpeg -llepcc -lpng
import "C"

type NodeData struct {
	m *C.struct__i3s_node_data_t
}

type RawMesh struct {
	m *C.struct__i3s_raw_mesh_t
}

type RawPoints struct {
	m *C.struct__i3s_raw_points_t
}

type CartesianTransformation struct {
	m *C.struct__i3s_cartesian_transformation_t
}

type WriterContext struct {
	m *C.struct__i3s_writer_context_t
}

type LayerWriter struct {
	m *C.struct__i3s_layer_writer_t
}
