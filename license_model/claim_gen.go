package license_model

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Claims) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "iat":
			z.Iat, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Iat")
				return
			}
		case "iss":
			z.Iss, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Iss")
				return
			}
		case "prods":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Products")
				return
			}
			if cap(z.Products) >= int(zb0002) {
				z.Products = (z.Products)[:zb0002]
			} else {
				z.Products = make([]string, zb0002)
			}
			for za0001 := range z.Products {
				z.Products[za0001], err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Products", za0001)
					return
				}
			}
		case "lictype":
			{
				var zb0003 string
				zb0003, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "LicenseType")
					return
				}
				z.LicenseType = LicenseType(zb0003)
			}
		case "lickeyh":
			z.LicenseKeyHash, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "LicenseKeyHash")
				return
			}
		case "licmaxv":
			z.LicenseMaxVersion, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "LicenseMaxVersion")
				return
			}
		case "lisname":
			z.LicenseeName, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "LicenseeName")
				return
			}
		case "lismail":
			z.LicenseeMail, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "LicenseeMail")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Claims) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 8
	// write "iat"
	err = en.Append(0x88, 0xa3, 0x69, 0x61, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Iat)
	if err != nil {
		err = msgp.WrapError(err, "Iat")
		return
	}
	// write "iss"
	err = en.Append(0xa3, 0x69, 0x73, 0x73)
	if err != nil {
		return
	}
	err = en.WriteString(z.Iss)
	if err != nil {
		err = msgp.WrapError(err, "Iss")
		return
	}
	// write "prods"
	err = en.Append(0xa5, 0x70, 0x72, 0x6f, 0x64, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Products)))
	if err != nil {
		err = msgp.WrapError(err, "Products")
		return
	}
	for za0001 := range z.Products {
		err = en.WriteString(z.Products[za0001])
		if err != nil {
			err = msgp.WrapError(err, "Products", za0001)
			return
		}
	}
	// write "lictype"
	err = en.Append(0xa7, 0x6c, 0x69, 0x63, 0x74, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(string(z.LicenseType))
	if err != nil {
		err = msgp.WrapError(err, "LicenseType")
		return
	}
	// write "lickeyh"
	err = en.Append(0xa7, 0x6c, 0x69, 0x63, 0x6b, 0x65, 0x79, 0x68)
	if err != nil {
		return
	}
	err = en.WriteString(z.LicenseKeyHash)
	if err != nil {
		err = msgp.WrapError(err, "LicenseKeyHash")
		return
	}
	// write "licmaxv"
	err = en.Append(0xa7, 0x6c, 0x69, 0x63, 0x6d, 0x61, 0x78, 0x76)
	if err != nil {
		return
	}
	err = en.WriteInt(z.LicenseMaxVersion)
	if err != nil {
		err = msgp.WrapError(err, "LicenseMaxVersion")
		return
	}
	// write "lisname"
	err = en.Append(0xa7, 0x6c, 0x69, 0x73, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.LicenseeName)
	if err != nil {
		err = msgp.WrapError(err, "LicenseeName")
		return
	}
	// write "lismail"
	err = en.Append(0xa7, 0x6c, 0x69, 0x73, 0x6d, 0x61, 0x69, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteString(z.LicenseeMail)
	if err != nil {
		err = msgp.WrapError(err, "LicenseeMail")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Claims) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 8
	// string "iat"
	o = append(o, 0x88, 0xa3, 0x69, 0x61, 0x74)
	o = msgp.AppendInt64(o, z.Iat)
	// string "iss"
	o = append(o, 0xa3, 0x69, 0x73, 0x73)
	o = msgp.AppendString(o, z.Iss)
	// string "prods"
	o = append(o, 0xa5, 0x70, 0x72, 0x6f, 0x64, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Products)))
	for za0001 := range z.Products {
		o = msgp.AppendString(o, z.Products[za0001])
	}
	// string "lictype"
	o = append(o, 0xa7, 0x6c, 0x69, 0x63, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, string(z.LicenseType))
	// string "lickeyh"
	o = append(o, 0xa7, 0x6c, 0x69, 0x63, 0x6b, 0x65, 0x79, 0x68)
	o = msgp.AppendString(o, z.LicenseKeyHash)
	// string "licmaxv"
	o = append(o, 0xa7, 0x6c, 0x69, 0x63, 0x6d, 0x61, 0x78, 0x76)
	o = msgp.AppendInt(o, z.LicenseMaxVersion)
	// string "lisname"
	o = append(o, 0xa7, 0x6c, 0x69, 0x73, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.LicenseeName)
	// string "lismail"
	o = append(o, 0xa7, 0x6c, 0x69, 0x73, 0x6d, 0x61, 0x69, 0x6c)
	o = msgp.AppendString(o, z.LicenseeMail)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Claims) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "iat":
			z.Iat, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Iat")
				return
			}
		case "iss":
			z.Iss, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Iss")
				return
			}
		case "prods":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Products")
				return
			}
			if cap(z.Products) >= int(zb0002) {
				z.Products = (z.Products)[:zb0002]
			} else {
				z.Products = make([]string, zb0002)
			}
			for za0001 := range z.Products {
				z.Products[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Products", za0001)
					return
				}
			}
		case "lictype":
			{
				var zb0003 string
				zb0003, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "LicenseType")
					return
				}
				z.LicenseType = LicenseType(zb0003)
			}
		case "lickeyh":
			z.LicenseKeyHash, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LicenseKeyHash")
				return
			}
		case "licmaxv":
			z.LicenseMaxVersion, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LicenseMaxVersion")
				return
			}
		case "lisname":
			z.LicenseeName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LicenseeName")
				return
			}
		case "lismail":
			z.LicenseeMail, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LicenseeMail")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Claims) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 4 + msgp.StringPrefixSize + len(z.Iss) + 6 + msgp.ArrayHeaderSize
	for za0001 := range z.Products {
		s += msgp.StringPrefixSize + len(z.Products[za0001])
	}
	s += 8 + msgp.StringPrefixSize + len(string(z.LicenseType)) + 8 + msgp.StringPrefixSize + len(z.LicenseKeyHash) + 8 + msgp.IntSize + 8 + msgp.StringPrefixSize + len(z.LicenseeName) + 8 + msgp.StringPrefixSize + len(z.LicenseeMail)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *LicenseType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 string
		zb0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = LicenseType(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z LicenseType) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteString(string(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z LicenseType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LicenseType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 string
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = LicenseType(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z LicenseType) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}
