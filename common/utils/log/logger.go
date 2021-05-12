package log

import (
	"GinDemo/config"
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	rotateLogs "github.com/lestrrat/go-file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"time"
)

var (
	Logger = log.New()
)

// Setup initialize the log instance
func Setup() {
	_ = getLogFilePath()
	fileName := getLogFileName()

	writer, _ := rotateLogs.New(config.ServerSetting.LogSavePath+
		"%Y-%m-%d_%H-%M"+fileName,
		//rotateLogs.WithLinkName(logFile),
		rotateLogs.WithMaxAge(config.ServerSetting.LogMaxAge*24*time.Hour),
		rotateLogs.WithRotationTime(config.ServerSetting.LogRotationTime*time.Hour),
	)
	mw := io.MultiWriter(os.Stdout, writer)
	Logger.SetOutput(mw)
	Logger.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Logger.SetReportCaller(true)
	//logger.Hooks.Add(NewContextHook())
}

func getLogFilePath() string {
	var (
		logPath                 = config.ServerSetting.LogSavePath
		currentPath, logAbsPath string
		err                     error
	)
	if currentPath, err = filepath.Abs("."); err != nil {
		panic(err)
	}
	logAbsPath = filepath.Join(currentPath, logPath)
	if _, err := os.Stat(logAbsPath); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(logAbsPath, 0755)
			if err != nil {
				panic(err)
			}
		}
	}
	return logPath
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s.%s",
		time.Now().Format(setting.AppSetting.TimeFormat),
		"log",
	)
}
