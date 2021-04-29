# Copyright 2020 Esri
# 
# Licensed under the Apache License, Version 2.0 (the "License"); you may not use
# this file except in compliance with the License. You may obtain a copy of
# the License at http://www.apache.org/licenses/LICENSE-2.0
# Unless required by applicable law or agreed to in writing, software distributed
# under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
# CONDITIONS OF ANY KIND, either express or implied. See the License for the
# specific language governing permissions and limitations under the License.
# 
# For additional information, contact:
# Environmental Systems Research Institute, Inc.
# Attn: Contracts Dept
# 380 New York Street
# Redlands, California, USA 92373
# email: contracts@esri.com

cmake_minimum_required(VERSION 3.10)

set(DEFAULT_BUILD_TYPE "Release")
if(NOT CMAKE_BUILD_TYPE AND NOT CMAKE_CONFIGURATION_TYPES)
  message(STATUS "Setting build type to '${DEFAULT_BUILD_TYPE}' as none was specified.")
  set(CMAKE_BUILD_TYPE ${DEFAULT_BUILD_TYPE} CACHE STRING "Choose the type of build." FORCE)
endif()

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF) # Please give us -std=c++17, not -std=gnu++17 or something

option(NO_ETC2_SUPPORT "Disable ETC2 support.")

project(i3s)

set(I3SLIB_SOURCES
  src/i3s/i3s_context_impl.cpp
  src/i3s/i3s_enums_generated.cpp
  src/i3s/i3s_layer_dom.cpp
  src/i3s/i3s_legacy_mesh.cpp
  src/i3s/i3s_writer_impl.cpp
  src/utils/dxt/IntelDXTCompressor.cpp
  src/utils/json/jsoncpp.cpp
  src/utils/utl_base64.cpp
  src/utils/utl_buffer.cpp
  src/utils/utl_bvh.cpp
  src/utils/utl_colors.cpp
  src/utils/dxt/utl_dxt_mipmap_dds.cpp
  src/utils/utl_fs.cpp
  src/utils/utl_gzip.cpp
  src/utils/utl_image_2d.cpp
  src/utils/utl_image_resize.cpp
  src/utils/utl_jpeg.cpp
  src/utils/utl_libdraco_api.cpp
  src/utils/utl_md5.cpp
  src/utils/utl_mime.cpp
  src/utils/utl_obb.cpp
  src/utils/utl_png.cpp
  src/utils/utl_quaternion.cpp
  src/utils/utl_resource_strings.cpp
  src/utils/utl_serialize_json.cpp
  src/utils/utl_serialize_json_dom.cpp
  src/utils/utl_slpk_writer_impl.cpp
  src/utils/utl_string.cpp
  src/utils/utl_variant.cpp
  src/utils/utl_zip_archive_impl.cpp
  src/utils/utl_etc2comp_c_api.cpp)

add_library(i3s STATIC ${I3SLIB_SOURCES})

# If we don't do this, CMake would define i3s_EXPORTS symbol by default. 
set_target_properties(i3s PROPERTIES DEFINE_SYMBOL "I3S_LIB")

if(WIN32)
  target_compile_definitions(i3s PRIVATE -D_UNICODE -DUNICODE)
else()
  target_compile_options(i3s PRIVATE -fvisibility=hidden -Wall)

  # Suppress complaints on MSVS-specific pragmas.
  target_compile_options(i3s PRIVATE -Wno-unknown-pragmas)
endif()

if(CMAKE_CXX_COMPILER_ID STREQUAL "Clang")
  # Disable this really useless but annoying Clang-specific parentheses warning.
  target_compile_options(i3s PRIVATE -Wno-logical-op-parentheses)
elseif(CMAKE_CXX_COMPILER_ID STREQUAL "GNU")
  # GCC can warn about usage of undefined macros.
  target_compile_options(i3s PRIVATE -Wundef)
  # GCC also has this weird parentheses warning.
  target_compile_options(i3s PRIVATE -Wno-parentheses)
  # FIXME: we have too much of those, fix'em all and remove this
  target_compile_options(i3s PRIVATE -Wno-sign-compare)
endif()

#target_compile_definitions(i3s PRIVATE -DI3S_EXPORT_API $<$<CONFIG:Release>:-DNDEBUG>)

target_include_directories(i3s PUBLIC include) 

target_include_directories(i3s PRIVATE
  src
  ${CMAKE_CURRENT_SOURCE_DIR}/../
  ${CMAKE_CURRENT_SOURCE_DIR}/../zlib
  ${CMAKE_CURRENT_SOURCE_DIR}/../libpng
  ${CMAKE_CURRENT_SOURCE_DIR}/../libjpeg
  ${CMAKE_CURRENT_BINARY_DIR}/../libjpeg
  ${CMAKE_CURRENT_SOURCE_DIR}/../draco/src
  ${CMAKE_CURRENT_BINARY_DIR}/../draco/
  ${CMAKE_CURRENT_SOURCE_DIR}/../lepcc/include
  ${CMAKE_CURRENT_SOURCE_DIR}/../lerc/include
  ${CMAKE_CURRENT_SOURCE_DIR}/../etc2comp/EtcLib/Etc
  ${CMAKE_CURRENT_SOURCE_DIR}/../etc2comp/EtcLib/EtcCodec)

target_link_libraries(i3s stdc++fs)
set(CMAKE_SHARED_LINKER_FLAGS "${CMAKE_SHARED_LINKER_FLAGS} -Wl,--no-undefined")

target_link_libraries(i3s zlib libpng libjpeg draco lepcc)
target_link_libraries(i3s EtcLib)

IF(FLYWAVE_ENABLE_SOLUTION_FOLDERS)
  SET_TARGET_PROPERTIES(i3s PROPERTIES FOLDER external)
ENDIF(FLYWAVE_ENABLE_SOLUTION_FOLDERS)
