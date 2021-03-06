package handlers

import (
	"log"
	"strconv"
	"strings"
	"time"
	mdl "virtus/models"
)

func registrarCronogramaComponente(produto mdl.ProdutoComponente, currentUser mdl.User, tipoData string) {
	sqlStatement := "UPDATE produtos_componentes SET "
	if tipoData == "iniciaEm" {
		sqlStatement += " inicia_em ='" + produto.IniciaEm + "', "
	} else {
		sqlStatement += " termina_em ='" + produto.TerminaEm + "', "
	}
	sqlStatement += " motivacao_reprogramacao ='" + produto.Motivacao + "'" +
		" WHERE entidade_id= " + strconv.FormatInt(produto.EntidadeId, 10) +
		" AND ciclo_id= " + strconv.FormatInt(produto.CicloId, 10) +
		" AND pilar_id= " + strconv.FormatInt(produto.PilarId, 10) +
		" AND componente_id= " + strconv.FormatInt(produto.ComponenteId, 10)
	log.Println(sqlStatement)
	updtForm, _ := Db.Prepare(sqlStatement)
	_, err := updtForm.Exec()
	if err != nil {
		log.Println(err.Error())
	}
}

func registrarAuditorComponente(produto mdl.ProdutoComponente, currentUser mdl.User) {
	sqlStatement := "UPDATE produtos_componentes SET " +
		" auditor_id=" + strconv.FormatInt(produto.AuditorId, 10) + ", justificativa='" + produto.Motivacao + "'" +
		" WHERE entidade_id= " + strconv.FormatInt(produto.EntidadeId, 10) +
		" AND ciclo_id= " + strconv.FormatInt(produto.CicloId, 10) +
		" AND pilar_id= " + strconv.FormatInt(produto.PilarId, 10) +
		" AND componente_id= " + strconv.FormatInt(produto.ComponenteId, 10)
	log.Println(sqlStatement)
	updtForm, _ := Db.Prepare(sqlStatement)
	_, err := updtForm.Exec()
	if err != nil {
		log.Println(err.Error())
	}
}

func registrarNotaElemento(produto mdl.ProdutoElemento, currentUser mdl.User) mdl.NotasAtuais {
	sqlStatement := "UPDATE produtos_elementos a SET nota = " + strconv.Itoa(produto.Nota) + ", " +
		" motivacao_nota = $1 , " +
		" tipo_pontuacao_id = (SELECT case when b.supervisor_id = " + strconv.FormatInt(currentUser.Id, 10) +
		" then 3 when 2 = " + strconv.FormatInt(currentUser.Role, 10) + " then 3 else 1 end " +
		" FROM produtos_componentes b where " +
		" a.entidade_id = b.entidade_id and " +
		" a.ciclo_id = b.ciclo_id and " +
		" a.pilar_id = b.pilar_id and " +
		// " a.plano_id = b.plano_id and " +
		" a.componente_id = b.componente_id) " +
		" WHERE a.entidade_id = " + strconv.FormatInt(produto.EntidadeId, 10) +
		" AND a.ciclo_id = " + strconv.FormatInt(produto.CicloId, 10) +
		" AND a.pilar_id = " + strconv.FormatInt(produto.PilarId, 10) +
		" AND a.plano_id = " + strconv.FormatInt(produto.PlanoId, 10) +
		" AND a.componente_id = " + strconv.FormatInt(produto.ComponenteId, 10) +
		" AND a.elemento_id = " + strconv.FormatInt(produto.ElementoId, 10) +
		" AND a.nota <> " + strconv.Itoa(produto.Nota)
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	updtForm.Exec(produto.Motivacao)
	atualizarTipoNotaNota(produto)
	atualizarPlanoNota(produto)
	atualizarComponenteNota(produto)
	atualizarPilarNota(produto)
	atualizarCicloNota(produto)

	// NOTAS ATUAIS
	notasAtuais := loadNotasAtuais(produto)
	return notasAtuais
}

func atualizarPilarNota(produto mdl.ProdutoElemento) {
	// PRODUTOS_PILARES
	sqlStatement := "UPDATE produtos_pilares a " +
		" SET nota = (select  " +
		" round(CAST(sum(nota*peso)/sum(peso) as numeric),2) AS media " +
		" FROM produtos_componentes b " +
		" WHERE " +
		" a.entidade_id = b.entidade_id " +
		" AND a.ciclo_id = b.ciclo_id  " +
		" AND a.pilar_id = b.pilar_id " +
		" GROUP BY b.entidade_id,  " +
		" b.ciclo_id, " +
		" b.pilar_id " +
		" HAVING sum(peso)>0) " +
		" WHERE a.entidade_id = $1 " +
		" AND a.ciclo_id = $2 "
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	updtForm.Exec(produto.EntidadeId, produto.CicloId)
}

func atualizarComponenteNota(produto mdl.ProdutoElemento) {
	// PRODUTOS_COMPONENTES
	sqlStatement := "UPDATE produtos_componentes a " +
		" set nota = (select  " +
		" round(CAST(sum(nota*peso)/100 as numeric),2) as media " +
		" FROM produtos_planos b " +
		" WHERE " +
		" a.entidade_id = b.entidade_id " +
		" and a.ciclo_id = b.ciclo_id  " +
		" and a.pilar_id = b.pilar_id " +
		" and a.componente_id = b.componente_id " +
		" and a.id_versao_origem is null " +
		" GROUP BY b.entidade_id,  " +
		" b.ciclo_id, " +
		" b.pilar_id, " +
		" b.componente_id " +
		" HAVING sum(peso)>0) " +
		" WHERE a.entidade_id = $1 " +
		" AND a.ciclo_id = $2 "
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	updtForm.Exec(produto.EntidadeId, produto.CicloId)
}

