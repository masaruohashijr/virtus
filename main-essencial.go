package main

import (
	"database/sql"
	"fmt"

	"github.com/BurntSushi/toml"
	_ "github.com/alexbrainman/odbc"

	//_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	//_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
	dpk "virtus-essencial/db"
	hd "virtus-essencial/handlers"
	mdl "virtus-essencial/models"
	route "virtus-essencial/routes"
	sec "virtus-essencial/security"
)

var db *sql.DB
var dbConfig Config
var eConfig Config
var sConfig Config

type Config struct {
	DbServer          string
	DbPort            string
	DbUser            string
	DbPassword        string
	DbName            string
	DbMaxOpenConns    string
	DbMaxIdleConns    string
	DbConnMaxLifetime string
	DbDriver          string
	DbDSN             string
	Ambiente          string
	EncryptionKey     string
	ServerPort        string
	EmailFrom         string
	EmailPassword     string
	EmailSubject      string
	SmtpHost          string
	SmtpPort          string
}

type ConfigType int

const (
	SERVER ConfigType = iota
	EMAIL
	BANCO
)

func dbConn() *sql.DB {
	var dbase *sql.DB
	dbConfig := ReadConfig(BANCO)
	//port, _ := strconv.Atoi(dbConfig.DbPort)
	//connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", dbConfig.DbServer, dbConfig.DbUser, dbConfig.DbPassword, port, dbConfig.DbName)
	connString := fmt.Sprintf("%s;UID=%s;PWD=%s", dbConfig.DbDSN, dbConfig.DbUser, dbConfig.DbPassword)
	log.Println(connString)
	//dbase, _ = sql.Open(dbConfig.DbDriver, connString)
	dbase, _ = sql.Open("odbc", connString)
	dbase.SetMaxOpenConns(20)
	dbase.SetMaxIdleConns(20)
	dbase.SetConnMaxLifetime(10 * time.Minute)
	err := dbase.Ping()
	if err != nil {
		log.Fatal("VIRTUS >>>>> : " + err.Error())
	} else {
		log.Printf("VIRTUS >>>>> CONECTADO COM SUCESSO!\n")
	}
	return dbase
}

