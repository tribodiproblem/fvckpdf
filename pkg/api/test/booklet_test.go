package test

import (
	"path/filepath"
	"testing"

	"github.com/tribodiproblem/fvckpdf/pkg/api"
	"github.com/tribodiproblem/fvckpdf/pkg/pdfcpu"
)

func testBooklet(t *testing.T, msg string, inFiles []string, outFile string, selectedPages []string, desc string, n int, isImg bool) {
	t.Helper()

	var (
		booklet *pdfcpu.NUp
		err     error
	)

	if isImg {
		if booklet, err = api.ImageBookletConfig(n, desc); err != nil {
			t.Fatalf("%s %s: %v\n", msg, outFile, err)
		}
	} else {
		if booklet, err = api.PDFBookletConfig(n, desc); err != nil {
			t.Fatalf("%s %s: %v\n", msg, outFile, err)
		}
	}

	if err := api.BookletFile(inFiles, outFile, selectedPages, booklet, nil); err != nil {
		t.Fatalf("%s %s: %v\n", msg, outFile, err)
	}
	if err := api.ValidateFile(outFile, nil); err != nil {
		t.Fatalf("%s: %v\n", msg, err)
	}
}

func TestBooklet(t *testing.T) {
	outDir := "../../samples/booklet"

	for _, tt := range []struct {
		msg           string
		inFiles       []string
		outFile       string
		selectedPages []string
		desc          string
		n             int
		isImg         bool
	}{

		// 2-up booklet from images on A4
		{"TestBookletFromImagesA42Up",
			imageFileNames(t, resDir),
			filepath.Join(outDir, "BookletFromImagesA4_2Up.pdf"),
			nil,
			"p:A4, border:false, g:on, ma:25, bgcol:#beded9",
			2,
			true,
		},

		// 4-up booklet from images on A4
		{"TestBookletFromImagesA44Up",
			imageFileNames(t, resDir),
			filepath.Join(outDir, "BookletFromImagesA4_4Up.pdf"),
			nil,
			"p:A4, border:false, g:on, ma:25, bgcol:#beded9",
			4,
			true,
		},

		// 2-up booklet from PDF on A4
		{"TestBookletFromPDF2UpA4",
			[]string{filepath.Join(inDir, "zineTest.pdf")},
			filepath.Join(outDir, "BookletFromPDFA4_2Up.pdf"),
			nil, // all pages
			"p:A4, border:false, g:on, ma:10, bgcol:#beded9",
			2,
			false,
		},

		// 4-up booklet from PDF on A4
		{"TestBookletFromPDF4UpA4",
			[]string{filepath.Join(inDir, "zineTest.pdf")},
			filepath.Join(outDir, "BookletFromPDFA4_4Up.pdf"),
			[]string{"1-"}, // all pages
			"p:A4, border:off, guides:on, ma:10, bgcol:#beded9",
			4,
			false,
		},

		// 4-up booklet from PDF on Ledger
		{"TestBookletFromPDF4UpLedger",
			[]string{filepath.Join(inDir, "bookletTest.pdf")},
			filepath.Join(outDir, "BookletFromPDFLedger_4Up.pdf"),
			[]string{"1-24"},
			"p:LedgerP, g:on, ma:10, bgcol:#f7e6c7",
			4,
			false,
		},

		// 4-up booklet from PDF on Ledger where the number of pages don't fill the whole sheet
		{"TestBookletFromPDF4UpLedgerWithTrailingBlankPages",
			[]string{filepath.Join(inDir, "bookletTest.pdf")},
			filepath.Join(outDir, "BookletFromPDFLedger_4UpWithTrailingBlankPages.pdf"),
			[]string{"1-21"},
			"p:LedgerP, g:on, ma:10, bgcol:#f7e6c7",
			4,
			false,
		},

		// 2-up booklet from PDF on Letter
		{"TestBookletFromPDF2UpLetter",
			[]string{filepath.Join(inDir, "bookletTest.pdf")},
			filepath.Join(outDir, "BookletFromPDFLetter_2Up.pdf"),
			[]string{"1-16"},
			"p:LetterP, g:on, ma:10, bgcol:#f7e6c7",
			2,
			false,
		},

		// 2-up booklet from PDF on Letter where the number of pages don't fill the whole sheet
		{"TestBookletFromPDF2UpLetterWithTrailingBlankPages",
			[]string{filepath.Join(inDir, "bookletTest.pdf")},
			filepath.Join(outDir, "BookletFromPDFLetter_2UpWithTrailingBlankPages.pdf"),
			[]string{"1-14"},
			"p:LetterP, g:on, ma:10, bgcol:#f7e6c7",
			2,
			false,
		},
	} {
		testBooklet(t, tt.msg, tt.inFiles, tt.outFile, tt.selectedPages, tt.desc, tt.n, tt.isImg)
	}
}
