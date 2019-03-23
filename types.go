package rpmdbinfo

import "encoding/xml"

// RPMHeader is the root structure from rpmdb per package installed
type RPMHeader struct {
	XMLName xml.Name `xml:"rpmHeader" json:"-"`
	Tags    []RPMTag `xml:"rpmTag"`
}

// RPMTag is the encapsulation for each set of info per package
type RPMTag struct {
	Name   string   `xml:"name,attr"`
	Values []string `xml:",any"`
}

// Tags show the names of the tags available on this package object
func (hdr RPMHeader) Tags() []string {
	tags := []string{}
	for _, tag := range hdr.Tags {
		tags = append(tags, tag.Name)
	}
	return tags
}

// Name is a convience for "Name" tag name
func (hdr RPMHeader) Name() string {
	return hdr.Tag("Name")[0]
}

// Version is a convience for "Version" tag name
func (hdr RPMHeader) Version() string {
	return hdr.Tag("Version")[0]
}

// Release is a convience for "Release" tag name
func (hdr RPMHeader) Release() string {
	return hdr.Tag("Release")[0]
}

// Arch is a convience for "Arch" tag name
func (hdr RPMHeader) Arch() string {
	return hdr.Tag("Arch")[0]
}

// License is a convience for "License" tag name
func (hdr RPMHeader) License() string {
	return hdr.Tag("License")[0]
}

// Sourcerpm is a convience for "Sourcerpm" tag name
func (hdr RPMHeader) Sourcerpm() string {
	return hdr.Tag("Sourcerpm")[0]
}

// Tag fetches the set of information for the provided tag name
func (hdr RPMHeader) Tag(tagname string) []string {
	for _, tag := range hdr.Tags {
		if tag.Name == tagname {
			return tag.Values
		}
	}
	return []string{""}
}
