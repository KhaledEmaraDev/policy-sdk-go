package capabilities

import (
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

type SigstorePubKeyVerifyType struct{}

func (_ SigstorePubKeyVerifyType) MarshalEasyJSON(w *jwriter.Writer) {
	w.String("SigstorePubKeyVerify")
}

type SigstoreKeylessVerifyType struct{}

func (_ SigstoreKeylessVerifyType) MarshalEasyJSON(w *jwriter.Writer) {
	w.String("SigstoreKeylessVerify")
}

type SigstoreKeylessPrefixVerifyType struct{}

func (_ SigstoreKeylessPrefixVerifyType) MarshalEasyJSON(w *jwriter.Writer) {
	w.String("SigstoreKeylessPrefixVerify")
}

type SigstoreGithubActionsVerifyType struct{}

func (_ SigstoreGithubActionsVerifyType) MarshalEasyJSON(w *jwriter.Writer) {
	w.String("SigstoreGithubActionsVerify")
}

type SigstoreCertificateVerifyType struct{}

func (_ SigstoreCertificateVerifyType) MarshalEasyJSON(w *jwriter.Writer) {
	w.String("SigstoreCertificateVerify")
}

type RuneString string

func (rs *RuneString) MarshalEasyJSON(w *jwriter.Writer) {
	runes := []rune(*rs)

	if runes == nil && (w.Flags&jwriter.NilSliceAsEmpty) == 0 {
		w.RawString("null")
	} else {
		w.RawByte('[')
		for v46, v47 := range runes {
			if v46 > 0 {
				w.RawByte(',')
			}
			w.Int32(int32(v47))
		}
		w.RawByte(']')
	}
}

func (rs *RuneString) UnmarshalEasyJSON(l *jlexer.Lexer) {
	var runes []rune

	if l.IsNull() {
		l.Skip()
		runes = nil
	} else {
		l.Delim('[')
		if runes == nil {
			if !l.IsDelim(']') {
				runes = make([]int32, 0, 16)
			} else {
				runes = []int32{}
			}
		} else {
			runes = (runes)[:0]
		}
		for !l.IsDelim(']') {
			var v45 int32 = int32(l.Int32())
			runes = append(runes, v45)
			l.WantComma()
		}
		l.Delim(']')
	}

	*rs = RuneString(runes)
}

// The encoding of the certificate
type CertificateEncoding int

const (
	Der CertificateEncoding = iota + 1
	Pem
)

func (e CertificateEncoding) MarshalEasyJSON(w *jwriter.Writer) {
	if e == Der {
		w.String("Der")
	} else if e == Pem {
		w.String("Pem")
	}
}
