package main

var pythonProtoLibraryRuleTemplate = mustTemplate(`load("@rules_proto_grpc//:defs.bzl", "bazel_build_rule_common_attrs", "proto_compile_attrs")
load("@rules_proto_grpc_python_pip_deps//:requirements.bzl", "requirement")
load("@rules_python//python:defs.bzl", "py_library")
load("//:{{ .Lang.Name }}_{{ .Rule.Kind }}_compile.bzl", "{{ .Lang.Name }}_{{ .Rule.Kind }}_compile")

def {{ .Rule.Name }}(name, **kwargs):
    # Compile protos
    name_pb = name + "_pb"
    python_proto_compile(
        name = name_pb,
        {{ .Common.CompileArgsForwardingSnippet }}
    )

    # Create {{ .Lang.Name }} library
    py_library(
        name = name,
        srcs = [name_pb],
        deps = PROTO_DEPS + kwargs.get("deps", []),
        data = kwargs.get("data", []),  # See https://github.com/rules-proto-grpc/rules_proto_grpc/issues/257 for use case
        imports = [name_pb],
        {{ .Common.LibraryArgsForwardingSnippet }}
    )

PROTO_DEPS = [
    Label(requirement("protobuf")),
]`)

var pythonGrpcLibraryRuleTemplate = mustTemplate(`load("@rules_proto_grpc//:defs.bzl", "bazel_build_rule_common_attrs", "proto_compile_attrs")
load("@rules_proto_grpc_python_pip_deps//:requirements.bzl", "requirement")
load("@rules_python//python:defs.bzl", "py_library")
load("//:{{ .Lang.Name }}_{{ .Rule.Kind }}_compile.bzl", "{{ .Lang.Name }}_{{ .Rule.Kind }}_compile")

def {{ .Rule.Name }}(name, **kwargs):
    # Compile protos
    name_pb = name + "_pb"
    python_grpc_compile(
        name = name_pb,
        {{ .Common.CompileArgsForwardingSnippet }}
    )

    # Create {{ .Lang.Name }} library
    py_library(
        name = name,
        srcs = [name_pb],
        deps = GRPC_DEPS + kwargs.get("deps", []),
        data = kwargs.get("data", []),  # See https://github.com/rules-proto-grpc/rules_proto_grpc/issues/257 for use case
        imports = [name_pb],
        {{ .Common.LibraryArgsForwardingSnippet }}
    )

GRPC_DEPS = [
    Label(requirement("grpcio")),
    Label(requirement("protobuf")),
]`)

var pythonGrpclibLibraryRuleTemplate = mustTemplate(`load("@rules_proto_grpc//:defs.bzl", "bazel_build_rule_common_attrs", "proto_compile_attrs")
load("@rules_proto_grpc_python_pip_deps//:requirements.bzl", "requirement")
load("@rules_python//python:defs.bzl", "py_library")
load("//:{{ .Lang.Name }}_grpclib_compile.bzl", "{{ .Lang.Name }}_grpclib_compile")

def {{ .Rule.Name }}(name, **kwargs):
    # Compile protos
    name_pb = name + "_pb"
    python_grpclib_compile(
        name = name_pb,
        {{ .Common.CompileArgsForwardingSnippet }}
    )

    # Create {{ .Lang.Name }} library
    py_library(
        name = name,
        srcs = [name_pb],
        deps = GRPCLIB_DEPS + kwargs.get("deps", []),
        data = kwargs.get("data", []),  # See https://github.com/rules-proto-grpc/rules_proto_grpc/issues/257 for use case
        imports = [name_pb],
        {{ .Common.LibraryArgsForwardingSnippet }}
    )

GRPCLIB_DEPS = [
    Label("@protobuf//:protobuf_python"),
    Label(requirement("grpclib")),
]`)

var pythonModuleExtraLines = `bazel_dep(name = "rules_python", version = "0.27.1")

python = use_extension("@rules_python//python/extensions:python.bzl", "python")
python.toolchain(python_version = "3.11")`

func makePython() *Language {
	return &Language{
		Name:  "python",
		DisplayName: "Python",
		Notes: mustTemplate("Rules for generating Python protobuf and gRPC ``.py`` files and libraries using standard Protocol Buffers and gRPC or `grpclib <https://github.com/vmagamedov/grpclib>`_. Libraries are created with ``py_library`` from ``rules_python``. To use the fast C++ Protobuf implementation, you can add ``--define=use_fast_cpp_protos=true`` to your build, but this requires you setup the path to your Python headers.\n\n.. note:: If you have proto libraries that produce overlapping import paths, be sure to set ``legacy_create_init=False`` on the top level ``py_binary`` or ``py_test`` to ensure all paths are importable."),
		ModuleExtraLines: pythonModuleExtraLines,
		Aliases: map[string]string{
			"py_proto_compile": "python_proto_compile",
			"py_grpc_compile": "python_grpc_compile",
			"py_grpclib_compile": "python_grpclib_compile",
			"py_proto_library": "python_proto_library",
			"py_grpc_library": "python_grpc_library",
			"py_grpclib_library": "python_grpclib_library",
		},
		Rules: []*Rule{
			&Rule{
				Name:             "python_proto_compile",
				Kind:             "proto",
				Implementation:   compileRuleTemplate,
				Plugins:          []string{"//:proto_plugin"},
				BuildExample:     protoCompileExampleTemplate,
				Doc:              "Generates Python protobuf ``.py`` files",
				Attrs:            compileRuleAttrs,
			},
			&Rule{
				Name:             "python_grpc_compile",
				Kind:             "grpc",
				Implementation:   compileRuleTemplate,
				Plugins:          []string{"//:proto_plugin", "//:grpc_plugin"},
				BuildExample:     grpcCompileExampleTemplate,
				Doc:              "Generates Python protobuf and gRPC ``.py`` files",
				Attrs:            compileRuleAttrs,
			},
			&Rule{
				Name:             "python_grpclib_compile",
				Kind:             "grpc",
				Implementation:   compileRuleTemplate,
				Plugins:          []string{"//:proto_plugin", "//:grpclib_plugin"},
				BuildExample:     grpcCompileExampleTemplate,
				Doc:              "Generates Python protobuf and grpclib ``.py`` files (supports Python 3 only)",
				Attrs:            compileRuleAttrs,
				SkipTestPlatforms: []string{"windows", "macos"},
			},
			&Rule{
				Name:             "python_proto_library",
				Kind:             "proto",
				Implementation:   pythonProtoLibraryRuleTemplate,
				BuildExample:     protoLibraryExampleTemplate,
				Doc:              "Generates a Python protobuf library using ``py_library`` from ``rules_python``",
				Attrs:            libraryRuleAttrs,
			},
			&Rule{
				Name:             "python_grpc_library",
				Kind:             "grpc",
				Implementation:   pythonGrpcLibraryRuleTemplate,
				BuildExample:     grpcLibraryExampleTemplate,
				Doc:              "Generates a Python protobuf and gRPC library using ``py_library`` from ``rules_python``",
				Attrs:            libraryRuleAttrs,
				SkipTestPlatforms: []string{"windows"},
			},
			&Rule{
				Name:             "python_grpclib_library",
				Kind:             "grpc",
				Implementation:   pythonGrpclibLibraryRuleTemplate,
				BuildExample:     grpcLibraryExampleTemplate,
				Doc:              "Generates a Python protobuf and grpclib library using ``py_library`` from ``rules_python`` (supports Python 3 only)",
				Attrs:            libraryRuleAttrs,
				SkipTestPlatforms: []string{"windows", "macos"},
			},
		},
	}
}
