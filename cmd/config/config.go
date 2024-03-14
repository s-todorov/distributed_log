package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

type Config struct {
	Logger  *slog.Logger
	AppName string `json:"appName"`
	LogPath string `json:"logPath"`
}

func NewConfig(cfgPath string) Config {

	var err error
	var conf Config
	filePath := filepath.Clean(cfgPath)
	f, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			slog.Error(err.Error(), err)
			panic(err)
		}
		slog.Error(err.Error(), err)
		panic(err)
	}

	err = json.NewDecoder(f).Decode(&conf)
	if err != nil {
		slog.Error(err.Error(), err)
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go conf.logChecker(&wg)
	wg.Wait()
	return conf
}

func (cfg *Config) logChecker(wg *sync.WaitGroup) {
	cfg.Logger = setLogOut(cfg.LogPath)
	slog.SetDefault(cfg.Logger)

	deleteOldFiles(cfg.LogPath)
	wg.Done()
	t := time.NewTicker(1 * time.Hour)
	for {
		<-t.C
		newLog := setLogOut(cfg.LogPath)
		if newLog != nil {
			*cfg.Logger = *newLog
			slog.SetDefault(cfg.Logger)
		}
		deleteOldFiles(cfg.LogPath)
	}
}

func deleteOldFiles(filesPath string) error {
	tmpfiles, err := os.ReadDir(filesPath)
	if err != nil {
		return err
	}
	for _, file := range tmpfiles {
		if !file.IsDir() {
			info, err := file.Info()
			if err != nil {
				return err
			}
			//TODO in live change days from config
			if isOlderThan(60, info.ModTime()) {
				err := os.Remove(filesPath + "/" + file.Name())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func isOlderThan(days int, t time.Time) bool {
	return time.Now().Sub(t) > time.Duration(days)*24*time.Hour
}

func setLogOut(filesPath string) *slog.Logger {
	timeNow := time.Now()
	fileName := strconv.Itoa(timeNow.Year()) + "-" + fmt.Sprintf("%02d", int(timeNow.Month())) + "-" + fmt.Sprintf("%02d", timeNow.Day())
	tempLogFile, err := os.OpenFile(filesPath+"/"+fileName+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		slog.Error(err.Error(), err)
	}
	var mw io.Writer
	mw = io.MultiWriter(tempLogFile, os.Stdout)
	textHandler := slog.NewTextHandler(mw, nil)
	return slog.New(textHandler)

}
