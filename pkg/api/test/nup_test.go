/*
Copyright 2020 The pdf Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package test

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tribodiproblem/fvckpdf/pkg/api"
	"github.com/tribodiproblem/fvckpdf/pkg/pdfcpu"
)

func testNUp(t *testing.T, msg string, inFiles []string, outFile string, selectedPages []string, desc string, n int, isImg bool) {
	t.Helper()

	var (
		nup *pdfcpu.NUp
		err error
	)

	if isImg {
		if nup, err = api.ImageNUpConfig(n, desc); err != nil {
			t.Fatalf("%s %s: %v\n", msg, outFile, err)
		}
	} else {
		if nup, err = api.PDFNUpConfig(n, desc); err != nil {
			t.Fatalf("%s %s: %v\n", msg, outFile, err)
		}
	}

	if err := api.NUpFile(inFiles, outFile, selectedPages, nup, nil); err != nil {
		t.Fatalf("%s %s: %v\n", msg, outFile, err)
	}
	if err := api.ValidateFile(outFile, nil); err != nil {
		t.Fatalf("%s: %v\n", msg, err)
	}
}

func imageFileNames(t *testing.T, dir string) []string {
	t.Helper()
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	fn := []string{}
	for _, fi := range files {
		if strings.HasSuffix(fi.Name(), "png") || strings.HasSuffix(fi.Name(), "jpg") {
			fn = append(fn, filepath.Join(dir, fi.Name()))
		}
	}
	return fn
}

func TestNUp(t *testing.T) {
	outDir := "../../samples/nup"

	for _, tt := range []struct {
		msg           string
		inFiles       []string
		outFile       string
		selectedPages []string
		desc          string
		n             int
		isImg         bool
	}{
		// 4-Up a PDF
		{"TestNUpFromPDF",
			[]string{filepath.Join(inDir, "WaldenFull.pdf")},
			filepath.Join(outDir, "NUpFromPDF.pdf"),
			nil,
			"m:10, bgcol:#f7e6c7",
			9,
			false},

		// 2-Up a PDF with CropBox
		{"TestNUpFromPdfWithCropBox",
			[]string{filepath.Join(inDir, "grid_example.pdf")},
			filepath.Join(outDir, "NUpFromPDFWithCropBox.pdf"),
			nil,
			"f:A5L, border:on, m:10, bgcol:#f7e6c7",
			2,
			false},

		// 16-Up an image
		{"TestNUpFromSingleImage",
			[]string{filepath.Join("../../../resources", "logoSmall.png")},
			filepath.Join(outDir, "NUpFromSingleImage.pdf"),
			nil,
			"f:A3P, m:10, bgcol:#f7e6c7",
			16,
			true},

		// 6-Up a sequence of images.
		{"TestNUpFromImages",
			imageFileNames(t, "../../../resources"),
			filepath.Join(outDir, "NUpFromImages.pdf"),
			nil,
			"f:Tabloid, border:on, m:10, bgcol:#f7e6c7",
			6,
			true},
	} {
		testNUp(t, tt.msg, tt.inFiles, tt.outFile, tt.selectedPages, tt.desc, tt.n, tt.isImg)
	}
}
