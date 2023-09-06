// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validate

import (
	"strings"
	"testing"

	pb "github.com/google/go-tdx-guest/proto/tdx"
	vr "github.com/google/go-tdx-guest/verify"

	"github.com/google/go-tdx-guest/abi"
	"github.com/google/go-tdx-guest/testing/testdata"
)

func convert(a []byte, x byte) []byte {
	for i := range a {
		a[i] = x
	}
	return a
}

func TestTdxQuote(t *testing.T) {
	if err := TdxQuote(nil, nil); err != vr.ErrOptionsNil {
		t.Error(err)
	}
	qeSvn := uint16(0)
	pceSvn := uint16(0)
	qeVendorID := []byte{0x93, 0x9a, 0x72, 0x33, 0xf7, 0x9c, 0x4c, 0xa9, 0x94, 0xa, 0xd, 0xb3, 0x95, 0x7f, 0x6, 0x7}
	teeTcbSvn := []byte{0x3, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	mrSeam := []byte{0x2f, 0xd2, 0x79, 0xc1, 0x61, 0x64, 0xa9, 0x3d, 0xd5, 0xbf, 0x37, 0x3d, 0x83, 0x43, 0x28, 0xd4,
		0x60, 0x8, 0xc2, 0xb6, 0x93, 0xaf, 0x9e, 0xbb, 0x86, 0x5b, 0x8, 0xb2, 0xce, 0xd3, 0x20, 0xc9,
		0xa8, 0x9b, 0x48, 0x69, 0xa9, 0xfa, 0xb6, 0xf, 0xbe, 0x9d, 0xc, 0x5a, 0x53, 0x63, 0xc6, 0x56}
	tdAttributes := []byte{0x0, 0x0, 0x0, 0x40, 0x0, 0x0, 0x0, 0x0}
	xfam := []byte{0xe7, 0x1a, 0x6, 0x0, 0x0, 0x0, 0x0, 0x0}
	mrTd := []byte{0x63, 0x63, 0xb8, 0x4, 0x36, 0x68, 0xa3, 0xad, 0x95, 0x32, 0x78, 0xe1, 0x3, 0x89, 0x57, 0x4d,
		0x32, 0x6c, 0x67, 0x49, 0xfb, 0x78, 0xaa, 0x81, 0xe, 0xcd, 0x93, 0x36, 0x92, 0x3d, 0xb8, 0x6f,
		0x22, 0xfc, 0x0, 0xb8, 0xdc, 0xd4, 0x4, 0xbc, 0x10, 0xd5, 0xe1, 0x19, 0xd7, 0x21, 0x5c, 0xbb}
	mrConfigID := []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	mrOwner := []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	mrOwnerConfig := []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	rtmr0 := []byte{0x29, 0x27, 0xda, 0x70, 0x46, 0x1c, 0xd6, 0x32, 0x66, 0xf4, 0x32, 0x30, 0xcc, 0x18, 0x49, 0xc0,
		0x3e, 0xf2, 0x5e, 0xbe, 0x49, 0x0, 0x62, 0xa8, 0x1, 0xd8, 0xfc, 0xc8, 0xa, 0xf4, 0x29, 0x76,
		0x82, 0x3a, 0xdf, 0x8, 0xf8, 0x33, 0xc1, 0xe5, 0xb, 0x51, 0x77, 0x9c, 0x65, 0x93, 0xf3, 0x2a}
	rtmr1 := []byte{0x2c, 0x70, 0xb, 0x8b, 0xa9, 0xb8, 0x57, 0x83, 0xf8, 0xbe, 0x9f, 0xb9, 0x44, 0x36, 0x47, 0xbd,
		0xc0, 0xbb, 0x3c, 0x50, 0x74, 0x7f, 0x6, 0x29, 0x7c, 0xc6, 0x53, 0x8c, 0x25, 0xa5, 0xf5, 0x89,
		0xc4, 0xb5, 0x6d, 0x3, 0x5c, 0x59, 0x10, 0x7c, 0x6b, 0xc5, 0x80, 0xd, 0xb2, 0xca, 0xcb, 0x61}
	rtmr2 := []byte{0x86, 0x52, 0xf0, 0xca, 0xab, 0xa7, 0xe2, 0x15, 0xea, 0x44, 0x2d, 0xc3, 0x6a, 0x44, 0x99, 0xd8,
		0xfe, 0xc3, 0x36, 0x2f, 0x3a, 0xb, 0x2c, 0xa1, 0x51, 0xcb, 0xe4, 0xb3, 0xe6, 0x46, 0x6f, 0xe5,
		0x9c, 0x73, 0x68, 0xb3, 0xc2, 0x28, 0x7f, 0xc7, 0xc3, 0xbf, 0x5c, 0x92, 0x4e, 0xb4, 0x42, 0x4e}
	rtmr3 := []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	reportData := []byte{0x6c, 0x62, 0xde, 0xc1, 0xb8, 0x19, 0x17, 0x49, 0xa3, 0x1d, 0xab, 0x49, 0xb, 0xe5, 0x32, 0xa3,
		0x59, 0x44, 0xde, 0xa4, 0x7c, 0xae, 0xf1, 0xf9, 0x80, 0x86, 0x39, 0x93, 0xd9, 0x89, 0x95, 0x45,
		0xeb, 0x74, 0x6, 0xa3, 0x8d, 0x1e, 0xed, 0x31, 0x3b, 0x98, 0x7a, 0x46, 0x7d, 0xac, 0xea, 0xd6,
		0xf0, 0xc8, 0x7a, 0x6d, 0x76, 0x6c, 0x66, 0xf6, 0xf2, 0x9f, 0x8a, 0xcb, 0x28, 0x1f, 0x11, 0x13}

	mknonce := func(front []byte) []byte {
		result := make([]byte, 64)
		copy(result[:], front)
		return result
	}

	make2darray := func(size int) [][]byte {
		a := make([][]byte, 4)
		for i := range a {
			a[i] = make([]byte, size)
		}
		return a
	}

	nonce12345 := mknonce([]byte{1, 2, 3, 4, 5})

	quoteFn := func(nonce []byte) *pb.QuoteV4 {
		quote, err := abi.QuoteToProto(testdata.RawQuote)
		if err != nil {
			t.Fatal(err)
		}
		data := make([]byte, abi.ReportDataSize)
		copy(data, nonce[:])
		quote.TdQuoteBody.ReportData = data

		return quote
	}
	quoteSample := quoteFn(reportData)
	quote12345 := quoteFn(nonce12345)

	type testCase struct {
		name    string
		quote   *pb.QuoteV4
		opts    *Options
		wantErr string
	}
	tests := []testCase{
		{
			name:  "deep check",
			quote: quoteSample,
			opts: &Options{
				HeaderOptions: HeaderOptions{
					MinimumQeSvn:  qeSvn,
					MinimumPceSvn: pceSvn,
					QeVendorID:    qeVendorID,
				},
				TdQuoteBodyOptions: TdQuoteBodyOptions{
					MinimumTeeTcbSvn: teeTcbSvn,
					MrSeam:           mrSeam,
					TdAttributes:     tdAttributes,
					Xfam:             xfam,
					MrTd:             mrTd,
					MrConfigID:       mrConfigID,
					MrOwner:          mrOwner,
					MrOwnerConfig:    mrOwnerConfig,
					Rtmrs:            [][]byte{rtmr0, rtmr1, rtmr2, rtmr3},
					ReportData:       reportData,
				},
			},
		},
		{
			name:  "min QE security-version check",
			quote: quote12345,
			opts: &Options{
				HeaderOptions: HeaderOptions{
					MinimumQeSvn: uint16(2),
				},
			},
			wantErr: "QE security-version number 0 is less than the required minimum 2",
		},
		{
			name:  "min Pce security-version check",
			quote: quote12345,
			opts: &Options{
				HeaderOptions: HeaderOptions{
					MinimumPceSvn: uint16(2),
				},
			},
			wantErr: "PCE security-version number 0 is less than the required minimum 2",
		},
		{
			name:  "min TEE TCB security-version check",
			quote: quote12345,
			opts: &Options{
				TdQuoteBodyOptions: TdQuoteBodyOptions{
					MinimumTeeTcbSvn: []byte{0x4, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
				},
			},
			wantErr: "TEE TCB security-version number [3 0 4 0 0 0 0 0 0 0 0 0 0 0 0 0] is less than the required minimum [4 0 4 0 0 0 0 0 0 0 0 0 0 0 0 0]",
		},
		{
			name:  "Test incorrect MR_SEAM",
			quote: quote12345,
			opts: &Options{
				TdQuoteBodyOptions: TdQuoteBodyOptions{
					MrSeam: make([]byte, abi.MrSeamSize),
				},
			},
			wantErr: "quote field MR_SEAM",
		},
		{
			name:  "Test incorrect TD_ATTRIBUTES",
			quote: quote12345,
			opts: &Options{
				TdQuoteBodyOptions: TdQuoteBodyOptions{
					TdAttributes: make([]byte, abi.TdAttributesSize),
				},
			},
			wantErr: "quote field TD_ATTRIBUTES",
		},
		{
			name:  "Test incorrect XFAM",
			quote: quote12345,
			opts: &Options{
				TdQuoteBodyOptions: TdQuoteBodyOptions{
					Xfam: make([]byte, abi.XfamSize),
				},
			},
			wantErr: "quote field XFAM",
		},
		{
			name:  "Test incorrect MR_TD",
			quote: quote12345,
			opts: &Options{
				TdQuoteBodyOptions: TdQuoteBodyOptions{
					MrTd: make([]byte, abi.MrTdSize),
				},
			},
			wantErr: "quote field MR_TD",
		},
		{
			name:  "Test incorrect MR_CONFIG_ID",
			quote: quote12345,
			opts: &Options{
				TdQuoteBodyOptions: TdQuoteBodyOptions{
					MrConfigID: convert(make([]byte, abi.MrConfigIDSize), 1),
				},
			},
			wantErr: "quote field MR_CONFIG_ID",
		},
		{
			name:  "Test incorrect MR_OWNER",
			quote: quote12345,
			opts: &Options{
				TdQuoteBodyOptions: TdQuoteBodyOptions{
					MrOwner: convert(make([]byte, abi.MrOwnerSize), 1),
				},
			},
			wantErr: "quote field MR_OWNER",
		},
		{
			name:  "Test incorrect MR_OWNER_CONFIG",
			quote: quote12345,
			opts: &Options{
				TdQuoteBodyOptions: TdQuoteBodyOptions{
					MrOwnerConfig: convert(make([]byte, abi.MrOwnerConfigSize), 1),
				},
			},
			wantErr: "quote field MR_OWNER_CONFIG",
		},
		{name: "Test incorrect RTMRS",
			quote: quote12345,
			opts: &Options{TdQuoteBodyOptions: TdQuoteBodyOptions{
				Rtmrs: make2darray(abi.RtmrSize),
			},
			},
			wantErr: "quote field RTMRS",
		},
		{
			name:  "Test incorrect REPORT_DATA",
			quote: quote12345,
			opts: &Options{
				TdQuoteBodyOptions: TdQuoteBodyOptions{
					ReportData: make([]byte, abi.ReportDataSize),
				},
			},
			wantErr: "quote field REPORT_DATA",
		},
		{
			name:  "Test incorrect QE_VENDOR_ID",
			quote: quote12345,
			opts: &Options{
				HeaderOptions: HeaderOptions{QeVendorID: make([]byte, abi.QeVendorIDSize)},
			},
			wantErr: "quote field QE_VENDOR_ID"},
	}

	for _, tc := range tests {
		if err := TdxQuote(tc.quote, tc.opts); (err == nil && tc.wantErr != "") ||
			(err != nil && (tc.wantErr == "" || !strings.Contains(err.Error(), tc.wantErr))) {
			t.Errorf("%s: TdxQuote() errored unexpectedly. Got '%v', want '%s'", tc.name, err, tc.wantErr)
		}
	}
}
