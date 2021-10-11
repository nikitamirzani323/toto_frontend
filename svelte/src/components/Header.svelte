<script>
    import { Col, Card, CardBody } from "sveltestrap";
    import dayjs from "dayjs";
    import utc from "dayjs/plugin/utc";
    import timezone from "dayjs/plugin/timezone";
    import Modal from "../components/Modalfull2.svelte"
    import PanelFull from "../components/Panelfull.svelte"
    dayjs.extend(utc);
    dayjs.extend(timezone);

    export let client_token = "";
    export let client_username = "";
    export let client_credit = 0;
    export let client_ipaddress = "";
    export let client_timezone = "";
    export let client_device = "";
    let modal_table_fontsize_header = "13px";
    let modal_table_fontsize_body = "12px";
    if (client_device == "MOBILE") {
        modal_table_fontsize_header = "11px";
        modal_table_fontsize_body = "11px";
    }
    let display_credit = 0;
    let clockmachine = "";
    let record = "";
    let filterBukuMimpi = [];
    let listBukumimpi = [];
    let searchbukumimpi = "";
    let tipe = "";
    function updateClock() {
        let endtime = dayjs()
            .tz(client_timezone)
            .format("DD MMM YYYY | HH:mm:ss");

        clockmachine = endtime;
    }
    async function fetch_bukumimpi() {
        const res = await fetch("/api/bookdream", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                tipe: tipe,
                nama: searchbukumimpi.toLowerCase(),
            }),
        });

        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listBukumimpi = [
                        ...listBukumimpi,
                        {
                            bukumimpi_tipe: record[i]["bukumimpi_tipe"],
                            bukumimpi_nama: record[i]["bukumimpi_nama"],
                            bukumimpi_nomor: record[i]["bukumimpi_nomor"],
                        },
                    ];
                }
            } else {
                alert("Error");
            }
        } else {
            alert("Error");
        }
    }
    $: {
        setInterval(updateClock, 100);
        if (searchbukumimpi) {
            filterBukuMimpi = listBukumimpi.filter((item) =>
                item.bukumimpi_nama
                    .toLowerCase()
                    .includes(searchbukumimpi.toLowerCase())
            );
        } else {
            filterBukuMimpi = [...listBukumimpi];
        }
    }
    display_credit = new Intl.NumberFormat().format(client_credit);

    let hasilkeluaran = [
        {no:1,tanggal:"2021-10-12",pasaran:"CAMBODIA",periode:"CAM-01",result:"1234"},
        {no:2,tanggal:"2021-10-12",pasaran:"BULLSEYE",periode:"BE-01",result:"2345"},
        {no:3,tanggal:"2021-10-12",pasaran:"SYDNEY",periode:"SY-01",result:"3456"},
        {no:4,tanggal:"2021-10-12",pasaran:"PCSO",periode:"PCSO-01",result:"7890"},
        {no:5,tanggal:"2021-10-12",pasaran:"HONGKONG",periode:"HK-01",result:"5678"},
    ]
    let hasilinvoice = [
        {status:"RUNNING",tanggal:"2021-10-12",pasaran:"CAMBODIA",periode:"CAM-01",winlose:"500.000"},
        {status:"RUNNING",tanggal:"2021-10-12",pasaran:"BULLSEYE",periode:"BE-01",winlose:"100.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"SYDNEY",periode:"SY-01",winlose:"20.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"PCSO",periode:"PCSO-01",winlose:"-20.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"HONGKONG",periode:"HK-01",winlose:"-10.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"HONGKONG",periode:"HK-01",winlose:"-10.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"HONGKONG",periode:"HK-01",winlose:"-10.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"HONGKONG",periode:"HK-01",winlose:"-10.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"HONGKONG",periode:"HK-01",winlose:"-10.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"HONGKONG",periode:"HK-01",winlose:"-10.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"HONGKONG",periode:"HK-01",winlose:"-10.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"HONGKONG",periode:"HK-01",winlose:"-10.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"HONGKONG",periode:"HK-01",winlose:"-10.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"HONGKONG",periode:"HK-01",winlose:"-10.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"HONGKONG",periode:"HK-01",winlose:"-10.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"HONGKONG",periode:"HK-01",winlose:"-10.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"HONGKONG",periode:"HK-01",winlose:"-10.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"HONGKONG",periode:"HK-01",winlose:"-10.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"HONGKONG",periode:"HK-01",winlose:"-10.000"},
        {status:"COMPLETED",tanggal:"2021-10-12",pasaran:"HONGKONG",periode:"HK-01",winlose:"-10.000"},
    ]
    
    const handleClickButtonTop = (e) => {
        let idmodal = ""
        switch(e){
            case "result":
                idmodal = "modalhasilkeluaran"
                break;
            case "invoice":
                idmodal = "modalhasilinvoice"
                break;
            case "bukumimpi":
                idmodal = "modalbukumimpi"
                fetch_bukumimpi()
                break;
        }
        let myModal = new bootstrap.Modal(
            document.getElementById(idmodal)
        );
        myModal.show(); 
    }
    const handleClickBukuMimpi = (e) => {
        filterBukuMimpi = [];
        listBukumimpi = [];
        switch (e) {
            case "ALL":
                tipe = "";
                searchbukumimpi = "";
                break;
            case "4D":
                tipe = "4D";
                searchbukumimpi = "";
                break;
            case "3D":
                tipe = "3D";
                searchbukumimpi = "";
                break;
            case "2D":
                tipe = "2D";
                searchbukumimpi = "";
                break;
        }
        fetch_bukumimpi();
    }
    const handleKeyboardbukumimpi_checkenter = (e) => {
        let keyCode = e.which || e.keyCode;
        if (keyCode === 13) {
            filterBukuMimpi = [];
            listBukumimpi = [];
            fetch_bukumimpi();
        }
    };
</script>
{#if client_device == "WEBSITE"}
    <nav class="navbar fixed-top " style="background-color: #e91e63;">
        <div class="container">
            <a href="/?token={client_token}" title="totoapp">
                <img
                    id="imglogo"
                    alt="totoapp"
                    style="margin-top:0px;"
                    width="80"
                    src="https://dataset.b-cdn.net/assets/images/companies/isb388/isb388.png"
                />
            </a>
            <form class="d-flex">
                <button
                    on:click={() => {
                        handleClickButtonTop("result");
                    }} 
                    id="btn1" class="btn btn-secondary" type="button">RESULT</button>
                &nbsp;
                <button
                    on:click={() => {
                        handleClickButtonTop("invoice");
                    }}  
                    id="btn1" class="btn btn-secondary" type="button">INVOICE</button>
                &nbsp;
                <button
                    on:click={() => {
                        handleClickButtonTop("bukumimpi");
                    }}  
                    id="btn1" class="btn btn-secondary" type="button">BUKU MIMPI</button>
            </form>
        </div>
    </nav>
    <Col xxl="3" xl="3" lg="4" md="12" sm="12" style="padding:5px;margin:0px;margin-top:50px;">
        <Card color="dark" style="border:none;">
            <CardBody style="background-color:#2b3038;">
                <span style="font-size:13px;">Welcome, {client_username}</span
                ><br />
                <span style="font-size:13px;"
                    >Credit : IDR <span id="style_text">{display_credit}</span
                    ></span
                ><br />
                <span style="font-size:13px;"
                    >IPADDRESS : <span id="style_text">{client_ipaddress}</span
                    ></span
                ><br />
                <span style="font-size:13px;"
                    >TIMEZONE : <span id="style_text">{client_timezone}</span
                    ></span
                ><br />
                <span style="font-size:13px;"
                    >CLOCK : <span id="style_text">{clockmachine}</span></span
                >
            </CardBody>
        </Card>
    </Col>
{:else}
    <center style="margin-bottom:5px;">
        <a href="/?token={client_token}" title="isbtoto">
            <img
                id="imglogo"
                alt=""
                style="margin-top:10px;"
                width="100"
                src="https://dataset.b-cdn.net/assets/images/companies/isb388/isb388.png"
            />
        </a>
    </center>
    <Col
        xxl="12"
        xl="12"
        lg="12"
        md="12"
        sm="12"
        style="padding:5px;margin:0px;"
    >
        <Card color="dark">
            <CardBody style="background-color:#2b3038;">
                <span style="font-size:12px;">Welcome, {client_username}</span
                ><br />
                <span style="font-size:12px;"
                    >Credit : IDR <span id="style_text">{display_credit}</span
                    ></span
                ><br />
                <span style="font-size:12px;"
                    >IPADDRESS : <span id="style_text">{client_ipaddress}</span
                    ></span
                ><br />
                <span style="font-size:12px;"
                    >TIMEZONE : <span id="style_text">{client_timezone}</span
                    ></span
                ><br />
                <span style="font-size:12px;"
                    >CLOCK : <span id="style_text">{clockmachine}</span></span
                >
            </CardBody>
        </Card>
    </Col>
{/if}

<div class="clearfix mb-10" />
<Modal 
    modal_id={"modalhasilkeluaran"}
    modal_footer_flag={false} 
    modal_body_height={"height:350px;"}
    modal_size={"modal-dialog-centered"}>
    <slot:template slot="header">
        <h5 class="modal-title">
            RESULT
        </h5>
        
    </slot:template>
    <slot:template slot="body">
        <table class="table table-dark table-striped">
            <thead>
                <tr>
                    <th
                        width="1%"
                        style="text-align:center;vertical-align:top;background:#303030;font-size:{modal_table_fontsize_header};border-bottom:none;"
                        >NO</th
                    >
                    <th
                        width="15%"
                        style="text-align:center;vertical-align:top;background:#303030;font-size:{modal_table_fontsize_header};border-bottom:none;"
                        >TANGGAL</th
                    >
                    <th
                        width="*"
                        style="text-align:left;vertical-align:top;background:#303030;font-size:{modal_table_fontsize_header};border-bottom:none;"
                        >PASARAN</th
                    >
                    <th
                        width="15%"
                        style="text-align:center;vertical-align:top;background:#303030;font-size:{modal_table_fontsize_header};border-bottom:none;"
                        >PERIODE</th
                    >
                    <th
                        width="25%"
                        style="text-align:center;vertical-align:top;background:#303030;font-size:{modal_table_fontsize_header};border-bottom:none;"
                        >HASIL</th
                    >
                </tr>
            </thead>
            <tbody>
                {#each hasilkeluaran as rec }
                <tr>
                    <td style="text-align: center;vertical-align: top;font-size:{modal_table_fontsize_body};">{rec.no}</td>
                    <td style="text-align: center;vertical-align: top;font-size:{modal_table_fontsize_body};">{rec.tanggal}</td>
                    <td style="text-align: left;vertical-align: top;font-size:{modal_table_fontsize_body};">{rec.pasaran}</td>
                    <td style="text-align: center;vertical-align: top;font-size:{modal_table_fontsize_body};">{rec.periode}</td>
                    <td style="text-align: center;vertical-align: top;font-size:{modal_table_fontsize_body};color:rgb(255, 204, 0);">{rec.result}</td>
                </tr>
                {/each}
            </tbody>
        </table>
    </slot:template>
</Modal>
<Modal 
    modal_id={"modalhasilinvoice"}
    modal_footer_flag={false} 
    modal_body_height={"height:350px;"}
    modal_size={"modal-dialog-centered"}>
    <slot:template slot="header">
        <h5 class="modal-title">
            INVOICE
        </h5>
        
    </slot:template>
    <slot:template slot="body">
        <table class="table table-dark table-striped">
            <thead>
                <tr>
                    <th
                        width="1%"
                        style="text-align:center;vertical-align:top;background:#303030;font-size:{modal_table_fontsize_header};border-bottom:none;"
                        >STATUS</th
                    >
                    <th
                        width="15%"
                        style="text-align:center;vertical-align:top;background:#303030;font-size:{modal_table_fontsize_header};border-bottom:none;"
                        >TANGGAL</th
                    >
                    <th
                        width="*"
                        style="text-align:left;vertical-align:top;background:#303030;font-size:{modal_table_fontsize_header};border-bottom:none;"
                        >PASARAN</th
                    >
                    <th
                        width="15%"
                        style="text-align:center;vertical-align:top;background:#303030;font-size:{modal_table_fontsize_header};border-bottom:none;"
                        >PERIODE</th
                    >
                    <th
                        width="25%"
                        style="text-align:right;vertical-align:top;background:#303030;font-size:{modal_table_fontsize_header};border-bottom:none;"
                        >WINLOSE</th
                    >
                </tr>
            </thead>
            <tbody>
                {#each hasilinvoice as rec }
                <tr>
                    <td style="text-align: center;vertical-align: top;font-size:{modal_table_fontsize_body};">{rec.status}</td>
                    <td style="text-align: center;vertical-align: top;font-size:{modal_table_fontsize_body};">{rec.tanggal}</td>
                    <td style="text-align: left;vertical-align: top;font-size:{modal_table_fontsize_body};">{rec.pasaran}</td>
                    <td style="text-align: center;vertical-align: top;font-size:{modal_table_fontsize_body};">{rec.periode}</td>
                    <td style="text-align: right;vertical-align: top;font-size:{modal_table_fontsize_body};color:rgb(255, 204, 0);">{rec.winlose}</td>
                </tr>
                {/each}
            </tbody>
        </table>
    </slot:template>
</Modal>
<Modal 
    modal_id={"modalbukumimpi"}
    modal_headerbootom_flag={true} 
    modal_footer_flag={false} 
    modal_body_height={"height:550px;"}
    modal_size={"modal-dialog-centered"}>
    <slot:template slot="header">
        <h5 class="modal-title">
            BUKU MIMPI
        </h5>
    </slot:template>
    <slot:template slot="headerbottom">
        <ul
            class="nav nav-pills mb-3"
            id="pills-tab"
            role="tablist"
            style="background-color: #181818;"
        >
            <li
                on:click={() => {
                    handleClickBukuMimpi("ALL");
                }}
                class="nav-item"
                role="presentation"
            >
                <button
                    class="nav-link active"
                    id="pills-home-tab"
                    data-bs-toggle="pill"
                    data-bs-target="#pills-bukumimpiall"
                    type="button"
                    role="tab"
                    aria-controls="pills-bukumimpiall"
                    aria-selected="true">ALL</button
                >
            </li>
            <li
                on:click={() => {
                    handleClickBukuMimpi("4D");
                }}
                class="nav-item"
                role="presentation"
            >
                <button
                    class="nav-link"
                    id="pills-profile-tab"
                    data-bs-toggle="pill"
                    data-bs-target="#pills-bukumimpi4d"
                    type="button"
                    role="tab"
                    aria-controls="pills-bukumimpi4d"
                    aria-selected="false">4D</button
                >
            </li>
            <li
                on:click={() => {
                    handleClickBukuMimpi("3D");
                }}
                class="nav-item"
                role="presentation"
            >
                <button
                    class="nav-link"
                    id="pills-contact-tab"
                    data-bs-toggle="pill"
                    data-bs-target="#pills-bukumimpi3d"
                    type="button"
                    role="tab"
                    aria-controls="pills-bukumimpi3d"
                    aria-selected="false">3D</button
                >
            </li>
            <li
                on:click={() => {
                    handleClickBukuMimpi("2D");
                }}
                class="nav-item"
                role="presentation"
            >
                <button
                    class="nav-link"
                    id="pills-contact-tab"
                    data-bs-toggle="pill"
                    data-bs-target="#pills-bukumimpi2d"
                    type="button"
                    role="tab"
                    aria-controls="pills-bukumimpi2d"
                    aria-selected="false">2D</button
                >
            </li>
        </ul>
        <input
            bind:value={searchbukumimpi}
            on:keypress={handleKeyboardbukumimpi_checkenter}
            style="border-radius: none;border: none; background: rgb(48, 48, 48) none repeat scroll 0% 0%; color: white; font-size: 15px; "
            placeholder="Ketik apa yang telah kamu impikan"
            class="form-control"
            type="text"
        />
    </slot:template>
    <slot:template slot="body">
        <div class="tab-content" id="pills-tabContent">
            <div
                class="tab-pane fade show active"
                id="pills-bukumimpiall"
                role="tabpanel"
                aria-labelledby="pills-bukumimpiall-tab"
            >
                <PanelFull
                    header={false}
                    footer={false}
                    body_style="padding:0px;margin:0px;background:#121212;border:1px solid #0e0c13;height:620px;"
                >
                    <slot:template slot="body">
                        <table>
                            <tbody>
                                {#each filterBukuMimpi as rec}
                                    <tr>
                                        <td
                                            NOWRAP
                                            width="30px"
                                            style="text-align:center;vertical-align:top;font-size:14px;color:#fc0;"
                                            >{rec.bukumimpi_tipe}</td
                                        >
                                        <td
                                            width="*"
                                            style="text-align:left;vertical-align:top;font-size:15px;color:#8b8989;"
                                            >{rec.bukumimpi_nama}
                                            <br />
                                            <span
                                                style="color:#fc0;font-size:14px;"
                                                >{rec.bukumimpi_nomor}</span
                                            >
                                        </td>
                                    </tr>
                                {/each}
                            </tbody>
                        </table>
                    </slot:template>
                </PanelFull>
            </div>
            <div
                class="tab-pane fade"
                id="pills-bukumimpi4d"
                role="tabpanel"
                aria-labelledby="pills-bukumimpi4d-tab"
            >
                <PanelFull
                    header={false}
                    footer={false}
                    body_style="padding:0px;margin:0px;background:#121212;border:1px solid #0e0c13;height:620px;"
                >
                    <slot:template slot="body">
                        <table>
                            <tbody>
                                {#each filterBukuMimpi as rec}
                                    <tr>
                                        <td
                                            NOWRAP
                                            width="30px"
                                            style="text-align:center;vertical-align:top;font-size:14px;color:#fc0;"
                                            >{rec.bukumimpi_tipe}</td
                                        >
                                        <td
                                            width="*"
                                            style="text-align:left;vertical-align:top;font-size:15px;color:#8b8989;"
                                            >{rec.bukumimpi_nama}
                                            <br />
                                            <span
                                                style="color:#fc0;font-size:14px;"
                                                >{rec.bukumimpi_nomor}</span
                                            >
                                        </td>
                                    </tr>
                                {/each}
                            </tbody>
                        </table>
                    </slot:template>
                </PanelFull>
            </div>
            <div
                class="tab-pane fade"
                id="pills-bukumimpi3d"
                role="tabpanel"
                aria-labelledby="pills-bukumimpi3d-tab"
            >
                <PanelFull
                    header={false}
                    footer={false}
                    body_style="padding:0px;margin:0px;background:#121212;border:1px solid #0e0c13;height:620px;"
                >
                    <slot:template slot="body">
                        <table>
                            <tbody>
                                {#each filterBukuMimpi as rec}
                                    <tr>
                                        <td
                                            NOWRAP
                                            width="30px"
                                            style="text-align:center;vertical-align:top;font-size:14px;color:#fc0;"
                                            >{rec.bukumimpi_tipe}</td
                                        >
                                        <td
                                            width="*"
                                            style="text-align:left;vertical-align:top;font-size:15px;color:#8b8989;"
                                            >{rec.bukumimpi_nama}
                                            <br />
                                            <span
                                                style="color:#fc0;font-size:14px;"
                                                >{rec.bukumimpi_nomor}</span
                                            >
                                        </td>
                                    </tr>
                                {/each}
                            </tbody>
                        </table>
                    </slot:template>
                </PanelFull>
            </div>
            <div
                class="tab-pane fade"
                id="pills-bukumimpi2d"
                role="tabpanel"
                aria-labelledby="pills-bukumimpi2d-tab"
            >
                <PanelFull
                    header={false}
                    footer={false}
                    body_style="padding:0px;margin:0px;background:#121212;border:1px solid #0e0c13;height:620px;"
                >
                    <slot:template slot="body">
                        <table>
                            <tbody>
                                {#each filterBukuMimpi as rec}
                                    <tr>
                                        <td
                                            NOWRAP
                                            width="30px"
                                            style="text-align:center;vertical-align:top;font-size:14px;color:#fc0;"
                                            >{rec.bukumimpi_tipe}</td
                                        >
                                        <td
                                            width="*"
                                            style="text-align:left;vertical-align:top;font-size:15px;color:#8b8989;"
                                            >{rec.bukumimpi_nama}
                                            <br />
                                            <span
                                                style="color:#fc0;font-size:14px;"
                                                >{rec.bukumimpi_nomor}</span
                                            >
                                        </td>
                                    </tr>
                                {/each}
                            </tbody>
                        </table>
                    </slot:template>
                </PanelFull>
            </div>
        </div>
    </slot:template>
</Modal>