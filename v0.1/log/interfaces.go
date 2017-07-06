package log


type ILogging interface {
    Verbose(fmt string, l ...interface{})
    Info(fmt string, l ...interface{})
    Warning(fmt string, l ...interface{})
    Error(fmt string, l ...interface{})
    Fatal(fmt string, l ...interface{})
}

