var plano_tobe_deleted;

function updateConfigPlanos(){
	document.getElementById('config-planos-form').style.display='none';
	document.getElementById('motivar-reconfiguracao-form').style.display='none';
	let motivo = document.getElementById('motReconf_text').value;
	let splitted = document.getElementById('AcionadoPor').value.split('_');
	let entidadeId = splitted[1];
	let cicloId = splitted[2];
	let pilarId = splitted[3];
	let componenteId = splitted[4];
	let auditorId = splitted[5];
	let planosSelecionados = document.getElementById('ConfigPlanos');
	let selecionados = getSelectedOptions(planosSelecionados);
	let valores = '';
	for(n=0;n<selecionados.length;n++){
		valores = selecionados[n].value+'_'+ valores;
	}
	console.log("Planos_AuditorComponente_"+entidadeId+"_"+cicloId+"_"+pilarId+"_"+componenteId+"_"+auditorId);
	// atualizar os planos através de Ajax
	atualizarConfigPlanos(entidadeId, cicloId, pilarId, componenteId, valores, false, motivo);
	document.getElementsByName("Planos_AuditorComponente_"+entidadeId+"_"+cicloId+"_"+pilarId+"_"+componenteId+"_"+auditorId)[0].value = valores;
	console.log("Saindo");
}

function atualizarConfigPlanos(entidadeId, cicloId, pilarId, componenteId, valores, superUser, motivacao){
	console.log("atualizarConfigPlanos");
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var messageText = xmlhttp.responseText;
				if( ! messageText.includes("não pode ser removido")){
					callback_MudaCorBotaoAcionador();
				} else {
					if(isChefe){
						superUser = true;
						alterarOkConfirm(entidadeId, cicloId, pilarId, componenteId, valores, superUser);
					}
				}
				if(messageText!=''){
					document.getElementById("messageText").innerText = messageText;
					document.getElementById("message").style.display="block";		
				}
			}
	}
	xmlhttp.open("GET","/updateConfigPlanos?entidadeId="+
			entidadeId+"&cicloId="+cicloId+"&pilarId="+
			pilarId+"&componenteId="+componenteId+"&planos="+
			valores+"&superUser="+superUser+
			"&motivacao="+motivacao, true);
	xmlhttp.send();
}

var minhaFuncaoBtnOk;

function alterarOkConfirm(entidadeId, cicloId, pilarId, componenteId, valores, superUser){
	let divElement = document.getElementById('divOKConfirm');
	let btns = divElement.getElementsByTagName("button");
	existeBtnNao = false;
	let btnNao; 
	for(i=0;i<btns.length;i++){
		if(btns[i].innerText == "Não"){
			existeBtnNao = true;
			btnNao = btns[i]; 
			break;
		}
	}
	if(!existeBtnNao){
		btnNao = document.createElement("button");
		btnNao.id = "__btnNao";
		btnNao.classList = btns[0].classList;
		minhaFuncaoBtnSim = btns[0].onclick
		btnNao.onclick = minhaFuncaoBtnSim;
		btnNao.innerText = "Não"; 
		btnNao.style = "padding: 5px; margin: 5px;"; 
	}
	btns[0].innerText = "Sim";
	btns[0].addEventListener("click", function() {
	  	atualizarConfigPlanos(entidadeId,cicloId,pilarId,componenteId,valores,superUser);
		document.getElementById("__btnNao");
		btnNao.remove();
		btnOk = this.cloneNode(true);
		btnOk.innerText = "OK";
		btnOk.addEventListener("click",minhaFuncaoBtnOk);
		divElement.appendChild(btnOk);
		this.remove();
	});
	divElement.appendChild(btnNao)
}

function callback_MudaCorBotaoAcionador(){
	let selecionados = document.getElementById("ConfigPlanos").value;
	let acionadoPor = document.getElementById("AcionadoPor").value;
	console.log(document.getElementsByName(acionadoPor)[0]);
	let classList = document.getElementsByName(acionadoPor)[0].classList;
	if(selecionados.length>0){
		classList.remove('w3-red');
		classList.add('w3-green');
	} else {
		classList.remove('w3-green');
		classList.add('w3-red');
	}
}

function updateTodosConfigPlanos(){
	document.getElementById('config-planos-form').style.display='none';
	let planosSelecionados = document.getElementById('ConfigPlanos');
	let selecionados = getSelectedOptions(planosSelecionados);
	let valores = '';
	for(n=0;n<selecionados.length;n++){
		valores = selecionados[n].value+'_'+ valores;
	}
	let col = document.getElementsByTagName("input");
	for(n=0;n<col.length;n++){
		console.log(col[n].name);
		if(col[n].name.startsWith("Planos_AuditorComponente_")){
			col[n].value = valores;
		}
	}
	col = document.getElementsByTagName("button");
	for(n=0;n<col.length;n++){
		if(col[n].name.startsWith("BtnPlanos_")){
			document.getElementsByName(col[n].name)[0].classList.remove('w3-red');
			document.getElementsByName(col[n].name)[0].classList.add('w3-green');
		}
	}
}

