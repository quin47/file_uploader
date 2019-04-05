package main

import (
	"bufio"
	"file_uploader/notification"
	"file_uploader/uploader"
	"file_uploader/watcher"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var folder = os.Getenv("watch_folder")

func main() {

	watcher.RegHandFunc(fsnotify.Create, func(event fsnotify.Event) {

		file, e := os.Open(event.Name)
		if e != nil {
			log.Println("read file failed ", e)
			return
		}

		reader := bufio.NewReader(file)

		info, _ := file.Stat()
		bytes, _ := ioutil.ReadFile(event.Name)
		contentType := http.DetectContentType(bytes)
		httpPath := uploader.Upload(info.Name(), reader, info.Size(), contentType)

		notification.NotifyAndExportUrl(info.Name(),httpPath)

		defer file.Close()

	})

	//移到回收站 mac下是 rename 操作
	watcher.RegHandFunc(fsnotify.Rename, func(event fsnotify.Event) {
		//log.Print("remove", event)

	})

	watcher.RegHandFunc(fsnotify.Remove, func(event fsnotify.Event) {
		//log.Print("remove", event)

	})

	watcher.AddWatcherFolder(folder)

	select {}

}
