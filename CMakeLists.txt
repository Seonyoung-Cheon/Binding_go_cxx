
cmake_minimum_required(VERSION 3.20)

set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} "${CMAKE_CURRENT_SOURCE_DIR}/cmake")
# set(CMAKE_MODUL_PATH ${CMAKE_CURRENT_LIST_DIR}/)

project(go-make LANGUAGES Go CXX)

project(goCppBind LANGUAGES CXX)
include(cmake/golang.cmake)
include(cmake/flags.cmake)

add_subdirectory(src)
add_subdirectory(cpp)
add_dependencies(go-make goCppBind)
# TARGET_LINK_LIBRARIES(go-make PUBLIC goCppBind)
