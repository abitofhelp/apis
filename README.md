# apis

Recently, many exciting releases of key technologies used in microservice architectures have been released.  Unfortunately, I couldn't find any examples online that use the latest releases of those technologies, so I have created this sample, which works out the kinks I encountered.  
Bazel (v7-bzlmod)
Go (1.21.5)
Protobuf (v21.7)
gRPC (v1.48.1) 
GoogleApis

QUESTION:
When building, I receive a missing dependency error unless I copy the hello_world.pb.go file, which I manually generate from hello_world.proto, to the proto/hello_world folder.

??? Is there a way to automate or improve this process?
