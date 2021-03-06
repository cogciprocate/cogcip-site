package main

import (
	_ "github.com/lib/pq"
	"github.com/c0gent/unframed"
	"github.com/c0gent/unframed/log"
)

var net *unframed.NetHandle
var db *unframed.DbHandle

var cfgFile string = "/etc/cogcip-site/config.json"

func main() {
	cfg := unframed.ReadConfig(cfgFile)

	db = unframed.NewDB(cfg.DbType, cfg.ConnStr)
	defer db.Close()
	log.Message(cfg.Wd)
	net = unframed.NewNet()

	unframed.DefaultPageTitle = "Cogciprocate"
	net.TemplateFiles(
		"tmpl/_base.html.tmpl",
		// "tmpl/_side_table_1.html.tmpl",
	)
	//net.Get("/episodes/watch", episodesWatch)
	//net.Get("/episodes", episodesList)
	net.Dir("assets/")
	net.Dir("public/")

	registerControllers()

	db.PrepareStatements()
	net.LoadTemplates()

	log.Message("Serving cogcip-site")
	net.Serve(cfg.ListenPort)
}

func registerControllers() {
	// inhousesReg()
	// lfgsReg()
	// lfmsReg()
	homeReg()
	//codingChallengeReg()
}
