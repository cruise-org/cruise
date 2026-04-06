// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package enums

type PageType int

const (
	Home PageType = iota
	Containers
	Images
	Vulnerability
	Monitoring
	Networks
	Volumes
	Nav
	Registry
)

type ErrorType int

const (
	Fatal ErrorType = iota
	Warning
)

type Severity int

const (
	Critical Severity = iota
	High
	Medium
	Low
	Unknown
)
