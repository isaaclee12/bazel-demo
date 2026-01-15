load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name="stringutils",
    srcs=["stringutils.go"],
    importpath = "bazel-demo/stringutils",
    visibility = ["//visibility:public"],
)

go_test(
    name="stringutils_test",
    srcs=["stringutils_test.go"],
    embed=[":stringutils"]
)
