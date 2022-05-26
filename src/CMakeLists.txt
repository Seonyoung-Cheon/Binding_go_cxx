cmake_minimum_required(VERSION 3.18)

set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} ${CMAKE_CURRENT_LIST_DIR}/cmake)

set(SOURCE_FILES
  goCppBind.go)

# ADD_EXECUTABLE(go-make ${SOURCE_FICE})

add_go_executable(go-make ${SOURCE_FILES})



