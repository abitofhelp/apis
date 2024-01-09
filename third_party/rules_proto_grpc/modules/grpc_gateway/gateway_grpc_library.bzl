"""Generated definition of gateway_grpc_library."""

load("@rules_go//go:def.bzl", "go_library")
load("@rules_proto_grpc//:defs.bzl", "bazel_build_rule_common_attrs", "proto_compile_attrs")
load("@rules_proto_grpc_go//:defs.bzl", "GRPC_DEPS")
load("//:gateway_grpc_compile.bzl", "gateway_grpc_compile")

def gateway_grpc_library(name, **kwargs):
    # Compile protos
    name_pb = name + "_pb"
    gateway_grpc_compile(
        name = name_pb,
        prefix_path = kwargs.get("prefix_path", kwargs.get("importpath", "")),
        **{
            k: v
            for (k, v) in kwargs.items()
            if (k in proto_compile_attrs.keys() and k != "prefix_path") or
               k in bazel_build_rule_common_attrs
        }  # Forward args
    )

    # Create go library
    go_library(
        name = name,
        srcs = [name_pb],
        deps = kwargs.get("go_deps", []) + GATEWAY_DEPS + GRPC_DEPS + kwargs.get("deps", []),
        importpath = kwargs.get("importpath"),
        **{
            k: v
            for (k, v) in kwargs.items()
            if k in bazel_build_rule_common_attrs
        }  # Forward Bazel common args
    )

GATEWAY_DEPS = [
    Label("@com_github_grpc_ecosystem_grpc_gateway_v2//runtime:go_default_library"),
    Label("@com_github_grpc_ecosystem_grpc_gateway_v2//utilities:go_default_library"),
    # Label("@go_googleapis//google/api:annotations_go_proto"),  # https://github.com/bazelbuild/bazel-central-registry/issues/1113
]
