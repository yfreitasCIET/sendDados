package main

import (
	"fmt"
)

type conexao struct{
	host string
	drive string
	nome_banco string
	usuario string
	senha string
}



func (c conexao) salvar(cNovo conexao){
	c.drive = cNovo.drive
	c.host = cNovo.host
	c.nome_banco = cNovo.nome_banco
	c.senha = cNovo.senha
	c.usuario = cNovo.usuario
}

func main() {
	var hostN string 
	var driveN string 
	var bancoN string 
	var senhaN string 
	var usuarioN string 

	
	
}
