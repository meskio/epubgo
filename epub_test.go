// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import "testing"

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"os"
)

const (
	bookPath          = "testdata/a_dogs_tale.epub"
	bookTitle         = "A Dog's Tale"
	bookLang          = "en"
	bookIdentifier    = "http://www.gutenberg.org/ebooks/3174"
	bookCreator       = "Mark Twain"
	bookSubject       = "Dogs -- Fiction"
	bookRights        = "Public domain in the USA."
	lenMetadatafields = 9
	identifierScheme  = "URI"
	creatorFileAs     = "Twain, Mark"
	metaName          = "cover"
	htmlFile          = "@public@vhost@g@gutenberg@html@files@3174@3174-h@3174-h-0.htm.html"
	fileId            = "item8"
	htmlPath          = "3174/" + htmlFile
	noNCXPath         = "testdata/noncx.epub"
	invalidNCXPath    = "testdata/invalidncx.epub"
	fileCapsPath      = "testdata/fileCaps.epub"
)

func TestOpenClose(t *testing.T) {
	f, err := Open(bookPath)
	if err != nil {
		t.Errorf("Open(%v) return an error: %v", bookPath, err)
	}

	f.Close()
}

func TestLoad(t *testing.T) {
	file, _ := os.Open(bookPath)
	fileInfo, _ := file.Stat()
	f, err := Load(file, fileInfo.Size())
	if err != nil {
		t.Errorf("Open(%v) return an error: %v", bookPath, err)
	}

	f.Close()
}

func TestOpenFile(t *testing.T) {
	f, _ := Open(bookPath)
	defer f.Close()

	html, err := f.OpenFile(htmlFile)
	if err != nil {
		t.Errorf("OpenFile(%v) return an error: %v", htmlFile, err)
		return
	}
	defer html.Close()

	zipFile, _ := zip.OpenReader(bookPath)
	defer zipFile.Close()
	var file *zip.File
	for _, file = range zipFile.Reader.File {
		if file.Name == htmlPath {
			break
		}
	}
	zipHTML, _ := file.Open()

	buff1, err := ioutil.ReadAll(html)
	if err != nil {
		t.Errorf("Error reading the opened file: %v", err)
		return
	}
	buff2, _ := ioutil.ReadAll(zipHTML)
	if !bytes.Equal(buff1, buff2) {
		t.Errorf("The files on zip and OpenFile are not equal")
		return
	}
}

func TestOpenFileId(t *testing.T) {
	f, _ := Open(bookPath)
	defer f.Close()

	html, err := f.OpenFileId(fileId)
	if err != nil {
		t.Errorf("OpenFileId(%v) return an error: %v", fileId, err)
		return
	}
	defer html.Close()

	zipFile, _ := zip.OpenReader(bookPath)
	defer zipFile.Close()
	var file *zip.File
	for _, file = range zipFile.Reader.File {
		if file.Name == htmlPath {
			break
		}
	}
	zipHTML, _ := file.Open()

	buff1, err := ioutil.ReadAll(html)
	if err != nil {
		t.Errorf("Error reading the opened file: %v", err)
		return
	}
	buff2, _ := ioutil.ReadAll(zipHTML)
	if !bytes.Equal(buff1, buff2) {
		t.Errorf("The files on zip and OpenFile are not equal")
		return
	}
}

func TestNoNCX(t *testing.T) {
	f, err := Open(noNCXPath)
	if err != nil {
		t.Errorf("Open(%v) return an error: %v", noNCXPath, err)
	}
	f.Close()
}

func TestInvalidNCX(t *testing.T) {
	f, err := Open(invalidNCXPath)
	if err != nil {
		t.Errorf("Open(%v) return an error: %v", invalidNCXPath, err)
	}
	f.Close()
}

func TestFileCaps(t *testing.T) {
	f, err := Open(fileCapsPath)
	if err != nil {
		t.Errorf("Open(%v) return an error: %v", fileCapsPath, err)
	}
	f.Close()
}

func TestMetadata(t *testing.T) {
	f, _ := Open(bookPath)
	defer f.Close()

	if title, _ := f.Metadata("title"); title[0] != bookTitle {
		t.Errorf("Metadata title '%v', the expected was '%v'", title[0], bookTitle)
	}
	if language, _ := f.Metadata("language"); language[0] != bookLang {
		t.Errorf("Metadata language '%v', the expected was '%v'", language[0], bookLang)
	}
	if identifier, _ := f.Metadata("identifier"); identifier[0] != bookIdentifier {
		t.Errorf("Metadata identifier '%v', the expected was '%v'", identifier[0], bookIdentifier)
	}
	if creator, _ := f.Metadata("creator"); creator[0] != bookCreator {
		t.Errorf("Metadata creator '%v', the expected was '%v'", creator[0], bookCreator)
	}
	if subject, _ := f.Metadata("subject"); subject[0] != bookSubject {
		t.Errorf("Metadata subject '%v', the expected was '%v'", subject[0], bookSubject)
	}
	if rights, _ := f.Metadata("rights"); rights[0] != bookRights {
		t.Errorf("Metadata rights '%v', the expected was '%v'", rights[0], bookRights)
	}
}

func TestMetadataFields(t *testing.T) {
	f, _ := Open(bookPath)
	defer f.Close()

	fields := f.MetadataFields()
	if len(fields) != lenMetadatafields {
		t.Errorf("len(MetadataFields()) should be %v, but was %v", lenMetadatafields, len(fields))
	}

	isTitle := false
	for _, field := range fields {
		if field == "title" {
			isTitle = true
		}
	}
	if !isTitle {
		t.Errorf("title is not in the metadata fields")
	}
}

func TestMetadataAttr(t *testing.T) {
	f, _ := Open(bookPath)
	defer f.Close()

	if identifier, _ := f.MetadataAttr("identifier"); identifier[0]["scheme"] != identifierScheme {
		t.Errorf("Metadata identifier attr scheme '%v', the expected was '%v'", identifier[0]["scheme"], identifierScheme)
	}
	if creator, _ := f.MetadataAttr("creator"); creator[0]["file-as"] != creatorFileAs {
		t.Errorf("Metadata creator attr file-as '%v', the expected was '%v'", creator[0]["file-as"], creatorFileAs)
	}
	if meta, _ := f.MetadataAttr("meta"); meta[0]["name"] != metaName {
		t.Errorf("Metadata meta attr name '%v', the expected was '%v'", meta[0]["name"], metaName)
	}
}