function getSelectedOptions(sel) {
    var opts = [], opt;
    for (var i=0, len=sel.options.length; i<len; i++) {
        opt = sel.options[i];
        if ( opt.selected ) {
            opts.push(opt);
        }
    }
    return opts;
}

function openConfigPlanos(btn){
	document.getElementById("AcionadoPor").value = btn.name;
	//btn.disabled = true;
	let entidadeId = btn.name.split("_")[1];
	//console.log(entidadeId);
	let cicloId = btn.name.split("_")[2];
	//console.log(cicloId);
	let pilarId = btn.name.split("_")[3];
	//console.log(pilarId);
	let componenteId = btn.name.split("_")[4];
	//console.log(componenteId);
	//console.log(document.getElementById('Planos_AuditorComponente'+componenteId));
	document.getElementById('ConfigPlanos').selectedIndex = -1;
	let soPGA = document.getElementById('SomentePGA_'+entidadeId+'_'+cicloId+'_'+pilarId+'_'+componenteId).value;
	configurarOpcoesNaoPGA('ConfigPlanos', soPGA);
	document.getElementById('EntidadeConfigPlanos').value = entidadesMap.get(entidadeId);
	document.getElementById('CicloConfigPlanos').value = ciclosMap.get(cicloId);
	document.getElementById('PilarConfigPlanos').value = pilaresMap.get(pilarId);
	document.getElementById('ComponenteConfigPlanos').value = componentesMap.get(componenteId);
	document.getElementById('ComponenteRefConfigPlanos').value = entidadeId+'_'+cicloId+'_'+pilarId+'_'+componenteId;
	document.getElementById('config-planos-form').style.display='block';
	loadConfigPlanos(entidadeId, cicloId, pilarId, componenteId);
}

function configurarOpcoesNaoPGA(nomeCampoSelect, soPGA){
	let campo = document.getElementById(nomeCampoSelect);
	let opcoes = campo.options;
	for(i=0;i<opcoes.length;i++){
		txt = opcoes[i].text;
		console.log('cnpb: '+txt);
		if(txt.substr(0,1)!='9'){
			if(soPGA == 'S'){
				console.log('S');
				opcoes[i].disabled = true;
				opcoes[i].style = "display:none";
			} else {
				console.log('N');
				opcoes[i].disabled = false;
				opcoes[i].style = "display:outline";
			}
		}
	}
}

function loadConfigPlanos(entidadeId, cicloId, pilarId, componenteId){
	var xmlhttp;
	xmlhttp=new XMLHttpRequest();
	xmlhttp.onreadystatechange=function()
	{
			if (xmlhttp.readyState==4 && xmlhttp.status==200)
			{
				var configPlanosJson = JSON.parse(xmlhttp.responseText);
				var cfg = document.getElementById('ConfigPlanos');
				if(configPlanosJson.length == 0){
					callback_MudaCorBotaoAcionador();
				} else {
					for(i=0;i<configPlanosJson.length;i++){
						for(j=0;j<cfg.options.length;j++){
							if(cfg.options[j].value == configPlanosJson[i].planoId){
								cfg.options[j].selected = true; 
								console.log(configPlanosJson[i]+' :: '+cfg.options[j].selected);
							}
						}
					}
				}
				return planos;
			}
	}
	xmlhttp.open("GET","/loadConfigPlanos?entidadeId="+entidadeId+"&cicloId="+cicloId+"&pilarId="+pilarId+"&componenteId="+componenteId,true);
	xmlhttp.send();
}

class ConfigPlano {
	constructor(id, entidadeId, planoId) {
		this.id = id;
		this.entidadeId = entidadeId;
		this.planoId = planoId;
	}
}
	
class Plano {
	constructor(order, id, entidadeId, nome, descricao, cnpb, c_recursoGarantidor, recursoGarantidor, modalidade, autorId, autorNome, criadoEm, status, cStatus) {
		this.order = order;
		this.id = id;
		this.entidadeId = entidadeId;
		this.nome = nome;
		this.descricao = descricao;
		this.cnpb = cnpb;
		this.c_recursoGarantidor = c_recursoGarantidor;
		this.recursoGarantidor = recursoGarantidor;
		this.modalidade = modalidade;
		this.autorId = autorId;
		this.autorNome = autorNome;
		this.criadoEm = criadoEm;
		this.status = status;
		this.cStatus = cStatus;
	}
}

