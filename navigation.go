// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

type NavigationIterator struct {
	depth  []navpoint
	navMap []navpoint
	index  int
}

func newNavigationIterator(navMap []navpoint) *NavigationIterator {
	var nav NavigationIterator
	nav.navMap = navMap
	nav.index = 0
	return &nav
}

func (nav NavigationIterator) Title() string {
	return nav.item().Text
}

func (nav NavigationIterator) HasChildren() bool {
	return nav.item().NavPoint != nil
}

func (nav NavigationIterator) item() *navpoint {
	return &nav.navMap[nav.index]
}
