cmake_minimum_required(VERSION 3.6)

project(LEPCC CXX C)
string(TOLOWER ${PROJECT_NAME} PROJECT_NAME_LOWER)

set(ROOT_DIR "${PROJECT_SOURCE_DIR}")

file(GLOB BASE_SRCS
    src/*.cpp
    )
set (CMAKE_CXX_STANDARD 11)

INCLUDE_DIRECTORIES("${CMAKE_CURRENT_SOURCE_DIR}/include")

set(LEPCC_INCLUDE_DIR "${PROJECT_SOURCE_DIR}/include")

file(GLOB TEST_SRCS
    ${PROJECT_SOURCE_DIR}/src/Test_C_Api.cpp)

list(REMOVE_ITEM BASE_SRCS ${TEST_SRCS})

add_library(lepcc STATIC "${BASE_SRCS}")

SOURCE_GROUP("src" FILES ${BASE_SRCS})

SET_TARGET_PROPERTIES(lepcc
		                  PROPERTIES PUBLIC_HEADER "${headers}")

IF(FLYWAVE_ENABLE_SOLUTION_FOLDERS)
  SET_TARGET_PROPERTIES(lepcc PROPERTIES FOLDER external)
ENDIF(FLYWAVE_ENABLE_SOLUTION_FOLDERS)
