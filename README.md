# apis

With the recent release of Bazel 7 and the preference for bzlmod dependency management,
I've created a simple project that uses the most current versions of Bazel (v7-bzlmod), Go (1.21.5), Protobuf (v21.7), and gRPC (v1.48.1).
The goal is to have a simple app that shows how to use the latest technologies to generate protobuf and grpc Go files from proto files
using GoogleApi (i.e. annotations.proto).

QUESTION:

At this time, it seems that I have to manually generate from hello_world.proto to hellow_world.pb.go and copy the hello_world.pb.go file 
to the proto's folder: proto/hello_world. If I don't copy it, then the build fails due to a missing dependency on the generated file.

??? Is there a way to automate or improve this process?