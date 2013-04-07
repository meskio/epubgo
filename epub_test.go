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
	book_path         = "testdata/a_dogs_tale.epub"
	book_title        = "A Dog's Tale"
	book_lang         = "en"
	book_identifier   = "http://www.gutenberg.org/ebooks/3174"
	book_creator      = "Mark Twain"
	book_subject      = "Dogs -- Fiction"
	book_rights       = "Public domain in the USA."
	len_metadafields  = 9
	identifier_scheme = "URI"
	creator_file_as   = "Twain, Mark"
	meta_name         = "cover"
	html_file         = "@public@vhost@g@gutenberg@html@files@3174@3174-h@3174-h-0.htm.html"
	html_path         = "3174/" + html_file
)

func TestOpenClose(t *testing.T) {
	f, err := Open(book_path)
	if err != nil {
		t.Errorf("Open(%v) return an error: %v", book_path, err)
	}

	f.Close()
}

func TestLoad(t *testing.T) {
	file, _ := os.Open(book_path)
	fileInfo, _ := file.Stat()
	f, err := Load(file, fileInfo.Size())
	if err != nil {
		t.Errorf("Open(%v) return an error: %v", book_path, err)
	}

	f.Close()
}

func TestOpenFile(t *testing.T) {
	f, _ := Open(book_path)
	defer f.Close()

	html, err := f.OpenFile(html_file)
	if err != nil {
		t.Errorf("OpenFile(%v) return an error: %v", html_file, err)
		return
	}
	defer html.Close()

	zipFile, _ := zip.OpenReader(book_path)
	defer zipFile.Close()
	var file *zip.File
	for _, file = range zipFile.Reader.File {
		if file.Name == html_path {
			break
		}
	}
	zipHtml, _ := file.Open()

	buff1, err := ioutil.ReadAll(html)
	if err != nil {
		t.Errorf("Error reading the opened file: %v", err)
		return
	}
	buff2, _ := ioutil.ReadAll(zipHtml)
	if !bytes.Equal(buff1, buff2) {
		t.Errorf("The files on zip and OpenFile are not equal")
		return
	}
}

func TestMetadata(t *testing.T) {
	f, _ := Open(book_path)
	defer f.Close()

	if title, _ := f.Metadata("title"); title[0] != book_title {
		t.Errorf("Metadata title '%v', the expected was '%v'", title[0], book_title)
	}
	if language, _ := f.Metadata("language"); language[0] != book_lang {
		t.Errorf("Metadata language '%v', the expected was '%v'", language[0], book_lang)
	}
	if identifier, _ := f.Metadata("identifier"); identifier[0] != book_identifier {
		t.Errorf("Metadata identifier '%v', the expected was '%v'", identifier[0], book_identifier)
	}
	if creator, _ := f.Metadata("creator"); creator[0] != book_creator {
		t.Errorf("Metadata creator '%v', the expected was '%v'", creator[0], book_creator)
	}
	if subject, _ := f.Metadata("subject"); subject[0] != book_subject {
		t.Errorf("Metadata subject '%v', the expected was '%v'", subject[0], book_subject)
	}
	if rights, _ := f.Metadata("rights"); rights[0] != book_rights {
		t.Errorf("Metadata rights '%v', the expected was '%v'", rights[0], book_rights)
	}
}

func TestMetadataFields(t *testing.T) {
	f, _ := Open(book_path)
	defer f.Close()

	fields := f.MetadataFields()
	if len(fields) != len_metadafields {
		t.Errorf("len(MetadataFields()) should be %v, but was %v", len_metadafields, len(fields))
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
	f, _ := Open(book_path)
	defer f.Close()

	if identifier, _ := f.MetadataAttr("identifier"); identifier[0]["scheme"] != identifier_scheme {
		t.Errorf("Metadata identifier attr scheme '%v', the expected was '%v'", identifier[0]["scheme"], identifier_scheme)
	}
	if creator, _ := f.MetadataAttr("creator"); creator[0]["file-as"] != creator_file_as {
		t.Errorf("Metadata creator attr file-as '%v', the expected was '%v'", creator[0]["file-as"], creator_file_as)
	}
	if meta, _ := f.MetadataAttr("meta"); meta[0]["name"] != meta_name {
		t.Errorf("Metadata meta attr name '%v', the expected was '%v'", meta[0]["name"], meta_name)
	}
}
