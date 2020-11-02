package tool

import (
	gredis2 "finders-server/pkg/gredis"
	"finders-server/st"
	"os"
)

func TestCache(){
	_, err := gredis2.Get("hello")
	if err != nil {
		st.Debug(err)
		os.Exit(-1)
	}

}