func atualizarPlanoNota(produto mdl.ProdutoElemento) {
	// PRODUTOS_PLANOS
	sqlStatement := "UPDATE produtos_planos a " +
		" set nota = (select  " +
		" round(CAST(sum(nota*peso)/sum(peso) as numeric),2) as media " +
		" FROM produtos_tipos_notas b " +
		" WHERE " +
		" a.entidade_id = b.entidade_id " +
		" AND a.ciclo_id = b.ciclo_id  " +
		" AND a.pilar_id = b.pilar_id " +
		" AND a.componente_id = b.componente_id " +
		" AND a.plano_id = b.plano_id " +
		" GROUP BY b.entidade_id,  " +
		" b.ciclo_id, " +
		" b.pilar_id, " +
		" b.plano_id, " +
		" b.componente_id " +
		" HAVING sum(peso)>0) " +
		" WHERE a.entidade_id = " + strconv.FormatInt(produto.EntidadeId, 10) +
		" AND a.ciclo_id = " + strconv.FormatInt(produto.CicloId, 10) +
		" AND a.pilar_id = " + strconv.FormatInt(produto.PilarId, 10) +
		" AND a.componente_id = " + strconv.FormatInt(produto.ComponenteId, 10) +
		" AND a.plano_id = " + strconv.FormatInt(produto.PlanoId, 10)
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	updtForm.Exec()

}

func atualizarTipoNotaNota(produto mdl.ProdutoElemento) {
	// PRODUTOS_TIPOS_NOTAS
	sqlStatement := "UPDATE produtos_tipos_notas a " +
		" set nota = (select  " +
		" round(CAST(sum(nota*peso)/sum(peso) as numeric),2) as media " +
		" FROM produtos_elementos b " +
		" WHERE " +
		" a.entidade_id = b.entidade_id " +
		" and a.ciclo_id = b.ciclo_id  " +
		" and a.pilar_id = b.pilar_id " +
		" and a.plano_id = b.plano_id " +
		" and a.componente_id = b.componente_id " +
		" and a.tipo_nota_id = b.tipo_nota_id " +
		" GROUP BY b.entidade_id,  " +
		" b.ciclo_id, " +
		" b.pilar_id, " +
		" b.plano_id, " +
		" b.componente_id, " +
		" b.tipo_nota_id " +
		" HAVING sum(peso)>0) " +
		" WHERE a.entidade_id = " + strconv.FormatInt(produto.EntidadeId, 10) +
		" AND a.ciclo_id = " + strconv.FormatInt(produto.CicloId, 10) +
		" AND a.pilar_id = " + strconv.FormatInt(produto.PilarId, 10) +
		" AND a.componente_id = " + strconv.FormatInt(produto.ComponenteId, 10) +
		" AND a.plano_id = " + strconv.FormatInt(produto.PlanoId, 10) +
		" AND a.tipo_nota_id = " + strconv.FormatInt(produto.TipoNotaId, 10)
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	updtForm.Exec()
}

func atualizarCicloNota(produto mdl.ProdutoElemento) {
	// PRODUTOS_CICLOS
	sqlStatement := "UPDATE produtos_ciclos a " +
		" SET nota = (select  " +
		" round(CAST(sum(nota*peso)/sum(peso) as numeric),2) AS media " +
		" FROM produtos_pilares b " +
		" WHERE " +
		" a.entidade_id = b.entidade_id " +
		" AND a.ciclo_id = b.ciclo_id  " +
		" AND EXISTS (SELECT 1 FROM produtos_itens c WHERE b.ciclo_id = c.ciclo_id " +
		" AND b.entidade_id = c.entidade_id " +
		" AND b.pilar_id = c.pilar_id) " +
		" GROUP BY b.entidade_id,  " +
		" b.ciclo_id " +
		" HAVING sum(peso)>0) " +
		" WHERE a.entidade_id = $1 " +
		" AND a.ciclo_id = $2 "
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	updtForm.Exec(produto.EntidadeId, produto.CicloId)

}

func registrarTiposPontuacao(produto mdl.ProdutoElemento, currentUser mdl.User) {
	sqlStatement := "UPDATE produtos_tipos_notas a SET " +
		" tipo_pontuacao_id = (SELECT case when b.supervisor_id = $1 " +
		" then 3 else 2 end FROM produtos_componentes b where a.id = b.id) " +
		" WHERE entidade_id = $2 " +
		" AND  ciclo_id = $3 " +
		" AND  pilar_id = $4 " +
		" AND  componente_id = $5 " +
		" AND  tipo_nota_id = $6 "
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	updtForm.Exec(
		currentUser.Id,
		produto.EntidadeId,
		produto.CicloId,
		produto.PilarId,
		produto.ComponenteId,
		produto.TipoNotaId)
	sqlStatement = "UPDATE produtos_componentes a SET " +
		" tipo_pontuacao_id = (SELECT case when b.supervisor_id = $1 " +
		" then 3 else 2 end FROM produtos_componentes b where a.id = b.id) " +
		" WHERE entidade_id = $2 " +
		" AND  ciclo_id = $3 " +
		" AND  pilar_id = $4 " +
		" AND  componente_id = $5 "
	log.Println(sqlStatement)
	updtForm, err = Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	updtForm.Exec(
		currentUser.Id,
		produto.EntidadeId,
		produto.CicloId,
		produto.PilarId,
		produto.ComponenteId)
	sqlStatement = "UPDATE produtos_pilares a SET " +
		" tipo_pontuacao_id = (SELECT case when b.supervisor_id = $1 " +
		" then 3 else 2 end FROM produtos_pilares b where a.id = b.id) " +
		" WHERE entidade_id = $2 " +
		" AND  ciclo_id = $3 " +
		" AND  pilar_id = $4 "
	log.Println(sqlStatement)
	updtForm, err = Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	updtForm.Exec(
		currentUser.Id,
		produto.EntidadeId,
		produto.CicloId,
		produto.PilarId)
	sqlStatement = "UPDATE produtos_ciclos a SET " +
		" tipo_pontuacao_id = (SELECT case when b.supervisor_id = $1 " +
		" then 3 else 2 end FROM produtos_ciclos b where a.id = b.id) " +
		" WHERE entidade_id = $2 " +
		" AND  ciclo_id = $3 "
	log.Println(sqlStatement)
	updtForm, err = Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	updtForm.Exec(
		currentUser.Id,
		produto.EntidadeId,
		produto.CicloId)
}

