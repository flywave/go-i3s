#pragma once

#include <stdbool.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

#if defined(WIN32) || defined(WINDOWS) || defined(_WIN32) || defined(_WINDOWS)
#define FLYWAVE_I3S_API __declspec(dllexport)
#else
#define FLYWAVE_I3S_API
#endif

typedef struct _i3s_material_data_t i3s_material_data_t;
typedef struct _i3s_texture_buffer_t i3s_texture_buffer_t;
typedef struct _i3s_ctx_properties_t i3s_ctx_properties_t;
typedef struct _i3s_writer_context_t i3s_writer_context_t;
typedef struct _i3s_layer_writer_t i3s_layer_writer_t;
typedef struct _i3s_node_data_t i3s_node_data_t;
typedef struct _i3s_raw_mesh_t i3s_raw_mesh_t;
typedef struct _i3s_raw_points_t i3s_raw_points_t;
typedef struct _i3s_cartesian_transformation_t i3s_cartesian_transformation_t;
typedef struct _i3s_layer_meta_t i3s_layer_meta_t;
typedef struct _i3s_attribute_meta_t i3s_attribute_meta_t;
typedef struct _i3s_mesh_data_t i3s_mesh_data_t;
typedef struct _i3s_texture_meta_t i3s_texture_meta_t;

FLYWAVE_I3S_API i3s_ctx_properties_t *ctx_properties_create();
FLYWAVE_I3S_API void ctx_properties_free(i3s_ctx_properties_t *prop);

FLYWAVE_I3S_API void
ctx_properties_set_max_texture_size(i3s_ctx_properties_t *prop, uint16_t max);
FLYWAVE_I3S_API void ctx_properties_is_drop_normals(i3s_ctx_properties_t *prop,
                                                    _Bool flag);
FLYWAVE_I3S_API void
ctx_properties_set_geom_encoding_support(i3s_ctx_properties_t *prop,
                                         uint32_t tp);
FLYWAVE_I3S_API void
ctx_properties_set_geom_decoding_support(i3s_ctx_properties_t *prop,
                                         uint32_t tp);
FLYWAVE_I3S_API void
ctx_properties_set_gpu_tex_encoding_support(i3s_ctx_properties_t *prop,
                                            uint32_t tp);
FLYWAVE_I3S_API void
ctx_properties_set_gpu_tex_rendering_support(i3s_ctx_properties_t *prop,
                                             uint32_t tp);

FLYWAVE_I3S_API i3s_writer_context_t *
writer_context_create(i3s_ctx_properties_t *prop,
                      i3s_cartesian_transformation_t *tran);
FLYWAVE_I3S_API void writer_context_free(i3s_writer_context_t *ctx);

FLYWAVE_I3S_API i3s_layer_meta_t *layer_meta_create();
FLYWAVE_I3S_API void layer_meta_free(i3s_layer_meta_t *nd);
FLYWAVE_I3S_API void layer_meta_set_type(i3s_layer_meta_t *nd, int type);
FLYWAVE_I3S_API void layer_meta_set_name(i3s_layer_meta_t *nd,
                                         const char *name);
FLYWAVE_I3S_API void layer_meta_set_desc(i3s_layer_meta_t *nd,
                                         const char *desc);
FLYWAVE_I3S_API void layer_meta_set_copyright(i3s_layer_meta_t *nd,
                                              const char *copyright);

FLYWAVE_I3S_API void layer_meta_set_spatial_reference_wkid(i3s_layer_meta_t *nd,
                                                           int wkid);
FLYWAVE_I3S_API void
layer_meta_set_spatial_reference_latest_wkid(i3s_layer_meta_t *nd,
                                             int latest_wkid);
FLYWAVE_I3S_API void
layer_meta_set_spatial_reference_vcs_id(i3s_layer_meta_t *nd, int vcs_id);
FLYWAVE_I3S_API void
layer_meta_set_spatial_reference_latest_vcs_id(i3s_layer_meta_t *nd,
                                               int latest_vcs_id);
FLYWAVE_I3S_API void layer_meta_set_spatial_reference_wkt(i3s_layer_meta_t *nd,
                                                          const char *wkt);

FLYWAVE_I3S_API void layer_meta_set_uid(i3s_layer_meta_t *nd, const char *uid);
FLYWAVE_I3S_API void layer_meta_set_drawing_info(i3s_layer_meta_t *nd,
                                                 const char *drawing_info);
