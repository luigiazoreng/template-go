package log

import (
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	SystemTag     = color.New(color.FgBlue, color.Bold).SprintFunc()
	TicketTag     = color.New(color.FgYellow, color.Bold).SprintFunc()
	DecoderTag    = color.New(color.FgRed, color.Bold).SprintFunc()
	middlewareTag = color.New(color.FgCyan, color.Bold).SprintFunc()
	routerTag     = color.New(color.FgGreen, color.Bold).SprintFunc()
)

type Logger struct {
	*log.Logger
}

func StartLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}
func (l *Logger) System(format string, v ...interface{}) {
	l.Printf(SystemTag("[System] ")+format, v...)
}

func (l *Logger) SystemFatal(format string, v ...interface{}) {
	l.Printf(SystemTag("[System] ")+format, v...)
	os.Exit(1)
}

func (l *Logger) Product(format string, v ...interface{}) {
	l.Printf(TicketTag("[Ticket] ")+format, v...)
}
func (l *Logger) ProductFatal(format string, v ...interface{}) {
	l.Printf(TicketTag("[Ticket] ")+format, v...)
	os.Exit(1)
}

func (l *Logger) Decoder(format string, v ...interface{}) {
	l.Printf(DecoderTag("[Json Decoder] ")+format, v...)
}
func (l *Logger) DecoderFatal(format string, v ...interface{}) {
	l.Printf(DecoderTag("[Json Decoder] ")+format, v...)
	os.Exit(1)
}

func (l *Logger) Middleware(format string, v ...interface{}) {
	l.Printf(middlewareTag("[Middleware] ")+format, v...)
}
func (l *Logger) MiddlewareFatalf(format string, v ...interface{}) {
	l.Printf(middlewareTag("[Middleware] ")+format, v...)
	os.Exit(1)
}

func (l *Logger) Router(format string, v ...interface{}) {
	l.Printf(routerTag("[Router] ")+format, v...)
}
func (l *Logger) RouterFatal(format string, v ...interface{}) {
	l.Printf(routerTag("[Router] ")+format, v...)
	os.Exit(1)
}
