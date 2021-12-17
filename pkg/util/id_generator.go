package util

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

var (
	ID  *snowflake.Node
	err error
)

func init() {
	ID, err = snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetIntegerValue(str string) int64 {
	parseBase58, err := snowflake.ParseBase58([]byte(str))
	if err != nil {
		return 0
	}
	return parseBase58.Int64()
}

func GetStringValue(id int64) string {
	return snowflake.ParseInt64(id).Base58()
}
