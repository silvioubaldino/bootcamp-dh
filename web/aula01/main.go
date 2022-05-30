package main

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

type usuario struct {
	Id          int     `json:"id"`
	Nome        string  `json:"nome"`
	Sobrenome   string  `json:"sobrenome"`
	Email       string  `json:"email"`
	Idade       int     `json:"idade"`
	Altura      float64 `json:"altura"`
	Ativo       bool    `json:"ativo"`
	DataCriacao string  `json:"dataCriacao"`
}

func main() {
	router := gin.Default()

	router.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
	group := router.Group("/usuarios")
	{
		group.GET("/getAll", getAll)
		group.GET("/", getWithParam)
	}

	router.Run()
}

func getAll(c *gin.Context) {
	c.BindJSON("/Users/sineto/Documents/bootcamp-dh/web/usuarios.json")
}

func readJSON() ([]usuario, error) {
	content, err := ioutil.ReadFile("/Users/sineto/Documents/bootcamp-dh/web/usuarios.json")
	if err != nil {
		return nil, errors.New("erro ao ler arquivo")
	}

	var usuarios []usuario
	if err := json.Unmarshal(content, &usuarios); err != nil {
		return nil, errors.New("erro ao converter objetos para json")
	}
	return usuarios, nil
}

func getWithParam(c *gin.Context) {
	usuarios, err := readJSON()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	var filtrados []usuario
	for _, usuario := range usuarios {
		if filtrar(c, &usuario) {
			filtrados = append(filtrados, usuario)
		}
	}
	if len(filtrados) > 0 {
		c.JSON(200, filtrados)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "nenhum usuario encontrado",
		})
	}
}

func filtrar(c *gin.Context, usuario *usuario) bool {
	//queries := c.Request.URL.Query()
	//for _, q := range queries {
	//}
	if c.Query("nome") != "" && !strings.Contains(usuario.Nome, c.Query("nome")) {
		return false
	} else if c.Query("sobrenome") != "" && !strings.Contains(usuario.Sobrenome, c.Query("sobrenome")) {
		return false
	}
	return true
}
