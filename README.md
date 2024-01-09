# apis

Recently, many exciting releases of key technologies used in microservice architectures have been released.
Unfortunately, I couldn't find any examples online that use the latest releases of those technologies, so I have created
this sample, which works out the kinks I encountered. Although it is a simple project, it does present real-world
requirements and usage.

This system is composed of gRPC and HTTP proxy servers. You use the gRPC client or an HTTP client, such as curl or
Postman, to interact with the servers.

* Bazel (v7-bzlmod)
* Go (1.21.5)
* Protobuf (v21.7)
* gRPC (v1.48.1)
* gRPC-Gateway HTTP Proxy
* GoogleApis

## DEPENDENCY UPDATES

For Go-based projects, the go.mod file is the "source of truth" for dependencies in Bazel.
The latest version of rules_go eliminated the need to create and use a Gazelle rule, "gazelle-update-repos", to update
dependencies in a Go application. After adding, updating, or removing a dependency in go.mod, do a 'go mod tidy' and
verify that the application can be built by:

```go build ./..```

After you've verified that you can build the application using Go Build, there are a few steps to draw the changes into
Bazel.

* If you've added or removed a DIRECT dependency from go.mod, then you need to amend the MODULE.bazel file. This file
  has a use_repo() block, which contains entries for all DIRECT dependencies, so add or remove the dependency from the
  block, as required.

* After you've adjusted MODULE.bazel, you will use Gazelle to update dependencies for Bazel. I recommend committing to
  git (local) prior to performing this step so you can easily identify any files that are changed and review the
  changes. Gazelle will sometimes undo changes that I've made in BUILD.bazel files. Here is the command to use for this
  step:

```bazel run //:gazelle```

## ISSUE

One of Bazel's core tenets is that it must be able to perform reproducible builds,
so it builds in a hermetically sealed environment ("sandbox"), and the sources are immutable during a build.
Hence, during a build, the generated files cannot be copied to the sources involved.

This is an issue with a Go project because the module system will not locate the generated files in the sandbox,
which results in missing dependency errors. This issue has been open for years, and there is no immediate solution
on the horizon, so most teams will use two distinct steps -- generate files and build the application.

## CD/CD & LOCAL DEVELOPMENT PROCEDURE

So, the workaround is the following:

Create a file named .bazelignore in the proto/hello_world/v1 folder; its content is "hello_world.pb.go".

1) Invoke Bazel on your generation target;
2) Copy the generated file to a location in your sources; and
3) Build your application;

For this project, the process is the following:

1. ```
   bazel build //proto/hello_world/v1:hello_world
   ```


2. Manually copy the generated files:
   ```
   from
   
   bazel-bin/proto/hello_world/v1/hello_world_pb/github.com/abitofhelp/apis/proto/hello_world/v1/proto/hello_world/v1/hello_world.pb.go
   bazel-bin/proto/hello_world/v1/hello_world_pb/github.com/abitofhelp/apis/proto/hello_world/v1/proto/hello_world/v1/hello_world.pb.gw.go
   bazel-bin/proto/hello_world/v1/hello_world_pb/github.com/abitofhelp/apis/proto/hello_world/v1/proto/hello_world/v1/hello_world_grpc.pb.go
   
   to
   
   proto/hello_world/v1/hello_world.pb.go
   proto/hello_world/v1/hello_world.pb.gw.go
   proto/hello_world/v1/hello_world_grpc.pb.go
   
   AND from
   
   bazel-bin/proto/enum/access_tier/v5/access_tier_pb/github.com/abitofhelp/apis/proto/enum/access_tier/v5/proto/enum/access_tier/v5/access_tier.pb.go
   bazel-bin/proto/enum/access_tier/v5/access_tier_pb/github.com/abitofhelp/apis/proto/enum/access_tier/v5/proto/enum/access_tier/v5/access_tier.pb.gw.go
   bazel-bin/proto/enum/access_tier/v5/access_tier_pb/github.com/abitofhelp/apis/proto/enum/access_tier/v5/proto/enum/access_tier/v5/access_tier_grpc.pb.go
   
   TO
   
   proto/enum/access_tier/v5/access_tier.pb.go
   proto/enum/access_tier/v5/access_tier.pb.gw.go
   proto/enum/access_tier/v5/access_tier_grpc.pb.go
   ```

3. ```
   bazel build //...
   ```

To test the build, use three separate terminal windows:

```
  # Start the gRPC server:
  bazel run //server:server
  
  # Start the HTTP Proxy server:
  bazel run //proxy:proxy
  
  # Use the gRPC client to send the request.
  bazel run //client:client  
  
  # Use the HTTP client to send the request.
  curl --location 'http://localhost:50052/v1/hello' \
    --header 'Content-Type: application/json' \
    --data '{"name": "Mike","sentUtc": "2024-01-03T01:33:31Z"}'
```

## LOCAL DEVELOPMENT PROCESS IMPROVEMENT

The aforementioned is a procedure that can be used for CI/CD as well as local development. However, for local work, it
may be possible to avoid the two-step process and manual copying of generated files. If you use Visual Studio Code,
there is a plug-in available that integrates generated files into Go's builds without having to copy them. For Goland
developers, there is a Bazel plug-in that provides the same capability. Still, you have to use the following menu for
Goland to configure the build of Bazel projects properly.
**File -> Import Bazel Project**
After I imported this project, I did not have to copy generated files, so my development cycle was fast.

## Please note that for CI/CD builds, I am unaware of any way to avoid the two-step process I've described.