func registrarPesoElemento(produto mdl.ProdutoElemento, currentUser mdl.User) mdl.PesosAtuais {
	// PESOS ELEMENTOS
	sqlStatement := "UPDATE produtos_elementos SET peso = $1, motivacao_peso = $2 " +
		" WHERE entidade_id = $3 AND " +
		" ciclo_id = $4 AND " +
		" pilar_id = $5 AND " +
		// " plano_id = $6 AND " +
		" componente_id = $6 AND " +
		" elemento_id = $7 "
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	updtForm.Exec(produto.Peso,
		produto.Motivacao,
		produto.EntidadeId,
		produto.CicloId,
		produto.PilarId,
		//produto.PlanoId,
		produto.ComponenteId,
		produto.ElementoId)
	atualizarPesoTiposNotas(produto, currentUser)
	atualizarPesoPlanos(produto, currentUser)
	atualizarPesoComponentes(produto, currentUser)
	// PESOS ATUAIS
	pesosAtuais := loadPesosAtuais(produto)
	pesosAtuais.ElementoPeso = produto.Peso
	return pesosAtuais
}

func atualizarPesoTiposNotas(produto mdl.ProdutoElemento, currentUser mdl.User) {
	// PESOS TIPOS NOTAS
	sqlStatement := "UPDATE produtos_tipos_notas as p SET peso = R1.peso " +
		"FROM " +
		"(WITH TMP AS " +
		"     (SELECT entidade_id, " +
		"             ciclo_id, " +
		"             pilar_id, " +
		"             plano_id, " +
		"             componente_id, " +
		"             round(CAST(sum(peso) AS numeric), 2) AS TOTAL " +
		"      FROM produtos_elementos " +
		"      WHERE componente_id = " + strconv.FormatInt(produto.ComponenteId, 10) +
		"        AND pilar_id = " + strconv.FormatInt(produto.PilarId, 10) +
		"        AND plano_id = " + strconv.FormatInt(produto.PlanoId, 10) +
		"        AND ciclo_id = " + strconv.FormatInt(produto.CicloId, 10) +
		"        AND entidade_id = " + strconv.FormatInt(produto.EntidadeId, 10) +
		"      GROUP BY entidade_id, " +
		"               ciclo_id, " +
		"               pilar_id, " +
		"               plano_id, " +
		"               componente_id)  " +
		"   SELECT r.entidade_id, r.ciclo_id, r.pilar_id, r.plano_id, r.componente_id, r.tipo_nota_id, round(CAST((sum(r.peso)/(sum(t.TOTAL)/count(1)))*100 AS numeric), 2) AS peso " +
		"   FROM " +
		"     (SELECT b.entidade_id, " +
		"             b.ciclo_id, " +
		"             b.pilar_id, " +
		"             b.plano_id, " +
		"             b.componente_id, " +
		"             b.tipo_nota_id, " +
		"             b.peso " +
		"      FROM produtos_elementos b " +
		"      WHERE componente_id = " + strconv.FormatInt(produto.ComponenteId, 10) +
		"        AND pilar_id = " + strconv.FormatInt(produto.PilarId, 10) +
		"        AND plano_id = " + strconv.FormatInt(produto.PlanoId, 10) +
		"        AND ciclo_id = " + strconv.FormatInt(produto.CicloId, 10) +
		"        AND entidade_id = " + strconv.FormatInt(produto.EntidadeId, 10) +
		"   ) r " +
		"   LEFT JOIN TMP t ON r.entidade_id = t.entidade_id " +
		"   AND r.ciclo_id = t.ciclo_id " +
		"   AND r.pilar_id = t.pilar_id " +
		"   AND r.plano_id = t.plano_id " +
		"   AND r.componente_id = t.componente_id " +
		"   GROUP BY r.entidade_id, " +
		"            r.ciclo_id, " +
		"            r.pilar_id, " +
		"            r.plano_id, " +
		"            r.componente_id, " +
		"            r.tipo_nota_id) AS R1 " +
		"	WHERE " +
		"		 p.tipo_nota_id = R1.tipo_nota_id " +
		"	 	 AND p.componente_id = R1.componente_id " +
		"        AND p.plano_id = R1.plano_id " +
		"        AND p.pilar_id = R1.pilar_id " +
		"        AND p.ciclo_id = R1.ciclo_id " +
		"        AND p.entidade_id = R1.entidade_id "
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	updtForm.Exec()
}

func atualizarPesoPlanos(produto mdl.ProdutoElemento, currentUser mdl.User) {
	// PESOS PLANOS
	sqlStatement := "UPDATE produtos_planos as p SET peso = R1.peso_percentual " +
		" FROM " +
		" ( WITH total AS " +
		"   (SELECT a.entidade_id, " +
		"           a.ciclo_id, " +
		"           a.pilar_id, " +
		"           a.componente_id, " +
		"           sum(p.recurso_garantidor) AS total " +
		"    FROM produtos_planos a " +
		"    INNER JOIN planos p ON (p.entidade_id = a.entidade_id " +
		"                            AND p.id = a.plano_id) " +
		"    WHERE a.entidade_id = " + strconv.FormatInt(produto.EntidadeId, 10) +
		"      AND a.ciclo_id = " + strconv.FormatInt(produto.CicloId, 10) +
		"      AND a.pilar_id = " + strconv.FormatInt(produto.PilarId, 10) +
		"      AND a.componente_id = " + strconv.FormatInt(produto.ComponenteId, 10) +
		"    GROUP BY 1, " +
		"             2, " +
		"             3, " +
		"             4) " +
		" SELECT a.entidade_id, a.ciclo_id, a.pilar_id,  a.plano_id, a.componente_id, " +
		" round(CAST(p.recurso_garantidor/t.total AS numeric), 2)*100 AS peso_percentual " +
		" FROM produtos_planos a " +
		" INNER JOIN planos p ON (p.entidade_id = a.entidade_id " +
		"                         AND p.id = a.plano_id) " +
		" INNER JOIN total t ON (a.entidade_id = t.entidade_id " +
		"                        AND a.ciclo_id = t.ciclo_id " +
		"                        AND a.pilar_id = t.pilar_id " +
		"                        AND a.componente_id = t.componente_id) ) AS R1" +
		"	WHERE " +
		"        p.plano_id = R1.plano_id " +
		"	 	 AND p.componente_id = R1.componente_id " +
		"        AND p.pilar_id = R1.pilar_id " +
		"        AND p.ciclo_id = R1.ciclo_id " +
		"        AND p.entidade_id = R1.entidade_id "
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	updtForm.Exec()
}

