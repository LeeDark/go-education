// Copyright 2013 HORISEN AG. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cdbeasy

import (
	"container/list"
	"io"

	"github.com/jbarham/cdb"
)

// CdbEasy provides wrapper functions to CDB library.
type CdbEasy struct {
	FileName string
	Cdbh     *cdb.Cdb
}

// Open opens CDB file.
func Open(filename string) (*CdbEasy, error) {
	cdbh, err := cdb.Open(filename)
	if err != nil {
		return nil, err
	}
	cdbe := new(CdbEasy)
	cdbe.FileName = filename
	cdbe.Cdbh = cdbh
	return cdbe, nil
}

// FindOne finds given key if exists. If not returns empty string.
func (cdbe *CdbEasy) FindOne(key string) (string, error) {
	keyBytes := []byte(key)
	valueBytes, err := cdbe.Cdbh.Data(keyBytes)
	if err != nil && err != io.EOF {
		return "", err
	}
	if valueBytes == nil || err == io.EOF {
		return "", nil
	}
	value := string(valueBytes)
	return value, nil
}

// FindMultiple finds given key and returns list of values found (may be multiple).
func (cdbe *CdbEasy) FindMultiple(key string) (*list.List, error) {
	values := new(list.List)
	keyBytes := []byte(key)
	cdbe.Cdbh.FindStart()
	for {
		valueBytesSr, err := cdbe.Cdbh.FindNext(keyBytes)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if err == io.EOF || valueBytesSr.Size() == 0 {
			break
		}
		//fmt.Printf("Size=%d\n", valueBytes_sr.Size())
		var valueBytesWhole [1024]byte
		n, err := valueBytesSr.Read(valueBytesWhole[:])
		if err == io.EOF {
			// I believe this is bug in CDB library. FindNext returns err:= null but
			// valueBytes_sr is empty so actually we get io.EOF while reading it
			break
		}
		if err != nil {
			return nil, err
		}
		valueBytes := valueBytesWhole[0:n]
		value := string(valueBytes)
		values.PushBack(value)
	}
	return values, nil
}

// Close closes CDB wrapper
func (cdbe *CdbEasy) Close() error {
	return cdbe.Cdbh.Close()
}

// FindOne finds given key in cdb library in given file.
// If no key found returns empty string
func FindOne(filename string, key string) (string, error) {
	cdbe, err := Open(filename)
	if err != nil {
		return "", err
	}
	defer cdbe.Close()
	return cdbe.FindOne(key)
}

// FindMultiple finds given key in cdb library in given file,
// and returns list of values found (may be multiple).
func FindMultiple(filename string, key string) (*list.List, error) {
	cdbe, err := Open(filename)
	if err != nil {
		return nil, err
	}
	defer cdbe.Close()
	return cdbe.FindMultiple(key)
}
