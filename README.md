# CloudFlow API

![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/ysicing/cloudflow?filename=go.mod&style=flat-square)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/ysicing/cloudflow?style=flat-square)
![GitHub](https://img.shields.io/badge/license-AGPL-blue)
[![Releases](https://img.shields.io/github/release-pre/ysicing/cloudflow.svg)](https://github.com/ysicing/cloudflow/releases)
[![docs](https://img.shields.io/badge/docs-done-green)](https://blog.ysicing.net/)

Schema of the API types that are served by `CloudFlow`

## Purpose

This library is the canonical location of the CloudFlow API definition.

We recommend using the go types in this repo. You may serialize them directly to JSON.


| Kubernetes Version in your Project | CloudFlow Version in your Project  | Import cloudflow-api  |
|------------------------------------|----------------------------|----------------------------|
| >= 1.23                            | v0.x.y         |                               >= v0.1       |

## Where does it come from?

`cloudflow-api` is synced from [https://github.com/ysicing/cloudflow/tree/master/apis](https://github.com/ysicing/cloudflow/tree/master/apis).
Code changes are made in that location, merged into `ysicing/cloudflow` and later synced here.

## Things you should NOT do

[https://github.com/ysicing/cloudflow/tree/master/apis](https://github.com/ysicing/cloudflow/tree/master/apis) is synced to here.
All changes must be made in the former. The latter is read-only.