func atualizarPesoComponentes(produto mdl.ProdutoElemento, currentUser mdl.User) {
	// PESOS COMPONENTES
	log.Println("*** ATUALIZAR PESO COMPONENTE ***")
	sqlStatement := "UPDATE produtos_componentes a " +
		" SET peso = (SELECT round(CAST(avg(b.peso) as numeric),2) " +
		" FROM produtos_elementos b " +
		" WHERE b.componente_id = a.componente_id AND " +
		" b.pilar_id = a.pilar_id AND " +
		" b.ciclo_id = a.ciclo_id AND " +
		" b.entidade_id = a.entidade_id " +
		" GROUP BY b.entidade_id, b.ciclo_id, b.pilar_id, b.componente_id) " +
		" WHERE entidade_id = $1 AND " +
		" ciclo_id = $2 AND " +
		" pilar_id = $3 AND " +
		" componente_id = $4 "
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = updtForm.Exec(
		produto.EntidadeId,
		produto.CicloId,
		produto.PilarId,
		produto.ComponenteId)
	if err != nil {
		log.Println(err.Error())
	}
}

func registrarProdutosCiclos(currentUser mdl.User, entidadeId string, cicloId string) {
	sqlStatement := "INSERT INTO produtos_ciclos ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" nota, " +
		" tipo_pontuacao_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT " +
		entidadeId + ", " +
		cicloId + ", " +
		" 1, " +
		" $1, " +
		" $2, " +
		" $3 " +
		" FROM ciclos_entidades a " +
		" WHERE NOT EXISTS " +
		"  (SELECT 1 " +
		"   FROM produtos_ciclos b " +
		"   WHERE b.entidade_id = a.entidade_id " +
		"     AND b.ciclo_id = a.ciclo_id) RETURNING id "
	log.Println(sqlStatement)
	produtoCicloId := 0
	err := Db.QueryRow(
		sqlStatement,
		mdl.Calculada,
		currentUser.Id,
		time.Now()).Scan(&produtoCicloId)
	if err != nil {
		log.Println(err)
	}
	sqlStatement = "INSERT INTO produtos_pilares " +
		" (entidade_id, ciclo_id, pilar_id, peso, nota, tipo_pontuacao_id, author_id, criado_em) " +
		" SELECT " +
		entidadeId + ", " +
		cicloId + ", " +
		" a.pilar_id, " +
		" a.peso_padrao, " +
		" 1, " +
		" $1, " +
		" $2, " +
		" $3 " +
		" FROM pilares_ciclos a " +
		" WHERE NOT EXISTS " +
		"  (SELECT 1 " +
		"   FROM produtos_pilares b " +
		"   WHERE b.entidade_id = " + entidadeId +
		"     AND b.ciclo_id = a.ciclo_id " +
		"     AND b.pilar_id = a.pilar_id) RETURNING id"
	log.Println(sqlStatement)
	produtoPilarId := 0
	err = Db.QueryRow(
		sqlStatement,
		mdl.Calculada,
		currentUser.Id,
		time.Now()).Scan(&produtoPilarId)
	if err != nil {
		log.Println(err)
	}

	sqlStatement = "INSERT INTO produtos_componentes ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" pilar_id, " +
		" componente_id, " +
		" peso, " +
		" nota, " +
		" tipo_pontuacao_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT " + entidadeId + ", " + cicloId + ", a.pilar_id, b.componente_id, " +
		" round(CAST(avg(c.peso_padrao) AS numeric),2), 1, " +
		" $1, $2, $3 " +
		" FROM " +
		" PILARES_CICLOS a " +
		" LEFT JOIN COMPONENTES_PILARES b ON (a.pilar_id = b.pilar_id) " +
		" LEFT JOIN ELEMENTOS_COMPONENTES c ON (b.componente_id = c.componente_id) " +
		" WHERE  " +
		" a.ciclo_id = " + cicloId +
		" AND NOT EXISTS " +
		"  (SELECT 1 " +
		"   FROM produtos_componentes c " +
		"   WHERE c.entidade_id = " + entidadeId +
		"     AND c.ciclo_id = a.ciclo_id " +
		"     AND c.pilar_id = a.pilar_id " +
		"     AND c.componente_id = b.componente_id) " +
		" GROUP BY 1,2,3,4 ORDER BY 1,2,3,4" +
		" RETURNING id"
	log.Println(sqlStatement)
	produtoComponenteId := 0
	err = Db.QueryRow(
		sqlStatement,
		mdl.Calculada,
		currentUser.Id,
		time.Now()).Scan(&produtoComponenteId)
	if err != nil {
		log.Println(err)
	}
}

