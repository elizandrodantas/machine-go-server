package main

import (
	"flag"
	"log"
)

type flags struct {
	drop   bool
	create bool
}

func main() {
	var f flags
	flag.BoolVar(&f.create, "c", false, "create tables database")
	flag.BoolVar(&f.drop, "d", false, "drop tables database")
	flag.Parse()

	if f.create {
		err := f.do_create()

		if err != nil {
			panic(err)
		}

		log.Println("tables created successfully")
		return
	}

	if f.drop {
		err := f.do_drop()

		if err != nil {
			panic(err)
		}

		log.Println("tables have been successfully dropped")
		return
	}

	f.do_server()
}
