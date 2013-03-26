// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import (
	"errors"
)

type NavigationIterator struct {
	parents []navCursor
	curr    navCursor
}
type navCursor struct {
	navMap []navpoint
	index  int
}

func newNavigationIterator(navMap []navpoint) *NavigationIterator {
	var nav NavigationIterator
	nav.curr.navMap = navMap
	nav.curr.index = 0
	return &nav
}

func (nav NavigationIterator) Title() string {
	return nav.item().Title()
}

func (nav NavigationIterator) Url() string {
	return nav.item().Url()
}

func (nav NavigationIterator) HasChildren() bool {
	return nav.item().Children() != nil
}

func (nav NavigationIterator) HasParents() bool {
	return nav.parents != nil
}

func (nav NavigationIterator) IsFirst() bool {
	return nav.curr.index == 0
}

func (nav NavigationIterator) IsLast() bool {
	return nav.curr.index == len(nav.curr.navMap)-1
}

func (nav *NavigationIterator) Next() error {
	if nav.IsLast() {
		return errors.New("It is the last entry")
	}
	nav.curr.index++
	return nil
}

func (nav *NavigationIterator) Previous() error {
	if nav.IsFirst() {
		return errors.New("It is the first entry")
	}
	nav.curr.index--
	return nil
}

func (nav *NavigationIterator) In() error {
	if !nav.HasChildren() {
		return errors.New("It has no children")
	}
	nav.parents = append(nav.parents, nav.curr)
	nav.curr.navMap = nav.item().Children()
	nav.curr.index = 0
	return nil
}

func (nav *NavigationIterator) Out() error {
	if !nav.HasParents() {
		return errors.New("It has no parents")
	}
	nav.curr = nav.parents[len(nav.parents)-1]
	nav.parents = nav.parents[:len(nav.parents)-1]
	return nil
}

func (nav NavigationIterator) item() *navpoint {
	return &nav.curr.navMap[nav.curr.index]
}