func registrarProdutosPlanos(param mdl.ProdutoPlano, planos string, currentUser mdl.User) int {
	produtoPlanoId := 0
	sqlStatement := "INSERT INTO produtos_planos ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" pilar_id, " +
		" componente_id, " +
		" plano_id, " +
		//" peso, " +
		" nota, " +
		" tipo_pontuacao_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT " + strconv.FormatInt(param.EntidadeId, 10) + ", " +
		strconv.FormatInt(param.CicloId, 10) + ", " +
		strconv.FormatInt(param.PilarId, 10) + ", " +
		strconv.FormatInt(param.ComponenteId, 10) + ", " +
		" p.id, " +
		//" p.recurso_garantidor as peso," +
		" 0 as nota, " +
		" $1, $2, $3 " +
		" FROM " +
		" PILARES_CICLOS a " +
		" LEFT JOIN COMPONENTES_PILARES b ON (a.pilar_id = b.pilar_id) " +
		" LEFT JOIN ELEMENTOS_COMPONENTES c ON (b.componente_id = c.componente_id) " +
		" LEFT JOIN PLANOS p ON (p.entidade_id = " + strconv.FormatInt(param.EntidadeId, 10) +
		" AND p.id IN (" + planos + ")) " +
		" WHERE  " +
		" p.entidade_id = " + strconv.FormatInt(param.EntidadeId, 10) +
		" AND p.id IN (" + planos + ") " +
		" AND NOT EXISTS " +
		"  (SELECT 1 " +
		"   FROM produtos_planos e " +
		"   WHERE e.entidade_id = " + strconv.FormatInt(param.EntidadeId, 10) +
		"     AND e.ciclo_id = " + strconv.FormatInt(param.CicloId, 10) +
		"     AND e.pilar_id = " + strconv.FormatInt(param.PilarId, 10) +
		"     AND e.componente_id = " + strconv.FormatInt(param.ComponenteId, 10) +
		"     AND e.plano_id IN (" + planos + ")) " +
		" GROUP BY 1,2,3,4,5,6 ORDER BY 1,2,3,4,5,6" +
		" RETURNING id"
	log.Println(sqlStatement)
	err := Db.QueryRow(
		sqlStatement,
		mdl.Calculada,
		currentUser.Id,
		time.Now()).Scan(&produtoPlanoId)
	if err != nil {
		log.Println(err.Error())
	}
	if produtoPlanoId == 0 {
		return produtoPlanoId
	}
	sqlStatement = "INSERT INTO produtos_tipos_notas ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" pilar_id, " +
		" componente_id, " +
		" plano_id, " +
		" tipo_nota_id, " +
		" peso, " +
		" nota, " +
		" tipo_pontuacao_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT p.entidade_id, " +
		" p.ciclo_id, " +
		" p.pilar_id, " +
		" p.componente_id, " +
		" p.plano_id, " +
		" d.tipo_nota_id, " +
		" round(CAST(avg(d.peso_padrao) AS numeric),2), " +
		" 0, $1, $2, $3 " +
		" FROM " +
		" PILARES_CICLOS a " +
		" LEFT JOIN COMPONENTES_PILARES b ON a.pilar_id = b.pilar_id " +
		" LEFT JOIN TIPOS_NOTAS_COMPONENTES d ON b.componente_id = d.componente_id " +
		" LEFT JOIN PRODUTOS_PLANOS p ON b.componente_id = p.componente_id " +
		" WHERE  " +
		" a.ciclo_id = " + strconv.FormatInt(param.CicloId, 10) +
		" AND p.entidade_id = " + strconv.FormatInt(param.EntidadeId, 10) +
		" AND p.pilar_id = " + strconv.FormatInt(param.PilarId, 10) +
		" AND p.componente_id = " + strconv.FormatInt(param.ComponenteId, 10) +
		" AND p.plano_id IN (" + planos + ") " +
		" AND NOT EXISTS " +
		"  (SELECT 1 " +
		"   FROM produtos_tipos_notas e " +
		"   WHERE e.entidade_id = " + strconv.FormatInt(param.EntidadeId, 10) +
		"     AND e.ciclo_id = a.ciclo_id " +
		"     AND e.pilar_id = a.pilar_id " +
		"     AND e.plano_id IN (" + planos + ") " +
		"     AND e.tipo_nota_id = d.tipo_nota_id " +
		"     AND e.componente_id = b.componente_id) " +
		" GROUP BY 1,2,3,4,5,6 ORDER BY 1,2,3,4,5,6" +
		" RETURNING id"
	log.Println(sqlStatement)
	produtoTipoNotaId := 0
	err = Db.QueryRow(
		sqlStatement,
		mdl.Calculada,
		currentUser.Id,
		time.Now()).Scan(&produtoTipoNotaId)
	if err != nil {
		log.Println(err)
	}

	sqlStatement = "INSERT INTO produtos_elementos ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" pilar_id, " +
		" componente_id, " +
		" plano_id, " +
		" tipo_nota_id, " +
		" elemento_id, " +
		" peso," +
		" nota," +
		" tipo_pontuacao_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT d.entidade_id, " +
		" d.ciclo_id, " +
		" d.pilar_id, " +
		" d.componente_id, " +
		" d.plano_id, " +
		" c.tipo_nota_id, " +
		" c.elemento_id, " +
		" c.peso_padrao, " +
		" 0, $1, $2, $3 " +
		" FROM " +
		" pilares_ciclos a " +
		" INNER JOIN " +
		" componentes_pilares b ON a.pilar_id = b.pilar_id " +
		" INNER JOIN " +
		" elementos_componentes c ON b.componente_id = c.componente_id " +
		" INNER JOIN " +
		" produtos_planos d ON b.componente_id = d.componente_id " +
		" WHERE " +
		" a.ciclo_id = " + strconv.FormatInt(param.CicloId, 10) +
		" AND d.entidade_id = " + strconv.FormatInt(param.EntidadeId, 10) +
		" AND d.pilar_id = " + strconv.FormatInt(param.PilarId, 10) +
		" AND d.componente_id = " + strconv.FormatInt(param.ComponenteId, 10) +
		" AND d.plano_id IN (" + planos + ") " +
		" AND NOT EXISTS " +
		"  (SELECT 1 " +
		"   FROM produtos_elementos e " +
		"   WHERE e.entidade_id = " + strconv.FormatInt(param.EntidadeId, 10) +
		"     AND e.ciclo_id = a.ciclo_id " +
		"     AND e.pilar_id = a.pilar_id " +
		"     AND e.componente_id = b.componente_id " +
		"     AND e.plano_id IN (" + planos + ") " +
		"     AND e.elemento_id = c.elemento_id) RETURNING id"
	log.Println(sqlStatement)
	produtoElementoId := 0
	err = Db.QueryRow(
		sqlStatement,
		mdl.Calculada,
		currentUser.Id,
		time.Now()).Scan(&produtoElementoId)
	if err != nil {
		log.Println(err)
	}

	sqlStatement = "INSERT INTO produtos_itens ( " +
		" entidade_id, " +
		" ciclo_id, " +
		" pilar_id, " +
		" componente_id, " +
		" plano_id, " +
		" tipo_nota_id, " +
		" elemento_id, " +
		" item_id, " +
		" author_id, " +
		" criado_em ) " +
		" SELECT p.entidade_id, " +
		" p.ciclo_id, " +
		" p.pilar_id, " +
		" p.componente_id, " +
		" p.plano_id, " +
		" c.tipo_nota_id, " +
		" c.elemento_id, d.id, $1, $2 " +
		" FROM pilares_ciclos a " +
		" INNER JOIN componentes_pilares b ON a.pilar_id = b.pilar_id " +
		" INNER JOIN elementos_componentes c ON b.componente_id = c.componente_id " +
		" INNER JOIN itens d ON c.elemento_id = d.elemento_id " +
		" INNER JOIN PRODUTOS_PLANOS p ON b.componente_id = p.componente_id " +
		" WHERE " +
		" a.ciclo_id = " + strconv.FormatInt(param.CicloId, 10) +
		" AND p.entidade_id = " + strconv.FormatInt(param.EntidadeId, 10) +
		" AND p.pilar_id = " + strconv.FormatInt(param.PilarId, 10) +
		" AND p.componente_id = " + strconv.FormatInt(param.ComponenteId, 10) +
		" AND p.plano_id IN ( " + planos + ") " +
		" AND NOT EXISTS (SELECT 1 " +
		"     FROM produtos_itens e " +
		"     WHERE e.entidade_id = " + strconv.FormatInt(param.EntidadeId, 10) +
		"       AND e.plano_id IN ( " + planos + ") " +
		"       AND e.ciclo_id = a.ciclo_id " +
		"       AND e.pilar_id = a.pilar_id " +
		"       AND e.componente_id = b.componente_id " +
		"	   AND e.elemento_id = c.elemento_id " +
		"	   AND e.item_id = d.id) RETURNING id"
	log.Println(sqlStatement)
	produtoItemId := 0
	err = Db.QueryRow(
		sqlStatement,
		currentUser.Id,
		time.Now()).Scan(&produtoItemId)
	if err != nil {
		log.Println(err)
	}

	log.Println("INICIANDO CICLO --  UPDATE NOTA")
	var produto mdl.ProdutoElemento
	produto.EntidadeId = param.EntidadeId
	produto.CicloId = param.CicloId
	produto.PilarId = param.PilarId
	produto.ComponenteId = param.ComponenteId
	produto.PlanoId = param.PlanoId
	// Atualizando os pesos
	atualizarPesoTiposNotas(produto, currentUser)
	atualizarPesoPlanos(produto, currentUser)
	atualizarPesoComponentes(produto, currentUser)
	sqlStatement = "UPDATE produtos_componentes a " +
		" set nota = (select  " +
		" sum(nota*peso)/sum(peso) as media " +
		" FROM produtos_elementos b " +
		" WHERE " +
		" a.entidade_id = b.entidade_id " +
		" and a.ciclo_id = b.ciclo_id  " +
		" and a.pilar_id = b.pilar_id " +
		" and a.componente_id = b.componente_id " +
		" GROUP BY b.entidade_id,  " +
		" b.ciclo_id, " +
		" b.pilar_id, " +
		" b.componente_id " +
		" HAVING sum(peso)>0) " +
		" WHERE a.entidade_id = $1 " +
		" AND a.ciclo_id = $2 "
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	updtForm.Exec(strconv.FormatInt(param.EntidadeId, 10), strconv.FormatInt(param.CicloId, 10))

	// PRODUTOS TIPOS NOTAS
	sqlStatement = "UPDATE produtos_tipos_notas a " +
		" set nota = (select  " +
		" sum(nota*peso)/sum(peso) as media " +
		" FROM produtos_elementos b " +
		" WHERE " +
		" a.entidade_id = b.entidade_id " +
		" and a.ciclo_id = b.ciclo_id  " +
		" and a.pilar_id = b.pilar_id " +
		" and a.componente_id = b.componente_id " +
		" and a.tipo_nota_id = b.tipo_nota_id " +
		" GROUP BY b.entidade_id,  " +
		" b.ciclo_id, " +
		" b.pilar_id, " +
		" b.componente_id, " +
		" b.tipo_nota_id " +
		" HAVING sum(peso)>0) " +
		" WHERE a.entidade_id = $1 " +
		" AND a.ciclo_id = $2 "
	log.Println(sqlStatement)
	updtForm, err = Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	updtForm.Exec(strconv.FormatInt(param.EntidadeId, 10), strconv.FormatInt(param.CicloId, 10))
	return produtoPlanoId
}

