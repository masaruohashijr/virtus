sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'alessandro', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ALESSANDRO VAINE', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'alessandro')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'alexandre', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ALEXANDRE BISPO NUNES GRECO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'alexandre')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'alfredo', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ALFREDO SULZBACHER WONDRACEK', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'alfredo')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'alvaro', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ALVARO RODRIGUES DOS SANTOS NETO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'alvaro')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'andre.silva', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ANDRE LUIS SOUZA DA SILVA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'andre.silva')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'andre.goncalves', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ANDRE MACHADO GONÇALVES', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'andre.goncalves')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'angelica', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ANGELICA DE ALMEIDA FONSECA CAMPINHO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'angelica')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'annette', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ANNETTE LOPES PINTO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'annette')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'antonio.portes', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ANTONIO ALBERTO GROSSI PORTES', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'antonio.portes')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'antonio.garcia', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ANTONIO AUGUSTO GARCIA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'antonio.garcia')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'antonio.frainer', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ANTONIO SEVERO FRAINER', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'antonio.frainer')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'arnaldo', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ARNALDO FERREIRA LEITE BURLE NETO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'arnaldo')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'carlos.silveira', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'CARLOS EDUARDO BUCCOS SILVEIRA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'carlos.silveira')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'carlos.neves', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'CARLOS EUGÊNIO D'AVILA NEVES', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'carlos.neves')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'carlos', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'CARLOS MARNE DIAS ALVES', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'carlos')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'charles.silva', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'CHARLES DINIZ BOTELHO DA SILVA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'charles.silva')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'charles.dantas', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'CHARLES SILVA DANTAS', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'charles.dantas')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'chow', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'CHOW CHI KWAN', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'chow')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'clovis', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'CLOVIS GUIMARAES COELHO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'clovis')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'dagomar', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'DAGOMAR ALECIO ANHE', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'dagomar')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'dauto', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'DAUTO BARBOSA DE SOUSA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'dauto')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'david', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'DAVID PRATES COUTINHO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'david')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'delma', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'DELMA PIRES GALVÃO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'delma')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'diogo', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'DIOGO BORBA DE ARAUJO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'diogo')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'douglas', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'DOUGLAS CARVALHO DUARTE', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'douglas')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'eduardo', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'EDUARDO MENEZES MEIRELES', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'eduardo')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'eliane', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ELIANE OLIVEIRA DA COSTA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'eliane')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'elyson', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ELYSON MARCEL TOME SCAFATI', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'elyson')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'enaide', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ENAIDE MARIA TEIXEIRA DE SOUZA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'enaide')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'estevam', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ESTEVAM BRAYN', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'estevam')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'evelyn', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'EVELYN DE QUEIROZ ITO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'evelyn')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'felipe', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'FELIPE SPOLAVORI MARTINS', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'felipe')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'francisco.coelho', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'FRANCISCO HELIO ARRUDA COELHO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'francisco.coelho')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'francisco.junior', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'FRANCISCO RODRIGUES BRAGA JÚNIOR', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'francisco.junior')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'germano', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'GERMANO CAVALCANTI BARREIRA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'germano')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'giselle', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'GISELLE LOPES MIRANDA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'giselle')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'hamilton', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'HAMILTON NOLETO MOREIRA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'hamilton')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'helvio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'HELVIO ANTONIO PEREIRA MARINHO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'helvio')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'hilton', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'HILTON DE ENZO MITSUNAGA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'hilton')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'humberto', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'HUMBERTO DA SILVA JUNIOR', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'humberto')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'isabel', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ISABEL CRISTINA MAIA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'isabel')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'izabel', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'IZABEL CRISTINA REZENDE NEVES', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'izabel')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'james', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'JAMES TAYLOR FARIA CHAVES', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'james')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'jorge', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'JORGE LUIZ FONSECA FRISCHEISEN', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'jorge')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'jose.chedeak', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'JOSÉ CARLOS SAMPAIO CHEDEAK', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'jose.chedeak')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'jose.pires', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'JOSE ERALDO PIRES', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'jose.pires')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'jose.cestari', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'JOSE MARCOS CARVALHO CESTARI', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'jose.cestari')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'jose.fernanes', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'JOSÉ RICARDO FERREIRA FERNANDES', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'jose.fernanes')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'jucinea', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'JUCINEA DAS MERCES NASCIMENTO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'jucinea')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'juliana', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'JULIANA HELENA DE PAIVA PEREIRA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'juliana')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'leonardo', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'LEONARDO FROTA ALVES', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'leonardo')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'luciano.draghetti', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'LUCIANO DRAGHETTI', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'luciano.draghetti')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'luciano.pinheiro', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'LUCIANO VILELA PINHEIRO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'luciano.pinheiro')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'luis.pugliese', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'LUIS ANTONIO ALVES PUGLIESE', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'luis.pugliese')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'luis.barbosa', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'LUIS GUSTAVO DA CUNHA BARBOSA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'luis.barbosa')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'luis.angoti', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'LUÍS RONALDO MARTINS ANGOTI', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'luis.angoti')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'luiz', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'LUIZ ALBERTO GONÇALVES FIALHO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'luiz')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'maique', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'MAIQUE PEREIERA AGNES', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'maique')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'marcelo.melo', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'MARCELO FREITAS TOLEDO DE MELO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'marcelo.melo')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'marcelo', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'MARCELO ZELIK WAJSENZON', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'marcelo')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'marcia', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'MÁRCIA VIVAS DE ARAÚJO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'marcia')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'marcio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'MARCIO MATHEUS GUIMARAES MACHADO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'marcio')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'marcus', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'MARCUS VINÍCIUS PINTO DE SOUZA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'marcus')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'maria', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'MARIA BATISTA DA SILVA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'maria')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'maria.silva', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'MARIA DA GLÓRIA FERREIRA PIMENTA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'maria.silva')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'maria.pimenta', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'MARIA JULIETA CHERULLI MACHADO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'maria.pimenta')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'marina', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'MARINA MARCONDES DA SILVA ALVES', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'marina')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'mauricio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'MAURICIO DE AGUIRRE NAKATA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'mauricio')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'mauricio.lundgren', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'MAURICIO TIGRE VALOIS LUNDGREN', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'mauricio.lundgren')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'maury', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'MAURY COELHO DE OLIVEIRA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'maury')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'nercilia', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'NERCILIA BARROS DE SOUZA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'nercilia')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'nivea', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'NIVEA CLEIDE FERREIRA DOS SANTOS', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'nivea')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'otavio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'OTAVIO LIMA REIS', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'otavio')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'patricia', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'PATRICIA CERQUEIRA MONTEIRO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'patricia')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'paulo.matsumoto', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'PAULO ANDRÉ HIDEAKI MATSUMOTO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'paulo.matsumoto')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'paulo.diniz', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'PAULO NOBILE DINIZ', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'paulo.diniz')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'paulo.vitorino', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'PAULO VITORINO SILVA DE SOUSA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'paulo.vitorino')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'pedro', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'PEDRO IWAO KAKITANI', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'pedro')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'pedro.eugenio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'PEDRO PAULO EUGENIO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'pedro.eugenio')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'peterson', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'PETERSON GONCALVES', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'peterson')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'rafael', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'RAFAEL MAGALHÃES', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'rafael')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'raquel', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'RAQUEL GERHARDT', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'raquel')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'rita', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'RITA DE CASSIA CORREA DA SILVA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'rita')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'roberto', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ROBERTO SAKAMOTO', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'roberto')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'rodrigo.oliveira', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'RODRIGO AIRES DE OLIVEIRA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'rodrigo.oliveira')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'rodrigo.abreu', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'RODRIGO MARTINS ABREU', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'rodrigo.abreu')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'rodrigo.andrade', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'RODRIGO RANGEL DE ANDRADE', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'rodrigo.andrade')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'romulo', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'ROMULO GONÇALVES DA SILVA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'romulo')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'sergio', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'SERGIO DJUNDI TANIGUCHI', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'sergio')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'simone', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'SIMONE CORRÊA REIS', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'simone')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'vandeisa', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'VANDEISA MOURA ALMEIDA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'vandeisa')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'vanessa', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'VANESSA DA ROCHA VIER', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'vanessa')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'veronica', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'VERONICA SOUSA SILVEIRA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'veronica')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'vitor', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'VITOR FERNANDES RIBEIRO DE OLIVEIRA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'vitor')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'walter', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'WALTER DE CARVALHO PARENTE', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'walter')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'wander', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'WANDER RICARDO MINGARDI', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'wander')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'wania', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'WÂNIA MARIA FRANÇA CAPPARELLI', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'wania')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'wellington.marques', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'WELINGTON RODRIGUES MARQUES', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'wellington.marques')"
 db.Exec(sql)
sql = "INSERT INTO virtus.users (username, password, email, mobile, name, id_role, id_author, criado_em) SELECT 'wellington.pereira', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy','', '', 'WELLINGTON MARCOS ASSIS PEREIRA', 4, 1, GETDATE() WHERE NOT EXISTS (SELECT id FROM virtus.users WHERE username = 'wellington.pereira')"