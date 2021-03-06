package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	//"strings"
	mdl "virtus/models"
	sec "virtus/security"
)

func ListMatrizesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Matrizes Handler")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "viewMatriz") {
		log.Println("--------------")
		currentUser := GetUserInCookie(w, r)
		var page mdl.PageEntidadesCiclos
		// Entidades da jurisdição do Escritório ao qual pertenço
		sql := "SELECT b.entidade_id, d.nome " +
			" FROM escritorios a " +
			" LEFT JOIN jurisdicoes b ON a.id = b.escritorio_id " +
			" LEFT JOIN membros c ON c.escritorio_id = b.escritorio_id " +
			" LEFT JOIN entidades d ON d.id = b.entidade_id " +
			" INNER JOIN ciclos_entidades e ON e.entidade_id = b.entidade_id " +
			" WHERE c.usuario_id = $1 OR a.chefe_id = $2 " +
			" GROUP BY 1,2"
		log.Println(sql)
		rows, _ := Db.Query(sql, currentUser.Id, currentUser.Id)
		defer rows.Close()
		var entidades []mdl.Entidade
		var entidade mdl.Entidade
		var i = 1
		for rows.Next() {
			rows.Scan(
				&entidade.Id,
				&entidade.Nome)
			entidade.Order = i
			i++
			sql = "SELECT b.id, b.nome " +
				" FROM ciclos_entidades a " +
				" LEFT JOIN ciclos b ON a.ciclo_id = b.id " +
				" WHERE a.entidade_id = $1 " +
				" ORDER BY id asc"
			rows, _ = Db.Query(sql, entidade.Id)
			defer rows.Close()
			var ciclosEntidade []mdl.CicloEntidade
			var cicloEntidade mdl.CicloEntidade
			i = 1
			for rows.Next() {
				rows.Scan(&cicloEntidade.Id, &cicloEntidade.Nome)
				cicloEntidade.Order = i
				i++
				ciclosEntidade = append(ciclosEntidade, cicloEntidade)
			}
			entidade.CiclosEntidade = ciclosEntidade
			entidades = append(entidades, entidade)
		}
		page.Entidades = entidades
		page.AppName = mdl.AppName
		page.Title = "Matriz de Trabalho"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/matrizes/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Entidades-Matrizes", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func loadElementosDaMatriz(entidadeId string, cicloId string, pilarId string, componenteId string) []mdl.ElementoDaMatriz {
	sql := " SELECT " +
		" 	     coalesce(R3.author_id,0) as author_id, " +
		" 	     coalesce(q.name,'') as author_name, " +
		" 	     coalesce(R3.motivacao_peso,'') as motivacao_peso, " +
		" 	     coalesce(R3.motivacao_nota,'') as motivacao_nota, " +
		" 	     coalesce(CO.supervisor_id,0) as super_id, coalesce(o.name,'') as supervisor_nome, " +
		" 	     coalesce(CO.auditor_id,0) as auditor_id, coalesce(p.name,'') as auditor_nome, " +
		"		 R1.elemento_id as elemento_id, " +
		"		 COALESCE((SELECT count(1) FROM elementos_componentes " +
		"				WHERE componente_id = " + componenteId + "),0) AS qtdElementos, " +
		"        R1.ciclo_id, " +
		"        COALESCE(R1.ciclo_nome, ''), " +
		"        (SELECT coalesce(nota,0) from produtos_ciclos where  ciclo_id = " + cicloId + " AND entidade_id = " + entidadeId + ") AS ciclo_nota, " +
		"   (SELECT count(1) " +
		"    FROM " +
		"      (SELECT pilar_id " +
		"       FROM pilares_ciclos " +
		"       WHERE ciclo_id = R1.ciclo_id " +
		"       GROUP BY pilar_id) R) AS qtdPilares, " +
		"        R1.pilar_id, " +
		"        COALESCE(R1.pilar_nome, ''), " +
		"        COALESCE(PI.peso, 0) AS pilar_peso, " +
		"        COALESCE(PI.nota, 0) AS pilar_nota, " +
		"   (SELECT count(1) " +
		"    FROM " +
		"      (SELECT componente_id " +
		"       FROM componentes_pilares " +
		"       WHERE pilar_id = R1.pilar_id " +
		"       GROUP BY componente_id) R) AS qtdComponentes, " +
		"        R1.componente_id, " +
		"        COALESCE(R1.componente_nome, ''), " +
		"        COALESCE(CO.peso, 0) AS componente_peso, " +
		"        COALESCE(CO.nota, 0) AS componente_nota, " +
		"   (SELECT count(1) " +
		"    FROM " +
		"      (SELECT tipo_nota_id " +
		"       FROM elementos_componentes " +
		"       WHERE componente_id = R1.componente_id " +
		"       GROUP BY tipo_nota_id) R) AS qtdTiposNotas, " +
		"        R1.tipo_nota_id, " +
		"        COALESCE(m.nome,'') AS tipo_nota_nome, " +
		"        COALESCE(m.letra,'') AS tipo_nota_letra, " +
		"        COALESCE(m.cor_letra,'') as tipo_nota_cor_letra, " +
		"        COALESCE(TN.peso, 0) AS tipo_nota_peso, " +
		"        COALESCE(TN.nota, 0) AS tipo_nota_nota, " +
		"      	 COALESCE(n.nome, '') AS elemento_nome, " +
		"     	 COALESCE(EL.peso, 0) AS elemento_peso, " +
		"    	 COALESCE(EL.nota, 0) AS elemento_nota,	" +
		"        " + entidadeId + " AS entidade_id, " +
		"        COALESCE(y.nome,'') as entidade_nome, " +
		"        COALESCE(R2.plano_id, 0) AS plano_id, " +
		"        COALESCE(z.cnpb,'') AS cnpb, " +
		"        CASE WHEN z.recurso_garantidor < 1000000000 THEN z.recurso_garantidor::numeric::MONEY/1000000||' mi' ELSE z.recurso_garantidor::numeric::MONEY/1000000000||' bi' END as rg, " +
		"        COALESCE(z.modalidade_id,'') as modalidade, " +
		"        COALESCE((SELECT count(1) " +
		"   		FROM " +
		"     		(SELECT DISTINCT b.elemento_id " +
		"      			FROM produtos_planos a " +
		"	   			LEFT JOIN elementos_componentes b ON " +
		"	  			a.componente_id = b.componente_id " +
		"	  			WHERE a.ciclo_id = " + cicloId +
		"      			AND a.entidade_id = " + entidadeId +
		"	  			AND a.pilar_id = " + pilarId +
		"	  			AND a.componente_id = " + componenteId +
		"      				GROUP BY b.elemento_id) R),0) as EntidadeQtdPlanos " +
		" FROM " +
		"   (SELECT a.id AS ciclo_id, " +
		"           a.nome AS ciclo_nome, " +
		"           b.id AS pilar_id, " +
		"           c.nome AS pilar_nome, " +
		"           e.id AS componente_id, " +
		"           e.nome AS componente_nome, " +
		"           g.id AS tipo_nota_id, " +
		"           g.nome AS tipo_nota_nome, " +
		"			i.id AS elemento_id, " +
		"			i.nome AS elemento_nome " +
		"    FROM ciclos a " +
		"    INNER JOIN pilares_ciclos b ON b.ciclo_id = a.id " +
		"    INNER JOIN pilares c ON c.id = b.pilar_id " +
		"    INNER JOIN componentes_pilares d ON d.pilar_id = c.id " +
		"    INNER JOIN componentes e ON d.componente_id = e.id " +
		"    INNER JOIN tipos_notas_componentes f ON e.id = f.componente_id " +
		"    INNER JOIN tipos_notas g ON g.id = f.tipo_nota_id " +
		" 	 INNER JOIN elementos_componentes h ON e.id = h.componente_id and g.id = h.tipo_nota_id " +
		"	 INNER JOIN elementos i ON i.id = h.elemento_id " +
		"    WHERE a.id = 1 "
	if pilarId != "" {
		sql += " AND b.pilar_id = " + pilarId
	}
	if componenteId != "" {
		sql += " AND d.componente_id = " + componenteId
	}
	sql += " ) R1 " +
		" LEFT JOIN " +
		"   (SELECT DISTINCT entidade_id, " +
		"                    ciclo_id, " +
		"                    PILAR_ID, " +
		"                    COMPONENTE_ID, " +
		"                    plano_id " +
		"    FROM produtos_planos " +
		"    WHERE ciclo_id = " + cicloId +
		"      AND entidade_id = " + entidadeId +
		"    ORDER BY 1) R2 ON (R1.CICLO_id = R2.ciclo_id " +
		"                       AND R1.PILAR_ID = R2.PILAR_ID " +
		"                       AND R1.COMPONENTE_ID = R2.COMPONENTE_ID) " +
		" LEFT JOIN produtos_elementos EL ON (R1.ciclo_id = EL.CICLO_ID " +
		"                                     AND R1.pilar_id = EL.pilar_id " +
		"                                     AND R1.componente_id = EL.componente_id " +
		"                                     AND R1.tipo_nota_id = EL.tipo_nota_id " +
		"								   	  AND R1.elemento_id = EL.elemento_id " +
		"                                     AND EL.entidade_id = R2.entidade_id " +
		"                                     AND EL.plano_id = R2.plano_id) " +
		" LEFT JOIN produtos_tipos_notas TN ON (R1.ciclo_id = TN.CICLO_ID " +
		"                                       AND R1.pilar_id = TN.pilar_id " +
		"                                       AND R1.componente_id = TN.componente_id " +
		"                                       AND R1.tipo_nota_id = TN.tipo_nota_id " +
		"                                       AND TN.entidade_id = R2.entidade_id " +
		"                                       AND TN.plano_id = R2.plano_id) " +
		" LEFT JOIN produtos_componentes CO ON (R1.componente_id = CO.componente_id " +
		"                                       AND R1.pilar_id = CO.pilar_id " +
		"                                       AND R1.ciclo_id = CO.ciclo_id " +
		"                                       AND CO.entidade_id = " + entidadeId + ") " +
		" LEFT JOIN produtos_pilares PI ON (R1.pilar_id = PI.pilar_id " +
		"                                   AND R1.ciclo_id = PI.ciclo_id " +
		"                                   AND PI.entidade_id = R2.entidade_id) " +
		" LEFT JOIN produtos_ciclos CI ON (R1.ciclo_id = CI.ciclo_id " +
		"                                  AND R2.entidade_id = CI.entidade_id) " +
		" LEFT JOIN tipos_notas m ON R1.tipo_nota_id = m.id " +
		" LEFT JOIN elementos n ON R1.elemento_id = n.id " +
		" LEFT JOIN planos z ON R2.plano_id = z.id " +
		" LEFT JOIN entidades y ON y.id = " + entidadeId +
		" LEFT JOIN users o ON co.supervisor_id = o.id " +
		" LEFT JOIN users p ON co.auditor_id = p.id " +
		" LEFT JOIN (SELECT R1.id, R1.entidade_id, R1.ciclo_id, R1.pilar_id, R1.plano_id, R1.componente_id, R1.tipo_nota_id, R1.elemento_id, motivacao_peso, motivacao_nota, author_id, criado_em " +
		" FROM produtos_elementos_historicos R1 " +
		" INNER JOIN (SELECT peh.entidade_id, peh.ciclo_id, peh.pilar_id, peh.componente_id, peh.plano_id, peh.elemento_id, max(peh.id) as id " +
		" FROM produtos_elementos_historicos PEH group by 1,2,3,4,5,6) R2 " +
		" ON R1.id = R2.id) R3 ON (R3.ciclo_id = EL.CICLO_ID " +
		" AND R3.pilar_id = EL.pilar_id " +
		" AND R3.componente_id = EL.componente_id " +
		" AND R3.tipo_nota_id = EL.tipo_nota_id " +
		" AND R3.entidade_id = EL.entidade_id " +
		" AND R3.plano_id = EL.plano_id) " +
		" LEFT JOIN users q ON R3.author_id = q.id " +
		" ORDER BY ciclo_id, pilar_id, componente_id, elemento_id, tipo_nota_id, rg "
	log.Println(sql)
	rows, _ := Db.Query(sql)
	defer rows.Close()
	var elementosMatriz []mdl.ElementoDaMatriz
	var elementoMatriz mdl.ElementoDaMatriz
	var i = 1
	for rows.Next() {
		rows.Scan(
			&elementoMatriz.AutorNotaId,
			&elementoMatriz.AutorNotaName,
			&elementoMatriz.MotivacaoPeso,
			&elementoMatriz.MotivacaoNota,
			&elementoMatriz.SupervisorId,
			&elementoMatriz.SupervisorName,
			&elementoMatriz.AuditorId,
			&elementoMatriz.AuditorName,
			&elementoMatriz.ElementoId,
			&elementoMatriz.ComponenteQtdElementos,
			&elementoMatriz.CicloId,
			&elementoMatriz.CicloNome,
			&elementoMatriz.CicloNota,
			&elementoMatriz.CicloQtdPilares,
			&elementoMatriz.PilarId,
			&elementoMatriz.PilarNome,
			&elementoMatriz.PilarPeso,
			&elementoMatriz.PilarNota,
			&elementoMatriz.PilarQtdComponentes,
			&elementoMatriz.ComponenteId,
			&elementoMatriz.ComponenteNome,
			&elementoMatriz.ComponentePeso,
			&elementoMatriz.ComponenteNota,
			&elementoMatriz.ComponenteQtdTiposNotas,
			&elementoMatriz.TipoNotaId,
			&elementoMatriz.TipoNotaNome,
			&elementoMatriz.TipoNotaLetra,
			&elementoMatriz.TipoNotaCorLetra,
			&elementoMatriz.TipoNotaPeso,
			&elementoMatriz.TipoNotaNota,
			&elementoMatriz.ElementoNome,
			&elementoMatriz.ElementoPeso,
			&elementoMatriz.ElementoNota,
			&elementoMatriz.EntidadeId,
			&elementoMatriz.EntidadeNome,
			&elementoMatriz.PlanoId,
			&elementoMatriz.CNPB,
			&elementoMatriz.RecursoGarantidor,
			&elementoMatriz.Modalidade,
			&elementoMatriz.EntidadeQtdPlanos)
		elementoMatriz.Order = i
		i++
		elementosMatriz = append(elementosMatriz, elementoMatriz)
	}
	return elementosMatriz
}
func loadTiposNotasMatriz(entidadeId string, cicloId string, pilarId string) []mdl.ElementoDaMatriz {
	sql := " SELECT R1.ciclo_id, " +
		"        COALESCE(R1.ciclo_nome, ''), " +
		"        (SELECT coalesce(nota,0) from produtos_ciclos where  ciclo_id = " + cicloId + " AND entidade_id = " + entidadeId + ") AS ciclo_nota, " +
		"   (SELECT count(1) " +
		"    FROM " +
		"      (SELECT pilar_id " +
		"       FROM pilares_ciclos " +
		"       WHERE ciclo_id = R1.ciclo_id " +
		"       GROUP BY pilar_id) R) AS qtdPilares, " +
		"        R1.pilar_id, " +
		"        COALESCE(R1.pilar_nome, ''), " +
		"        COALESCE(PI.peso, 0) AS pilar_peso, " +
		"        COALESCE(PI.nota, 0) AS pilar_nota, " +
		"   (SELECT count(1) " +
		"    FROM " +
		"      (SELECT componente_id " +
		"       FROM componentes_pilares " +
		"       WHERE pilar_id = R1.pilar_id " +
		"       GROUP BY componente_id) R) AS qtdComponentes, " +
		"   (SELECT count(1) " +
		"    FROM " +
		"      (SELECT tipo_nota_id " +
		"       FROM elementos_componentes " +
		"       WHERE componente_id = R1.componente_id " +
		"       GROUP BY tipo_nota_id) R) AS qtdTiposNotas, " +
		"        R1.componente_id, " +
		"        COALESCE(R1.componente_nome, ''), " +
		"        COALESCE(CO.peso, 0) AS componente_peso, " +
		"        COALESCE(CO.nota, 0) AS componente_nota, " +
		"        R1.tipo_nota_id, " +
		"        COALESCE(m.letra,'') AS tipo_nota_letra, " +
		"        COALESCE(m.cor_letra,'') as tipo_nota_cor_letra, " +
		"        COALESCE(TN.peso, 0) AS tipo_nota_peso, " +
		"        COALESCE(TN.nota, 0) AS tipo_nota_nota, " +
		"        " + entidadeId + " AS entidade_id, " +
		"        COALESCE(y.nome,'') as entidade_nome, " +
		"        COALESCE(R2.plano_id, 0) AS plano_id, " +
		"        COALESCE(z.cnpb,'') AS cnpb, " +
		"        CASE WHEN z.recurso_garantidor < 1000000000 THEN z.recurso_garantidor::numeric::MONEY/1000000||' mi' ELSE z.recurso_garantidor::numeric::MONEY/1000000000||' bi' END as rg, " +
		"        COALESCE(z.modalidade_id,'') as modalidade, " +
		"        (SELECT count(1) FROM (SELECT DISTINCT plano_id FROM produtos_planos WHERE entidade_id = " + entidadeId + " AND ciclo_id = " + cicloId + " GROUP BY plano_id) S) as EntidadeQtdPlanos " +
		" FROM " +
		"   (SELECT a.id AS ciclo_id, " +
		"           a.nome AS ciclo_nome, " +
		"           b.id AS pilar_id, " +
		"           c.nome AS pilar_nome, " +
		"           e.id AS componente_id, " +
		"           e.nome AS componente_nome, " +
		"           g.id AS tipo_nota_id, " +
		"           g.nome AS tipo_nota_nome " +
		"    FROM ciclos a " +
		"    INNER JOIN pilares_ciclos b ON b.ciclo_id = a.id " +
		"    INNER JOIN pilares c ON c.id = b.pilar_id " +
		"    INNER JOIN componentes_pilares d ON d.pilar_id = c.id " +
		"    INNER JOIN componentes e ON d.componente_id = e.id " +
		"    INNER JOIN tipos_notas_componentes f ON e.id = f.componente_id " +
		"    INNER JOIN tipos_notas g ON g.id = f.tipo_nota_id " +
		"    WHERE a.id = 1 "
	if pilarId != "" {
		sql += " AND b.pilar_id = " + pilarId
	}
	sql += " ) R1 " +
		" LEFT JOIN " +
		"   (SELECT DISTINCT entidade_id, " +
		"                    ciclo_id, " +
		"                    PILAR_ID, " +
		"                    COMPONENTE_ID, " +
		"                    plano_id " +
		"    FROM produtos_planos " +
		"    WHERE ciclo_id = " + cicloId +
		"      AND entidade_id = " + entidadeId +
		"    ORDER BY 1) R2 ON (R1.CICLO_id = R2.ciclo_id " +
		"                       AND R1.PILAR_ID = R2.PILAR_ID " +
		"                       AND R1.COMPONENTE_ID = R2.COMPONENTE_ID) " +
		" LEFT JOIN produtos_tipos_notas TN ON (R1.ciclo_id = TN.CICLO_ID " +
		"                                       AND R1.pilar_id = TN.pilar_id " +
		"                                       AND R1.componente_id = TN.componente_id " +
		"                                       AND R1.tipo_nota_id = TN.tipo_nota_id " +
		"                                       AND TN.entidade_id = R2.entidade_id " +
		"                                       AND TN.plano_id = R2.plano_id) " +
		" LEFT JOIN produtos_componentes CO ON (R1.componente_id = CO.componente_id " +
		"                                       AND R1.pilar_id = CO.pilar_id " +
		"                                       AND R1.ciclo_id = CO.ciclo_id " +
		"                                       AND CO.entidade_id = R2.entidade_id) " +
		" LEFT JOIN produtos_pilares PI ON (R1.pilar_id = PI.pilar_id " +
		"                                   AND R1.ciclo_id = PI.ciclo_id " +
		"                                   AND PI.entidade_id = R2.entidade_id) " +
		" LEFT JOIN produtos_ciclos CI ON (R1.ciclo_id = CI.ciclo_id " +
		"                                  AND R2.entidade_id = CI.entidade_id) " +
		" LEFT JOIN tipos_notas m ON R1.tipo_nota_id = m.id " +
		" LEFT JOIN planos z ON R2.plano_id = z.id " +
		" LEFT JOIN entidades y ON y.id = " + entidadeId +
		" ORDER BY ciclo_id,pilar_id,componente_id,tipo_nota_id,rg "
	log.Println(sql)
	rows, _ := Db.Query(sql)
	defer rows.Close()
	var elementosMatriz []mdl.ElementoDaMatriz
	var elementoMatriz mdl.ElementoDaMatriz
	var i = 1
	for rows.Next() {
		rows.Scan(
			&elementoMatriz.CicloId,
			&elementoMatriz.CicloNome,
			&elementoMatriz.CicloNota,
			&elementoMatriz.CicloQtdPilares,
			&elementoMatriz.PilarId,
			&elementoMatriz.PilarNome,
			&elementoMatriz.PilarPeso,
			&elementoMatriz.PilarNota,
			&elementoMatriz.PilarQtdComponentes,
			&elementoMatriz.ComponenteQtdTiposNotas,
			&elementoMatriz.ComponenteId,
			&elementoMatriz.ComponenteNome,
			&elementoMatriz.ComponentePeso,
			&elementoMatriz.ComponenteNota,
			&elementoMatriz.TipoNotaId,
			&elementoMatriz.TipoNotaLetra,
			&elementoMatriz.TipoNotaCorLetra,
			&elementoMatriz.TipoNotaPeso,
			&elementoMatriz.TipoNotaNota,
			&elementoMatriz.EntidadeId,
			&elementoMatriz.EntidadeNome,
			&elementoMatriz.PlanoId,
			&elementoMatriz.CNPB,
			&elementoMatriz.RecursoGarantidor,
			&elementoMatriz.Modalidade,
			&elementoMatriz.EntidadeQtdPlanos)
		elementoMatriz.Order = i
		i++
		//log.Println(elementoMatriz)
		elementosMatriz = append(elementosMatriz, elementoMatriz)
	}
	return elementosMatriz
}

func ExecutarMatrizHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Matriz Handler")
	if sec.IsAuthenticated(w, r) {

		entidadeId := r.FormValue("EntidadeId")
		cicloId := r.FormValue("CicloId")
		pilarId := r.FormValue("PilarId")
		componenteId := r.FormValue("ComponenteId")
		definicaoTemplate := "Main-Matriz"
		title := "Matriz"
		var elementosMatriz []mdl.ElementoDaMatriz
		var page mdl.PageMatriz
		if componenteId != "" {
			title = "Visão do Componente"
			definicaoTemplate = "Main-Componente"
			elementosMatriz = loadElementosDaMatriz(entidadeId, cicloId, pilarId, componenteId)
		} else if pilarId != "" {
			title = "Visão do Pilar"
			definicaoTemplate = "Main-Pilar"
			elementosMatriz = loadTiposNotasMatriz(entidadeId, cicloId, pilarId)
		} else {
			elementosMatriz = loadTiposNotasMatriz(entidadeId, cicloId, pilarId)
		}
		page.ElementosDaMatriz = preencherColspans(elementosMatriz, cicloId)
		page.ComponenteQtdTiposNotas = elementosMatriz[0].ComponenteQtdTiposNotas

		sql := " SELECT " +
			" a.usuario_id, " +
			" coalesce(b.name,'') " +
			" FROM integrantes a " +
			" LEFT JOIN users b " +
			" ON a.usuario_id = b.id " +
			" WHERE " +
			" a.entidade_id = " + entidadeId +
			" AND a.ciclo_id = " + cicloId +
			" AND b.role_id = 3 "
		log.Println(sql)
		rows, _ := Db.Query(sql)
		defer rows.Close()
		var supervisores []mdl.User
		var supervisor mdl.User
		for rows.Next() {
			rows.Scan(&supervisor.Id, &supervisor.Name)
			supervisores = append(supervisores, supervisor)
		}
		page.Supervisores = supervisores

		sql = " SELECT " +
			" a.usuario_id, " +
			" b.name " +
			" FROM integrantes a " +
			" LEFT JOIN users b " +
			" ON a.usuario_id = b.id " +
			" WHERE " +
			" a.entidade_id = " + entidadeId +
			" AND a.ciclo_id = " + cicloId +
			" AND b.role_id = 4 "
		log.Println(sql)
		rows, _ = Db.Query(sql)
		defer rows.Close()
		var auditores []mdl.User
		var auditor mdl.User
		for rows.Next() {
			rows.Scan(&auditor.Id, &auditor.Name)
			auditores = append(auditores, auditor)
		}
		page.Supervisores = supervisores
		page.Auditores = auditores
		page.AppName = mdl.AppName
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		page.Dec = func(i int) int {
			return i - 1
		}
		page.Inc = func(i int) int {
			return i + 1
		}
		page.MulTxt = func(i int, j string) int {
			k, _ := strconv.Atoi(j)
			return i * k
		}
		page.Mul = func(i int, j int) int {
			return i * j
		}
		page.SomarTxt = func(i int, j string) int {
			k, _ := strconv.Atoi(j)
			return i + k
		}
		page.Somar = func(i int, j int) int {
			return i + j
		}
		var tmpl = template.Must(template.ParseGlob("tiles/matrizes/*"))
		tmpl.ParseGlob("tiles/*")
		page.Title = title
		tmpl.ExecuteTemplate(w, definicaoTemplate, page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func calcularColspan(tipo string, identificador int64) int {
	sql := ""
	if tipo == "ciclo" {
		tipo = "pc.ciclo"
	} else if tipo == "pilar" {
		tipo = "pc.pilar"
	} else if tipo == "componente" {
		tipo = "cp.componente"
	}
	sql = " SELECT COUNT(1) FROM ( " +
		" SELECT ciclo_id, pc.pilar_id, cp.componente_id, tnc.tipo_nota_id " +
		" FROM tipos_notas_componentes tnc " +
		" LEFT JOIN componentes c ON tnc.componente_id = c.id " +
		" LEFT JOIN componentes_pilares cp ON c.id = cp.componente_id " +
		" LEFT JOIN pilares p ON p.id = cp.pilar_id " +
		" LEFT JOIN pilares_ciclos pc ON p.id = pc.pilar_id " +
		" WHERE " + tipo + "_id = $1) R "
	log.Println(sql)
	rows, _ := Db.Query(sql, identificador)
	defer rows.Close()
	resultado := 0
	if rows.Next() {
		rows.Scan(&resultado)
		return resultado
	}
	return 0
}

func preencherColspans(elementosMatriz []mdl.ElementoDaMatriz, cicloId string) []mdl.ElementoDaMatriz {
	sql := " SELECT ciclo_id, 0 AS pilar_id, 0 as componente_id, sum(qtdCelula) FROM ( " +
		" SELECT ciclo_id, pilar_id, 0 as componente_id, sum(qtdCelula) AS qtdCelula FROM ( " +
		" SELECT ciclo_id, pilar_id, componente_id, count(qtdCelula) AS qtdCelula FROM ( " +
		" SELECT a.ciclo_id, a.pilar_id, b.componente_id, c.tipo_nota_id, COUNT(c.tipo_nota_id) AS qtdCelula " +
		" FROM pilares_ciclos a " +
		" LEFT JOIN componentes_pilares b ON b.pilar_id = a.pilar_id " +
		" LEFT JOIN tipos_notas_componentes c ON b.componente_id = c.componente_id " +
		" INNER JOIN elementos_componentes ec ON ec.tipo_nota_id = c.tipo_nota_id " +
		" AND b.componente_id = ec.componente_id " +
		" WHERE a.ciclo_id = $1 " +
		" GROUP BY 1, 2, 3, 4 " +
		" ORDER BY 1, 2, 3, 4) R1 " +
		" GROUP BY 1, 2, 3 ) R2 " +
		" GROUP BY 1, 2, 3 ) R3 " +
		" GROUP BY 1, 2, 3 " +
		" UNION " +
		" SELECT ciclo_id, pilar_id, 0 as componente_id, sum(qtdCelula) AS qtdCelula FROM ( " +
		" SELECT ciclo_id, pilar_id, componente_id, count(qtdCelula) AS qtdCelula FROM ( " +
		" SELECT a.ciclo_id, a.pilar_id, b.componente_id, c.tipo_nota_id, COUNT(c.tipo_nota_id) AS qtdCelula " +
		" FROM pilares_ciclos a " +
		" LEFT JOIN componentes_pilares b ON b.pilar_id = a.pilar_id " +
		" LEFT JOIN tipos_notas_componentes c ON b.componente_id = c.componente_id " +
		" INNER JOIN elementos_componentes ec ON ec.tipo_nota_id = c.tipo_nota_id " +
		" AND b.componente_id = ec.componente_id " +
		" WHERE a.ciclo_id = $2 " +
		" GROUP BY 1, 2, 3, 4 " +
		" ORDER BY 1, 2, 3, 4) R1 " +
		" GROUP BY 1, 2, 3 ) R2 " +
		" GROUP BY 1, 2, 3  " +
		" UNION " +
		" SELECT ciclo_id, pilar_id, componente_id, count(qtdCelula) AS qtdCelula FROM ( " +
		" SELECT a.ciclo_id, a.pilar_id, b.componente_id, c.tipo_nota_id, COUNT(c.tipo_nota_id) AS qtdCelula " +
		" FROM pilares_ciclos a " +
		" LEFT JOIN componentes_pilares b ON b.pilar_id = a.pilar_id " +
		" LEFT JOIN tipos_notas_componentes c ON b.componente_id = c.componente_id " +
		" INNER JOIN elementos_componentes ec ON ec.tipo_nota_id = c.tipo_nota_id " +
		" AND b.componente_id = ec.componente_id " +
		" WHERE a.ciclo_id = $3 " +
		" GROUP BY 1, 2, 3, 4 " +
		" ORDER BY 1, 2, 3, 4) R1 " +
		" GROUP BY 1, 2, 3 " +
		" ORDER BY 1, 2, 3 "
	rows, _ := Db.Query(sql, cicloId, cicloId, cicloId)
	defer rows.Close()
	log.Println(sql)
	var cols []mdl.ColSpan
	var col mdl.ColSpan
	for rows.Next() {
		rows.Scan(&col.CicloId, &col.PilarId, &col.ComponenteId, &col.Qtd)
		cols = append(cols, col)
		// log.Println(col)
	}
	var novosElementos []mdl.ElementoDaMatriz
	for _, elemento := range elementosMatriz {
		for _, col := range cols {
			if col.PilarId == 0 && col.ComponenteId == 0 {
				elemento.CicloColSpan = col.Qtd
			} else if elemento.PilarId == col.PilarId && col.ComponenteId == 0 {
				elemento.PilarColSpan = col.Qtd
			} else if elemento.PilarId == col.PilarId && elemento.ComponenteId == col.ComponenteId {
				elemento.ComponenteColSpan = col.Qtd
			}
		}
		novosElementos = append(novosElementos, elemento)
	}
	return novosElementos
}
