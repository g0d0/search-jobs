 package main


 import (
	"fmt"
	"strings"
	"github.com/PuerkitoBio/goquery"
 )

 func main() {
	doc, _ := goquery.NewDocument("https://www.bauruempregos.com.br/home/vagas")
	
	var data_corrente string
	doc.Find("#content-left div").Each(func(index int, item *goquery.Selection) {
		var vaga Vaga

		if item.HasClass("data-vagas") {
			data_corrente = item.Text()
		}
	
		if item.HasClass("vaga") {
			a := item.Find(".descricao-vaga")
			url, _ := a.Attr("href")
			vaga.nome = strings.Trim(a.Text(), " \t\n")
			vaga.url = url
			vaga.cidade = strings.Trim(item.Find(".cidade-vaga").Text(), "\n ")
		}
	
		if vaga.nome != "" {
			vaga.data = data_corrente
			fmt.Println(vaga)
		}
	})

 }
