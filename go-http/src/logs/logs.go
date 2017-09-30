package logs

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

const (
	_debug = 0
	_info  = 1
	_warn  = 2
	_error = 3
)

var default_logger *logger = nil

type logger struct {
	level   int
	console *log.Logger
	file    *log.Logger
}

func init() {
	//fmt.Println("init logs ...")
	var _logger = logger{
		level: 0,
	}

	var consoles = []io.Writer{
		os.Stdout,
	}
	_logger.console = log.New(io.MultiWriter(consoles...), "[devel] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)

	default_logger = &_logger
}

const (
	color_text_black = iota + 30
	color_text_red
	color_text_green
	color_text_yellow
	color_text_blue
	color_text_magenta
	color_text_cyan
	color_text_white
)

func color_text(color int, str string) string {
	if runtime.GOOS == "windows" {
		return str
	}

	var format = "\x1b[0;%dm%s\x1b[0m"
	switch color {
	case color_text_black,
		color_text_red,
		color_text_green,
		color_text_yellow,
		color_text_blue,
		color_text_magenta,
		color_text_cyan,
		color_text_white:
		return fmt.Sprintf(format, color, str)
	default:
		return str
	}
}

func Debug(format string, args ...interface{}) {
	if default_logger.level > _debug {
		return
	}

	var _format = "[DEBUG]▸%s"
	if default_logger.console != nil {
		default_logger.console.Output(2,
			color_text(
				color_text_green,
				fmt.Sprintf(_format, fmt.Sprintf(format, args...))))
	}

	if default_logger.file != nil {
		default_logger.file.Output(
			2,
			fmt.Sprintf(_format, fmt.Sprintf(format, args...)))
	}
}
func Info(format string, args ...interface{}) {
	if default_logger.level > _info {
		return
	}

	var _format = "[INFO]▸%s"
	if default_logger.console != nil {
		default_logger.console.Output(
			2,
			color_text(
				color_text_white,
				fmt.Sprintf(_format, fmt.Sprintf(format, args...))))
	}

	if default_logger.file != nil {
		default_logger.file.Output(
			2,
			fmt.Sprintf(_format, fmt.Sprintf(format, args...)))
	}
}
func Warn(format string, args ...interface{}) {
	if default_logger.level > _warn {
		return
	}

	var _format = "[WARN]▸%s"
	if default_logger.console != nil {
		default_logger.console.Output(
			2,
			color_text(
				color_text_yellow,
				fmt.Sprintf(_format, fmt.Sprintf(format, args...))))
	}

	if default_logger.file != nil {
		default_logger.file.Output(
			2,
			fmt.Sprintf(_format, fmt.Sprintf(format, args...)))
	}
}
func Error(format string, args ...interface{}) {
	if default_logger.level > _error {
		return
	}

	var _format = "[ERROR]▸%s"
	if default_logger.console != nil {
		var log_info = color_text(
			color_text_red,
			fmt.Sprintf(_format, fmt.Sprintf(format, args...)))
		default_logger.console.Output(
			2,
			log_info)

		panic(log_info)
	}

	if default_logger.file != nil {
		var log_info = fmt.Sprintf(_format, fmt.Sprintf(format, args...))
		default_logger.file.Output(
			2,
			log_info)

		panic(log_info)
	}
}
