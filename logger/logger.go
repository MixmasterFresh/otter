package logger

import(
    "fmt"
)

type Level int
type OutputType int

type loggerConfiguration struct {
    level          Level
    outputType     OutputType
    fileLocation   string
    buffer         bool
    bufferInterval int
    stream         string
}

type streamConfig struct {
    address string // more to be added later
}

const(
    DEBUG Level = iota
    INFO  Level = iota
    WARN  Level = iota
    ERROR Level = iota
)

const(
    STDOUT OutputType = iota
    FILE   OutputType = iota
    STREAM OutputType = iota
)

var config loggerConfiguration

func Initialize(level Level, outputType OutputType) {
    config = loggerConfiguration{level: level, outputType: outputType}
}

func Log(level Level, logData ...interface{}) {
    if level >= config.level{
        fmt.Println(logData)
    }
    //set up logging function here
}
