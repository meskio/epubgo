package epub

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"io"
)

/* TODO for full support on metadata: 
 *   - identifier: id & opf:scheme
 *   - creator: opf:file-as="King, Martin Luther Jr." opf:role="aut"
 *   - contributor: opf:file-as & opf:role
 *   - date: opf:event & convert to date type
 */

// Metadata of the epub
// TODO: the public struct should not have the `xml`
type Meta struct {
	Title       []string `xml:"metadata>title"`
	Language    []string `xml:"metadata>language"`
	Identifier  []string `xml:"metadata>identifier"`
	Creator     []string `xml:"metadata>creator"`
	Subject     []string `xml:"metadata>subject"`
	Description []string `xml:"metadata>description"`
	Publisher   []string `xml:"metadata>publisher"`
	Contributor []string `xml:"metadata>contributor"`
	Date        []string `xml:"metadata>date"`
	Type        []string `xml:"metadata>type"`
	Format      []string `xml:"metadata>format"`
	Source      []string `xml:"metadata>source"`
	Relation    []string `xml:"metadata>relation"`
	Coverage    []string `xml:"metadata>coverage"`
	Rights      []string `xml:"metadata>rights"`
}

func openFile(file *zip.ReadCloser, path string) (io.ReadCloser, error) {
	for _, f := range file.File {
		if f.Name == path {
			return f.Open()
		}
	}
	return nil, errors.New("File " + path + " not found")
}

type rootfile struct {
	Path string `xml:"full-path,attr"`
}
type container_xml struct {
	// FIXME: only support for one rootfile, can it be more than one?
	Rootfile rootfile `xml:"rootfiles>rootfile"`
}

func contentPath(file *zip.ReadCloser) (string, error) {
	f, err := openFile(file, "META-INF/container.xml")
	if err != nil {
		return "", err
	}
	defer f.Close()

	var c container_xml
	decoder := xml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		return "", err
	}
	return c.Rootfile.Path, nil
}

func parseMetadata(file *zip.ReadCloser) (metadata Meta, err error) {
	path, err := contentPath(file)
	if err != nil {
		return
	}

	f, err := openFile(file, path)
	if err != nil {
		return
	}
	defer f.Close()

	decoder := xml.NewDecoder(f)
	err = decoder.Decode(&metadata)
	if err != nil {
		return
	}

	return
}
