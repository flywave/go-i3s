#include "i3s_api.h"

#include <algorithm>
#include <fstream>
#include <stdint.h>
#include <vector>

#include <i3s/i3s_common_.h>
#include <i3s/i3s_writer.h>
#include <utils/utl_i3s_resource_defines.h>
#include <utils/utl_lock.h>
#include <utils/utl_slpk_writer_api.h>

#ifdef __cplusplus
extern "C" {
#endif

struct _i3s_material_data_t {
  i3slib::i3s::Material_data *data;
};

struct _i3s_mesh_data_t {
  i3slib::i3s::Mesh_data *data;
};

struct _i3s_texture_buffer_t {
  i3slib::i3s::Texture_buffer buf;
};

struct _i3s_ctx_properties_t {
  i3slib::i3s::Ctx_properties prop;
};

struct _i3s_writer_context_t {
  i3slib::i3s::Writer_context::Ptr ctx;
};

struct _i3s_layer_writer_t {
  i3slib::i3s::Layer_writer::Ptr ptr;
};

struct _i3s_node_data_t {
  i3slib::i3s::Simple_node_data data;
};

struct _i3s_raw_mesh_t {
  i3slib::i3s::Simple_raw_mesh mesh;
};

struct _i3s_raw_points_t {
  i3slib::i3s::Simple_raw_points pts;
};

struct _i3s_cartesian_transformation_t {
  i3slib::i3s::Cartesian_transformation::Ptr ptr;
};

struct _i3s_layer_meta_t {
  i3slib::i3s::Layer_meta meta;
};

struct _i3s_attribute_meta_t {
  i3slib::i3s::Attribute_meta meta;
};

struct _i3s_texture_meta_t {
  i3slib::i3s::Texture_meta *meta;
};

struct _i3s_spatial_reference_t {
  const i3slib::i3s::Spatial_reference *sr;
};

struct _i3s_slpk_writer_t {
  i3slib::utl::Slpk_writer::Ptr ptr;
};

FLYWAVE_I3S_API i3s_ctx_properties_t *ctx_properties_create() {
  return new i3s_ctx_properties_t{};
}

FLYWAVE_I3S_API void ctx_properties_free(i3s_ctx_properties_t *prop) {
  if (prop) {
    delete prop;
  }
}

FLYWAVE_I3S_API void
ctx_properties_set_max_texture_size(i3s_ctx_properties_t *prop, uint16_t max) {
  prop->prop.max_texture_size = max;
}

FLYWAVE_I3S_API void ctx_properties_is_drop_normals(i3s_ctx_properties_t *prop,
                                                    _Bool flag) {
  prop->prop.is_drop_normals = flag;
}

FLYWAVE_I3S_API void
ctx_properties_set_geom_encoding_support(i3s_ctx_properties_t *prop,
                                         uint32_t tp) {
  prop->prop.geom_encoding_support =
      static_cast<i3slib::i3s::Geometry_compression_flags>(tp);
}

FLYWAVE_I3S_API void
ctx_properties_set_geom_decoding_support(i3s_ctx_properties_t *prop,
                                         uint32_t tp) {
  prop->prop.geom_decoding_support =
      static_cast<i3slib::i3s::Geometry_compression_flags>(tp);
}

FLYWAVE_I3S_API void
ctx_properties_set_gpu_tex_encoding_support(i3s_ctx_properties_t *prop,
                                            uint32_t tp) {
  prop->prop.gpu_tex_encoding_support =
      static_cast<i3slib::i3s::GPU_texture_compression_flags>(tp);
}

FLYWAVE_I3S_API void
ctx_properties_set_gpu_tex_rendering_support(i3s_ctx_properties_t *prop,
                                             uint32_t tp) {
  prop->prop.gpu_tex_rendering_support =
      static_cast<i3slib::i3s::GPU_texture_compression_flags>(tp);
}

FLYWAVE_I3S_API i3s_writer_context_t *
writer_context_create(i3s_ctx_properties_t *prop,
                      i3s_cartesian_transformation_t *tran) {
  return new i3s_writer_context_t{i3slib::i3s::create_i3s_writer_context(
      prop->prop, (tran == nullptr) ? nullptr : tran->ptr)};
}

FLYWAVE_I3S_API void writer_context_free(i3s_writer_context_t *ctx) {
  if (ctx) {
    delete ctx;
  }
}

FLYWAVE_I3S_API i3s_layer_meta_t *layer_meta_create() {
  return new i3s_layer_meta_t{};
}

FLYWAVE_I3S_API void layer_meta_free(i3s_layer_meta_t *nd) {
  if (nd) {
    delete nd;
  }
}

FLYWAVE_I3S_API void layer_meta_set_type(i3s_layer_meta_t *nd, int type) {
  nd->meta.type = static_cast<i3slib::i3s::Layer_type>(type);
}

FLYWAVE_I3S_API void layer_meta_set_name(i3s_layer_meta_t *nd,
                                         const char *name) {
  nd->meta.name = name;
}

FLYWAVE_I3S_API void layer_meta_set_desc(i3s_layer_meta_t *nd,
                                         const char *desc) {
  nd->meta.desc = desc;
}

FLYWAVE_I3S_API void layer_meta_set_copyright(i3s_layer_meta_t *nd,
                                              const char *copyright) {
  nd->meta.copyright = copyright;
}

FLYWAVE_I3S_API void layer_meta_set_spatial_reference_wkid(i3s_layer_meta_t *nd,
                                                           int wkid) {
  nd->meta.sr.wkid = wkid;
}

FLYWAVE_I3S_API void
layer_meta_set_spatial_reference_latest_wkid(i3s_layer_meta_t *nd,
                                             int latest_wkid) {
  nd->meta.sr.latest_wkid = latest_wkid;
}

FLYWAVE_I3S_API void
layer_meta_set_spatial_reference_vcs_id(i3s_layer_meta_t *nd, int vcs_id) {
  nd->meta.sr.vcs_id = vcs_id;
}

FLYWAVE_I3S_API void
layer_meta_set_spatial_reference_latest_vcs_id(i3s_layer_meta_t *nd,
                                               int latest_vcs_id) {
  nd->meta.sr.latest_vcs_id = latest_vcs_id;
}

FLYWAVE_I3S_API void layer_meta_set_spatial_reference_wkt(i3s_layer_meta_t *nd,
                                                          const char *wkt) {
  nd->meta.sr.wkt = wkt;
}

FLYWAVE_I3S_API void layer_meta_set_uid(i3s_layer_meta_t *nd, const char *uid) {
  nd->meta.uid = uid;
}

FLYWAVE_I3S_API void layer_meta_set_drawing_info(i3s_layer_meta_t *nd,
                                                 const char *drawing_info) {
  nd->meta.drawing_info = drawing_info;
}

FLYWAVE_I3S_API void layer_meta_set_elevation_info(i3s_layer_meta_t *nd,
                                                   const char *elevation_info) {
  nd->meta.elevation_info = elevation_info;
}

FLYWAVE_I3S_API void layer_meta_set_popup_info(i3s_layer_meta_t *nd,
                                               const char *popup_info) {
  nd->meta.popup_info = popup_info;
}

FLYWAVE_I3S_API void layer_meta_set_timestamp(i3s_layer_meta_t *nd,
                                              uint64_t timestamp) {
  nd->meta.timestamp = timestamp;
}

FLYWAVE_I3S_API void
layer_meta_set_normal_reference_frame(i3s_layer_meta_t *nd,
                                      int normal_reference_frame) {
  nd->meta.normal_reference_frame =
      static_cast<i3slib::i3s::Normal_reference_frame>(normal_reference_frame);
}

FLYWAVE_I3S_API void
layer_meta_set_height_model_info_height_model(i3s_layer_meta_t *nd,
                                              int height_model) {
  nd->meta.height_model_info.height_model =
      static_cast<i3slib::i3s::Height_model>(height_model);
}

FLYWAVE_I3S_API void
layer_meta_set_height_model_info_vert_crs(i3s_layer_meta_t *nd,
                                          const char *vert_crs) {
  nd->meta.height_model_info.vert_crs = vert_crs;
}

FLYWAVE_I3S_API void
layer_meta_set_height_model_info_height_unit(i3s_layer_meta_t *nd,
                                             int height_unit) {
  nd->meta.height_model_info.height_unit =
      static_cast<i3slib::i3s::Height_unit>(height_unit);
}

FLYWAVE_I3S_API i3s_attribute_meta_t *attribute_meta_create() {
  return new i3s_attribute_meta_t{};
}

FLYWAVE_I3S_API void attribute_meta_free(i3s_attribute_meta_t *nd) {
  if (nd) {
    delete nd;
  }
}

FLYWAVE_I3S_API void attribute_meta_set_key(i3s_attribute_meta_t *nd,
                                            const char *key) {
  nd->meta.key = key;
}

FLYWAVE_I3S_API void attribute_meta_set_name(i3s_attribute_meta_t *nd,
                                             const char *name) {
  nd->meta.name = name;
}

FLYWAVE_I3S_API void attribute_meta_set_alias(i3s_attribute_meta_t *nd,
                                              const char *alias) {
  nd->meta.alias = alias;
}

FLYWAVE_I3S_API i3s_layer_writer_t *
layer_writer_create(i3s_writer_context_t *ctx, const char *path) {
  return new i3s_layer_writer_t{
      i3slib::i3s::Layer_writer::Ptr(i3slib::i3s::create_mesh_layer_builder(
          ctx->ctx, std::filesystem::path(path)))};
}

FLYWAVE_I3S_API i3s_layer_writer_t *
layer_writer_create_with_slpk_writer(i3s_writer_context_t *ctx,
                                     i3s_slpk_writer_t *writer) {
  return new i3s_layer_writer_t{i3slib::i3s::Layer_writer::Ptr(
      i3slib::i3s::create_mesh_layer_builder(ctx->ctx, writer->ptr))};
}

FLYWAVE_I3S_API void layer_writer_free(i3s_layer_writer_t *lw) {
  if (lw) {
    delete lw;
  }
}

FLYWAVE_I3S_API void layer_writer_set_layer_meta(i3s_layer_writer_t *lw,
                                                 i3s_layer_meta_t *lm) {
  lw->ptr->set_layer_meta(lm->meta);
}

FLYWAVE_I3S_API void layer_writer_set_attribute_meta(i3s_layer_writer_t *lw,
                                                     int idx,
                                                     i3s_attribute_meta_t *am,
                                                     int schema_id) {
  lw->ptr->set_attribute_meta(idx, am->meta, schema_id);
}

FLYWAVE_I3S_API void layer_writer_create_output_node(i3s_layer_writer_t *lw,
                                                     i3s_node_data_t *nd,
                                                     size_t node_id) {
  lw->ptr->create_output_node(nd->data, node_id);
}

FLYWAVE_I3S_API void layer_writer_create_node(i3s_layer_writer_t *lw,
                                              i3s_node_data_t *nd,
                                              size_t node_id) {
  lw->ptr->create_node(nd->data, node_id);
}

FLYWAVE_I3S_API void layer_writer_process_children(i3s_layer_writer_t *lw,
                                                   i3s_node_data_t *nd,
                                                   size_t node_id) {
  lw->ptr->process_children(nd->data, node_id);
}

FLYWAVE_I3S_API void layer_writer_create_mesh_from_raw(i3s_layer_writer_t *lw,
                                                       i3s_raw_mesh_t *meshs,
                                                       i3s_node_data_t *nd) {
  lw->ptr->create_mesh_from_raw(meshs->mesh, nd->data.mesh);
}

FLYWAVE_I3S_API void layer_writer_create_points_from_raw(
    i3s_layer_writer_t *lw, i3s_raw_points_t *points, i3s_node_data_t *nd) {
  lw->ptr->create_mesh_from_raw(points->pts, nd->data.mesh);
}

FLYWAVE_I3S_API void layer_writer_save(i3s_layer_writer_t *lw) {
  lw->ptr->save();
}

FLYWAVE_I3S_API i3s_node_data_t *node_data_create() {
  return new i3s_node_data_t{};
}

FLYWAVE_I3S_API void node_data_free(i3s_node_data_t *nd) {
  if (nd) {
    delete nd;
  }
}

FLYWAVE_I3S_API i3s_mesh_data_t *node_data_get_mesh_data(i3s_node_data_t *nd) {
  return new i3s_mesh_data_t{&nd->data.mesh};
}

FLYWAVE_I3S_API void mesh_data_free(i3s_mesh_data_t *md) {
  if (md) {
    delete md;
  }
}

FLYWAVE_I3S_API i3s_material_data_t *
mesh_data_get_material(i3s_mesh_data_t *md) {
  return new i3s_material_data_t{&md->data->material};
}

FLYWAVE_I3S_API void material_data_free(i3s_material_data_t *md) {
  if (md) {
    delete md;
  }
}

FLYWAVE_I3S_API void
material_data_set_material_properties_alpha_mode(i3s_material_data_t *md,
                                                 int alpha_mode) {
  md->data->properties.alpha_mode =
      static_cast<i3slib::i3s::Alpha_mode>(alpha_mode);
}

FLYWAVE_I3S_API void
material_data_set_material_properties_alpha_cut_off(i3s_material_data_t *md,
                                                    int alpha_cut_off) {
  md->data->properties.alpha_cut_off = alpha_cut_off;
}

FLYWAVE_I3S_API void
material_data_set_material_properties_double_sided(i3s_material_data_t *md,
                                                   bool double_sided) {
  md->data->properties.double_sided = double_sided;
}

FLYWAVE_I3S_API void
material_data_set_material_properties_emissive_factor(i3s_material_data_t *md,
                                                      float *emissive_factor) {
  md->data->properties.emissive_factor.x = emissive_factor[0];
  md->data->properties.emissive_factor.y = emissive_factor[1];
  md->data->properties.emissive_factor.z = emissive_factor[2];
}

FLYWAVE_I3S_API void
material_data_set_material_properties_cull_face(i3s_material_data_t *md,
                                                int cull_face) {
  md->data->properties.cull_face =
      static_cast<i3slib::i3s::Face_culling_mode>(cull_face);
}

FLYWAVE_I3S_API void material_data_set_metallic_roughness_base_color_factor(
    i3s_material_data_t *md, float *base_color_factor) {
  md->data->metallic_roughness.base_color_factor.x = base_color_factor[0];
  md->data->metallic_roughness.base_color_factor.y = base_color_factor[1];
  md->data->metallic_roughness.base_color_factor.z = base_color_factor[2];
  md->data->metallic_roughness.base_color_factor.w = base_color_factor[3];
}

FLYWAVE_I3S_API void
material_data_set_metallic_roughness_metallic_factor(i3s_material_data_t *md,
                                                     float metallic_factor) {
  md->data->metallic_roughness.metallic_factor = metallic_factor;
}

FLYWAVE_I3S_API void
material_data_set_metallic_roughness_roughness_factor(i3s_material_data_t *md,
                                                      float roughness_factor) {
  md->data->metallic_roughness.roughness_factor = roughness_factor;
}

FLYWAVE_I3S_API void material_data_append_metallic_roughness_base_color_tex(
    i3s_material_data_t *md, i3s_texture_buffer_t *tex) {
  md->data->metallic_roughness.base_color_tex.emplace_back(tex->buf);
}

FLYWAVE_I3S_API void
material_data_append_metallic_roughness_metallic_roughness_tex(
    i3s_material_data_t *md, i3s_texture_buffer_t *tex) {
  md->data->metallic_roughness.metallic_roughness_tex.emplace_back(tex->buf);
}

FLYWAVE_I3S_API void
material_data_append_normal_map_tex(i3s_material_data_t *md,
                                    i3s_texture_buffer_t *tex) {
  md->data->normal_map_tex.emplace_back(tex->buf);
}

FLYWAVE_I3S_API i3s_texture_buffer_t *texture_buffer_create(int width,
                                                            int height,
                                                            int channel_count,
                                                            const char *data) {
  auto ptr = new i3s_texture_buffer_t{};
  if (i3slib::i3s::create_texture_from_image(width, height, channel_count, data,
                                             ptr->buf))
    return ptr;
  delete ptr;
  return nullptr;
}

FLYWAVE_I3S_API void texture_buffer_free(i3s_texture_buffer_t *nd) {
  if (nd) {
    delete nd;
  }
}

FLYWAVE_I3S_API i3s_texture_meta_t *
texture_buffer_get_meta(i3s_texture_buffer_t *nd) {
  return new i3s_texture_meta_t{&nd->buf.meta};
}

FLYWAVE_I3S_API void texture_meta_free(i3s_texture_meta_t *nd) {
  if (nd) {
    delete nd;
  }
}

FLYWAVE_I3S_API void texture_meta_set_mip0_width(i3s_texture_meta_t *nd,
                                                 int mip0_width) {
  nd->meta->mip0_width = mip0_width;
}

FLYWAVE_I3S_API void texture_meta_set_mip0_height(i3s_texture_meta_t *nd,
                                                  int mip0_height) {
  nd->meta->mip0_height = mip0_height;
}

FLYWAVE_I3S_API void texture_meta_set_mip_count(i3s_texture_meta_t *nd,
                                                int mip_count) {
  nd->meta->mip_count = mip_count;
}

FLYWAVE_I3S_API void texture_meta_set_uv_set(i3s_texture_meta_t *nd,
                                             int uv_set) {
  nd->meta->uv_set = uv_set;
}

FLYWAVE_I3S_API void texture_meta_set_alpha_status(i3s_texture_meta_t *nd,
                                                   int alpha_status) {
  nd->meta->alpha_status =
      static_cast<i3slib::i3s::Texture_meta::Alpha_status>(alpha_status);
}

FLYWAVE_I3S_API void texture_meta_set_wrap_mode(i3s_texture_meta_t *nd,
                                                int wrap_mode) {
  nd->meta->wrap_mode =
      static_cast<i3slib::i3s::Texture_meta::Wrap_mode>(wrap_mode);
}

FLYWAVE_I3S_API void texture_meta_set_format(i3s_texture_meta_t *nd,
                                             int format) {
  nd->meta->format = static_cast<i3slib::i3s::Image_format>(format);
}

FLYWAVE_I3S_API void texture_meta_set_is_atlas(i3s_texture_meta_t *nd,
                                               _Bool is_atlas) {
  nd->meta->is_atlas = is_atlas;
}

FLYWAVE_I3S_API void node_data_set_lod_threshold(i3s_node_data_t *nd,
                                                 double lod_threshold) {
  nd->data.lod_threshold = lod_threshold;
}

FLYWAVE_I3S_API void node_data_set_node_depth(i3s_node_data_t *nd,
                                              int node_depth) {
  nd->data.node_depth = node_depth;
}

FLYWAVE_I3S_API void node_data_append_children(i3s_node_data_t *nd,
                                               size_t node_id) {
  nd->data.children.emplace_back(node_id);
}

FLYWAVE_I3S_API i3s_raw_mesh_t *raw_mesh_create() {
  return new i3s_raw_mesh_t{};
}

FLYWAVE_I3S_API void raw_mesh_free(i3s_raw_mesh_t *nd) {
  if (nd) {
    delete nd;
  }
}

FLYWAVE_I3S_API void raw_mesh_set_vertex(i3s_raw_mesh_t *nd,
                                         const double *vertexs,
                                         const float *uvs, size_t vertex_count,
                                         const uint32_t *indices,
                                         size_t index_count) {
  nd->mesh.vertex_count = vertex_count;
  nd->mesh.abs_xyz = reinterpret_cast<const i3slib::utl::Vec3d *>(vertexs);
  nd->mesh.uv = reinterpret_cast<const i3slib::utl::Vec2f *>(uvs);
  nd->mesh.index_count = index_count;
  nd->mesh.indices = indices;
}

FLYWAVE_I3S_API _Bool raw_mesh_set_texture(i3s_raw_mesh_t *nd, int width,
                                           int height, int channel_count,
                                           const char *data) {
  if (!i3slib::i3s::create_texture_from_image(width, height, channel_count,
                                              data, nd->mesh.img))
    return false;
  return true;
}

FLYWAVE_I3S_API i3s_raw_points_t *raw_points_create() {
  return new i3s_raw_points_t{};
}

FLYWAVE_I3S_API void raw_points_free(i3s_raw_points_t *nd) {
  if (nd) {
    delete nd;
  }
}

FLYWAVE_I3S_API void raw_points_set_points(i3s_raw_points_t *nd,
                                           const double *vertexs,
                                           const uint32_t *fids,
                                           uint64_t vertex_count) {
  nd->pts.count = vertex_count;
  nd->pts.abs_xyz = reinterpret_cast<const i3slib::utl::Vec3d *>(vertexs);
  nd->pts.fids = fids;
}

FLYWAVE_I3S_API char *get_message_string(int string_id) {
  std::string msg = i3slib::utl::get_message_string(string_id);
  char *ret = (char *)malloc(msg.size());
  memcpy(ret, msg.c_str(), msg.size());
  return ret;
}

FLYWAVE_I3S_API void spatial_reference_free(i3s_spatial_reference_t *nd) {
  if (nd) {
    delete nd;
  }
}

FLYWAVE_I3S_API int spatial_reference_get_wkid(i3s_spatial_reference_t *sr) {
  return sr->sr->wkid;
}

FLYWAVE_I3S_API int
spatial_reference_get_latest_wkid(i3s_spatial_reference_t *sr) {
  return sr->sr->latest_wkid;
}

FLYWAVE_I3S_API int spatial_reference_get_vcs_id(i3s_spatial_reference_t *sr) {
  return sr->sr->vcs_id;
}

FLYWAVE_I3S_API int
spatial_reference_get_latest_vcs_id(i3s_spatial_reference_t *sr) {
  return sr->sr->latest_vcs_id;
}

FLYWAVE_I3S_API const char *
spatial_reference_get_wkt(i3s_spatial_reference_t *sr) {
  return sr->sr->wkt.c_str();
}

extern _Bool toCartesian(void *, i3s_spatial_reference_t *, float *, int);
extern _Bool fromCartesian(void *, i3s_spatial_reference_t *, float *, int);

struct go_cartesian_transformation
    : public i3slib::i3s::Cartesian_transformation {
public:
  void *ctx;
  go_cartesian_transformation(void *_ctx) : ctx(_ctx) {}

  virtual bool to_cartesian(const i3slib::i3s::Spatial_reference &sr,
                            i3slib::utl::Vec3d *xyz, int count) override {
    i3s_spatial_reference_t csr{&sr};
    return toCartesian(ctx, &csr, reinterpret_cast<float *>(xyz), count);
  }

  virtual bool from_cartesian(const i3slib::i3s::Spatial_reference &sr,
                              i3slib::utl::Vec3d *xyz, int count) override {
    i3s_spatial_reference_t csr{&sr};
    return fromCartesian(ctx, &csr, reinterpret_cast<float *>(xyz), count);
  }
};

FLYWAVE_I3S_API i3s_cartesian_transformation_t *
cartesian_transformation_create(void *ctx) {
  return new i3s_cartesian_transformation_t{
      i3slib::i3s::Cartesian_transformation::Ptr{
          new go_cartesian_transformation(ctx)}};
}

FLYWAVE_I3S_API void
cartesian_transformation_free(i3s_cartesian_transformation_t *nd) {
  if (nd) {
    delete nd;
  }
}

extern _Bool createArchive(void *, char *, uint32_t);
extern _Bool appendFile(void *, char *, uint8_t *, int, uint32_t, uint32_t);
extern uint8_t *getFile(void *, char *, size_t *);
extern _Bool finalize(void *);

class go_slpk_writer : public i3slib::utl::Slpk_writer {
public:
  DECL_PTR(go_slpk_writer);
  go_slpk_writer(void *_ctx) : ctx(_ctx) {}
  ~go_slpk_writer() override = default;

  void *ctx;

  virtual bool create_archive(const std::filesystem::path &path,
                              Create_flags flags) override {
    return createArchive(ctx, const_cast<char *>(path.c_str()), flags);
  }

  virtual bool append_file(const std::string &path_in_archive,
                           const char *buffer, int n_bytes,
                           i3slib::utl::Mime_type type,
                           i3slib::utl::Mime_encoding pack) override {
    return appendFile(ctx, const_cast<char *>(path_in_archive.c_str()),
                      (uint8_t *)(buffer), n_bytes, uint32_t(type),
                      uint32_t(pack));
  }

  virtual bool get_file(const std::string &path_in_archive,
                        std::string *content) override {
    size_t size;
    uint8_t *buf =
        getFile(ctx, const_cast<char *>(path_in_archive.c_str()), &size);
    content->resize(size);
    memcpy(&(*content)[0], buf, size);
    ::free(buf);
    return true;
  }

  virtual bool finalize() override { return ::finalize(ctx); }
};

FLYWAVE_I3S_API i3s_slpk_writer_t *slpk_writer_create(void *ctx) {
  return new i3s_slpk_writer_t{go_slpk_writer::Ptr(new go_slpk_writer(ctx))};
}

FLYWAVE_I3S_API void slpk_writer_free(i3s_slpk_writer_t *nd) {
  if (nd) {
    delete nd;
  }
}

#ifdef __cplusplus
}
#endif
