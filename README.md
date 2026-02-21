# Taho HCL Go Project

A Go Module for working with HCL.

## Why this Module Exists

In `taho@0.0.x`, we used file based parsing to access comments because at the
time, the [Hashicorp HCLv2 Toolkit] package did not provide a way of working
with comments or whitespace. The code in `taho@0.0.x` is a mess with some parts
of the Hcl files parsed using the toolkit while other parts are parsed using
rudimentary file operations. Content is often parsed using the HCLv2 toolkit then
individual blocks are written out to the file system so they can be reparsed
usin file system based techniques.

This module was created to allow fugure versions of Taho to work with HCL
through the abstraction layer provided by it's interfaces and avoid performing
file based parsing.

## What this Module Provides

This module provides a method to read HCL files and directories and returns an
interface that allows callars to interact with the content in an object oriented
way.

Various implementations of the interface defined by this module are planned.
The [taho-go-hcl-thin] implementation aims to parse content without using the
[Hashicorp HCLv2 Toolkit] while the [taho-go-hcl-hcl-v2] will be based on the
`taho@0.0.x`.

<!-- links -->
[Hashicorp HCLv2 Toolkit]: https://pkg.go.dev/github.com/hashicorp/hcl/v2
