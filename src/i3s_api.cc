#include "i3s_api.h"

#include <i3s/i3s_writer.h>
#include <i3s/i3s_common_.h>
#include <utils/utl_slpk_writer_api.h>

#ifdef __cplusplus
extern "C" {
#endif

struct _i3s_material_data_t {
  i3slib::i3s::Material_data data;
};

struct _i3s_mesh_data_t {
  i3slib::i3s::Mesh_data data;
};

struct _i3s_attribute_buffer_t {
  i3slib::i3s::Attribute_buffer::Ptr data;
};
 
struct _i3s_geometry_buffer_t {
  i3slib::i3s::Geometry_buffer::Ptr data;
};

struct _i3s_texture_buffer_t {
  i3slib::i3s::Texture_buffer::Ptr ptr;
};

struct _i3s_ctx_properties_t {
  i3slib::i3s::Ctx_properties prop;
};

struct _i3s_context_t {
  i3slib::i3s::Context::Ptr ctx;
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

struct _i3s_slpk_writer_t {
  i3slib::utl::Slpk_writer::Ptr ptr;
};

#ifdef __cplusplus
}
#endif
