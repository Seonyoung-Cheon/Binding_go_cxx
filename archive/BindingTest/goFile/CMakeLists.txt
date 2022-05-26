set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} "${CMAKE_CURRENT_SOURCE_DIR}/cmake")

cmake_minimum_required(VERSION 3.0)

project(gotest Go)

include(cmake/golang.cmake)
include(cmake/flags.cmake)

set(SOURCE_FILE
	num.go
	num_test.go
	lib.go
	)

# add_custom_command(OUTPUT ${CMAKE_CURRENT_BINARY_DIR}/${LIB}
#		DEPENDS ${SOURCE_FILE}
#		WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
#		COMMAND env GOPATH=${GOPATH} go build -buildmode=c-archive
#		-o "${CMAKE_CURRENT_BINARY_DIR}/${LIB}"
#		${CMAKE_GO_FLAGS} ./...
#		COMMENT "Building Go library")

# add_custom_target(${TARGET} DEPENDS ${LIB} ${HEADER})

# add_go_executable(gotest ${SOURCE_FILE})
# set_target_properties(gotest
# PROPERTIES
# IMPORTED_LOCATION ${CMAKE_CURRENT_BINARY_DIR}/${LIB}
# INTERFACE_INCLUDE_DIRECTORIES ${CMAKE_CURRENT_BINARY_DIR})
