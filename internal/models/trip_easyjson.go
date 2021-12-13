// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
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

func easyjson30497fe1DecodeSnakealiveMInternalModels(in *jlexer.Lexer, out *UserInfo) {
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
		case "id":
			out.Id = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "surname":
			out.Surname = string(in.String())
		case "avatar":
			out.Avatar = string(in.String())
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
func easyjson30497fe1EncodeSnakealiveMInternalModels(out *jwriter.Writer, in UserInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"surname\":"
		out.RawString(prefix)
		out.String(string(in.Surname))
	}
	{
		const prefix string = ",\"avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson30497fe1EncodeSnakealiveMInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson30497fe1EncodeSnakealiveMInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson30497fe1DecodeSnakealiveMInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson30497fe1DecodeSnakealiveMInternalModels(l, v)
}
func easyjson30497fe1DecodeSnakealiveMInternalModels1(in *jlexer.Lexer, out *TripWithUserInfo) {
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
		case "id":
			out.Id = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "sights":
			if in.IsNull() {
				in.Skip()
				out.Sights = nil
			} else {
				in.Delim('[')
				if out.Sights == nil {
					if !in.IsDelim(']') {
						out.Sights = make([]Place, 0, 0)
					} else {
						out.Sights = []Place{}
					}
				} else {
					out.Sights = (out.Sights)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Place
					(v1).UnmarshalEasyJSON(in)
					out.Sights = append(out.Sights, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "albums":
			if in.IsNull() {
				in.Skip()
				out.Albums = nil
			} else {
				in.Delim('[')
				if out.Albums == nil {
					if !in.IsDelim(']') {
						out.Albums = make([]Album, 0, 0)
					} else {
						out.Albums = []Album{}
					}
				} else {
					out.Albums = (out.Albums)[:0]
				}
				for !in.IsDelim(']') {
					var v2 Album
					(v2).UnmarshalEasyJSON(in)
					out.Albums = append(out.Albums, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]UserInfo, 0, 1)
					} else {
						out.Users = []UserInfo{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v3 UserInfo
					(v3).UnmarshalEasyJSON(in)
					out.Users = append(out.Users, v3)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson30497fe1EncodeSnakealiveMInternalModels1(out *jwriter.Writer, in TripWithUserInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"sights\":"
		out.RawString(prefix)
		if in.Sights == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.Sights {
				if v4 > 0 {
					out.RawByte(',')
				}
				(v5).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"albums\":"
		out.RawString(prefix)
		if in.Albums == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Albums {
				if v6 > 0 {
					out.RawByte(',')
				}
				(v7).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"users\":"
		out.RawString(prefix)
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Users {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TripWithUserInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson30497fe1EncodeSnakealiveMInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TripWithUserInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson30497fe1EncodeSnakealiveMInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TripWithUserInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson30497fe1DecodeSnakealiveMInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TripWithUserInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson30497fe1DecodeSnakealiveMInternalModels1(l, v)
}
func easyjson30497fe1DecodeSnakealiveMInternalModels2(in *jlexer.Lexer, out *TripUser) {
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
		case "email":
			out.Email = string(in.String())
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
func easyjson30497fe1EncodeSnakealiveMInternalModels2(out *jwriter.Writer, in TripUser) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix[1:])
		out.String(string(in.Email))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TripUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson30497fe1EncodeSnakealiveMInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TripUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson30497fe1EncodeSnakealiveMInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TripUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson30497fe1DecodeSnakealiveMInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TripUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson30497fe1DecodeSnakealiveMInternalModels2(l, v)
}
func easyjson30497fe1DecodeSnakealiveMInternalModels3(in *jlexer.Lexer, out *TripSight) {
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
		case "id":
			out.Id = int(in.Int())
		case "lng":
			out.Lng = float32(in.Float32())
		case "lat":
			out.Lat = float32(in.Float32())
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
func easyjson30497fe1EncodeSnakealiveMInternalModels3(out *jwriter.Writer, in TripSight) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"lng\":"
		out.RawString(prefix)
		out.Float32(float32(in.Lng))
	}
	{
		const prefix string = ",\"lat\":"
		out.RawString(prefix)
		out.Float32(float32(in.Lat))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TripSight) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson30497fe1EncodeSnakealiveMInternalModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TripSight) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson30497fe1EncodeSnakealiveMInternalModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TripSight) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson30497fe1DecodeSnakealiveMInternalModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TripSight) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson30497fe1DecodeSnakealiveMInternalModels3(l, v)
}
func easyjson30497fe1DecodeSnakealiveMInternalModels4(in *jlexer.Lexer, out *Trip) {
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
		case "id":
			out.Id = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "sights":
			if in.IsNull() {
				in.Skip()
				out.Sights = nil
			} else {
				in.Delim('[')
				if out.Sights == nil {
					if !in.IsDelim(']') {
						out.Sights = make([]Place, 0, 0)
					} else {
						out.Sights = []Place{}
					}
				} else {
					out.Sights = (out.Sights)[:0]
				}
				for !in.IsDelim(']') {
					var v10 Place
					(v10).UnmarshalEasyJSON(in)
					out.Sights = append(out.Sights, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "albums":
			if in.IsNull() {
				in.Skip()
				out.Albums = nil
			} else {
				in.Delim('[')
				if out.Albums == nil {
					if !in.IsDelim(']') {
						out.Albums = make([]Album, 0, 0)
					} else {
						out.Albums = []Album{}
					}
				} else {
					out.Albums = (out.Albums)[:0]
				}
				for !in.IsDelim(']') {
					var v11 Album
					(v11).UnmarshalEasyJSON(in)
					out.Albums = append(out.Albums, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]int, 0, 8)
					} else {
						out.Users = []int{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v12 int
					v12 = int(in.Int())
					out.Users = append(out.Users, v12)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson30497fe1EncodeSnakealiveMInternalModels4(out *jwriter.Writer, in Trip) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"sights\":"
		out.RawString(prefix)
		if in.Sights == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v13, v14 := range in.Sights {
				if v13 > 0 {
					out.RawByte(',')
				}
				(v14).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"albums\":"
		out.RawString(prefix)
		if in.Albums == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v15, v16 := range in.Albums {
				if v15 > 0 {
					out.RawByte(',')
				}
				(v16).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"users\":"
		out.RawString(prefix)
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v17, v18 := range in.Users {
				if v17 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v18))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Trip) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson30497fe1EncodeSnakealiveMInternalModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Trip) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson30497fe1EncodeSnakealiveMInternalModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Trip) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson30497fe1DecodeSnakealiveMInternalModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Trip) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson30497fe1DecodeSnakealiveMInternalModels4(l, v)
}
func easyjson30497fe1DecodeSnakealiveMInternalModels5(in *jlexer.Lexer, out *Place) {
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
		case "id":
			out.Id = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "country":
			out.Country = string(in.String())
		case "rating":
			out.Rating = float32(in.Float32())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]string, 0, 4)
					} else {
						out.Tags = []string{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v19 string
					v19 = string(in.String())
					out.Tags = append(out.Tags, v19)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "description":
			out.Description = string(in.String())
		case "photos":
			if in.IsNull() {
				in.Skip()
				out.Photos = nil
			} else {
				in.Delim('[')
				if out.Photos == nil {
					if !in.IsDelim(']') {
						out.Photos = make([]string, 0, 4)
					} else {
						out.Photos = []string{}
					}
				} else {
					out.Photos = (out.Photos)[:0]
				}
				for !in.IsDelim(']') {
					var v20 string
					v20 = string(in.String())
					out.Photos = append(out.Photos, v20)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "day":
			out.Day = int(in.Int())
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
func easyjson30497fe1EncodeSnakealiveMInternalModels5(out *jwriter.Writer, in Place) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"country\":"
		out.RawString(prefix)
		out.String(string(in.Country))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.Float32(float32(in.Rating))
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v21, v22 := range in.Tags {
				if v21 > 0 {
					out.RawByte(',')
				}
				out.String(string(v22))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"photos\":"
		out.RawString(prefix)
		if in.Photos == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v23, v24 := range in.Photos {
				if v23 > 0 {
					out.RawByte(',')
				}
				out.String(string(v24))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"day\":"
		out.RawString(prefix)
		out.Int(int(in.Day))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Place) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson30497fe1EncodeSnakealiveMInternalModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Place) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson30497fe1EncodeSnakealiveMInternalModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Place) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson30497fe1DecodeSnakealiveMInternalModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Place) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson30497fe1DecodeSnakealiveMInternalModels5(l, v)
}
func easyjson30497fe1DecodeSnakealiveMInternalModels6(in *jlexer.Lexer, out *Album) {
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
		case "id":
			out.Id = int(in.Int())
		case "trip_id":
			out.TripId = int(in.Int())
		case "user_id":
			out.UserId = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "photos":
			if in.IsNull() {
				in.Skip()
				out.Photos = nil
			} else {
				in.Delim('[')
				if out.Photos == nil {
					if !in.IsDelim(']') {
						out.Photos = make([]string, 0, 4)
					} else {
						out.Photos = []string{}
					}
				} else {
					out.Photos = (out.Photos)[:0]
				}
				for !in.IsDelim(']') {
					var v25 string
					v25 = string(in.String())
					out.Photos = append(out.Photos, v25)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson30497fe1EncodeSnakealiveMInternalModels6(out *jwriter.Writer, in Album) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"trip_id\":"
		out.RawString(prefix)
		out.Int(int(in.TripId))
	}
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix)
		out.Int(int(in.UserId))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"photos\":"
		out.RawString(prefix)
		if in.Photos == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v26, v27 := range in.Photos {
				if v26 > 0 {
					out.RawByte(',')
				}
				out.String(string(v27))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Album) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson30497fe1EncodeSnakealiveMInternalModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Album) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson30497fe1EncodeSnakealiveMInternalModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Album) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson30497fe1DecodeSnakealiveMInternalModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Album) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson30497fe1DecodeSnakealiveMInternalModels6(l, v)
}