function criarPlano(){
	console.log('criarPlano');
	var cnpb = document.getElementById('CNPBPlanoForInsert').value;
	var nome = document.getElementById('NomePlanoForInsert').value;
	var modalidade = document.getElementById('ModalidadePlanoForInsert').value;
	var recursoGarantidor = document.getElementById('RecursoGarantidorPlanoForInsert').value;
	var descricao = document.getElementById('DescricaoPlanoForInsert').value;
	var erros = '';
	if(cnpb==''){
		erros += 'Falta preencher o CNPB.\n';
		alert(erros);
		return;
	}
	planoId = getMaxId(planos);
	plano = new Plano(0, planoId, 0, nome, descricao, cnpb, '', recursoGarantidor, modalidade,'','','','','');
	planos.push(plano);
	console.log('contexto: '+contexto);
	//console.log("table-planos-"+contexto);
	addPlanoRow("table-planos-"+contexto);
	limparCamposPlanoForm();
	document.getElementById('create-plano-form').style.display='none';
}

function addPlanoRow(tableID) {
	console.log('addPlanoRow');
	let tableRef = document.getElementById(tableID);
	let newRow = tableRef.childNodes[1].insertRow(-1);
	order = planos.length-1;
	plano = planos[order];
	let newCell = newRow.insertCell(0);
	let newText = document.createTextNode(plano.cnpb);
	let rg = plano.c_recursoGarantidor;
	plano.c_recursoGarantidor = '';
	let json = JSON.stringify(plano);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	newCell.appendChild(newText);
	newCell.style = 'vertical-align:middle';
	newCell.innerHTML = '<input type="hidden" name="plano'+plano.id+'" value="'+json+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="entidadeId" value="'+plano.entidadeId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="id" value="'+plano.id+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+newCell.innerHTML;
	// modalidade
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(plano.modalidade);
	newCell.appendChild(newText);
	newCell.style = 'vertical-align:middle';
	// recurso garantidor
	newCell = newRow.insertCell(2);
	if(plano.c_recursoGarantidor == ''){
		plano.c_recursoGarantidor = plano.recursoGarantidor;
	}
	newText = document.createTextNode(rg);
	newCell.appendChild(newText);
	newCell.style = 'vertical-align:middle';
	// nome
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(plano.nome);
	newCell.appendChild(newText);
	newCell.style = 'display:none';
	// descricao
	newCell = newRow.insertCell(4);
	newText = document.createTextNode(plano.descricao);
	newCell.appendChild(newText);
	newCell.style = 'display:none';
	// recursoGarantidor NUMERO
	newCell = newRow.insertCell(5);
	newText = document.createTextNode(plano.recursoGarantidor);
	newCell.appendChild(newText);
	newCell.style = 'display:none';
	// Botões
	newCell = newRow.insertCell(6);
	newCell.style = 'vertical-align:middle';
	// Botão Editar
	let btnEditar = document.createElement('input');
	btnEditar.type = "button";
	btnEditar.className = "w3-btn w3-teal";
	btnEditar.style = "margin-right: 10px";
	btnEditar.value = "Editar";
	btnEditar.onclick = function() {editPlano(btnEditar)};
	newCell.appendChild(btnEditar);
	// Botão Apagar
	let btnApagar = document.createElement('input');
	btnApagar.type = "button";
	btnApagar.className = "w3-btn w3-red";
	btnApagar.value = "Apagar";
	btnApagar.onclick = function() {showDeletePlanoForm(btnApagar)};
	newCell.appendChild(btnApagar);
}

function editPlano(e) {
	console.log('editPlano');
	var editPlanoForm = document.getElementById('edit-plano-form');
	editPlanoForm.style.display = 'block';
	var order = e.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var id = e.parentNode.parentNode.childNodes[0].childNodes[1].value;
	var entidadeId = e.parentNode.parentNode.childNodes[0].childNodes[2].value;
	var cnpb = e.parentNode.parentNode.childNodes[0].innerText;
	var modalidade = e.parentNode.parentNode.childNodes[1].innerText;
	var recursoGarantidor = e.parentNode.parentNode.childNodes[5].innerText;
	var nome = e.parentNode.parentNode.childNodes[3].innerText;
	var descricao = e.parentNode.parentNode.childNodes[4].innerText;
	// Atribuindo os valores de edit-item-form
	document.getElementById('id-edit').value=id;
	document.getElementById('order-edit').value=order;
	document.getElementById('entidadeId-edit').value=entidadeId;
	document.getElementById('NomePlanoForUpdate').value=nome;
	document.getElementById('CNPBPlanoForUpdate').value=cnpb;
	document.getElementById('RecursoGarantidorPlanoForUpdate').value=recursoGarantidor;
	document.getElementById('ModalidadePlanoForUpdate').value=modalidade;
	document.getElementById('DescricaoPlanoForUpdate').value=descricao;
}

