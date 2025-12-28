# XSD2Go - Automatically generate golang xml parser based on XSD

:warning: **This is a fork of [gocomply/xsd2go](https://github.com/gocomply/xsd2go) for internal purposes.** :warning:

Special thanks to the original author, [Lukáš Zapletal](https://github.com/lukas-zapletal) ([isimluk](http://isimluk.com/)), for his great work on this module.

[![Go Report Card](https://goreportcard.com/badge/github.com/inovacc/xsd2go)](https://goreportcard.com/report/github.com/inovacc/xsd2go)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/inovacc/xsd2go)](https://pkg.go.dev/github.com/inovacc/xsd2go)

:warning: **You should run xsd2go, before ever importing `encoding/xml` to your project.** :warning:

You may want to start reading [blog introduction](http://isimluk.com/posts/2020/05/xsd2go-automatically-generate-golang-xml-parsers/) to this project.

## Motivation

Did you ever got to implement XML parser? Perhaps for atom, or scap? You may have got XSD
(XML Schema Definition) files to verify adherence to given xml application? Wouldn't it be
cool to automatically generate XML parser based on XSD definition? That's exactly what we
are up to here.

### Related projects:
 - ![Metaschema](https://github.com/gocomply/metaschema) - generate golang code based on NIST metaschema input
 - ![SCAP](https://github.com/gocomply/scap) - parsers of NIST SCAP family of standards
