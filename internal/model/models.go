package model

type (
	Clube struct {
		ID     int
		Nome   string
		Classe int
	}

	Jogo struct {
		Rodada           int
		Mandante         Clube
		Visitante        Clube
		MandantePosicao  int
		VisitantePosicao int
		Pontuacao        float32
		EntrouEmCampo    bool
		Valido           bool
	}

	Jogador struct {
		ID      int
		Nome    string
		Posicao string
		Clube   Clube
		Jogos   []Jogo
	}

	Estatisticas struct {
		Jogos       int
		Media       float32
		Forma       float32
		TotalPontos float32
	}

	PontuacaoJogador struct {
		Apelido       string  `json:"apelido"`
		Valor         float32 `json:"pontuacao"`
		Posicao       int     `json:"posicao_id"`
		Clube         int     `json:"clube_id"`
		EntrouEmCampo bool    `json:"entrou_em_campo"`
	}

	PontuacaoRodada struct {
		Atletas map[string]PontuacaoJogador `json:"atletas"`
		Rodada  int                         `json:"rodada"`
	}

	ResultadoPartida struct {
		Visitante        int `json:"clube_visitante_id"`
		VisitantePosicao int `json:"clube_visitante_posicao"`
		Mandante         int `json:"clube_casa_id"`
		MandantePosicao  int `json:"clube_casa_posicao"`
	}

	Scout struct {
		Rodada            int
		Jogador           Jogador
		Geral             Estatisticas
		ComoMandante      Estatisticas
		ComoVisitante     Estatisticas
		ContraTimeMaior   Estatisticas
		ContraTimeIgual   Estatisticas
		ContraTimeMenor   Estatisticas
		ContraTimesAcima  Estatisticas
		ContraTimesAbaixo Estatisticas
	}
)
