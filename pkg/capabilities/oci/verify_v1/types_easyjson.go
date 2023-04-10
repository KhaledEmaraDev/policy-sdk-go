// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package verify_v1

import (
	json "encoding/json"
	oci "github.com/kubewarden/policy-sdk-go/pkg/capabilities/oci"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV1(in *jlexer.Lexer, out *sigstorePubKeysVerifyRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "SigstorePubKeyVerify":
			(out.SigstorePubKeysVerify).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV1(out *jwriter.Writer, in sigstorePubKeysVerifyRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"SigstorePubKeyVerify\":"
		out.RawString(prefix[1:])
		(in.SigstorePubKeysVerify).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v sigstorePubKeysVerifyRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v sigstorePubKeysVerifyRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *sigstorePubKeysVerifyRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *sigstorePubKeysVerifyRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV1(l, v)
}
func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV11(in *jlexer.Lexer, out *sigstorePubKeysVerify) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "image":
			out.Image = string(in.String())
		case "pub_keys":
			if in.IsNull() {
				in.Skip()
				out.PubKeys = nil
			} else {
				in.Delim('[')
				if out.PubKeys == nil {
					if !in.IsDelim(']') {
						out.PubKeys = make([]string, 0, 4)
					} else {
						out.PubKeys = []string{}
					}
				} else {
					out.PubKeys = (out.PubKeys)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.PubKeys = append(out.PubKeys, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "annotations":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Annotations = make(map[string]string)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 string
					v2 = string(in.String())
					(out.Annotations)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV11(out *jwriter.Writer, in sigstorePubKeysVerify) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"image\":"
		out.RawString(prefix[1:])
		out.String(string(in.Image))
	}
	{
		const prefix string = ",\"pub_keys\":"
		out.RawString(prefix)
		if in.PubKeys == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.PubKeys {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"annotations\":"
		out.RawString(prefix)
		if in.Annotations == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v5First := true
			for v5Name, v5Value := range in.Annotations {
				if v5First {
					v5First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v5Name))
				out.RawByte(':')
				out.String(string(v5Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v sigstorePubKeysVerify) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v sigstorePubKeysVerify) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *sigstorePubKeysVerify) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *sigstorePubKeysVerify) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV11(l, v)
}
func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV12(in *jlexer.Lexer, out *sigstoreKeylessVerifyRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "SigstoreKeylessVerify":
			(out.SigstoreKeylessVerify).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV12(out *jwriter.Writer, in sigstoreKeylessVerifyRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"SigstoreKeylessVerify\":"
		out.RawString(prefix[1:])
		(in.SigstoreKeylessVerify).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v sigstoreKeylessVerifyRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v sigstoreKeylessVerifyRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *sigstoreKeylessVerifyRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *sigstoreKeylessVerifyRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV12(l, v)
}
func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV13(in *jlexer.Lexer, out *sigstoreKeylessVerify) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "image":
			out.Image = string(in.String())
		case "keyless":
			if in.IsNull() {
				in.Skip()
				out.Keyless = nil
			} else {
				in.Delim('[')
				if out.Keyless == nil {
					if !in.IsDelim(']') {
						out.Keyless = make([]oci.KeylessInfo, 0, 2)
					} else {
						out.Keyless = []oci.KeylessInfo{}
					}
				} else {
					out.Keyless = (out.Keyless)[:0]
				}
				for !in.IsDelim(']') {
					var v6 oci.KeylessInfo
					(v6).UnmarshalEasyJSON(in)
					out.Keyless = append(out.Keyless, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "annotations":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Annotations = make(map[string]string)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v7 string
					v7 = string(in.String())
					(out.Annotations)[key] = v7
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV13(out *jwriter.Writer, in sigstoreKeylessVerify) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"image\":"
		out.RawString(prefix[1:])
		out.String(string(in.Image))
	}
	{
		const prefix string = ",\"keyless\":"
		out.RawString(prefix)
		if in.Keyless == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Keyless {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"annotations\":"
		out.RawString(prefix)
		if in.Annotations == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v10First := true
			for v10Name, v10Value := range in.Annotations {
				if v10First {
					v10First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v10Name))
				out.RawByte(':')
				out.String(string(v10Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v sigstoreKeylessVerify) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v sigstoreKeylessVerify) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *sigstoreKeylessVerify) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *sigstoreKeylessVerify) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciVerifyV13(l, v)
}
