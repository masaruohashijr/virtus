package models

import (
	"github.com/lib/pq"
)

var AppName = "Virtus"
var Ambiente = ""

type TipoPontuacao int

const (
	Manual TipoPontuacao = iota
	Calculada
	Ajustada
	Definitiva
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ColSpan struct {
	CicloId      int64 `json:"cicloId"`
	PilarId      int64 `json:"pilarId"`
	ComponenteId int64 `json:"componenteId"`
	Qtd          int   `json:"qtd"`
}

type Action struct {
	Order          int
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	OriginId       int64  `json:"originid"`
	Origin         string `json:"originName"`
	DestinationId  int64  `json:"destinationid"`
	Destination    string `json:"destinationName"`
	OtherThan      bool   `json:"otherthan"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CreatedAt      string `json:"createdAt"`
	C_CreatedAt    string `json:"c_createdAt"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
	Selected       bool
	Roles          []Role
}

type Activity struct {
	Order                int    `json:"order"`
	Id                   int64  `json:"id"`
	WorkflowId           int64  `json:"wid"`
	ActionId             int64  `json:"actionId"`
	ActionName           string `json:"actionName"`
	ExpirationActionId   int64  `json:"expActionId"`
	ExpirationActionName string `json:"expActionName"`
	ExpirationTimeDays   int    `json:"expTime"`
	CStartAt             string `json:"startAt"`
	CEndAt               string `json:"endAt"`
	CRoles               string `json:"roles"`
	CRoleNames           string `json:"roleNames"`
	Roles                []Role `json:"roles_array"`
	CFeatures            string `json:"features"`
	CFeatureNames        string `json:"featureNames"`
}

type Chamado struct {
	Order           int
	Id              int64  `json:"id"`
	Titulo          string `json:"titulo"`
	Descricao       string `json:"descricao"`
	Acompanhamento  string `json:"acompanhamento"`
	IniciaEm        string `json:"iniciaEm"`
	ProntoEm        string `json:"prontoEm"`
	TipoChamadoId   string `json:"tipoChamadoId"`
	PrioridadeId    string `json:"prioridadeId"`
	Estimativa      int    `json:"estimativa"`
	ResponsavelId   int64  `json:"responsavelId"`
	ResponsavelName string `json:"responsavelName"`
	RelatorId       int64  `json:"relatorId"`
	RelatorName     string `json:"relatorName"`
	AuthorId        int64  `json:"authorId"`
	AuthorName      string `json:"authorName"`
	CriadoEm        string `json:"criadoEm"`
	C_CriadoEm      string `json:"c_criadoEm"`
	IdVersaoOrigem  int64  `json:"idVersaoOrigem"`
	StatusId        int64  `json:"statusId"`
	CStatus         string `json:"cStatus"`
}

type Ciclo struct {
	Order          int
	Id             int64  `json:"id"`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	Referencia     string `json:"referencia"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CriadoEm       string `json:"criadoEm"`
	C_CriadoEm     string `json:"c_criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type CicloEntidade struct {
	Order          int
	Id             int64  `json:"id"`
	EntidadeId     int64  `json:"entidadeId"`
	CicloId        int64  `json:"cicloId"`
	Nome           string `json:"nome"`
	TipoMediaId    int    `json:"tipoMediaId"`
	TipoMedia      string `json:"tipoMedia"`
	IniciaEm       string `json:"iniciaEm"`
	TerminaEm      string `json:"terminaEm"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type Componente struct {
	Order          int
	Id             int64  `json:"id"`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	Referencia     string `json:"referencia"`
	PGA            string `json:"pga"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CriadoEm       string `json:"criadoEm"`
	C_CriadoEm     string `json:"c_criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type ComponentePilar struct {
	Order          int
	Id             int64  `json:"id"`
	PilarId        int64  `json:"pilarId"`
	ComponenteId   int64  `json:"componenteId"`
	ComponenteNome string `json:"componenteNome"`
	TipoMediaId    int    `json:"tipoMediaId"`
	TipoMedia      string `json:"tipoMedia"`
	Sonda          string `json:"sonda"`
	PesoPadrao     string `json:"pesoPadrao"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type Descricao struct {
	Texto string `json:"texto"`
	Link  string `json:"link"`
}

type Elemento struct {
	Order          int
	Id             int64  `json:"id"`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	Referencia     string `json:"referencia"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	Peso           int    `json:"weight"`
	CriadoEm       string `json:"criadoEm"`
	C_CriadoEm     string `json:"c_criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type ElementoComponente struct {
	Order          int
	Id             int64  `json:"id"`
	ComponenteId   int64  `json:"componenteId"`
	ElementoId     int64  `json:"elementoId"`
	ElementoNome   string `json:"elementoNome"`
	TipoNotaId     int64  `json:"tipoNotaId"`
	TipoNotaNome   string `json:"tipoNotaNome"`
	PesoPadrao     int    `json:"pesoPadrao"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type Entidade struct {
	Order                 int
	Id                    int64           `json:"id"`
	Nome                  string          `json:"nome"`
	Descricao             string          `json:"descricao"`
	Sigla                 string          `json:"sigla"`
	Codigo                string          `json:"codigo"`
	Situacao              string          `json:"situacao"`
	ESI                   bool            `json:"esi"`
	Municipio             string          `json:"municipio"`
	SiglaUF               string          `json:"siglaUF"`
	EscritorioId          int64           `json:"escritorioId"`
	EscritorioAbreviatura string          `json:"escritorioAbreviatura"`
	CicloNome             string          `json:"cicloNome"`
	ChefeId               int64           `json:"chefeId"`
	ChefeName             string          `json:"chefeName"`
	Escritorio            string          `json:"escritorio"`
	AuthorId              int64           `json:"authorId"`
	AuthorName            string          `json:"authorName"`
	CriadoEm              string          `json:"criadoEm"`
	C_CriadoEm            string          `json:"c_criadoEm"`
	IdVersaoOrigem        int64           `json:"idVersaoOrigem"`
	StatusId              int64           `json:"statusId"`
	CStatus               string          `json:"cStatus"`
	CiclosEntidade        []CicloEntidade `json:"ciclos"`
}

type Escritorio struct {
	Order          int
	Id             int64  `json:"id"`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	Abreviatura    string `json:"abreviatura"`
	ChefeId        int64  `json:"chefeId"`
	ChefeNome      string `json:"chefeNome"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CriadoEm       string `json:"criadoEm"`
	C_CriadoEm     string `json:"c_criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
	Selected       bool   `json:"selected"`
}

type Feature struct {
	Order          int
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Code           string `json:"code"`
	Description    string `json:"description"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CreatedAt      string `json:"createdAt"`
	C_CreatedAt    string `json:"c_createdAt"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
	Selected       bool
}

type Historico struct {
	Id                string `json:"id"`
	EntidadeId        string `json:"entidadeId"`
	CicloId           string `json:"cicloId"`
	PilarId           string `json:"pilarId"`
	PlanoId           string `json:"planoId"`
	ComponenteId      string `json:"componenteId"`
	ElementoId        string `json:"elementoId"`
	Nota              string `json:"nota"`
	Metodo            string `json:"metodo"`
	Peso              string `json:"peso"`
	IniciaEm          string `json:"iniciaEm"`
	IniciaEmAnterior  string `json:"iniciaEmAnterior"`
	TerminaEm         string `json:"terminaEm"`
	TerminaEmAnterior string `json:"terminaEmAnterior"`
	Config            string `json:"config"`
	ConfigAnterior    string `json:"configAnterior"`
	AuditorNovoId     string `json:"auditorNovoId"`
	AuditorAnteriorId string `json:"auditorAnteriorId"`
	SupervisorId      string `json:"supervisorId"`
	AutorId           string `json:"autorId"`
	AutorNome         string `json:"autorNome"`
	AlteradoEm        string `json:"alteradoEm"`
	Motivacao         string `json:"motivacao"`
	TipoAlteracao     string `json:"tipoAlteracao"`
}

type Item struct {
	Order          int    `json:"order"`
	Id             int64  `json:"id"`
	ElementoId     int64  `json:"elementoId"`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	Referencia     string `json:"referencia"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	C_CriadoEm     string `json:"c_criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"status"`
	CStatus        string `json:"cStatus"`
}

type Jurisdicao struct {
	Order          int
	Id             int64  `json:"id"`
	EscritorioId   int64  `json:"escritorioId"`
	EntidadeId     int64  `json:"entidadeId"`
	EntidadeNome   string `json:"entidadeNome"`
	EntidadeSigla  string `json:"entidadeSigla"`
	IniciaEm       string `json:"iniciaEm"`
	TerminaEm      string `json:"terminaEm"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type LoggedUser struct {
	User          User
	HasPermission func(string) bool
}

type Membro struct {
	Order          int
	Id             int64  `json:"id"`
	EscritorioId   int64  `json:"escritorioId"`
	UsuarioId      int64  `json:"usuarioId"`
	UsuarioNome    string `json:"usuarioNome"`
	UsuarioPerfil  string `json:"usuarioPerfil"`
	SubordinacaoId int64  `json:"subordinacaoId"`
	IniciaEm       string `json:"iniciaEm"`
	TerminaEm      string `json:"terminaEm"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type Integrante struct {
	Order          int
	Id             int64  `json:"id"`
	EntidadeId     int64  `json:"entidadeId"`
	CicloId        int64  `json:"cicloId"`
	UsuarioId      int64  `json:"usuarioId"`
	UsuarioNome    string `json:"usuarioNome"`
	UsuarioPerfil  string `json:"usuarioPerfil"`
	IniciaEm       string `json:"iniciaEm"`
	TerminaEm      string `json:"terminaEm"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type NotasAtuais struct {
	CicloNota      string `json:"cicloNota"`
	PilarNota      string `json:"pilarNota"`
	ComponenteNota string `json:"componenteNota"`
	PlanoNota      string `json:"planoNota"`
	TipoNotaNota   string `json:"tipoNotaNota"`
}

type NullTime struct {
	pq.NullTime
}

type ConfigPlano struct {
	Id         int64 `json:"id"`
	EntidadeId int64 `json:"entidadeId"`
	PlanoId    int64 `json:"planoId"`
}

type PesosAtuais struct {
	CicloPeso      string `json:"cicloPeso"`
	PilarPeso      string `json:"pilarPeso"`
	ComponentePeso string `json:"componentePeso"`
	PlanoPeso      string `json:"planoPeso"`
	TipoNotaPeso   string `json:"tipoNotaPeso"`
	ElementoPeso   string `json:"elementoPeso"`
}

type ValoresAtuais struct {
	CicloPeso        string `json:"cicloPeso"`
	PilarPeso        string `json:"pilarPeso"`
	ComponentePeso   string `json:"componentePeso"`
	PlanoPeso        string `json:"planoPeso"`
	TipoNotaPeso     string `json:"tipoNotaPeso"`
	ElementoPeso     string `json:"elementoPeso"`
	CicloNota        string `json:"cicloNota"`
	PilarNota        string `json:"pilarNota"`
	ComponenteNota   string `json:"componenteNota"`
	PlanoNota        string `json:"planoNota"`
	TipoNotaNota     string `json:"tipoNotaNota"`
	ComponenteStatus string `json:"componenteStatus"`
}

type Plano struct {
	Order               int    `json:"order"`
	Id                  int64  `json:"id"`
	Nome                string `json:"nome"`
	Descricao           string `json:"descricao"`
	EntidadeId          int64  `json:"entidadeId"`
	EntidadeNome        string `json:"entidadeNome"`
	CNPB                string `json:"cnpb"`
	C_RecursoGarantidor string `json:"c_recursoGarantidor"`
	RecursoGarantidor   string `json:"recursoGarantidor"`
	Modalidade          string `json:"modalidade"`
	AuthorId            int64  `json:"authorId"`
	AuthorName          string `json:"authorName"`
	CriadoEm            string `json:"criadoEm"`
	C_CriadoEm          string `json:"c_criadoEm"`
	IdVersaoOrigem      int64  `json:"idVersaoOrigem"`
	StatusId            int64  `json:"statusId"`
	CStatus             string `json:"cStatus"`
}

type Pilar struct {
	Order          int
	Id             int64  `json:"id"`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	Referencia     string `json:"referencia"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CriadoEm       string `json:"criadoEm"`
	C_CriadoEm     string `json:"c_criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type PilarCiclo struct {
	Order          int
	Id             int64  `json:"id"`
	CicloId        int64  `json:"cicloId"`
	PilarId        int64  `json:"pilarId"`
	PilarNome      string `json:"pilarNome"`
	TipoMediaId    int    `json:"tipoMediaId"`
	TipoMedia      string `json:"tipoMedia"`
	PesoPadrao     string `json:"pesoPadrao"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type Processo struct {
	Order          int
	Id             int64  `json:"id"`
	AnotacaoId     int64  `json:"anotacaoId"`
	Numero         string `json:"numero"`
	Descricao      string `json:"descricao"`
	Referencia     string `json:"referencia"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CriadoEm       string `json:"criadoEm"`
	C_CriadoEm     string `json:"c_criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type ProdutoCiclo struct {
	Order           int
	Id              int64   `json:"id"`
	EntidadeId      int64   `json:"entidadeId"`
	EntidadeNome    int64   `json:"entidadeNome"`
	CicloId         int64   `json:"cicloId"`
	CicloNome       string  `json:"cicloNome"`
	TipoPontuacaoId int     `json:"tipoPontuacaoId"`
	TipoPontuacao   string  `json:"tipoPontuacao"`
	Nota            float64 `json:"nota"`
	Motivacao       string  `json:"motivacao"`
	AuditorId       int64   `json:"auditorId"`
	AuditorName     string  `json:"auditorName"`
	SupervisorId    int64   `json:"supervisorId"`
	SupervisorName  string  `json:"supervisorName"`
	AuthorId        int64   `json:"autorId"`
	AuthorName      string  `json:"autorNome"`
	CriadoEm        string  `json:"criadoEm"`
	IdVersaoOrigem  int64   `json:"idVersaoOrigem"`
	StatusId        int64   `json:"statusId"`
	CStatus         string  `json:"cStatus"`
}
type ProdutoPilar struct {
	Order           int
	Id              int64   `json:"id"`
	EntidadeId      int64   `json:"entidadeId"`
	EntidadeNome    int64   `json:"entidadeNome"`
	CicloId         int64   `json:"cicloId"`
	CicloNome       string  `json:"cicloNome"`
	PilarId         int64   `json:"pilarId"`
	PilarNome       string  `json:"pilarNome"`
	TipoPontuacaoId int     `json:"tipoPontuacaoId"`
	TipoPontuacao   string  `json:"tipoPontuacao"`
	Peso            float64 `json:"peso"`
	Nota            float64 `json:"nota"`
	Motivacao       string  `json:"motivacao"`
	AuditorId       int64   `json:"auditorId"`
	AuditorName     string  `json:"auditorName"`
	SupervisorId    int64   `json:"supervisorId"`
	SupervisorName  string  `json:"supervisorName"`
	AuthorId        int64   `json:"autorId"`
	AuthorName      string  `json:"autorNome"`
	CriadoEm        string  `json:"criadoEm"`
	IdVersaoOrigem  int64   `json:"idVersaoOrigem"`
	StatusId        int64   `json:"statusId"`
	CStatus         string  `json:"cStatus"`
}

type ProdutoComponente struct {
	Order             int
	Id                int64   `json:"id"`
	EntidadeId        int64   `json:"entidadeId"`
	EntidadeNome      string  `json:"entidadeNome"`
	CicloId           int64   `json:"cicloId"`
	CicloNome         string  `json:"cicloNome"`
	PilarId           int64   `json:"pilarId"`
	PilarNome         string  `json:"pilarNome"`
	PlanoId           int64   `json:"planoId"`
	CNPB              string  `json:"cnpb"`
	ComponenteId      int64   `json:"componenteId"`
	ComponenteNome    string  `json:"componenteNome"`
	TipoPontuacaoId   int     `json:"tipoPontuacaoId"`
	TipoPontuacao     string  `json:"tipoPontuacao"`
	Nota              float64 `json:"nota"`
	Motivacao         string  `json:"motivacao"`
	AuditorId         int64   `json:"auditorId"`
	AuditorName       string  `json:"auditorName"`
	AuditorAnteriorId int64   `json:"auditorAnteriorId"`
	SupervisorId      int64   `json:"supervisorId"`
	SupervisorName    string  `json:"supervisorName"`
	IniciaEm          string  `json:"iniciaEm"`
	TerminaEm         string  `json:"terminaEm"`
	IniciaEmAnterior  string  `json:"iniciaEmAnterior"`
	TerminaEmAnterior string  `json:"terminaEmAnterior"`
	AuthorId          int64   `json:"autorId"`
	AuthorName        string  `json:"autorNome"`
	CriadoEm          string  `json:"criadoEm"`
	IdVersaoOrigem    int64   `json:"idVersaoOrigem"`
	StatusId          int64   `json:"statusId"`
	CStatus           string  `json:"cStatus"`
	CAction           string  `json:"CAction"`
	Configurado       string  `json:"configurado"`
	SomentePGA        string  `json:"somentePGA"`
}

type ProdutoPlano struct {
	Order           int
	Id              int64   `json:"id"`
	EntidadeId      int64   `json:"entidadeId"`
	EntidadeNome    string  `json:"entidadeNome"`
	CicloId         int64   `json:"cicloId"`
	CicloNome       string  `json:"cicloNome"`
	PilarId         int64   `json:"pilarId"`
	PilarNome       string  `json:"pilarNome"`
	ComponenteId    int64   `json:"componenteId"`
	ComponenteNome  string  `json:"componenteNome"`
	PlanoId         int64   `json:"planoId"`
	PlanoCNPB       string  `json:"planoCNPB"`
	PlanoModalidade string  `json:"planoModalidade"`
	TipoPontuacaoId int     `json:"tipoPontuacaoId"`
	TipoPontuacao   string  `json:"tipoPontuacao"`
	Nota            float64 `json:"nota"`
	Motivacao       string  `json:"motivacao"`
	AuditorId       int64   `json:"auditorId"`
	AuditorName     string  `json:"auditorName"`
	SupervisorId    int64   `json:"supervisorId"`
	SupervisorName  string  `json:"supervisorName"`
	AuthorId        int64   `json:"autorId"`
	AuthorName      string  `json:"autorNome"`
	CriadoEm        string  `json:"criadoEm"`
	IdVersaoOrigem  int64   `json:"idVersaoOrigem"`
	StatusId        int64   `json:"statusId"`
	CStatus         string  `json:"cStatus"`
}

type ProdutoElemento struct {
	Order           int
	Id              int64   `json:"id"`
	EntidadeId      int64   `json:"entidadeId"`
	EntidadeNome    string  `json:"entidadeNome"`
	CicloId         int64   `json:"cicloId"`
	CicloNome       string  `json:"cicloNome"`
	PilarId         int64   `json:"pilarId"`
	PilarNome       string  `json:"pilarNome"`
	PlanoId         int64   `json:"planoId"`
	CNPB            string  `json:"cnpb"`
	ComponenteId    int64   `json:"componenteId"`
	ComponenteNome  string  `json:"componenteNome"`
	ElementoId      int64   `json:"elementoId"`
	ElementoNome    string  `json:"elementoNome"`
	TipoNotaId      int64   `json:"tipoNotaId"`
	TipoNotaNome    string  `json:"tipoNotaNome"`
	TipoPontuacaoId int     `json:"tipoPontuacaoId"`
	TipoPontuacao   string  `json:"tipoPontuacao"`
	Peso            float64 `json:"peso"`
	Nota            int     `json:"nota"`
	Motivacao       string  `json:"motivacao"`
	AuditorId       int64   `json:"auditorId"`
	AuditorName     string  `json:"auditorName"`
	SupervisorId    int64   `json:"supervisorId"`
	SupervisorName  string  `json:"supervisorName"`
	AuthorId        int64   `json:"autorId"`
	AuthorName      string  `json:"autorNome"`
	CriadoEm        string  `json:"criadoEm"`
	IdVersaoOrigem  int64   `json:"idVersaoOrigem"`
	StatusId        int64   `json:"statusId"`
	CStatus         string  `json:"cStatus"`
}

type ProdutoItem struct {
	Order               int
	Id                  int64  `json:"id"`
	EntidadeId          int64  `json:"entidadeId"`
	EntidadeNome        string `json:"entidadeNome"`
	CicloId             int64  `json:"cicloId"`
	CicloNome           string `json:"cicloNome"`
	CicloNota           string `json:"cicloNota"`
	PilarId             int64  `json:"pilarId"`
	PilarNome           string `json:"pilarNome"`
	PilarPeso           string `json:"pilarPeso"`
	PilarNota           string `json:"pilarNota"`
	ComponenteId        int64  `json:"componenteId"`
	ComponenteNome      string `json:"componenteNome"`
	ComponentePeso      string `json:"componentePeso"`
	ComponenteNota      string `json:"componenteNota"`
	PlanoId             int64  `json:"planoId"`
	CNPB                string `json:"cnpb"`
	RecursoGarantidor   string `json:"recursoGarantidor"`
	PlanoPeso           string `json:"planoPeso"`
	PlanoNota           string `json:"planoNota"`
	PlanoModalidade     string `json:"planoModalidade"`
	ElementoId          int64  `json:"elementoId"`
	ElementoNome        string `json:"elementoNome"`
	ElementoPeso        string `json:"elementoPeso"`
	ElementoNota        string `json:"elementoNota"`
	TipoPontuacaoId     string `json:"tipoPontuacaoId"`
	TipoNotaId          int    `json:"tipoNotaId"`
	TipoNotaNome        string `json:"tipoNotaNome"`
	TipoNotaLetra       string `json:"tipoNotaLetra"`
	TipoNotaCorLetra    string `json:"tipoNotaCorLetra"`
	TipoNotaPeso        string `json:"tipoNotaPeso"`
	TipoNotaNota        string `json:"tipoNotaNota"`
	PesoPadraoEC        string `json:"pesoPadraoEC"`
	TipoMediaCPId       int    `json:"tipoMediaCPId"`
	TipoMediaCP         string `json:"tipoMediaCP"`
	PesoPadraoCP        string `json:"pesoPadraoCP"`
	TipoMediaPCId       int    `json:"tipoMediaPCId"`
	TipoMediaPC         string `json:"tipoMediaPC"`
	PesoPadraoPC        string `json:"pesoPadraoPC"`
	TipoMediaCEId       int    `json:"tipoMediaCEId"`
	TipoMediaCE         string `json:"tipoMediaCE"`
	IniciaEm            string `json:"iniciaEm"`
	TerminaEm           string `json:"terminaEm"`
	ItemId              int64  `json:"itemId"`
	ItemNome            string `json:"itemNome"`
	AuditorId           int64  `json:"auditorId"`
	AuditorName         string `json:"auditorName"`
	SupervisorId        int64  `json:"supervisorId"`
	SupervisorName      string `json:"supervisorName"`
	AuthorId            int64  `json:"autorId"`
	AuthorName          string `json:"autorNome"`
	CriadoEm            string `json:"criadoEm"`
	IdVersaoOrigem      int64  `json:"idVersaoOrigem"`
	StatusId            int64  `json:"statusId"`
	CStatus             string `json:"cStatus"`
	ProdutoComponenteId string `json:"produtoComponenteId"`
	PeriodoPermitido    bool   `json:"periodoPermitido"`
	PeriodoCiclo        bool   `json:"periodoCiclo"`
}

type ElementoDaMatriz struct {
	Order                   int
	Id                      int64  `json:"id"`
	EntidadeId              int64  `json:"entidadeId"`
	EntidadeNome            string `json:"entidadeNome"`
	EntidadeQtdPlanos       int    `json:"entidadeQtdPlanos"`
	EntidadeRowspan         int    `json:"entidadeRowspan"`
	PlanoId                 int64  `json:"planoId"`
	CNPB                    string `json:"cnpb"`
	RecursoGarantidor       string `json:"recursoGarantidor"`
	Modalidade              string `json:"modalidade"`
	CicloId                 int64  `json:"cicloId"`
	CicloNome               string `json:"cicloNome"`
	CicloNota               string `json:"cicloNota"`
	CicloQtdPilares         string `json:"cicloQtdPilares"`
	CicloColSpan            int    `json:"cicloColSpan"`
	PilarId                 int64  `json:"pilarId"`
	PilarNome               string `json:"pilarNome"`
	PilarPeso               string `json:"pilarPeso"`
	PilarNota               string `json:"pilarNota"`
	PilarColSpan            int    `json:"pilarColSpan"`
	PilarQtdComponentes     string `json:"pilarQtdComponentes"`
	ComponenteQtdElementos  string `json:"componenteQtdElementos"`
	ComponenteQtdPlanos     string `json:"componenteQtdPlanos"`
	ComponenteId            int64  `json:"componenteId"`
	ComponenteNome          string `json:"componenteNome"`
	ComponentePeso          string `json:"componentePeso"`
	ComponenteNota          string `json:"componenteNota"`
	ComponenteColSpan       int    `json:"componenteColSpan"`
	ComponenteQtdTiposNotas string `json:"componenteQtdTiposNotas"`
	ElementoId              int64  `json:"elementoId"`
	ElementoNome            string `json:"elementoNome"`
	ElementoPeso            string `json:"elementoPeso"`
	ElementoNota            string `json:"elementoNota"`
	TipoNotaId              int    `json:"tipoNotaId"`
	TipoNotaNome            string `json:"tipoNotaNome"`
	TipoNotaLetra           string `json:"tipoNotaLetra"`
	TipoNotaCorLetra        string `json:"tipoNotaCorLetra"`
	TipoNotaPeso            string `json:"tipoNotaPeso"`
	TipoNotaNota            string `json:"tipoNotaNota"`
	PesoPadraoEC            string `json:"pesoPadraoEC"`
	TipoMediaCPId           int    `json:"tipoMediaCPId"`
	TipoMediaCP             string `json:"tipoMediaCP"`
	PesoPadraoCP            string `json:"pesoPadraoCP"`
	TipoMediaPCId           int    `json:"tipoMediaPCId"`
	TipoMediaPC             string `json:"tipoMediaPC"`
	PesoPadraoPC            string `json:"pesoPadraoPC"`
	TipoMediaCEId           int    `json:"tipoMediaCEId"`
	TipoMediaCE             string `json:"tipoMediaCE"`
	IniciaEm                string `json:"iniciaEm"`
	TerminaEm               string `json:"terminaEm"`
	ItemId                  int64  `json:"itemId"`
	ItemNome                string `json:"itemNome"`
	AuditorId               int64  `json:"auditorId"`
	AuditorName             string `json:"auditorName"`
	SupervisorId            int64  `json:"supervisorId"`
	SupervisorName          string `json:"supervisorName"`
	AutorNotaId             int64  `json:"autorNotaId"`
	AutorNotaName           string `json:"autorNotaNome"`
	MotivacaoPeso           string `json:"motivacaoPeso"`
	MotivacaoNota           string `json:"motivacaoNota"`
	CriadoEm                string `json:"criadoEm"`
	IdVersaoOrigem          int64  `json:"idVersaoOrigem"`
	StatusId                int64  `json:"statusId"`
	CStatus                 string `json:"cStatus"`
}

type Publicacao struct {
	Order          int
	Id             int64  `json:"id"`
	Mensagem       string `json:"mensagem"`
	IniciaEm       string `json:"iniciaEm"`
	TerminaEm      string `json:"terminaEm"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CriadoEm       string `json:"criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type AnotacaoRadar struct {
	Order                 int
	Id                    int64  `json:"id"`
	RadarId               int64  `json:"radarId"`
	EntidadeId            int64  `json:"entidadeId"`
	AnotacaoId            int64  `json:"anotacaoId"`
	Observacoes           string `json:"observacoes"`
	RegistroAta           string `json:"registroEmAta"`
	AuthorId              int64  `json:"autorId"`
	AuthorName            string `json:"autorNome"`
	CriadoEm              string `json:"criadoEm"`
	StatusId              int64  `json:"statusId"`
	CStatus               string `json:"cStatus"`
	UltimaAtualizacao     string `json:"ultimaAtualizacao"`
	UltimoAtualizadorId   int    `json:"ultimoAtualizador"`
	UltimoAtualizadorNome string `json:"ultimoAtualizadorNome"`
	IdVersaoOrigem        int64  `json:"idVersaoOrigem"`
}

type Anotacao struct {
	Order           int
	Id              int64  `json:"id"`
	EntidadeId      int    `json:"entidadeId"`
	EntidadeSigla   string `json:"entidadeSigla"`
	CicloId         int    `json:"cicloId"`
	PilarId         int    `json:"pilarId"`
	ComponenteId    int    `json:"componenteId"`
	PlanoId         int    `json:"planoId"`
	TipoNotaId      int    `json:"tipoNotaId"`
	ElementoId      int    `json:"elementoId"`
	ItemId          int    `json:"itemId"`
	Assunto         string `json:"assunto"`
	Descricao       string `json:"descricao"`
	Matriz          string `json:"matriz"`
	Risco           string `json:"risco"`
	Tendencia       string `json:"tendencia"`
	ResponsavelId   int64  `json:"responsevelId"`
	ResponsavelName string `json:"responsavelName"`
	RelatorId       int64  `json:"relatorId"`
	RelatorName     string `json:"relatorName"`
	AuthorId        int64  `json:"authorId"`
	AuthorName      string `json:"authorName"`
	CriadoEm        string `json:"criadoEm"`
	C_CriadoEm      string `json:"c_criadoEm"`
	IdVersaoOrigem  int64  `json:"idVersaoOrigem"`
	StatusId        int64  `json:"statusId"`
	CStatus         string `json:"cStatus"`
}

type Radar struct {
	Order          int
	Id             int64  `json:"id"`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	Referencia     string `json:"referencia"`
	DataRadar      string `json:"dataRadar"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CriadoEm       string `json:"criadoEm"`
	C_CriadoEm     string `json:"c_criadoEm"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
}

type Role struct {
	Order          int
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CreatedAt      string `json:"createdAt"`
	C_CreatedAt    string `json:"c_createdAt"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
	Selected       bool
	Features       []Feature
}

type Versao struct {
	Order           int
	Id              int64  `json:"id"`
	Nome            string `json:"nome"`
	Objetivo        string `json:"objetivo"`
	DefinicaoPronto string `json:"definicaoPronto"`
	IniciaEm        string `json:"iniciaEm"`
	TerminaEm       string `json:"terminaEm"`
	AuthorId        int64  `json:"authorId"`
	AuthorName      string `json:"authorName"`
	CriadoEm        string `json:"criadoEm"`
	C_CriadoEm      string `json:"c_criadoEm"`
	IdVersaoOrigem  int64  `json:"idVersaoOrigem"`
	StatusId        int64  `json:"statusId"`
	CStatus         string `json:"cStatus"`
}

type Status struct {
	Order          int
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Stereotype     string `json:"stereotype"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CreatedAt      string `json:"createdAt"`
	C_CreatedAt    string `json:"c_createdAt"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
	Selected       bool
}

type TipoNota struct {
	Order          int
	Id             int64  `json:"id"`
	TipoNotaId     string `json:"tipoNotaId"'`
	ComponenteId   string `json:"componenteId"'`
	Nome           string `json:"nome"`
	Descricao      string `json:"descricao"`
	Referencia     string `json:"referencia"`
	Letra          string `json:"letra"`
	CorLetra       string `json:"corLetra"`
	PesoPadrao     string `json:"pesoPadrao"'`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CreatedAt      string `json:"createdAt"`
	C_CreatedAt    string `json:"c_createdAt"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
	Selected       bool
}

type User struct {
	Order          int       `json:"order"`
	Id             int64     `json:"id"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	Mobile         string    `json:"mobile"`
	Escritorio     int64     `json:"escritorio"`
	EscritorioNome string    `json:"escritorioNome"`
	Role           int64     `json:"role"`
	RoleName       string    `json:"roleName"`
	AuthorId       int64     `json:"authorId"`
	AuthorName     string    `json:"authorName"`
	CriadoEm       string    `json:"criadoEm"`
	C_CriadoEm     string    `json:"c_criadoEm"`
	IdVersaoOrigem int64     `json:"idVersaoOrigem"`
	StatusId       int64     `json:"statusId"`
	CStatus        string    `json:"cStatus"`
	Features       []Feature `json:"features"`
	Selected       bool      `json:"selected"`
}

type Workflow struct {
	Order          int
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	EntityType     string `json:"entity"`
	StartAt        string `json:"startAt`
	EndAt          string `json:"endAt"`
	Description    string `json:"description"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CreatedAt      string `json:"createdAt"`
	C_CreatedAt    string `json:"c_createdAt"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
	Selected       bool
}

// PÁGINAS
type PageAbout struct {
	ErrMsg     string
	Msg        string
	AppName    string
	Title      string
	Chamados   []Chamado
	LoggedUser LoggedUser
}

type PageActions struct {
	ErrMsg     string
	Msg        string
	AppName    string
	Title      string
	Statuss    []Status
	Actions    []Action
	LoggedUser LoggedUser
}

type PageChamados struct {
	ErrMsg       string
	Msg          string
	AppName      string
	Title        string
	Chamados     []Chamado
	Relatores    []User
	Responsaveis []User
	LoggedUser   LoggedUser
}

type PageCiclos struct {
	ErrMsg     string
	Msg        string
	AppName    string
	Title      string
	Ciclos     []Ciclo
	Entidades  []Entidade
	Pilares    []Pilar
	Users      []User
	LoggedUser LoggedUser
}

type PageComponentes struct {
	ErrMsg      string
	Msg         string
	AppName     string
	Title       string
	TiposNota   []TipoNota
	Componentes []Componente
	Elementos   []Elemento
	Users       []User
	LoggedUser  LoggedUser
}

type PageElementos struct {
	ErrMsg     string
	Msg        string
	AppName    string
	Title      string
	Elementos  []Elemento
	Users      []User
	LoggedUser LoggedUser
}

type PageEntidades struct {
	ErrMsg     string
	Msg        string
	AppName    string
	Title      string
	Escritorio string
	Entidades  []Entidade
	Ciclos     []Ciclo
	Planos     []Plano
	Users      []User
	LoggedUser LoggedUser
}

type PageChefe struct {
	ErrMsg              string
	Msg                 string
	AppName             string
	Title               string
	Escritorio          string
	Pendencias          []ProdutoComponente
	EntidadesPermanente []Entidade
	DemaisEntidades     []Entidade
	Ciclos              []Ciclo
	Planos              []Plano
	Users               []User
	LoggedUser          LoggedUser
}

type PageFeatures struct {
	ErrMsg     string
	Msg        string
	AppName    string
	Title      string
	Features   []Feature
	LoggedUser LoggedUser
}

type PageEscritorios struct {
	ErrMsg      string
	Msg         string
	AppName     string
	Title       string
	Escritorios []Escritorio
	Entidades   []Entidade
	Users       []User
	LoggedUser  LoggedUser
}

type PageEntidadesCiclos struct {
	WarnMsg      string
	ErrMsg       string
	Msg          string
	AppName      string
	Title        string
	Entidades    []Entidade
	Membros      []Membro
	Supervisores []User
	LoggedUser   LoggedUser
}

type PageLogin struct {
	WarnMsg string
	ErrMsg  string
	Msg     string
	AppName string
	Title   string
}

type PageMatriz struct {
	ErrMsg                  string
	Msg                     string
	AppName                 string
	Title                   string
	ComponenteQtdTiposNotas string
	ElementosDaMatriz       []ElementoDaMatriz
	Supervisores            []User
	Auditores               []User
	LoggedUser              LoggedUser
	Dec                     func(i int) int
	Inc                     func(i int) int
	MulTxt                  func(i int, j string) int
	Mul                     func(i int, j int) int
	SomarTxt                func(i int, j string) int
	Somar                   func(i int, j int) int
}

type PageProcessos struct {
	ErrMsg     string
	Msg        string
	AppName    string
	Title      string
	Processos  []Processo
	LoggedUser LoggedUser
}

type PageProdutosItens struct {
	ErrMsg       string
	Msg          string
	AppName      string
	Title        string
	PilarId      string
	ComponenteId string
	Inc          func(i int) int
	Produtos     []ProdutoItem
	Supervisores []User
	Auditores    []User
	LoggedUser   LoggedUser
}

type PageProdutosComponentes struct {
	WarnMsg      string
	ErrMsg       string
	Msg          string
	AppName      string
	Title        string
	Produtos     []ProdutoComponente
	Planos       []Plano
	Supervisores []User
	Auditores    []User
	LoggedUser   LoggedUser
}

type PagePilares struct {
	ErrMsg      string
	Msg         string
	AppName     string
	Title       string
	Pilares     []Pilar
	Componentes []Componente
	Users       []User
	LoggedUser  LoggedUser
}

type PagePlanos struct {
	ErrMsg     string
	Msg        string
	AppName    string
	Title      string
	Planos     []Plano
	Users      []User
	LoggedUser LoggedUser
}

type PageAnotacoes struct {
	ErrMsg       string
	Msg          string
	AppName      string
	Title        string
	Entidades    []Entidade
	Anotacoes    []Anotacao
	Relatores    []User
	Responsaveis []User
	LoggedUser   LoggedUser
}

type PageRadares struct {
	ErrMsg     string
	Msg        string
	AppName    string
	Title      string
	Radares    []Radar
	Anotacoes  []Anotacao
	LoggedUser LoggedUser
}

type PageRoles struct {
	ErrMsg     string
	Msg        string
	AppName    string
	Title      string
	Roles      []Role
	Features   []Feature
	LoggedUser LoggedUser
}

type PageSobre struct {
	ErrMsg      string
	Msg         string
	AppName     string
	Title       string
	Versoes     []Versao
	Publicacoes []Publicacao
	LoggedUser  LoggedUser
}

type PageVersoes struct {
	ErrMsg     string
	Msg        string
	AppName    string
	Title      string
	Versoes    []Versao
	LoggedUser LoggedUser
}

type PageStatus struct {
	ErrMsg     string
	Msg        string
	AppName    string
	Title      string
	Statuss    []Status
	LoggedUser LoggedUser
}

type PageTiposNotas struct {
	ErrMsg     string
	Msg        string
	AppName    string
	Title      string
	TiposNotas []TipoNota
	Users      []User
	LoggedUser LoggedUser
}

type PageUsers struct {
	WarnMsg     string
	ErrMsg      string
	Msg         string
	AppName     string
	Title       string
	Users       []User
	Escritorios []Escritorio
	Roles       []Role
	LoggedUser  LoggedUser
}

type PageWorkflows struct {
	ErrMsg     string
	Msg        string
	AppName    string
	Title      string
	Features   []Feature
	Actions    []Action
	Roles      []Role
	Workflows  []Workflow
	LoggedUser LoggedUser
}
