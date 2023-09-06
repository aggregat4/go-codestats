package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/boyter/scc/v3/processor"
)

type statsProcessor struct{}

func (p *statsProcessor) ProcessLine(job *processor.FileJob, currentLine int64, lineType processor.LineType) bool {
	switch lineType {
	case processor.LINE_BLANK:
		fmt.Println(currentLine, "lineType", "BLANK")
	case processor.LINE_CODE:
		fmt.Println(currentLine, "lineType", "CODE")
	case processor.LINE_COMMENT:
		fmt.Println(currentLine, "lineType", "COMMENT")
	}
	return true
}

// func main() {
// 	bts, _ := os.ReadFile("main.go")

// 	t := &statsProcessor{}
// 	filejob := &processor.FileJob{
// 		Filename: "test.go",
// 		Language: "Go",
// 		Content:  bts,
// 		Callback: t,
// 		Bytes:    int64(len(bts)),
// 	}

// 	processor.ProcessConstants() // Required to load the language information and need only be done once
// 	processor.CountStats(filejob)
// }

func main() {
	root := "/home/terzic/dev/cad/master/cadenza/"
	stack := []string{root}
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			stack = append(stack, path)
		}
		parentPath := ""
		if len(stack) > 1 {
			parentPath = stack[len(stack)-2]
		}
		fmt.Println(path, parentPath)
		if d.IsDir() {
			defer func() { stack = stack[:len(stack)-1] }()
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
