load("@rules_go//go:def.bzl", "go_binary")

# load("@gazelle//:def.bzl", "gazelle_binary")

# gazelle_binary(
#     name = "gazelle_binary",
#     languages = [
#         "@gazelle//language/proto",
#         "@gazelle//language/go",
#         "@gazelle_cc//language/cc",
#     ],
# )

go_binary(
    name = "main",
    srcs = ["main.go"],
    deps = ["//mathutils", "//stringutils"]
)