FLYWAVE_I3S_API void layer_meta_set_elevation_info(i3s_layer_meta_t *nd,
                                                   const char *elevation_info);
FLYWAVE_I3S_API void layer_meta_set_popup_info(i3s_layer_meta_t *nd,
                                               const char *popup_info);
FLYWAVE_I3S_API void layer_meta_set_timestamp(i3s_layer_meta_t *nd,
                                              uint64_t timestamp);
FLYWAVE_I3S_API void
layer_meta_set_normal_reference_frame(i3s_layer_meta_t *nd,
                                      int normal_reference_frame);
FLYWAVE_I3S_API void
layer_meta_set_height_model_info_height_model(i3s_layer_meta_t *nd,
                                              int height_model);
FLYWAVE_I3S_API void
layer_meta_set_height_model_info_vert_crs(i3s_layer_meta_t *nd,
                                          const char *vert_crs);
FLYWAVE_I3S_API void
layer_meta_set_height_model_info_height_unit(i3s_layer_meta_t *nd,
                                             int height_unit);

FLYWAVE_I3S_API i3s_attribute_meta_t *attribute_meta_create();
FLYWAVE_I3S_API void attribute_meta_free(i3s_attribute_meta_t *nd);
FLYWAVE_I3S_API void attribute_meta_set_key(i3s_attribute_meta_t *nd,
                                            const char *key);
FLYWAVE_I3S_API void attribute_meta_set_name(i3s_attribute_meta_t *nd,
                                             const char *name);
FLYWAVE_I3S_API void attribute_meta_set_alias(i3s_attribute_meta_t *nd,
                                              const char *alias);

FLYWAVE_I3S_API i3s_layer_writer_t *
layer_writer_create(i3s_writer_context_t *ctx, const char *path);
FLYWAVE_I3S_API void layer_writer_free(i3s_layer_writer_t *nd);
FLYWAVE_I3S_API void layer_writer_set_layer_meta(i3s_layer_writer_t *nd,
                                                 i3s_layer_meta_t *lm);
FLYWAVE_I3S_API void layer_writer_set_attribute_meta(i3s_layer_writer_t *nd,
                                                     int idx,
                                                     i3s_attribute_meta_t *am,
                                                     int schema_id);
FLYWAVE_I3S_API void layer_writer_create_output_node(i3s_layer_writer_t *lw,
                                                     i3s_node_data_t *nd,
                                                     uint64_t node_id);
FLYWAVE_I3S_API void layer_writer_create_node(i3s_layer_writer_t *lw,
                                              i3s_node_data_t *nd,
                                              uint64_t node_id);
FLYWAVE_I3S_API void layer_writer_process_children(i3s_layer_writer_t *lw,
                                                   i3s_node_data_t *nd,
                                                   uint64_t node_id);
FLYWAVE_I3S_API void layer_writer_create_mesh_from_raw(i3s_layer_writer_t *lw,
                                                       i3s_raw_mesh_t *mesh,
                                                       i3s_node_data_t *nd);
FLYWAVE_I3S_API void layer_writer_create_points_from_raw(
    i3s_layer_writer_t *lw, i3s_raw_points_t *points, i3s_node_data_t *nd);
FLYWAVE_I3S_API void layer_writer_save(i3s_layer_writer_t *lw);

FLYWAVE_I3S_API i3s_node_data_t *node_data_create();
FLYWAVE_I3S_API void node_data_free(i3s_node_data_t *nd);
FLYWAVE_I3S_API void node_data_set_lod_threshold(i3s_node_data_t *nd,
                                                 double lod_threshold);
FLYWAVE_I3S_API void node_data_set_node_depth(i3s_node_data_t *nd,
                                              int node_depth);
FLYWAVE_I3S_API void node_data_append_children(i3s_node_data_t *nd,
                                               uint64_t node_id);
FLYWAVE_I3S_API i3s_mesh_data_t *node_data_get_mesh_data(i3s_node_data_t *nd);

FLYWAVE_I3S_API void mesh_data_free(i3s_mesh_data_t *md);
FLYWAVE_I3S_API i3s_material_data_t *
mesh_data_get_material(i3s_mesh_data_t *md);

