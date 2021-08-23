package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

var (
	app *cli.App

	// flag* oprion指定の受け取り先
	flagOrg string
	flagIn  string
)

func main() {
	app = &cli.App{
		Usage: `-o [orgFile] -i [inFile]`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "o",
				Usage:       "org file",
				Destination: &flagOrg,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "i",
				Usage:       "insert file",
				Destination: &flagIn,
				Required:    true,
			},
		},
		Action: func(c *cli.Context) error {
			buf, err := ioutil.ReadFile(flagOrg)
			if err != nil {
				log.Printf("hoge %s", err)
				return err
			}

			data, err := ReadOnSliceMap(buf)
			if err != nil {
				log.Printf("fff %s", err)
				return err
			}
			for k, v := range data {
				if k.(string) == "paths" {
					fmt.Println(v)
					// for _, path := range v.() { // path
					// 	for _, method := range path { // method

					// 	}
					// }

					//   responses:
					// 	'200':
					// 	  content:
					// 		application/json:
					// 		  examples:
				}
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadOnSliceMap(fileBuffer []byte) (map[interface{}]interface{}, error) {
	data := make(map[interface{}]interface{})
	err := yaml.Unmarshal(fileBuffer, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
