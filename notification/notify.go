package notification

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/gen2brain/beeep"
)

func NotifyAndExportUrl(fname,fpath string)  {
	err := beeep.Notify(fmt.Sprintf("upload %v finished !",fname), "url in clipboard","")
	if err != nil {
		panic(err)
	}
	clipboard.WriteAll(fpath)
}

func SimpleNotify(title,desc string)  {
	err := beeep.Notify(title, desc,"")
	if err != nil {
		panic(err)
	}
}