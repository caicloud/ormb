<p align="center">
<img src="docs/images/logo.png" height="150">
</p>


[![Build Status](https://travis-ci.com/caicloud/ormb.svg?branch=master)](https://travis-ci.com/caicloud/ormb)
[![Coverage Status](https://coveralls.io/repos/github/caicloud/ormb/badge.svg?branch=master)](https://coveralls.io/github/caicloud/ormb?branch=master)

ormb is an open-source model registry to manage machine learning model. 

ormb helps you manage your Machine Learning/Deep Learning models. It makes your models easy to create, version, share and publish.

[![asciicast](https://asciinema.org/a/345812.svg)](https://asciinema.org/a/345812)

## Installation

You can install the pre-compiled binary, or compile from source.

### Install the pre-compiled binary

Download the pre-compiled binaries from [the releases](https://github.com/caicloud/ormb/releases) page and copy to the desired location.

### Compile from source

Clone:

```
$ git clone https://github.com/caicloud/ormb
$ cd ormb
```

Get the dependencies:

```
$ go mod tidy
```

Build:

```
$ make build-local
```

Verify it works:

```
$ ./bin/ormb --help
```

## Usage

### Save the model

```bash
$ ormb save ./resnet_v2_fp32_savedmodel_NCHW caicloud/resnetv2:v1
ref:     caicloud/resnetv2:v1
digest:  f51973c855608ab06d8f5e4333925a635f87f01ff992ffc5f9988f26d1da24e9
size:    90.6 MiB
format:  SavedModel
v1: saved
```

### Push the model to a remote registry

```bash
$ ormb push caicloud/resnetv2:v1
The push refers to repository [caicloud/resnetv2]
ref:     caicloud/resnetv2:v1
digest:  f51973c855608ab06d8f5e4333925a635f87f01ff992ffc5f9988f26d1da24e9
size:    90.6 MiB
format:  SavedModel
v1: pushed to remote (1 layer, 90.6 MiB total)
```

### Pull the model from a remote registry

```bash
$ ormb pull caicloud/resnetv2:v1 
v1: Pulling from caicloud/resnetv2
ref:     caicloud/resnetv2:v1
digest:  f51973c855608ab06d8f5e4333925a635f87f01ff992ffc5f9988f26d1da24e9
size:    90.6 MiB
format:  SavedModel
Status: Downloaded newer model for caicloud/resnetv2:v1
```

### Export the model to the current directory

```bash
$ ormb export caicloud/resnetv2:v1
ref:     localhost:5000/caicloud/resnetv2:v1
digest:  f51973c855608ab06d8f5e4333925a635f87f01ff992ffc5f9988f26d1da24e9
size:    90.6 MiB

$ tree ./resnet_v2_fp32_savedmodel_NCHW
resnet_v2_fp32_savedmodel_NCHW
├── 1538687196
│   ├── saved_model.pb
│   └── variables
│       ├── variables.data-00000-of-00001
│       └── variables.index

2 directories, 3 files
```

## Tutorial

### Training and serving

Please have a look at [docs/tutorial.md](docs/tutorial.md)

### Serving with Seldon Core

Please have a look at [docs/tutorial-serving-seldon.md](docs/tutorial-serving-seldon.md)

## OCI Model Configuration Specification

Please have a look at [docs/spec_v1alpha1.md](docs/spec-v1alpha1.md)

## Community

ormb project is part of Clever, a Cloud Native Machine Learning platform. We are going to open source a community edition soon.

The Clever slack workspace is caicloud-clever.slack.com. To join, click this [invitation to our Slack workspace](https://join.slack.com/t/caicloud-clever/shared_invite/zt-efz4rdrm-kcOg0Qvs_B8aIWGdZv9E6g).
