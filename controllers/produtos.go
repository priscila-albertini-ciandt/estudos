package controllers

import (
	"alura/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

const codeRedirect int = 301

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat, erro := strconv.ParseFloat(preco, 64)

		if erro != nil {
			log.Println("Erro na conversão do preço:", erro)
		}

		quantidadeConvertidaParaInt, erro := strconv.Atoi(quantidade)

		if erro != nil {
			log.Println("Erro na conversão da quantidade:", erro)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", codeRedirect)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	idDoProduto := r.URL.Query().Get("id")

	models.DeletarProduto(idDoProduto)

	http.Redirect(w, r, "/", codeRedirect)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")

	produto := models.BuscaProdutoPorId(idDoProduto)

	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idCovertidaParaInt, erro := strconv.Atoi(id)

		if erro != nil {
			log.Println("Erro na conversão do id:", erro)
		}

		precoConvertidoParaFloat, erro := strconv.ParseFloat(preco, 64)

		if erro != nil {
			log.Println("Erro na conversão do preço:", erro)
		}

		quantidadeConvertidaParaInt, erro := strconv.Atoi(quantidade)

		if erro != nil {
			log.Println("Erro na conversão da quantidade:", erro)
		}

		models.AtualizaProduto(idCovertidaParaInt, quantidadeConvertidaParaInt, nome, descricao, precoConvertidoParaFloat)

		http.Redirect(w, r, "/", codeRedirect)
	}
}
