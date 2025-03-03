package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bineyond/m3u8/dl"
)

var (
	url      string
	output   string
	chanSize int
	filename string
	mp4Type  bool
)

func init() {
	flag.StringVar(&url, "u", "", "M3U8 URL, required")
	flag.IntVar(&chanSize, "c", 25, "Maximum number of occurrences")
	flag.StringVar(&output, "o", "", "Output folder, required")
	flag.StringVar(&filename, "n", "", "Output filename, required")
	flag.BoolVar(&mp4Type, "mp4Type", false, "Output video mp4 type")
}

func main() {
	flag.Parse()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[error]", r)
			os.Exit(-1)
		}
	}()
	if url == "" {
		panicParameter("u")
	}
	if output == "" {
		panicParameter("o")
	}
	if chanSize <= 0 {
		panic("parameter 'c' must be greater than 0")
	}

	if filename != "" {
		dl.MergeTSFilename = filename + ".ts"
		dl.MergeMp4Filename = filename + ".mp4"
	}

	dl.MergeFileMp4Type = mp4Type

	downloader, err := dl.NewTask(output, url)
	if err != nil {
		panic(err)
	}
	if err := downloader.Start(chanSize); err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}

func panicParameter(name string) {
	panic("parameter '" + name + "' is required")
}
