# Taho HCL Go Project

A Go Module for working with HCL.

## Why this Module Exists

In [taho@0.1.0], we used file based parsing to access comments because at the
time, the [Hashicorp HCLv2 Toolkit] package did not provide a way of working
with comments or whitespace.

This module was created to allow [taho-go@0.1.0] to work with HCL through the
abstraction layer provided by it's interfaces and thus avoid performing file
based parsing.

## What this Module Provides

This module provides a method to read HCL files and directories and returns an
interface that allows callars to interact with the content in an object oriented
way.

<!-- links -->
[Hashicorp HCLv2 Toolkit]: https://pkg.go.dev/github.com/hashicorp/hcl/v2
[taho@0.1.0]: https://github.com/OpenTaho/taho/blob/main/CHANGELOG.md#v0.0.1
[taho-go@0.1.0]: https://github.com/OpenTaho/taho-go/blob/main/CHANGELOG.md#v0.1.0