func main() {
	sConfig := ReadConfig(SERVER)
	sec.Store = sessions.NewCookieStore([]byte(sConfig.EncryptionKey))
	hd.Db = dbConn()
	mdl.Ambiente = " [" + sConfig.Ambiente + " 1.3.96 init" + "]"
	mdl.AppName += mdl.Ambiente
	// injetando a variável Authenticated
	if true {
		dpk.Initialize()
		// dpk.MigracaoCiclos()
	}
	r := mux.NewRouter()
	// ----------------- SECURITY
	r.HandleFunc("/", hd.IndexHandler).Methods("GET")
	r.HandleFunc("/login", hd.LoginHandler)
	r.HandleFunc("/logout", hd.LogoutHandler)
	// ----------------- HOME
	r.HandleFunc("/about", hd.ListSobreHandler)
	r.HandleFunc("/admin", hd.AdminHomeHandler)
	r.HandleFunc("/chefe", hd.ChefeHomeHandler)
	r.HandleFunc("/supervisor", hd.SupervisorHomeHandler)
	r.HandleFunc("/auditor", hd.AuditorHomeHandler)
	r.HandleFunc("/visualizador", hd.VisualizadorHomeHandler)
	r.HandleFunc("/desenvolvedor", hd.DesenvolvedorHomeHandler)
	// ----------------- AVALIAÇÕES
	r.HandleFunc("/listAvaliarPlanos", hd.ListAvaliarPlanosHandler).Methods("GET")
	r.HandleFunc("/avaliarPlanos", hd.AvaliarPlanosHandler).Methods("POST")
	r.HandleFunc("/updateAvaliarPlanos", hd.UpdateAvaliarPlanosHandler).Methods("POST")
	// ----------------- CHAMADOS
	r.HandleFunc(route.ChamadosRoute, hd.ListChamadosHandler).Methods("GET")
	r.HandleFunc("/createChamado", hd.CreateChamadoHandler).Methods("POST")
	r.HandleFunc("/updateChamado", hd.UpdateChamadoHandler).Methods("POST")
	r.HandleFunc("/deleteChamado", hd.DeleteChamadoHandler).Methods("POST")
	// ----------------- MATRIZES
	r.HandleFunc("/listMatrizes", hd.ListMatrizesHandler).Methods("GET")
	r.HandleFunc("/executarMatriz", hd.ExecutarMatrizHandler).Methods("POST")
	// ----------------- EQUIPES
	r.HandleFunc("/listDistribuirAtividades", hd.ListDistribuirAtividadesHandler).Methods("GET")
	r.HandleFunc("/distribuirAtividades", hd.DistribuirAtividadesHandler).Methods("POST")
	r.HandleFunc("/updateDistribuirAtividades", hd.UpdateDistribuirAtividadesHandler).Methods("POST")
	// ----------------- ACTIONS
	r.HandleFunc(route.ActionsRoute, hd.ListActionsHandler).Methods("GET")
	r.HandleFunc("/createAction", hd.CreateActionHandler).Methods("POST")
	r.HandleFunc("/updateAction", hd.UpdateActionHandler).Methods("POST")
	r.HandleFunc("/deleteAction", hd.DeleteActionHandler).Methods("POST")
	// ----------------- CICLOS
	r.HandleFunc(route.CiclosRoute, hd.ListCiclosHandler).Methods("GET")
	r.HandleFunc("/iniciarCiclo", hd.IniciarCicloHandler).Methods("POST")
	r.HandleFunc("/createCiclo", hd.CreateCicloHandler).Methods("POST")
	r.HandleFunc("/updateCiclo", hd.UpdateCicloHandler).Methods("POST")
	r.HandleFunc("/deleteCiclo", hd.DeleteCicloHandler).Methods("POST")
	// ----------------- COMPONENTES
	r.HandleFunc(route.ComponentesRoute, hd.ListComponentesHandler).Methods("GET")
	r.HandleFunc("/createComponente", hd.CreateComponenteHandler).Methods("POST")
	r.HandleFunc("/updateComponente", hd.UpdateComponenteHandler).Methods("POST")
	r.HandleFunc("/deleteComponente", hd.DeleteComponenteHandler).Methods("POST")
	// ----------------- ELEMENTOS
	r.HandleFunc(route.ElementosRoute, hd.ListElementosHandler).Methods("GET")
	r.HandleFunc("/createElemento", hd.CreateElementoHandler).Methods("POST")
	r.HandleFunc("/updateElemento", hd.UpdateElementoHandler).Methods("POST")
	r.HandleFunc("/deleteElemento", hd.DeleteElementoHandler).Methods("POST")
	// ----------------- ENTIDADES
	r.HandleFunc(route.EntidadesRoute, hd.ListEntidadesHandler).Methods("GET")
	r.HandleFunc("/createEntidade", hd.CreateEntidadeHandler).Methods("POST")
	r.HandleFunc("/updateEntidade", hd.UpdateEntidadeHandler).Methods("POST")
	r.HandleFunc("/deleteEntidade", hd.DeleteEntidadeHandler).Methods("POST")
	// ----------------- ESCRITORIOS
	r.HandleFunc(route.EscritoriosRoute, hd.ListEscritoriosHandler).Methods("GET")
	r.HandleFunc("/createEscritorio", hd.CreateEscritorioHandler).Methods("POST")
	r.HandleFunc("/updateEscritorio", hd.UpdateEscritorioHandler).Methods("POST")
	r.HandleFunc("/updateJurisdicao", hd.UpdateJurisdicaoHandler).Methods("POST")
	r.HandleFunc("/updateMembrosEscritorio", hd.UpdateMembrosEscritorioHandler).Methods("POST")
	r.HandleFunc("/deleteEscritorio", hd.DeleteEscritorioHandler).Methods("POST")
	// ----------------- EQUIPES
	r.HandleFunc("/listDesignarEquipes", hd.ListDesignarEquipesHandler).Methods("GET")
	r.HandleFunc("/updateDesignarEquipe", hd.UpdateDesignarEquipeHandler).Methods("POST")
	// ----------------- FEATURES
	r.HandleFunc(route.FeaturesRoute, hd.ListFeaturesHandler).Methods("GET")
	r.HandleFunc("/createFeature", hd.CreateFeatureHandler).Methods("POST")
	r.HandleFunc("/updateFeature", hd.UpdateFeatureHandler).Methods("POST")
	r.HandleFunc("/deleteFeature", hd.DeleteFeatureHandler).Methods("POST")
	// ----------------- PILARES
	r.HandleFunc(route.PilaresRoute, hd.ListPilaresHandler).Methods("GET")
	r.HandleFunc("/createPilar", hd.CreatePilarHandler).Methods("POST")
	r.HandleFunc("/updatePilar", hd.UpdatePilarHandler).Methods("POST")
	r.HandleFunc("/deletePilar", hd.DeletePilarHandler).Methods("POST")
	// ----------------- QUESTOES
	r.HandleFunc(route.AnotacoesRoute, hd.ListAnotacoesHandler).Methods("GET")
	r.HandleFunc("/createAnotacao", hd.CreateAnotacaoHandler).Methods("POST")
	r.HandleFunc("/updateAnotacao", hd.UpdateAnotacaoHandler).Methods("POST")
	r.HandleFunc("/deleteAnotacao", hd.DeleteAnotacaoHandler).Methods("POST")
	// ----------------- RADARES
	r.HandleFunc(route.RadaresRoute, hd.ListRadaresHandler).Methods("GET")
	r.HandleFunc("/createRadar", hd.CreateRadarHandler).Methods("POST")
	r.HandleFunc("/updateRadar", hd.UpdateRadarHandler).Methods("POST")
	r.HandleFunc("/deleteRadar", hd.DeleteRadarHandler).Methods("POST")
	// ----------------- ROLES
	r.HandleFunc(route.RolesRoute, hd.ListPerfisHandler).Methods("GET")
	r.HandleFunc("/createRole", hd.CreateRoleHandler).Methods("POST")
	r.HandleFunc("/updateRole", hd.UpdateRoleHandler).Methods("POST")
	r.HandleFunc("/deleteRole", hd.DeleteRoleHandler).Methods("POST")
	// ----------------- VERSOES
	r.HandleFunc(route.VersoesRoute, hd.ListVersoesHandler).Methods("GET")
	r.HandleFunc("/createVersao", hd.CreateVersaoHandler).Methods("POST")
	r.HandleFunc("/updateVersao", hd.UpdateVersaoHandler).Methods("POST")
	r.HandleFunc("/deleteVersao", hd.DeleteVersaoHandler).Methods("POST")
	// ----------------- STATUS
	r.HandleFunc(route.StatusRoute, hd.ListStatusHandler).Methods("GET")
	r.HandleFunc("/createStatus", hd.CreateStatusHandler).Methods("POST")
	r.HandleFunc("/updateStatus", hd.UpdateStatusHandler).Methods("POST")
	r.HandleFunc("/deleteStatus", hd.DeleteStatusHandler).Methods("POST")
	// ----------------- TIPOS NOTAS
	r.HandleFunc(route.TiposNotasRoute, hd.ListTiposNotasHandler).Methods("GET")
	r.HandleFunc("/createTipoNota", hd.CreateTipoNotaHandler).Methods("POST")
	r.HandleFunc("/updateTipoNota", hd.UpdateTipoNotaHandler).Methods("POST")
	r.HandleFunc("/deleteTipoNota", hd.DeleteTipoNotaHandler).Methods("POST")
	// ----------------- USERS
	r.HandleFunc(route.UsersRoute, hd.ListUsersHandler).Methods("GET")
	r.HandleFunc("/signUpUser", hd.SignUpUserHandler).Methods("GET")
	r.HandleFunc("/registerNewUser", hd.RegisterNewUserHandler).Methods("POST")
	r.HandleFunc("/recoverPwd", hd.RecoverPasswordHandler).Methods("GET")
	r.HandleFunc("/sendPassword", hd.SendPasswordHandler).Methods("POST")
	r.HandleFunc("/changePassword", hd.ChangePasswordHandler).Methods("POST")
	r.HandleFunc("/password", hd.PasswordHandler).Methods("GET")
	r.HandleFunc("/createUser", hd.CreateUserHandler).Methods("POST")
	r.HandleFunc("/updateUser", hd.UpdateUserHandler).Methods("POST")
	r.HandleFunc("/deleteUser", hd.DeleteUserHandler).Methods("POST")
	// ----------------- WORKFLOWS
	r.HandleFunc(route.WorkflowsRoute, hd.ListWorkflowsHandler).Methods("GET")
	r.HandleFunc("/createWorkflow", hd.CreateWorkflowHandler).Methods("POST")
	r.HandleFunc("/updateWorkflow", hd.UpdateWorkflowHandler).Methods("POST")
	r.HandleFunc("/deleteWorkflow", hd.DeleteWorkflowHandler).Methods("POST")
	// ----------------- Microserviços
	r.HandleFunc("/salvarAnalise", hd.SalvarAnalise).Methods("POST")
	r.HandleFunc("/salvarPesoPilar", hd.SalvarPesoPilar).Methods("GET")
	r.HandleFunc("/salvarPesoElemento", hd.SalvarPesoElemento).Methods("GET")
	r.HandleFunc("/salvarNotaElemento", hd.SalvarNotaElemento).Methods("GET")
	r.HandleFunc("/salvarAuditorComponente", hd.SalvarAuditorComponente).Methods("GET")
	r.HandleFunc("/salvarReprogramacao", hd.SalvarReprogramacaoComponente).Methods("GET")
	r.HandleFunc("/loadAnotacoesRadaresByRadarId", hd.LoadAnotacoesRadaresByRadarId).Methods("GET")
	r.HandleFunc("/loadConfigPlanos", hd.LoadConfigPlanos).Methods("GET")
	r.HandleFunc("/updateConfigPlanos", hd.UpdateConfigPlanos).Methods("GET")
	r.HandleFunc("/loadAnalise", hd.LoadAnalise).Methods("GET")
	r.HandleFunc("/loadDescricao", hd.LoadDescricao).Methods("GET")
	r.HandleFunc("/loadHistoricosPilar", hd.LoadHistoricosPilar).Methods("GET")
	r.HandleFunc("/loadHistoricosComponente", hd.LoadHistoricosComponente).Methods("GET")
	r.HandleFunc("/loadHistoricosElemento", hd.LoadHistoricosElemento).Methods("GET")
	r.HandleFunc("/loadTiposNotaByComponenteId", hd.LoadTiposNotaByComponenteId).Methods("GET")
	r.HandleFunc("/loadSupervisorByEntidadeIdCicloId", hd.LoadSupervisorByEntidadeIdCicloId).Methods("GET")
	r.HandleFunc("/loadIntegrantesByEntidadeIdCicloId", hd.LoadIntegrantesByEntidadeIdCicloId).Methods("GET")
	r.HandleFunc("/loadMembrosByEscritorioId", hd.LoadMembrosByEscritorioId).Methods("GET")
	r.HandleFunc("/loadJurisdicoesByEscritorioId", hd.LoadJurisdicoesByEscritorioId).Methods("GET")
	r.HandleFunc("/loadElementosByComponenteId", hd.LoadElementosByComponenteId).Methods("GET")
	r.HandleFunc("/loadPlanosByEntidadeId", hd.LoadPlanosByEntidadeId).Methods("GET")
	r.HandleFunc("/loadComponentesByPilarId", hd.LoadComponentesByPilarId).Methods("GET")
	r.HandleFunc("/loadCiclosByEntidadeId", hd.LoadCiclosByEntidadeId).Methods("GET")
	r.HandleFunc("/loadPilaresByCicloId", hd.LoadPilaresByCicloId).Methods("GET")
	r.HandleFunc("/loadItensByElementoId", hd.LoadItensByElementoId).Methods("GET")
	r.HandleFunc("/loadFeaturesByRoleId", hd.LoadFeaturesByRoleId).Methods("GET")
	r.HandleFunc("/loadRolesByActionId", hd.LoadRolesByActionId).Methods("GET")
	r.HandleFunc("/loadActivitiesByWorkflowId", hd.LoadActivitiesByWorkflowId).Methods("GET")
	r.HandleFunc("/loadAllowedActions", hd.LoadAllowedActions).Methods("GET")
	r.HandleFunc("/loadAvailableFeatures", hd.LoadAvailableFeatures).Methods("GET")
	r.HandleFunc("/executeAction", hd.ExecuteActionHandler).Methods("GET")
	// ----------------- STATICS
	http.Handle("/statics/",
		http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics"))),
	)
	http.Handle("/", r)
	addr := ":" + sConfig.ServerPort
	log.Printf("Listening on %s...\n", addr)

	if err := http.ListenAndServeTLS(addr, "cert.pem", "privkey.pem", nil); err != nil {
		log.Println(err)
		http.ListenAndServe(addr, nil)
	}
	defer hd.Db.Close()
}

func check(e error) {
	if e != nil {
		log.Println(e)
	}
}

func ReadConfig(t ConfigType) Config {
	var configfile = "config/"
	if t == SERVER {
		configfile += "server.ini"
	} else if t == EMAIL {
		configfile += "email.ini"
	} else {
		configfile += "banco.ini"
	}
	log.Println(configfile)
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Arquivo de configuração "+configfile+" faltando: ", configfile)
	}
	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}
