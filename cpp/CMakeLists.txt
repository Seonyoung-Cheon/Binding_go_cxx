cmake_minimum_required(VERSION 3.18)

set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} ${CMAKE_CURRENT_LIST_DIR}/cmake)

add_library(goCppBind SHARED
  goCppBind.cpp
  )

target_link_directories(goCppBind PUBLIC /usr/lib)

target_compile_features(goCppBind PUBLIC cxx_std_17)
target_compile_options(goCppBind PRIVATE
  $<$<COMPILE_LANGUAGE:CXX>:-Wall>
  )
