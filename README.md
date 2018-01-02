# gorai

[![Build Status](https://travis-ci.org/lukealonso/gorai.svg?branch=master)](https://travis-ci.org/lukealonso/gorai)

A Golang RaiBlocks library, node and wallet.

This is an early work-in-progress. The initial goal is to create an offline wallet for airgapped use, followed by a full node and online wallet.

##### Requirements

 - [Bazel](https://docs.bazel.build/versions/master/install.html) >= 0.8.0
 - All other dependencies are provided by the Bazel build.

##### Building

```sh
bazel build //...:all
```

##### Running Tests
```sh
bazel test //...:all
```
