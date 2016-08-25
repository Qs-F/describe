// Copyright 2016 de-liKeR. All right reserved.
// This package is released under the MIT license.

// Package for executable file for describe service. See also github.com/Qs-F/describe/
package main

import (
	"io/ioutil"
	"os"
)

// unique constants for this package
const (
	fileName_        = ".describe" // specificed file name
	unexpectedErrMsg = "An unexpected error occured. Please send feedback at `https://github.com/Qs-F/describe/issues`"
)

// words managing flags
const (
	unexpectedErr_    = iota // unexpected flag or value
	init_                    // init on this directory(create .describe file in user working directory)
	showDescription_         // show description(show this directory's description)
	showDescriptions_        // show descriptions in tree form(show all inner directories' description)
)

/* Tree form A type
.
|-- ABC  // Hello This is ABC Direcotry
|   |
|   |-- Inner  // Hello. Is this readability good?
|   |          // Console Design is Difficult...
|   |
|   |--EFG  // This
|   |       // is
|   |       // Note
|   |       // Trying
|   |       // Description!
|   |
|   |-- HIJ
|   |   |-- KLM
|   |   `-- NOP
|   `-- QRS
`-- TUV
*/

/* Tree form B type
./ABC  // Hello. Is this cool?
       // Not so much...

./ABC/Inner  // Wow This is awesome...

./ABC/EFG  // This
           // is
					 // Note
					 // Trying
					 // Description!
*/

/* Tree form C type
./ABC
   Hello. Is this cool?
   Not so much...

./ABC/Inner
   Wow This is awesome...

./ABC/Inner/startup
   Wow This is awesome...

./ABC/Inner/DONOTUSE
   Wow This is awesome...
   But this is unavailable directory

./ABC/Inner
   Wow This is awesome...

./ABC/Inner
   Wow This is awesome...

./ABC/Inner
   Wow This is awesome...
*/

// public value
var (
	currentDir string // save current user working dir.
	workingDir string // move system working directory.
)

// Description struct
type Description struct {
	Path    string // from current user working dir
	Info    *os.FileInfo
	Content string
}

// Descriptions struct
type Descriptions []Description

// IsDescribed returns whether is there a written .describe file on the arg directory in bool type.
func (d *Description) IsDescribed(path string) bool {
	r := false
	return r
}

// HasDir returns true when given path is directories. pls use FileInfo.IsDir directory instead.
// func (d *Description) IsDir(path string) bool {
// 	return d.Info.IsDir
// }

// ListDir returns slice of given path's inner directories list.
func ListDir(path string) ([]string, error) {
	r := []string{}
	objs, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}, err
	}
	for _, obj := range objs {
		if obj.IsDir() == true {
			r = append(r, "/"+path+"/"+obj.Name())
		} else if obj.IsDir() == false && obj.Name() == fileName { // This is described directory
		}
	}
	return r, nil
}

func Init(basePath string) (d *Description) {}

// ShowDescription shows arg path's .describe file content.
func ShowDescription(path string) {}

// Init creates a .describe file in working directory
func InitService(path string) {}

// main
func main() {}
