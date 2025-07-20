# glua-kit-example

This repository contains a few examples of how to concisely place a remote dependency on `glua-kit` in your Bazel project. Some example Bazel targets are below.

At the heart of this example is the `MODULE.bazel` file which trivially adds `glua_kit` Bazel repo to the project.

## Standalone bytecode compilation

This uses a `glua-kit` Bazel rule to build a single Lua file into Lua bytecode. Run the following:

```
bazel build //lua:test.luac
```

You can verify the file:

```
file bazel-bin/lua/test.luac
bazel-bin/lua/test.luac: Lua bytecode, version 5.4
```

## Using in a Go application

```
bazel run //go:demo
```

or if you're having linker issues:

```
bazel run --linkopt=-fuse-ld=gold //go:demo
```

The Go binary simply declares an external dependency on the `glua-kit` Go Lua VM library:

```
deps = [
    "@glua_kit//go/vm",
],
```