package db

import (
	"encoding/binary"
	"fmt"

	"github.com/jackc/pgx/pgtype"
)

func HelperC(rp *int, src []byte) []byte {
	elemLen := int(int32(binary.BigEndian.Uint32(src[*rp:])))
	*rp += 4
	var elemSrc []byte
	if elemLen >= 0 {
		elemSrc = src[*rp : *rp+elemLen]
		*rp += elemLen
	}
	return elemSrc
}

func HelperA(ci *pgtype.ConnInfo, src []byte) (count int32, rpp int, err error) {
	var arrayHeader pgtype.ArrayHeader
	rp, err := arrayHeader.DecodeBinary(ci, src)
	if err != nil {
		return 0, 0, err
	}

	if len(arrayHeader.Dimensions) == 0 {

		return 0, rp, nil
	}

	elementCount := arrayHeader.Dimensions[0].Length
	for _, d := range arrayHeader.Dimensions[1:] {
		elementCount *= d.Length
	}
	fmt.Println("testarrrr", len(src), arrayHeader, rp)

	return elementCount, rp, nil

}

func Helper(ci *pgtype.ConnInfo, src []byte, fields []interface{}) error {
	if src == nil {
		//	*dst = Weburl{Status: pgtype.Null}
		return nil
	}

	//fmt.Println("testtext1111", len(src))

	rp := 0

	if len(src[rp:]) < 4 {
		return fmt.Errorf("Record incomplete %v", src)
	}
	fieldCount := int(int32(binary.BigEndian.Uint32(src[rp:])))
	rp += 4

	//fields := make([]pgtype.Value, fieldCount)

	for i := 0; i < fieldCount; i++ {
		if len(src[rp:]) < 8 {
			return fmt.Errorf("Record incomplete %v", src)
		}
		fieldOid := pgtype.Oid(binary.BigEndian.Uint32(src[rp:]))
		rp += 4

		fieldLen := int(int32(binary.BigEndian.Uint32(src[rp:])))
		rp += 4

		var binaryDecoder pgtype.BinaryDecoder
		if dt, ok := ci.DataTypeForOid(fieldOid); ok {
			if binaryDecoder, ok = dt.Value.(pgtype.BinaryDecoder); !ok {
				return fmt.Errorf("unknown oid while decoding record: %v", fieldOid)
			}

		}

		var fieldBytes []byte
		if fieldLen >= 0 {
			if len(src[rp:]) < fieldLen {
				return fmt.Errorf("Record incomplete %v", src)
			}
			fieldBytes = src[rp : rp+fieldLen]
			rp += fieldLen
		}

		if err := binaryDecoder.DecodeBinary(ci, fieldBytes); err != nil {
			return err
		}

		ga := binaryDecoder.(pgtype.Value)
		la := ga.Get()
		if la != nil {

			err := ga.AssignTo(fields[i])

			if err != nil {
				fmt.Println(err)
				return err
			}

		}
		//	fmt.Println("field", i, fieldOid, stru)

	}

	//	fmt.Println("url", fields[1].Get(), fields[1])
	//fmt.Printf("str %s %s\n", *stru.Url, *stru.W_cr_uid)
	//*dst = *stru
	return nil
}
