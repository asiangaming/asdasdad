<script>
    import { Input } from "sveltestrap";
	import Button from "../../components/Button.svelte";



	export let table_header_font
	export let table_body_font
	export let token = ""
	export let tokenprediksi = ""
	export let listHome = []
	export let totalrecord = 0
    let sData = "";
    let result = [];
    let pasaran = "";
    let nomorprediksi = "";
    let subtotal_company = 0;
    let subtotal_company_css = "";
    let css_loader = "display: none;";
    let msgloader = "";
   
    async function handleSave() {
        result = [];
        let flag = true
        let msg = ""
        if(pasaran == ""){
            flag = false
            msg = "The Pasaran is required"
        }
        if(nomorprediksi == ""){
            flag = false
            msg = "The Nomor Prediksi is required"
        }
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/prediksi", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    master:"sdsb",
                    token:tokenprediksi,
                    idpasarantogel:pasaran,
                    nomorprediksi:nomorprediksi,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                msgloader = json.message;
                let record = json.record;
                let no = 0
                subtotal_company= 0;
                for (var i = 0; i < record.length; i++) {
                    let css_bet = ""
                    let css_subtotal = ""
                    let css_subtotalwin = ""
                    let company_total = parseInt(record[i]["prediksi_subtotalwin"]);
                    subtotal_company = subtotal_company + company_total;
                    if(parseInt(record[i]["prediksi_totalbet"]) > 0 ){
                        css_bet = "color:blue;font-weight:bold;";
                    }else{
                        css_bet = "color:red;font-weight:bold;";
                    }
                    if(parseInt(record[i]["prediksi_subtotal"]) > 0 ){
                        css_subtotal = "color:blue;font-weight:bold;";
                    }else{
                        css_subtotal = "color:red;font-weight:bold;";
                    }
                    if(parseInt(record[i]["prediksi_subtotalwin"]) > 0 ){
                        css_subtotalwin = "color:blue;font-weight:bold;";
                    }else{
                        css_subtotalwin = "color:red;font-weight:bold;";
                    }
                    no = no + 1;
                    result = [
                        ...result,
                        {
                            prediksi_no: no,
                            prediksi_idcompany: record[i]["prediksi_idcompany"],
                            prediksi_nmcompany: record[i]["prediksi_nmcompany"],
                            prediksi_invoice: record[i]["prediksi_invoice"],
                            prediksi_invoicedate: record[i]["prediksi_invoicedate"],
                            prediksi_invoiceperiode: record[i]["prediksi_invoiceperiode"],
                            prediksi_totalbet: record[i]["prediksi_totalbet"],
                            prediksi_totalbetcss: css_bet,
                            prediksi_subtotal: record[i]["prediksi_subtotal"],
                            prediksi_subtotalcss: css_subtotal,
                            prediksi_subtotalwin: record[i]["prediksi_subtotalwin"],
                            prediksi_subtotalwincss: css_subtotalwin,
                        },
                    ];
                }
                if(parseInt(subtotal_company)>0){
                    subtotal_company_css = "color:blue;font-weight:bold;"
                }else{
                    subtotal_company_css = "color:red;font-weight:bold;"
                }
                
            } else if(json.status == 403){
                alert(json.message)
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
   
    function generate(field){
        let numbergenerate = (Math.floor(Math.random() * 10000) + 10000).toString().substring(1);
        switch(field){
            case "nomorprediksi":
                 nomorprediksi = numbergenerate
                break;
        }
    }
    function callFunction(event){
        switch(event.detail){
            case "GENERATE":
             handleSave();break;
        }
    }
    const handleKeyboard_format = () => {
		let numbera;
		for (let i = 0; i < nomorprediksi.length; i++) {
			numbera = parseInt(nomorprediksi[i]);
			if (isNaN(numbera)) {
				nomorprediksi = "";
			}
		}
        
    }
</script>
<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container-fluid" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-3">
            <div class="card mt-1">
                <div class="card-header">
                    <div class="float-end">
                        <Button
                            on:click={callFunction}
                            button_function="GENERATE"
                            button_title="Check"
                            button_css="btn-warning"/>
                    </div>
                    <h5 class="card-title">Prediksi</h5>
                </div>
                <div class="card-body overflow-auto" style="padding: 5px;margin:0px;height:300px;">
                    <div class="mb-3">
                        <label for="exampleForm" class="form-label">Pasaran</label>
                        <select 
                            bind:value={pasaran}
                            class="form-control">
                            {#each listHome as rec }
                            <option value="{rec.pasaranlist_idpasarantogel}">{rec.pasaranlist_nmpasarantogel}</option>
                            {/each}
                        </select>
                    </div>
                    <div class="mb-3">
                        <label for="exampleForm" class="form-label">Nomor Prediksi</label>
                        <div class="input-group mb-3">
                            <Input
                                bind:value={nomorprediksi}
                                on:keyup={handleKeyboard_format}   
                                type="text"
                                minlength="4"
                                maxlength="4"
                                placeholder="Nomor Prediksi"/>
                            <button
                                on:click={() => {
                                    generate("nomorprediksi");
                                }}
                                type="button" class="btn btn-info">Generate</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="col-sm-9">
            <div class="card mt-1">
                <div class="card-header">
                    <h5 class="card-title">Result</h5>
                </div>
                <div class="card-body overflow-auto" style="padding: 5px;margin:0px;height:700px;">
                    <table class="table">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align:left;vertical-align:top;font-size:14px;border:none;">NO</th>
                                <th NOWRAP width="*" style="text-align:left;vertical-align:top;font-size:14px;border:none;">AGEN</th>
                                <th NOWRAP width="7%" style="text-align:left;vertical-align:top;font-size:14px;border:none;">INVOICE</th>
                                <th NOWRAP width="7%" style="text-align:center;vertical-align:top;font-size:14px;border:none;">DATE</th>
                                <th NOWRAP width="7%" style="text-align:left;vertical-align:top;font-size:14px;border:none;">PERIODE</th>
                                <th NOWRAP width="20%" style="text-align:right;vertical-align:top;font-size:14px;border:none;">TOTAL BET</th>
                                <th NOWRAP width="20%" style="text-align:right;vertical-align:top;font-size:14px;border:none;">MEMBER WINLOSE</th>
                                <th NOWRAP width="20%" style="text-align:right;vertical-align:top;font-size:14px;border:none;">COMPANY WINLOSE</th>
                            </tr>
                        </thead>
                        <tbody>
                            {#each result as rec }
                            <tr>
                                <td style="text-align:center;vertical-align:top;font-size:13px;border:none;">{rec.prediksi_no}</td>
                                <td style="text-align:left;vertical-align:top;font-size:13px;border:none;">{rec.prediksi_nmcompany}</td>
                                <td style="text-align:left;vertical-align:top;font-size:13px;border:none;">{rec.prediksi_invoice}</td>
                                <td style="text-align:center;vertical-align:top;font-size:13px;border:none;">{rec.prediksi_invoicedate}</td>
                                <td style="text-align:left;vertical-align:top;font-size:13px;border:none;">{rec.prediksi_invoiceperiode}</td>
                                <td style="text-align:right;vertical-align:top;font-size:13px;border:none;{rec.prediksi_totalbetcss}">{rec.prediksi_totalbet}</td>
                                <td style="text-align:right;vertical-align:top;font-size:13px;border:none;{rec.prediksi_subtotalcss}">
                                    {new Intl.NumberFormat().format(
                                        rec.prediksi_subtotal
                                    )}
                                </td>
                                <td style="text-align:right;vertical-align:top;font-size:13px;border:none;{rec.prediksi_subtotalwincss}">
                                    {new Intl.NumberFormat().format(
                                        rec.prediksi_subtotalwin
                                    )}
                                </td>
                            </tr>
                            {/each}
                        </tbody>
                    </table>
                </div>
                <div class="card-footer">
                    <div class="float-end">
                        <table>
                            <tr>
                                <td>Subtotal Company</td>
                                <td>:</td>
                                <td style="{subtotal_company_css}">
                                    {new Intl.NumberFormat().format(subtotal_company)}
                                </td>
                            </tr>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>