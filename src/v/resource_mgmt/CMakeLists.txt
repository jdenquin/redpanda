v_cc_library(
  NAME resource_mgmt
  SRCS
    available_memory.cc
    memory_sampling.cc
    cpu_profiler.cc
    logger.cc
  DEPS
    Seastar::seastar
    v::ssx
    v::config
  )

v_cc_library(
  NAME storage_resource_mgmt
  SRCS
    storage.cc
  DEPS
    Seastar::seastar
    v::cloud_storage
    v::cluster
  )

add_subdirectory(tests)
