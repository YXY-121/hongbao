package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"runtime"
)

func main() {

	// 默认并发数
//	reader:=strings.NewReader("a b c")
//	f, err := os.OpenFile("goWriter.txt", os.O_APPEND, 0764)
//	defer f.Close()
//	if err != nil {
//		fmt.Println("文件不存在,正在创建文件")
//		f, _ = os.Create("goWriter.txt")
//	}
//	req,_:=http.NewRequest("get","https://apache.claz.org/zookeeper/zookeeper-3.7.0/apache-zookeeper-3.7.0-bin.tar.gz",nil)
//	rep,_:=http.DefaultClient.Do(req)
//	buf:=make([]byte,100)
//	io.CopyBuffer(f,rep.Body,buf)
//	fmt.Println(buf)
	concurrencyN := runtime.NumCPU()

	app := &cli.App{
		Name:  "downloader",
		Usage: "File concurrency downloader",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				Usage:    "`URL` to download",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Output `filename`",
			},
			&cli.IntFlag{
				Name:    "concurrency",
				Aliases: []string{"n"},
				Value:   concurrencyN,
				Usage:   "Concurrency `number`",
			},
			&cli.BoolFlag{
				Name:    "resume",
				Aliases: []string{"r"},
				Value:   true,
				Usage:   "Resume download",
			},
		},
		Action: func(c *cli.Context) error {
			strURL := c.String("url")
			filename := c.String("output")
			concurrency := c.Int("concurrency")
			resume := c.Bool("resume")
			return NewDownloader(concurrency, resume).Download(strURL, filename)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}