func loadNotasAtuais(produto mdl.ProdutoElemento) mdl.NotasAtuais {
	var notasAtuais mdl.NotasAtuais
	sql := " SELECT a.nota as tipo_nota, b.nota as plano, c.nota as componente, d.nota as pilar, e.nota as ciclo " +
		"    FROM produtos_tipos_notas a " +
		"    JOIN produtos_planos b ON (a.entidade_id = b.entidade_id  " +
		" 	AND a.ciclo_id = b.ciclo_id " +
		" 	AND a.pilar_id = b.pilar_id  " +
		" 	AND a.componente_id = b.componente_id " +
		" 	AND a.plano_id = b.plano_id) " +
		"    JOIN produtos_componentes c ON (a.entidade_id = c.entidade_id  " +
		" 	AND a.ciclo_id = c.ciclo_id " +
		" 	AND a.pilar_id = c.pilar_id  " +
		" 	AND a.componente_id = c.componente_id) " +
		"    JOIN produtos_pilares d ON (a.entidade_id = d.entidade_id  " +
		" 	AND a.ciclo_id = d.ciclo_id " +
		" 	AND a.pilar_id = d.pilar_id) " +
		"    JOIN produtos_ciclos e ON (a.entidade_id = e.entidade_id  " +
		" 	AND a.ciclo_id = e.ciclo_id) " +
		"    WHERE a.entidade_id = " + strconv.FormatInt(produto.EntidadeId, 10) +
		"      AND a.ciclo_id = " + strconv.FormatInt(produto.CicloId, 10) +
		"      AND a.pilar_id = " + strconv.FormatInt(produto.PilarId, 10) +
		"      AND a.componente_id = " + strconv.FormatInt(produto.ComponenteId, 10) +
		"      AND a.plano_id = " + strconv.FormatInt(produto.PlanoId, 10) +
		"      AND a.tipo_nota_id = " + strconv.FormatInt(produto.TipoNotaId, 10)
	log.Println(sql)
	rows, _ := Db.Query(sql)
	defer rows.Close()
	if rows.Next() {
		rows.Scan(
			&notasAtuais.TipoNotaNota,
			&notasAtuais.PlanoNota,
			&notasAtuais.ComponenteNota,
			&notasAtuais.PilarNota,
			&notasAtuais.CicloNota)
	}
	return notasAtuais
}

