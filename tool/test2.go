package tool

import (
	"finders-server/service/gredis"
	"finders-server/st"
	"os"
)

func TestCache(){
	_, err := gredis.Get("hello")
	if err != nil {
		st.Debug(err)
		os.Exit(-1)
	}

}