FLYWAVE_I3S_API void material_data_free(i3s_material_data_t *md);
FLYWAVE_I3S_API void
material_data_set_material_properties_alpha_mode(i3s_material_data_t *md,
                                                 int alpha_mode);
FLYWAVE_I3S_API void
material_data_set_material_properties_alpha_cut_off(i3s_material_data_t *md,
                                                    int alpha_cut_off);
FLYWAVE_I3S_API void
material_data_set_material_properties_double_sided(i3s_material_data_t *md,
                                                   bool double_sided);
FLYWAVE_I3S_API void
material_data_set_material_properties_emissive_factor(i3s_material_data_t *md,
                                                      float *emissive_factor);
FLYWAVE_I3S_API void
material_data_set_material_properties_cull_face(i3s_material_data_t *md,
                                                int cull_face);

FLYWAVE_I3S_API void material_data_set_metallic_roughness_base_color_factor(
    i3s_material_data_t *md, float *base_color_factor);
FLYWAVE_I3S_API void
material_data_set_metallic_roughness_metallic_factor(i3s_material_data_t *md,
                                                     float metallic_factor);
FLYWAVE_I3S_API void
material_data_set_metallic_roughness_roughness_factor(i3s_material_data_t *md,
                                                      float roughness_factor);
FLYWAVE_I3S_API void material_data_append_metallic_roughness_base_color_tex(
    i3s_material_data_t *md, i3s_texture_buffer_t *tex);
FLYWAVE_I3S_API void
material_data_append_metallic_roughness_metallic_roughness_tex(
    i3s_material_data_t *md, i3s_texture_buffer_t *tex);

FLYWAVE_I3S_API void
material_data_append_normal_map_tex(i3s_material_data_t *md,
                                    i3s_texture_buffer_t *tex);

FLYWAVE_I3S_API i3s_texture_buffer_t *texture_buffer_create(int width,
                                                            int height,
                                                            int channel_count,
                                                            const char *data);
FLYWAVE_I3S_API void texture_buffer_free(i3s_texture_buffer_t *nd);
FLYWAVE_I3S_API i3s_texture_meta_t *
texture_buffer_get_meta(i3s_texture_buffer_t *nd);
FLYWAVE_I3S_API void texture_meta_free(i3s_texture_meta_t *nd);

FLYWAVE_I3S_API void texture_meta_set_mip0_width(i3s_texture_meta_t *nd,
                                                 int mip0_width);
FLYWAVE_I3S_API void texture_meta_set_mip0_height(i3s_texture_meta_t *nd,
                                                  int mip0_height);
FLYWAVE_I3S_API void texture_meta_set_mip_count(i3s_texture_meta_t *nd,
                                                int mip_count);
FLYWAVE_I3S_API void texture_meta_set_uv_set(i3s_texture_meta_t *nd,
                                             int uv_set);
FLYWAVE_I3S_API void texture_meta_set_alpha_status(i3s_texture_meta_t *nd,
                                                   int alpha_status);
FLYWAVE_I3S_API void texture_meta_set_wrap_mode(i3s_texture_meta_t *nd,
                                                int wrap_mode);
FLYWAVE_I3S_API void texture_meta_set_format(i3s_texture_meta_t *nd,
                                             int format);
FLYWAVE_I3S_API void texture_meta_set_is_atlas(i3s_texture_meta_t *nd,
                                               _Bool is_atlas);

FLYWAVE_I3S_API i3s_raw_mesh_t *raw_mesh_create();
FLYWAVE_I3S_API void raw_mesh_free(i3s_raw_mesh_t *nd);
FLYWAVE_I3S_API void
raw_mesh_set_vertex(i3s_raw_mesh_t *nd, const double *vertexs, const float *uvs,
                    uint64_t vertex_count, const uint32_t *indices,
                    uint64_t index_count);
FLYWAVE_I3S_API _Bool raw_mesh_set_texture(i3s_raw_mesh_t *nd, int width,
                                           int height, int channel_count,
                                           const char *data);

FLYWAVE_I3S_API i3s_raw_points_t *raw_points_create();
FLYWAVE_I3S_API void raw_points_free(i3s_raw_points_t *nd);
FLYWAVE_I3S_API void raw_points_set_points(i3s_raw_points_t *nd,
                                           const double *vertexs,
                                           const uint32_t *fids,
                                           uint64_t vertex_count);

#ifdef __cplusplus
}
#endif
