package main

import (
	"cartola-scout/internal/model"
	"cartola-scout/internal/requests"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	posicao := map[int]string{
		1: "Goleiro",
		2: "Lateral",
		3: "Zagueiro",
		4: "Meia",
		5: "Atacante",
		6: "Técnico",
	}

	rodadaAlvo := 16
	clubes := carregarClubes()
	scouts := map[string]*model.Scout{}
	pontuacoesJogadores := make([]model.PontuacaoRodada, rodadaAlvo+1)
	resultadosRodada := make([][]requests.Partida, rodadaAlvo+1)

	for rodadaAtual := 1; rodadaAtual <= rodadaAlvo; rodadaAtual++ {
		reqPontuacao := requests.NewPontuacaoRodada(rodadaAtual)
		reqResultados := requests.NewResultadosRodada(rodadaAtual)

		download(rodadaAtual, "/Users/cicerohamonvitorinolopes/workspace/cartola-scout/arquivos/pontuacao/rodada-%d.json", *reqPontuacao)
		download(rodadaAtual, "/Users/cicerohamonvitorinolopes/workspace/cartola-scout/arquivos/resultados/rodada-%d.json", *reqResultados)

		pontuacoesJogadores[rodadaAtual] = carregarPontuacao(rodadaAtual)
		resultadosRodada[rodadaAtual] = carregarResultados(rodadaAtual)

		for idAtleta, pontuacaoAtleta := range pontuacoesJogadores[rodadaAtual].Atletas {
			if !pontuacaoAtleta.EntrouEmCampo {
				continue
			}

			var scoutAtual *model.Scout
			var ok bool

			if scoutAtual, ok = scouts[idAtleta]; !ok {
				scoutAtual = &model.Scout{}
				scoutAtual.Jogador = model.Jogador{
					Nome:    pontuacaoAtleta.Apelido,
					Posicao: posicao[pontuacaoAtleta.Posicao],
				}
				scoutAtual.Geral = model.Estatisticas{}
				scoutAtual.ContraTimeIgual = model.Estatisticas{}
				scoutAtual.ContraTimeMaior = model.Estatisticas{}
				scoutAtual.ContraTimeMenor = model.Estatisticas{}

				scouts[idAtleta] = scoutAtual
			}

			scoutAtual.Geral.Jogos++
			scoutAtual.Geral.TotalPontos += pontuacaoAtleta.Valor
			scoutAtual.Geral.Media = scoutAtual.Geral.TotalPontos / float32(scoutAtual.Geral.Jogos)

			for _, partida := range resultadosRodada[rodadaAtual] {
				if partida.ClubeCasaID != pontuacaoAtleta.Clube && partida.ClubeVisitanteID != pontuacaoAtleta.Clube {
					continue
				}

				if partida.ClubeCasaID == pontuacaoAtleta.Clube {
					incrementaEstatisticas(&scoutAtual.ComoMandante, pontuacaoAtleta)

					if partida.ClubeCasaPosicao > partida.ClubeVisitantePosicao {
						incrementaEstatisticas(&scoutAtual.ContraTimesAbaixo, pontuacaoAtleta)
					} else {
						incrementaEstatisticas(&scoutAtual.ContraTimesAcima, pontuacaoAtleta)
					}

					clubeCasa := clubes[pontuacaoAtleta.Clube]
					clubeVisitante := clubes[partida.ClubeVisitanteID]

					scoutAtual.Jogador.Clube.Nome = clubeCasa.Nome

					if clubeCasa.Classe > clubeVisitante.Classe {
						incrementaEstatisticas(&scoutAtual.ContraTimeMenor, pontuacaoAtleta)
					} else if clubeCasa.Classe < clubeVisitante.Classe {
						incrementaEstatisticas(&scoutAtual.ContraTimeMaior, pontuacaoAtleta)
					} else {
						incrementaEstatisticas(&scoutAtual.ContraTimeIgual, pontuacaoAtleta)
					}
				}

				if partida.ClubeVisitanteID == pontuacaoAtleta.Clube {
					incrementaEstatisticas(&scoutAtual.ComoVisitante, pontuacaoAtleta)

					if partida.ClubeCasaPosicao > partida.ClubeVisitantePosicao {
						incrementaEstatisticas(&scoutAtual.ContraTimesAbaixo, pontuacaoAtleta)
					} else {
						incrementaEstatisticas(&scoutAtual.ContraTimesAcima, pontuacaoAtleta)
					}

					clubeCasa := clubes[partida.ClubeCasaID]
					clubeVisitante := clubes[pontuacaoAtleta.Clube]

					scoutAtual.Jogador.Clube.Nome = clubeVisitante.Nome

					if clubeCasa.Classe > clubeVisitante.Classe {
						incrementaEstatisticas(&scoutAtual.ContraTimeMenor, pontuacaoAtleta)
					} else if clubeCasa.Classe < clubeVisitante.Classe {
						incrementaEstatisticas(&scoutAtual.ContraTimeMaior, pontuacaoAtleta)
					} else {
						incrementaEstatisticas(&scoutAtual.ContraTimeIgual, pontuacaoAtleta)
					}
				}

				break
			}
		}
	}

	sliceScouts := []*model.Scout{}

	for k, v := range scouts {
		id, _ := strconv.Atoi(k)
		v.Jogador.ID = id
		sliceScouts = append(sliceScouts, v)
	}

	j, _ := json.Marshal(sliceScouts)

	arquivoRodada := fmt.Sprintf("/Users/cicerohamonvitorinolopes/workspace/cartola-scout/arquivos/scouts/rodada-%d.json", rodadaAlvo)
	os.WriteFile(arquivoRodada, j, 0644)
}