func loadPesosAtuais(produto mdl.ProdutoElemento) mdl.PesosAtuais {
	var pesosAtuais mdl.PesosAtuais
	sql := " SELECT a.peso as tipo_nota, b.peso as plano, c.peso as componente, d.peso as pilar " +
		"    FROM produtos_tipos_notas a " +
		"    JOIN produtos_planos b ON (a.entidade_id = b.entidade_id  " +
		" 	AND a.ciclo_id = b.ciclo_id " +
		" 	AND a.pilar_id = b.pilar_id  " +
		" 	AND a.componente_id = b.componente_id " +
		" 	AND a.plano_id = b.plano_id) " +
		"    JOIN produtos_componentes c ON (a.entidade_id = c.entidade_id  " +
		" 	AND a.ciclo_id = c.ciclo_id " +
		" 	AND a.pilar_id = c.pilar_id  " +
		" 	AND a.componente_id = c.componente_id) " +
		"    JOIN produtos_pilares d ON (a.entidade_id = d.entidade_id  " +
		" 	AND a.ciclo_id = d.ciclo_id " +
		" 	AND a.pilar_id = d.pilar_id) " +
		"    WHERE a.entidade_id = " + strconv.FormatInt(produto.EntidadeId, 10) +
		"      AND a.ciclo_id = " + strconv.FormatInt(produto.CicloId, 10) +
		"      AND a.pilar_id = " + strconv.FormatInt(produto.PilarId, 10) +
		"      AND a.componente_id = " + strconv.FormatInt(produto.ComponenteId, 10) +
		"      AND a.plano_id = " + strconv.FormatInt(produto.PlanoId, 10) +
		"      AND a.tipo_nota_id = " + strconv.FormatInt(produto.TipoNotaId, 10)
	log.Println(sql)
	rows, _ := Db.Query(sql)
	defer rows.Close()
	if rows.Next() {
		rows.Scan(
			&pesosAtuais.TipoNotaPeso,
			&pesosAtuais.PlanoPeso,
			&pesosAtuais.ComponentePeso,
			&pesosAtuais.PilarPeso)
	}
	return pesosAtuais
}

func registrarPesoPilar(produto mdl.ProdutoPilar) {
	// PESOS PILARES
	sqlStatement := "UPDATE produtos_pilares SET peso = $1, motivacao_peso = $2 " +
		" WHERE entidade_id = $3 AND " +
		" ciclo_id = $4 AND " +
		" pilar_id = $5 "
	log.Println(sqlStatement)
	updtForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}
	updtForm.Exec(produto.Peso,
		produto.Motivacao,
		produto.EntidadeId,
		produto.CicloId,
		produto.PilarId)
}

func getAnalise(rota string) string {
	valores := strings.Split(rota, "_")
	sql := " SELECT 'analise' "
	if valores[1] == "Ciclo" {
		sql = " SELECT analise FROM produtos_ciclos WHERE entidade_id = " + valores[2] + " AND ciclo_id = " + valores[3]
	} else if valores[1] == "Pilar" {
		sql = " SELECT analise FROM produtos_pilares WHERE entidade_id = " + valores[2] + " AND ciclo_id = " + valores[3] +
			" AND pilar_id = " + valores[4]
	} else if valores[1] == "Componente" {
		sql = " SELECT analise FROM produtos_componentes WHERE entidade_id = " + valores[2] + " AND ciclo_id = " + valores[3] +
			" AND pilar_id = " + valores[4] + " AND componente_id = " + valores[5]
	} else if valores[1] == "Plano" {
		sql = " SELECT analise FROM produtos_planos WHERE entidade_id = " + valores[2] + " AND ciclo_id = " + valores[3] +
			" AND pilar_id = " + valores[4] + " AND componente_id = " + valores[5] + " AND plano_id = " + valores[6]
	} else if valores[1] == "TipoNota" {
		sql = " SELECT analise FROM produtos_tipos_notas WHERE entidade_id = " + valores[2] + " AND ciclo_id = " + valores[3] +
			" AND pilar_id = " + valores[4] + " AND componente_id = " + valores[5] + " AND plano_id = " + valores[6] +
			" AND tipo_nota_id = " + valores[7]
	} else if valores[1] == "Elemento" {
		sql = " SELECT analise FROM produtos_elementos WHERE entidade_id = " + valores[2] + " AND ciclo_id = " + valores[3] +
			" AND pilar_id = " + valores[4] + " AND componente_id = " + valores[5] + " AND plano_id = " + valores[6] +
			" AND tipo_nota_id = " + valores[7] + " AND elemento_id = " + valores[8]
	} else if valores[1] == "Item" {
		sql = " SELECT analise FROM produtos_itens WHERE entidade_id = " + valores[2] + " AND ciclo_id = " + valores[3] +
			" AND pilar_id = " + valores[4] + " AND componente_id = " + valores[5] + " AND plano_id = " + valores[6] +
			" AND tipo_nota_id = " + valores[7] + " AND elemento_id = " + valores[8] + " AND item_id = " + valores[9]
	}

	log.Println(sql)
	rows, _ := Db.Query(sql)
	defer rows.Close()
	retorno := ""
	if rows.Next() {
		rows.Scan(&retorno)
	}
	return retorno
}

