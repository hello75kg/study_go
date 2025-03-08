package main

import (
	"fmt"
)

type MyWriter interface {
	Write()
}
type MyClose interface {
	Close()
}

// 接口嵌套
type MyWriterAndClose interface {
	MyWriter
	MyClose
	WriteAndClose()
}

type WriterAndCloser struct {
	MyWriter
}

type WriterAndCloser2 struct{}

func (wac *WriterAndCloser2) Write() {
	fmt.Println("MyWriterAndClose write")
}

func (wac *WriterAndCloser2) Close() {
	fmt.Println("MyWriterAndClose close")
}

func (wac *WriterAndCloser2) WriteAndClose() {
	fmt.Println("WriteAndClose write and close")
}

type DatabaseWriter struct {
}

type FileWriter struct {
}

func (dw *DatabaseWriter) Write() {
	fmt.Println("Database Writer Write")
}

func (fw *FileWriter) Write() {
	fmt.Println("File Writer Write")
}

//	func (wac *WriterAndCloser) Write() {
//		fmt.Println("write")
//	}
func (wac *WriterAndCloser) Close() {
	fmt.Println("close")
}

func main() {
	//var mw MyWriter = &WriterAndCloser{}
	//mw.Write()
	var mc MyClose = &WriterAndCloser{}
	mc.Close()
	var fw MyWriter = &FileWriter{}
	fw.Write()
	var dw MyWriter = &DatabaseWriter{}
	dw.Write()

	var mw2 MyWriter = &WriterAndCloser{
		&FileWriter{},
	}
	mw2.Write()
	var mw3 MyWriter = &WriterAndCloser{
		&DatabaseWriter{},
	}
	mw3.Write()

	//接口嵌套
	var mwac MyWriterAndClose = &WriterAndCloser2{}
	mwac.Write()
	mwac.Close()
	mwac.WriteAndClose()
}
