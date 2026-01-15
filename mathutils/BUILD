load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name="mathutils",
    srcs=["mathutils.go"],
    importpath = "bazel-demo/mathutils",
    visibility = ["//visibility:public"],
)

go_test(
    name="mathutils_test",
    srcs=["mathutils_test.go"],
    embed=[":mathutils"]
)
