package requests

import "fmt"

func NewPontuacaoRodada(rodada int) *Request {
	r := &Request{
		Url:    "https://api.cartola.globo.com/atletas/pontuados/:rodada",
		Method: "GET",
	}

	r.queryParams(map[string]string{":rodada": fmt.Sprint(rodada)})

	return r
}

func NewResultadosRodada(rodada int) *Request {
	r := &Request{
		Url:    "https://api.cartola.globo.com/partidas/:rodada",
		Method: "GET",
	}

	r.queryParams(map[string]string{":rodada": fmt.Sprint(rodada)})

	return r
}
