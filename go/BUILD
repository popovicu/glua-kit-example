load("@rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "demo",
    srcs = [
        "demo.go",
    ],
    cgo = True,
    deps = [
        "@glua_kit//go/vm",
    ],
)
