#pragma once

#include <stdbool.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

#if defined(WIN32) || defined(WINDOWS) || defined(_WIN32) || defined(_WINDOWS)
#define FLYWAVE_VDB_API __declspec(dllexport)
#else
#define FLYWAVE_VDB_API
#endif

typedef struct _i3s_material_data_t i3s_material_data_t;
typedef struct _i3s_mesh_data_t i3s_mesh_data_t;
typedef struct _i3s_texture_buffer_t i3s_texture_buffer_t;
typedef struct _i3s_ctx_properties_t i3s_ctx_properties_t;
typedef struct _i3s_context_t i3s_context_t;
typedef struct _i3s_writer_context_t i3s_writer_context_t;
typedef struct _i3s_layer_writer_t i3s_layer_writer_t;
typedef struct _i3s_node_data_t i3s_node_data_t;
typedef struct _i3s_raw_mesh_t i3s_raw_mesh_t;
typedef struct _i3s_raw_points_t i3s_raw_points_t;
typedef struct _i3s_cartesian_transformation_t i3s_cartesian_transformation_t;
typedef struct _i3s_slpk_writer_t i3s_slpk_writer_t;

#ifdef __cplusplus
}
#endif