function updatePlano() {
	console.log('updatePlano');
	var id = document.getElementById('id-edit').value;
	var order = document.getElementById('order-edit').value;
	var entidadeId = document.getElementById('entidadeId-edit').value;
	var cnpb = document.getElementById('CNPBPlanoForUpdate').value;
	var nome = document.getElementById('NomePlanoForUpdate').value;
	var modalidade = document.getElementById('ModalidadePlanoForUpdate').value;
	var recursoGarantidor = document.getElementById('RecursoGarantidorPlanoForUpdate').value;
	var descricao = document.getElementById('DescricaoPlanoForUpdate').value;
	var erros = '';
	if(cnpb==''){
		erros += 'Falta preencher o CNPB.\n';
		alert(erros);
		return;
	}
	plano = new Plano(order, id, entidadeId, nome, descricao, cnpb, '', recursoGarantidor, modalidade,'','','','','');
	planos[order] = plano;
	console.log("table-planos-"+contexto);
	console.log('order: '+order);
	updatePlanoRow("table-planos-"+contexto,order);
	limparCamposPlanoForm();
	document.getElementById('edit-plano-form').style.display='none';
}

function updatePlanoRow(tableID, order){
	console.log('updatePlanoRow');
	let tbl = document.getElementById(tableID);
	let linhas = tbl.childNodes[1].childNodes;
	let row = tbl.childNodes[0];
	for(y=0;y<linhas.length;y++){
		if(linhas[y].childNodes[0]){
			let inputOrder = linhas[y].childNodes[0].childNodes[0];
			if(inputOrder && inputOrder.tagName=='INPUT'){ 
				if(inputOrder.value==order){
					row = linhas[y];
					break;
				}
			}
		}
	}
	let celula = row.childNodes[0];
	celula.innerText = planos[order].cnpb;
	var json = JSON.stringify(planos[order]);
	json = json.split(',').join('#');
	json = json.split('"').join('');
	json = json.split('{').join('');
	json = json.split('}').join('');
	celula.innerHTML = '<input type="hidden" name="plano'+order+'" value="'+json+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="entidadeId" value="'+plano.entidadeId+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="id" value="'+planos[order].id+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+celula.innerHTML;
	row.childNodes[1].innerText = planos[order].modalidade;
	if(planos[order].c_recursoGarantidor == ''){
		plano.c_recursoGarantidor = plano.recursoGarantidor;
	}
	row.childNodes[2].innerText = planos[order].c_recursoGarantidor;
	row.childNodes[3].innerText = planos[order].nome;
	row.childNodes[4].innerText = planos[order].descricao;
	row.childNodes[5].innerText = planos[order].recursoGarantidor;
}

function showDeletePlanoForm(e){
	console.log('showDeletePlanoForm');
	var deletePlanoForm = document.getElementById('delete-plano-form');
	deletePlanoForm.style.display = 'block';
	plano_tobe_deleted = e;
}

function deletePlano() {
	console.log('deletePlano');
	var order = plano_tobe_deleted.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var newPlanos = [];
	let tbl = plano_tobe_deleted.parentNode.parentNode.parentNode;
	let linhas = tbl.childNodes;
	contadorLinha = 1;
	for(y=0;y<linhas.length;y++){
		if(linhas[y].childNodes[0]){
			let inputOrder = linhas[y].childNodes[0].childNodes[0];
			if(inputOrder && inputOrder.tagName=='INPUT'){ 
				if(inputOrder.value==order){
					if(inputOrder.value == order){
						tbl.deleteRow(contadorLinha);
						break;
					}
				}
				contadorLinha ++;
			}
		}
	}
	for(i=0;i<planos.length;i++){
		if(i != order){
			newPlanos.push(planos[i]);
		}
	}
	planos = newPlanos;
	var deletePlanoForm = document.getElementById('delete-plano-form');
	deletePlanoForm.style.display = 'none';
}


function limparCamposPlanoForm(){
	console.log('limparCamposPlanoForm');
	document.getElementById('formulario-plano-create').reset()
	document.getElementById('formulario-plano-edit').reset()
}

function validarPlanosSelecionados(){
	let col = document.getElementsByTagName("input");
	for(n=0;n<col.length;n++){
		console.log(col[n].name);
		if(col[n].name.startsWith("Planos_AuditorComponente_")){
			if(col[n].value == ''){
				return false;
			}
		}
	}
}
