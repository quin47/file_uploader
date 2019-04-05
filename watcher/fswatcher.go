package watcher

import (
	"file_uploader/notification"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
)

type CmdType string

var watcher *fsnotify.Watcher

//Context 消息上下文环境,提供快捷提取消息数据的功能
type Context struct {
}

var handlerMap = make(map[fsnotify.Op][]HandleFunc)

//Handler event function
type HandleFunc func(event fsnotify.Event)

type Fswatcher struct {
	root    string
	watcher *fsnotify.Watcher
}

func RegHandFunc(op fsnotify.Op, hfunc HandleFunc) {
	opHandlers := append(handlerMap[op], hfunc)
	handlerMap[op] = opHandlers
}

func AddWatcherFolder(folder string) {

	_, e := os.Stat(folder)

	if e !=nil {

		notification.SimpleNotify("要监视的文件不存在,程序将退出",folder)
		os.Exit(1)

	}

	watcher.Add(folder)
}

func init() {
	var err error
	watcher, err = fsnotify.NewWatcher()

	if err != nil {
		log.Fatal("init fs watcher failed ")
	}

	go func() {
		for {
			select {

			case event, ok := <-watcher.Events:
				if ok {
					//log.Println(event)
					op := event.Op
					handlers := handlerMap[op]
					if handlers != nil {
						for _, handleFunc := range handlers {
							go handleFunc(event)
						}
					}

				}
			case err, ok := <-watcher.Errors:
				if ok {
					log.Println("error", err)
				}

			}

		}

	}()

}
