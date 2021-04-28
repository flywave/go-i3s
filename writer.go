package i3s

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