func setAnalise(rota string, analise string) string {
	valores := strings.Split(rota, "_")
	sqlStatement := " SELECT 'analise' "
	if valores[1] == "Ciclo" {
		sqlStatement = " UPDATE produtos_ciclos SET analise = '" + analise + "' WHERE entidade_id = " + valores[2] + " AND ciclo_id = " + valores[3]
	} else if valores[1] == "Pilar" {
		sqlStatement = " UPDATE produtos_pilares SET analise = '" + analise + "' WHERE entidade_id = " + valores[2] + " AND ciclo_id = " + valores[3] +
			" AND pilar_id = " + valores[4]
	} else if valores[1] == "Componente" {
		sqlStatement = " UPDATE produtos_componentes SET analise = '" + analise + "' WHERE entidade_id = " + valores[2] + " AND ciclo_id = " + valores[3] +
			" AND pilar_id = " + valores[4] + " AND componente_id = " + valores[5]
	} else if valores[1] == "Plano" {
		sqlStatement = " UPDATE produtos_planos SET analise = '" + analise + "' WHERE entidade_id = " + valores[2] + " AND ciclo_id = " + valores[3] +
			" AND pilar_id = " + valores[4] + " AND componente_id = " + valores[5] + " AND plano_id = " + valores[6]
	} else if valores[1] == "TipoNota" {
		sqlStatement = " UPDATE produtos_tipos_notas SET analise = '" + analise + "' WHERE entidade_id = " + valores[2] + " AND ciclo_id = " + valores[3] +
			" AND pilar_id = " + valores[4] + " AND componente_id = " + valores[5] + " AND plano_id = " + valores[6] +
			" AND tipo_nota_id = " + valores[7]
	} else if valores[1] == "Elemento" {
		sqlStatement = " UPDATE produtos_elementos SET analise = '" + analise + "' WHERE entidade_id = " + valores[2] + " AND ciclo_id = " + valores[3] +
			" AND pilar_id = " + valores[4] + " AND componente_id = " + valores[5] + " AND plano_id = " + valores[6] +
			" AND tipo_nota_id = " + valores[7] + " AND elemento_id = " + valores[8]
	} else if valores[1] == "Item" {
		sqlStatement = " UPDATE produtos_itens SET analise = '" + analise + "' WHERE entidade_id = " + valores[2] + " AND ciclo_id = " + valores[3] +
			" AND pilar_id = " + valores[4] + " AND componente_id = " + valores[5] + " AND plano_id = " + valores[6] +
			" AND tipo_nota_id = " + valores[7] + " AND elemento_id = " + valores[8] + " AND item_id = " + valores[9]
	}
	log.Println(sqlStatement)
	updtForm, _ := Db.Prepare(sqlStatement)
	_, err := updtForm.Exec()
	if err != nil {
		log.Println(err.Error())
	}
	return "OK"
}

func getDescricao(rota string) mdl.Descricao {
	valores := strings.Split(rota, "_")
	sql := " SELECT 'descricao' "
	if valores[1] == "Ciclo" {
		sql = " SELECT descricao, referencia FROM ciclos WHERE id = " + valores[3]
	} else if valores[1] == "Pilar" {
		sql = " SELECT descricao, referencia FROM pilares WHERE id = " + valores[4]
	} else if valores[1] == "Componente" {
		sql = " SELECT descricao, referencia FROM componentes WHERE id = " + valores[5]
	} else if valores[1] == "Plano" {
		sql = " SELECT descricao, referencia FROM planos WHERE id = " + valores[6]
	} else if valores[1] == "TipoNota" {
		sql = " SELECT descricao, referencia FROM tipos_notas WHERE id = " + valores[7]
	} else if valores[1] == "Elemento" {
		sql = " SELECT descricao, referencia FROM elementos WHERE id = " + valores[8]
	} else if valores[1] == "Item" {
		sql = " SELECT descricao, referencia FROM itens WHERE id = " + valores[9]
	}
	log.Println(sql)
	rows, _ := Db.Query(sql)
	defer rows.Close()
	var retorno mdl.Descricao
	if rows.Next() {
		rows.Scan(&retorno.Texto, &retorno.Link)
	}
	return retorno
}

func loadConfigPlanos(entidadeId string, cicloId string, pilarId string, componenteId string) string {
	sql := "SELECT planos_configurados FROM (SELECT a.componente_id, string_agg(b.cnpb,', ') planos_configurados " +
		" FROM produtos_planos a " +
		" INNER JOIN planos b ON a.plano_id = b.id " +
		" WHERE a.entidade_id = " + entidadeId + " AND a.ciclo_id = " + cicloId +
		" AND a.pilar_id = " + pilarId + " AND a.componente_id = " + componenteId +
		" GROUP BY 1) R "
	log.Println(sql)
	rows, _ := Db.Query(sql)
	defer rows.Close()
	var configuracaoPlanos string
	if rows.Next() {
		rows.Scan(&configuracaoPlanos)
	}
	return configuracaoPlanos
}
