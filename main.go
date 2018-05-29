package main

import (
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xprop"

	"os"
	"strconv"
)

func main() {
	conn, _ := xgbutil.NewConn()
	xprop.ChangeProp32(conn, conn.Screen().Root, "_XROOTPMAP_ID", "PIXMAP", func() uint { a,_:=strconv.ParseUint(os.Args[1], 0, 32); return uint(a) }())
}
