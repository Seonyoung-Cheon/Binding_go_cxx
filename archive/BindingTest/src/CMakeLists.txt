cmake_minimum_required(VERSION 3.18)

project(BindingTest LANGUAGES CXX )
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} "${CMAKE_CURRENT_SOURCE_DIR}/cmake")

cmake_minimum_required(VERSION 3.0)

add_library(BindingTest SHARED
  nummer.cpp
  num.cpp
  )

target_link_directories(BindingTest PUBLIC /usr/lib)

target_include_directories(BindingTest
  PRIVATE include
  )

target_compile_features(BindingTest PUBLIC cxx_std_17)
target_compile_options(BindingTest PRIVATE
  $<$<COMPILE_LANGUAGE:CXX>:-Wall>
  )

set(ENV{CGO_LDFLAGS} -L${CMAKE_CURRENT_SOURCE_DIR}/src )