func incrementaEstatisticas(estatisticas *model.Estatisticas, pontuacao model.PontuacaoJogador) {
	estatisticas.Jogos++
	estatisticas.TotalPontos += pontuacao.Valor
	estatisticas.Media = estatisticas.TotalPontos / float32(estatisticas.Jogos)
}

func carregarClubes() map[int]model.Clube {
	clubes := []model.Clube{}
	mapaClubes := map[int]model.Clube{}
	dados, _ := os.ReadFile("/Users/cicerohamonvitorinolopes/workspace/cartola-scout/arquivos/dados/clubes.json")
	json.Unmarshal(dados, &clubes)

	for _, clube := range clubes {
		mapaClubes[clube.ID] = clube
	}

	return mapaClubes
}

func carregarPontuacao(rodada int) model.PontuacaoRodada {
	pontuacaoRodada := model.PontuacaoRodada{}
	dados, _ := os.ReadFile(fmt.Sprintf("/Users/cicerohamonvitorinolopes/workspace/cartola-scout/arquivos/pontuacao/rodada-%d.json", rodada))
	json.Unmarshal(dados, &pontuacaoRodada)

	return pontuacaoRodada
}

func carregarResultados(rodada int) []requests.Partida {
	resultadosRodada := requests.Resultados{}
	dados, _ := os.ReadFile(fmt.Sprintf("/Users/cicerohamonvitorinolopes/workspace/cartola-scout/arquivos/resultados/rodada-%d.json", rodada))
	json.Unmarshal(dados, &resultadosRodada)

	return resultadosRodada.Partidas
}

func download(numeroRodada int, nomeArquivo string, request requests.Request) {
	arquivo := fmt.Sprintf(nomeArquivo, numeroRodada)

	if arquivoExiste(arquivo) {
		return
	}

	responsePontuacao, _ := http.Get(request.Url)
	dataPontuacao, _ := io.ReadAll(responsePontuacao.Body)

	os.WriteFile(arquivo, dataPontuacao, 0644)
}

func arquivoExiste(nome string) bool {
	info, err := os.Stat(nome)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
