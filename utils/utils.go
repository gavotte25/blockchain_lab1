package utils

import (
	"bytes"
	"encoding/binary"
	"time"
)

// used to convert an int64 to a byte array
func ConvertTimestampToByte(num int64) []byte {
	/*
		Function to convert an int64 to a byte array in Big Endian format.

		:param num: int64 number to be converted
		:return: byte array representing the hex value of the input number
	*/
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		// Handle the error in a way suitable for your application
		panic(err)
	}

	return buff.Bytes()
}

func GetTimestampFormat(timestamp int64) string {
	unixTimeUTC := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	return unixTimeUTC
}
