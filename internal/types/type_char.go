package types

type TypeChar struct {
	TypeByte
}

//func (t *TypeChar) GetLength() uint32 {
//	return 1
//}
//
//func (t *TypeChar) GetData() (byte, error) {
//	if !t.IsFull() {
//		return 0, errors.New("no data available")
//	}
//	return int8(t.data[0]), nil
//}
//
//func (t *TypeChar) GetBuffer(data int8) []byte {
//	return []byte{byte(data)}
//}
