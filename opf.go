// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import (
	"errors"
	"io"
	"reflect"
	"strings"
)

type xmlOPF struct {
	Metadata meta       `xml:"metadata"`
	Manifest []manifest `xml:"manifest>item"`
	Spine    spine      `xml:"spine"`
}
type meta struct {
	Title       []string     `xml:"title"`
	Language    []string     `xml:"language"`
	Identifier  []identifier `xml:"identifier"`
	Creator     []author     `xml:"creator"`
	Subject     []string     `xml:"subject"`
	Description []string     `xml:"description"`
	Publisher   []string     `xml:"publisher"`
	Contributor []author     `xml:"contributor"`
	Date        []date       `xml:"date"`
	Type        []string     `xml:"type"`
	Format      []string     `xml:"format"`
	Source      []string     `xml:"source"`
	Relation    []string     `xml:"relation"`
	Coverage    []string     `xml:"coverage"`
	Rights      []string     `xml:"rights"`
	Meta        []metafield  `xml:"meta"`
}
type identifier struct {
	Data   string `xml:",chardata"`
	ID     string `xml:"id,attr"`
	Scheme string `xml:"scheme,attr"`
}
type author struct {
	Data   string `xml:",chardata"`
	FileAs string `xml:"file-as,attr"`
	Role   string `xml:"role,attr"`
}
type date struct {
	// TODO: convert date to date type?
	Data  string `xml:",chardata"`
	Event string `xml:"event,attr"`
}
type metafield struct {
	Name    string `xml:"name,attr"`
	Content string `xml:"content,attr"`
}
type manifest struct {
	ID           string `xml:"id,attr"`
	Href         string `xml:"href,attr"`
	MediaType    string `xml:"media-type,attr"`
	Fallback     string `xml:"media-fallback,attr"`
	Properties   string `xml:"properties,attr"`
	MediaOverlay string `xml:"media-overlay,attr"`
}
type spine struct {
	ID              string      `xml:"id,attr"`
	Toc             string      `xml:"toc,attr"`
	PageProgression string      `xml:"page-progression-direction,attr"`
	Items           []spineItem `xml:"itemref"`
}
type spineItem struct {
	IDref      string `xml:"idref,attr"`
	Linear     string `xml:"linear,attr"`
	ID         string `xml:"id,attr"`
	Properties string `xml:"properties,attr"`
}

func parseOPF(opf io.Reader) (*xmlOPF, error) {
	var o xmlOPF
	err := decodeXML(opf, &o)
	if err != nil {
		return nil, err
	}

	return &o, nil
}

func (opf xmlOPF) ncxPath() string {
	if opf.Spine.Toc != "" {
		fileID := opf.Spine.Toc
		path := opf.filePath(fileID)
		if path != "" {
			return path
		}
	}

	fileID := "ncx"
	return opf.filePath(fileID)
}

func (opf xmlOPF) filePath(id string) string {
	for _, item := range opf.Manifest {
		if item.ID == id {
			return item.Href
		}
	}
	return ""
}

func (opf xmlOPF) toMData() mdata {
	m := opf.Metadata
	metadata := make(mdata)
	v := reflect.ValueOf(m)
	typeOf := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Len() == 0 {
			continue
		}

		fieldName := strings.ToLower(typeOf.Field(i).Name)
		data := make([]mdataElement, field.Len())
		for j := 0; j < field.Len(); j++ {
			element := field.Index(j).Interface()
			data[j] = elementToMData(element)
		}
		metadata[fieldName] = data
	}
	return metadata
}

func elementToMData(element interface{}) (result mdataElement) {
	result.attr = make(map[string]string)
	switch element.(type) {
	case string:
		result.content, _ = element.(string)
	case identifier:
		ident, _ := element.(identifier)
		result.content = ident.Data
		result.attr["id"] = ident.ID
		result.attr["scheme"] = ident.Scheme
	case author:
		auth, _ := element.(author)
		result.content = auth.Data
		result.attr["file-as"] = auth.FileAs
		result.attr["role"] = auth.Role
	case date:
		d, _ := element.(date)
		result.content = d.Data
		result.attr["event"] = d.Event
	case metafield:
		m, _ := element.(metafield)
		result.content = m.Content
		result.attr["name"] = m.Name
		result.attr["content"] = m.Content
	}
	return
}

func (opf xmlOPF) spineLength() int {
	return len(opf.Spine.Items)
}

func (opf xmlOPF) spineURL(index int) string {
	idref := opf.Spine.Items[index].IDref
	url, _ := opf.getURL(idref)
	return url
}

func (opf xmlOPF) getURL(id string) (string, error) {
	for _, item := range opf.Manifest {
		if item.ID == id {
			return item.Href, nil
		}
	}
	return "", errors.New("ID " + id + " not in the manifest")
}
