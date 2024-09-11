<script>
    import Home from "../prediksi/Home.svelte";
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let tokenprediksi = "";
    let akses_page = true;
    let listHome = [];
    let sData = "";
    let record = "";
    let record_message = "";
    let totalrecord = 0;

    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "PREDIKSI-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
            akses_page = false;
        } else {
            initHome();
        }
    }
    async function initHome() {
        const res = await fetch("/api/initprediksi", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            localStorage.setItem("tokenprediksi", json.token);
            tokenprediksi = json.token
            listpasaran(json.token)
        }else{
            logout();
        }
    }
    async function listpasaran(tokenprediksi) {
        const res = await fetch("/api/listpasaran", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: "root",
                token: tokenprediksi,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            record_message = json.message;
            if (record != null) {
                totalrecord = record.length;
                let no = 0
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listHome = [
                        ...listHome,
                        {
                            pasaranlist_idpasarantogel: record[i]["pasaranlist_idpasarantogel"],
                            pasaranlist_nmpasarantogel: record[i]["pasaranlist_nmpasarantogel"],
                        },
                    ];
                }
            }
        } else {
            // logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
   
    
    initapp()
</script>

<Home
 
   
    {token}
    {tokenprediksi}
    {table_header_font}
    {table_body_font}
    {listHome}
    {totalrecord}
/>