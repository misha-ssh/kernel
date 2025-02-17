package logger

type Logger interface {
	Log(key string, value interface{})
}

//const Skip = 1
//
//func GenerateFile(fl storage.File) error {
//	SetFile(fl)
//	logFile := GetFile()
//
//	if !logFile.IsExistFile() {
//		err := logFile.CreateFile()
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//// TODO переписать логику на пакет file
//func getOpenLogFile(fl storage.File) (*os.File, error) {
//	path := filepath.Join(fl.Path, fl.Name)
//
//	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
//	if err != nil {
//		return nil, err
//	}
//	return logFile, nil
//}
//
//func Danger(message string) {
//	logFile := GetFile()
//	openLogFile, err := getOpenLogFile(logFile)
//	if err != nil {
//		output.GetOutError("dont open log file")
//	}
//
//	defer func(openLogFile *os.File) {
//		err := openLogFile.Close()
//		if err != nil {
//			output.GetOutError("dont close opened log file")
//		}
//	}(openLogFile)
//
//	_, calledFile, line, ok := runtime.Caller(Skip)
//	if ok {
//		message = "file: " + calledFile + ", line: " + strconv.Itoa(line) + ", message: " + message
//	}
//
//	errorLog := log.New(openLogFile, "[error] ", log.LstdFlags|log.Lmicroseconds)
//	errorLog.Println(message)
//}
