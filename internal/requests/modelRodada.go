package requests

type (
	Jogador struct {
		Scout struct {
			Ds int `json:"DS"`
			Fc int `json:"FC"`
			Ff int `json:"FF"`
		} `json:"scout"`
		Apelido       string  `json:"apelido"`
		Foto          string  `json:"foto"`
		Pontuacao     float64 `json:"pontuacao"`
		PosicaoID     int     `json:"posicao_id"`
		ClubeID       int     `json:"clube_id"`
		EntrouEmCampo bool    `json:"entrou_em_campo"`
	}

	Posicao struct {
		ID         int    `json:"id"`
		Nome       string `json:"nome"`
		Abreviacao string `json:"abreviacao"`
	}

	Clube struct {
		Escudos struct {
			Six0X60   string `json:"60x60"`
			Four5X45  string `json:"45x45"`
			Three0X30 string `json:"30x30"`
		} `json:"escudos"`
		Nome         string `json:"nome"`
		Abreviacao   string `json:"abreviacao"`
		Slug         string `json:"slug"`
		Apelido      string `json:"apelido"`
		NomeFantasia string `json:"nome_fantasia"`
		ID           int    `json:"id"`
		URLEditoria  string `json:"url_editoria"`
	}

	Rodada struct {
		Atletas      map[string]Jogador `json:"atletas"`
		Clubes       map[string]Clube   `json:"clubes"`
		Posicoes     map[string]Posicao `json:"posicoes"`
		Numero       int                `json:"rodada"`
		TotalAtletas int                `json:"total_atletas"`
	}

	Partida struct {
		AproveitamentoVisitante []string `json:"aproveitamento_visitante"`
		AproveitamentoMandante  []string `json:"aproveitamento_mandante"`
		Transmissao             struct {
			Label string `json:"label"`
			URL   string `json:"url"`
		} `json:"transmissao"`
		Local                  string `json:"local"`
		StatusTransmissaoTr    string `json:"status_transmissao_tr"`
		StatusCronometroTr     string `json:"status_cronometro_tr"`
		PeriodoTr              string `json:"periodo_tr"`
		PartidaData            string `json:"partida_data"`
		InicioCronometroTr     string `json:"inicio_cronometro_tr"`
		PlacarOficialVisitante int    `json:"placar_oficial_visitante"`
		PlacarOficialMandante  int    `json:"placar_oficial_mandante"`
		PartidaID              int    `json:"partida_id"`
		ClubeVisitantePosicao  int    `json:"clube_visitante_posicao"`
		ClubeVisitanteID       int    `json:"clube_visitante_id"`
		ClubeCasaPosicao       int    `json:"clube_casa_posicao"`
		ClubeCasaID            int    `json:"clube_casa_id"`
		Timestamp              int    `json:"timestamp"`
		CampeonatoID           int    `json:"campeonato_id"`
		Valida                 bool   `json:"valida"`
	}

	Resultados struct {
		Partidas []Partida `json:"partidas"`
	}
)
