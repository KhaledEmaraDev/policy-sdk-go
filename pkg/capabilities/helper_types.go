package capabilities

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

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
