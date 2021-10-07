<script>
  import { Button, Col, TabContent, TabPane } from "sveltestrap";
  import Loader2 from "../components/Loader.svelte";
  import Modal from "../components/Modalfull.svelte";
  import Headerback from "../components/Headerback.svelte";
  import Panel from "../components/Panel.svelte";
  import PanelFull from "../components/Panelfull.svelte";
  import Form432d from "../permainan/Form432d.svelte";
  import Formcolok from "../permainan/Formcolok.svelte";
  import Form5050 from "../permainan/Form5050.svelte";
  import Formkombinasi from "../permainan/Formkombinasi.svelte";
  import Formdasar from "../permainan/Formdasar.svelte";
  import Formshio from "../permainan/Formshio.svelte";
  import { notifications } from "../components/Noti.svelte";

  export let client_token = "";
  export let client_company = "";
  export let client_username = "";
  export let client_credit = 0;
  export let client_ipaddress = "";
  export let client_timezone = "Asia/Jakarta";
  export let client_device = "";
  export let pasaran_code = "";
  export let pasaran_name = "";
  export let pasaran_periode = 0;
  export let permainan = "";
  let css_loader = "display:none;";
  let resulttogel = [];
  let resultinvoice = [];
  let resultslipbet = [];
  let filterBukuMimpi = [];
  let listBukumimpi = [];
  let record = "";
  let idcomppasaran = "";
  let idtrxkeluaran = "";
  let statuspasaran = "";
  let permainan_title = "4D / 3D / 2D";
  let permainan_periode = pasaran_periode + " - " + pasaran_code;
  let css_permainan_432 = "btn1";
  let css_permainan_colok = "btn1";
  let css_permainan_5050 = "btn1";
  let css_permainan_kombinasi = "btn1";
  let css_permainan_dasar = "btn1";
  let css_permainan_shio = "btn1";
  let totalbayar_invoice = 0;
  let totalbet_invoice = 0;
  let total4d_bayar = 0;
  let total3d_bayar = 0;
  let total2d_bayar = 0;
  let totalcolokbebas_bayar = 0;
  let totalcolokmacau_bayar = 0;
  let totalcoloknaga_bayar = 0;
  let totalcolokjitu_bayar = 0;
  let total5050umum_bayar = 0;
  let total5050special_bayar = 0;
  let total5050kombinasi_bayar = 0;
  let totalmacaukombinasi_bayar = 0;
  let totaldasar_bayar = 0;
  let totalshio_bayar = 0;
  let totalwin_4d = 0;
  let totalwin_3d = 0;
  let totalwin_2d = 0;
  let totalwin_colokbebas = 0;
  let totalwin_colokmacau = 0;
  let totalwin_coloknaga = 0;
  let totalwin_colokjitu = 0;
  let totalwin_5050umum = 0;
  let totalwin_5050special = 0;
  let totalwin_5050kombinasi = 0;
  let totalwin_macaukombinasi = 0;
  let totalwin_dasar = 0;
  let totalwin_shio = 0;
  let subtotal_bayar = 0;
  let subtotal_winner = 0;
  let total_winlose = 0;
  let defailheader = "";
  let searchbukumimpi = "";
  let tipe = "";
  switch (permainan) {
    case "4-3-2":
      css_permainan_432 = "btn2";
      break;
    case "colok":
      css_permainan_colok = "btn2";
      break;
    case "5050":
      css_permainan_5050 = "btn2";
      break;
    case "kombinasi":
      css_permainan_kombinasi = "btn2";
      break;
    case "dasar":
      css_permainan_dasar = "btn2";
      break;
    case "shio":
      css_permainan_shio = "btn2";
      break;
  }

  const changePermainan = (e) => {
    permainan = e;
    switch (e) {
      case "4-3-2":
        permainan_title = "4D / 3D / 2D";
        css_permainan_432 = "btn2";
        css_permainan_colok = "btn1";
        css_permainan_5050 = "btn1";
        css_permainan_kombinasi = "btn1";
        css_permainan_dasar = "btn1";
        css_permainan_shio = "btn1";
        break;
      case "colok":
        permainan_title = "COLOK";
        css_permainan_colok = "btn2";
        css_permainan_432 = "btn1";
        css_permainan_5050 = "btn1";
        css_permainan_kombinasi = "btn1";
        css_permainan_dasar = "btn1";
        css_permainan_shio = "btn1";
        break;
      case "5050":
        permainan_title = "50 - 50";
        css_permainan_5050 = "btn2";
        css_permainan_432 = "btn1";
        css_permainan_colok = "btn1";
        css_permainan_kombinasi = "btn1";
        css_permainan_dasar = "btn1";
        css_permainan_shio = "btn1";
        break;
      case "kombinasi":
        permainan_title = "MACAU / KOMBINASI";
        css_permainan_kombinasi = "btn2";
        css_permainan_432 = "btn1";
        css_permainan_colok = "btn1";
        css_permainan_5050 = "btn1";
        css_permainan_dasar = "btn1";
        css_permainan_shio = "btn1";
        break;
      case "dasar":
        permainan_title = "DASAR";
        css_permainan_dasar = "btn2";
        css_permainan_432 = "btn1";
        css_permainan_colok = "btn1";
        css_permainan_5050 = "btn1";
        css_permainan_kombinasi = "btn1";
        css_permainan_shio = "btn1";
        break;
      case "shio":
        permainan_title = "SHIO";
        css_permainan_shio = "btn2";
        css_permainan_432 = "btn1";
        css_permainan_colok = "btn1";
        css_permainan_5050 = "btn1";
        css_permainan_kombinasi = "btn1";
        css_permainan_dasar = "btn1";
        break;
    }
  };

  async function checkpasaran() {
    css_loader =
      "position:absolute;z-index: 1000;left: 0%;top: 35%;display:inline;";
    const res = await fetch("/api/checkpasaran", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        company: client_company,
        timezone: client_timezone,
        pasaran_code: pasaran_code,
      }),
    });

    const json = await res.json();
    record = json;
    if (json.status == "200") {
      css_loader = "display:none;";
    }
    idcomppasaran = record["pasaran_idcomp"];
    idtrxkeluaran = record["pasaran_idtransaction"];
    statuspasaran = record["pasaran_status"];
    pasaran_periode = record["pasaran_periode"];
    invoicebet("all");
  }
  async function resultpasaran() {
    const res = await fetch("/api/resulttogel", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        company: client_company,
        pasaran_code: pasaran_code,
      }),
    });

    const json = await res.json();
    record = json.record;
    if (json.status == 200) {
      record = json.record;
      if (record != null) {
        for (var i = 0; i < record.length; i++) {
          resulttogel = [
            ...resulttogel,
            {
              no: record[i]["no"],
              date: record[i]["date"],
              periode: record[i]["periode"],
              result: record[i]["result"],
            },
          ];
        }
      }
    }
  }
  async function invoicebet(e) {
    const res = await fetch("/api/invoicebet", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        company: client_company,
        username: client_username,
        pasaran_code: pasaran_code,
        pasaran_periode: pasaran_periode,
        permainan: e,
      }),
    });

    const json = await res.json();
    record = json.record;

    if (json.status == 200) {
      totalbayar_invoice = json.totalbayar;
      totalbet_invoice = json.totalbet;
      record = json.record;
      if (record != null) {
        for (var i = 0; i < record.length; i++) {
          resultinvoice = [
            ...resultinvoice,
            {
              permainan: record[i]["permainan"],
              periode: record[i]["periode"],
              nomor: record[i]["nomor"],
              bet: record[i]["bet"],
              diskon: record[i]["diskon"],
              kei: record[i]["kei"],
              bayar: record[i]["bayar"],
              win: record[i]["win"],
              menang: record[i]["menang"],
            },
          ];
        }
      }
    }
  }
  async function slipbet() {
    const res = await fetch("/api/slipperiode", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        company: client_company,
        username: client_username,
        pasaran_code: pasaran_code,
      }),
    });

    const json = await res.json();
    record = json.record;

    if (json.status == 200) {
      record = json.record;
      if (record != null) {
        for (var i = 0; i < record.length; i++) {
          resultslipbet = [
            ...resultslipbet,
            {
              status: record[i]["status"],
              tglkeluaran: record[i]["tglkeluaran"],
              idinvoice: record[i]["idinvoice"],
              periode: record[i]["periode"],
              totallose: record[i]["totallose"],
              backgroundstatus: record[i]["background"],
            },
          ];
        }
      }
    }
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
        notifications.push(
          "An error has occured, Please Contact Administrator"
        );
      }
    } else {
      notifications.push("An error has occured, Please Contact Administrator");
    }
  }
  const handleInvoice = (e) => {
    resultinvoice = [];
    resultslipbet = [];
    invoicebet("all");
    slipbet();
  };
  checkpasaran();

  const handleSelect = (event) => {
    changePermainan(event.target.value);
  };
  const handleClickTab = (e) => {
    switch (e) {
      case "SLIP":
        resultslipbet = [];
        slipbet();
        break;
      case "RESULT":
        resulttogel = [];
        resultpasaran();
        break;
      case "BUKUMIMPI":
        filterBukuMimpi = [];
        listBukumimpi = [];
        searchbukumimpi = "";
        tipe = "";
        fetch_bukumimpi();
        break;
    }
  };
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
  };
  const handleKeyboardbukumimpi_checkenter = (e) => {
    let keyCode = e.which || e.keyCode;
    if (keyCode === 13) {
      // tipe = "";
      filterBukuMimpi = [];
      listBukumimpi = [];
      fetch_bukumimpi();
    }
  };
  async function showDetail(e, periode) {
    defailheader = periode;
    const res = await fetch("/api/slipperiodedetail", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        company: client_company,
        username: client_username,
        idinvoice: e,
      }),
    });

    const json = await res.json();
    record = json.record;

    if (json.status == 200) {
      record = json.record;
      total4d_bayar = record["total4d_bayar"];
      total3d_bayar = record["total3d_bayar"];
      total2d_bayar = record["total2d_bayar"];
      totalcolokbebas_bayar = record["totalcolokbebas_bayar"];
      totalcolokmacau_bayar = record["totalcolokmacau_bayar"];
      totalcoloknaga_bayar = record["totalcoloknaga_bayar"];
      totalcolokjitu_bayar = record["totalcolokjitu_bayar"];
      total5050umum_bayar = record["total5050umum_bayar"];
      total5050special_bayar = record["total5050special_bayar"];
      total5050kombinasi_bayar = record["total5050kombinasi_bayar"];
      totalmacaukombinasi_bayar = record["totalmacaukombinasi_bayar"];
      totaldasar_bayar = record["totaldasar_bayar"];
      totalshio_bayar = record["totalshio_bayar"];
      totalwin_4d = record["totalwin_4d"];
      totalwin_3d = record["totalwin_3d"];
      totalwin_2d = record["totalwin_2d"];
      totalwin_colokbebas = record["totalwin_colokbebas"];
      totalwin_colokmacau = record["totalwin_colokmacau"];
      totalwin_coloknaga = record["totalwin_coloknaga"];
      totalwin_colokjitu = record["totalwin_colokjitu"];
      totalwin_5050umum = record["totalwin_5050umum"];
      totalwin_5050special = record["totalwin_5050special"];
      totalwin_5050kombinasi = record["totalwin_5050kombinasi"];
      totalwin_macaukombinasi = record["totalwin_macaukombinasi"];
      totalwin_dasar = record["totalwin_dasar"];
      totalwin_shio = record["totalwin_shio"];
      subtotal_bayar = record["subtotal_bayar"];
      subtotal_winner = record["subtotal_winner"];
      total_winlose = record["total_winlose"];
    }

    let myModal = new bootstrap.Modal(
      document.getElementById("modalSlipdetail")
    );
    myModal.show();
  }
  $: {
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
  let modal_table_fontsize_header = "13px";
  let modal_table_fontsize_body = "12px";
  if (client_device == "MOBILE") {
    modal_table_fontsize_header = "11px";
    modal_table_fontsize_body = "11px";
  }
</script>

{#if statuspasaran == "ONLINE"}
  <Headerback {client_token} {client_device} {client_company} />
  {#if client_device == "WEBSITE"}
    <Col xs="12" style="padding: 0px; margin:0px;margin-top:40px;">
      <Panel>
        <slot:template slot="header">
          <h1 style="padding:0px;margin:0px;color:white;font-size:13px;">
            JENIS PERMAINAN
          </h1>
        </slot:template>
        <slot:template slot="body">
          <ul class="nav nav-pills">
            <li class="active">
              <Button
                on:click={() => {
                  changePermainan("4-3-2");
                }}
                id={css_permainan_432}>4D/3D/2D</Button
              >&nbsp;
            </li>
            <li>
              <Button
                on:click={() => {
                  changePermainan("colok");
                }}
                id={css_permainan_colok}>Colok</Button
              >&nbsp;
            </li>
            <li>
              <Button
                on:click={() => {
                  changePermainan("5050");
                }}
                id={css_permainan_5050}>50-50</Button
              >&nbsp;
            </li>
            <li>
              <Button
                on:click={() => {
                  changePermainan("kombinasi");
                }}
                id={css_permainan_kombinasi}>KOMBINASI</Button
              >&nbsp;
            </li>
            <li>
              <Button
                on:click={() => {
                  changePermainan("dasar");
                }}
                id={css_permainan_dasar}>DASAR</Button
              >&nbsp;
            </li>
            <li>
              <Button
                on:click={() => {
                  changePermainan("shio");
                }}
                id={css_permainan_shio}>SHIO</Button
              >&nbsp;
            </li>
          </ul>
        </slot:template>
      </Panel>
    </Col>
    <div class="clearfix mt-2" />
    <Col
      xxl="7"
      xl="7"
      lg="7"
      md="7"
      sm="12"
      xs="12"
      style="padding:0px;padding-right:2px;"
    >
      {#if permainan == "4-3-2"}
        <Form432d
          on:handleInvoice={handleInvoice}
          {idcomppasaran}
          {idtrxkeluaran}
          {client_token}
          {client_company}
          {client_username}
          {client_timezone}
          {client_ipaddress}
          {client_device}
          {pasaran_name}
          {pasaran_code}
          {pasaran_periode}
          {permainan_title}
        />
      {/if}
      {#if permainan == "colok"}
        <Formcolok
          on:handleInvoice={handleInvoice}
          {idcomppasaran}
          {idtrxkeluaran}
          {client_token}
          {client_company}
          {client_username}
          {client_timezone}
          {client_ipaddress}
          {client_device}
          {pasaran_name}
          {pasaran_code}
          {pasaran_periode}
          {permainan_title}
          {permainan_periode}
        />
      {/if}
      {#if permainan == "5050"}
        <Form5050
          on:handleInvoice={handleInvoice}
          {idcomppasaran}
          {idtrxkeluaran}
          {client_token}
          {client_company}
          {client_username}
          {client_timezone}
          {client_ipaddress}
          {client_device}
          {pasaran_name}
          {pasaran_code}
          {pasaran_periode}
          {permainan_title}
          {permainan_periode}
        />
      {/if}
      {#if permainan == "kombinasi"}
        <Formkombinasi
          on:handleInvoice={handleInvoice}
          {idcomppasaran}
          {idtrxkeluaran}
          {client_token}
          {client_company}
          {client_username}
          {client_timezone}
          {client_ipaddress}
          {client_device}
          {pasaran_name}
          {pasaran_code}
          {pasaran_periode}
          {permainan_title}
          {permainan_periode}
        />
      {/if}
      {#if permainan == "dasar"}
        <Formdasar
          on:handleInvoice={handleInvoice}
          {idcomppasaran}
          {idtrxkeluaran}
          {client_token}
          {client_company}
          {client_username}
          {client_timezone}
          {client_ipaddress}
          {client_device}
          {pasaran_name}
          {pasaran_code}
          {pasaran_periode}
          {permainan_title}
          {permainan_periode}
        />
      {/if}
      {#if permainan == "shio"}
        <Formshio
          on:handleInvoice={handleInvoice}
          {idcomppasaran}
          {idtrxkeluaran}
          {client_token}
          {client_company}
          {client_username}
          {client_timezone}
          {client_ipaddress}
          {client_device}
          {pasaran_name}
          {pasaran_code}
          {pasaran_periode}
          {permainan_title}
          {permainan_periode}
        />
      {/if}
    </Col>
    <Col
      xxl="5"
      xl="5"
      lg="5"
      md="5"
      sm="12"
      xs="12"
      style="padding:0px;padding-left:2px;"
    >
      <ul
        class="nav nav-pills mb-3"
        id="pills-tab"
        role="tablist"
        style="background-color: #323030;"
      >
        <li class="nav-item" role="presentation">
          <button
            class="nav-link active"
            id="pills-home-tab"
            data-bs-toggle="pill"
            data-bs-target="#pills-pasangan"
            type="button"
            role="tab"
            aria-controls="pills-pasangan"
            aria-selected="true">PASANGAN</button
          >
        </li>
        <li
          on:click={() => {
            handleClickTab("SLIP");
          }}
          class="nav-item"
          role="presentation"
        >
          <button
            class="nav-link"
            id="pills-profile-tab"
            data-bs-toggle="pill"
            data-bs-target="#pills-slip"
            type="button"
            role="tab"
            aria-controls="pills-slip"
            aria-selected="false">INVOICE</button
          >
        </li>
        <li
          on:click={() => {
            handleClickTab("RESULT");
          }}
          class="nav-item"
          role="presentation"
        >
          <button
            class="nav-link"
            id="pills-contact-tab"
            data-bs-toggle="pill"
            data-bs-target="#pills-result"
            type="button"
            role="tab"
            aria-controls="pills-result"
            aria-selected="false">RESULT</button
          >
        </li>
        <li
          on:click={() => {
            handleClickTab("BUKUMIMPI");
          }}
          class="nav-item"
          role="presentation"
        >
          <button
            class="nav-link"
            id="pills-contact-tab"
            data-bs-toggle="pill"
            data-bs-target="#pills-bukumimpi"
            type="button"
            role="tab"
            aria-controls="pills-bukumimpi"
            aria-selected="false">BUKU MIMPI</button
          >
        </li>
      </ul>
      <div class="tab-content" id="pills-tabContent">
        <div
          class="tab-pane fade show active"
          id="pills-pasangan"
          role="tabpanel"
          aria-labelledby="pills-pasangan-tab"
        >
          <PanelFull
            header={true}
            footer={true}
            header_style="padding:0px;margin:0px;"
            body_style="padding:0px;margin:0px;background:#121212;border:1px solid #0e0c13;height:655px;"
          >
            <slot:template slot="header" />
            <slot:template slot="body">
              <table class="table table-dark table-striped">
                <thead>
                  <tr>
                    <th
                      width="10%"
                      style="text-align:center;vertical-align:top;background:#303030;font-size:13px;border-bottom:none;"
                      NOWRAP>NOMOR</th
                    >
                    <th
                      width="10%"
                      style="text-align:center;vertical-align:top;background:#303030;font-size:13px;border-bottom:none;"
                      NOWRAP>PERMINAN</th
                    >
                    <th
                      width="20%"
                      style="text-align:right;vertical-align:top;background:#303030;font-size:13px;border-bottom:none;"
                      NOWRAP>BET</th
                    >
                    <th
                      width="20%"
                      style="text-align:right;vertical-align:top;background:#303030;font-size:13px;border-bottom:none;"
                      NOWRAP>KEI(%)</th
                    >
                    <th
                      width="20%"
                      style="text-align:right;vertical-align:top;background:#303030;font-size:13px;border-bottom:none;"
                      NOWRAP>DIS(%)</th
                    >
                    <th
                      width="20%"
                      style="text-align:right;vertical-align:top;background:#303030;font-size:13px;border-bottom:none;"
                      NOWRAP>BAYAR</th
                    >
                  </tr>
                </thead>
                <tbody>
                  {#each resultinvoice as rec}
                    <tr>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;color:#fc0;"
                        NOWRAP>{rec.nomor}</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;color:#fc0;"
                        NOWRAP>{rec.permainan}</td
                      >
                      <td
                        style="text-align:right;vertical-align:top;font-size:12px;color:#fc0;"
                        NOWRAP>{new Intl.NumberFormat().format(rec.bet)}</td
                      >
                      <td
                        style="text-align:right;vertical-align:top;font-size:12px;color:#fc0;"
                        NOWRAP>{rec.kei.toFixed(1)}</td
                      >
                      <td
                        style="text-align:right;vertical-align:top;font-size:12px;color:#fc0;"
                        NOWRAP>{rec.diskon.toFixed(1)}</td
                      >
                      <td
                        style="text-align:right;vertical-align:top;font-size:12px;color:#fc0;"
                        NOWRAP>{new Intl.NumberFormat().format(rec.bayar)}</td
                      >
                    </tr>
                  {/each}
                </tbody>
              </table>
            </slot:template>
            <slot:template slot="footer">
              <table class="table" style="background:none;">
                <tr>
                  <td style="text-align:right;color:white;font-size:12px;"
                    >TOTAL BET</td
                  >
                  <td style="text-align:right;color:white;font-size:12px;">:</td
                  >
                  <td style="text-align:right;color:#fc0;font-size:12px;"
                    >{new Intl.NumberFormat().format(totalbet_invoice)}</td
                  >
                </tr>
                <tr>
                  <td style="text-align:right;color:white;font-size:12px;"
                    >TOTAL BAYAR</td
                  >
                  <td style="text-align:right;color:white;font-size:12px;">:</td
                  >
                  <td style="text-align:right;color:#fc0;font-size:12px;"
                    >{new Intl.NumberFormat().format(totalbayar_invoice)}</td
                  >
                </tr>
              </table>
            </slot:template>
          </PanelFull>
        </div>
        <div
          class="tab-pane fade"
          id="pills-slip"
          role="tabpanel"
          aria-labelledby="pills-slip-tab"
        >
          <PanelFull
            header={false}
            footer={false}
            body_style="padding:0px;margin:0px;background:#121212;border:1px solid #0e0c13;height:710px;"
          >
            <slot:template slot="body">
              <table class="table table-dark ">
                <thead>
                  <tr>
                    <th
                      width="10%"
                      style="text-align:center;vertical-align:top;background:#303030;font-size:13px;border-bottom:none;"
                      NOWRAP>STATUS</th
                    >
                    <th
                      width="20%"
                      style="text-align:center;vertical-align:top;background:#303030;font-size:13px;border-bottom:none;"
                      NOWRAP>PERIODE</th
                    >
                    <th
                      width="20%"
                      style="text-align:center;vertical-align:top;background:#303030;font-size:13px;border-bottom:none;"
                      NOWRAP>TANGGAL</th
                    >

                    <th
                      width="*"
                      style="text-align:right;vertical-align:top;background:#303030;font-size:13px;border-bottom:none;"
                      NOWRAP>WINLOSE</th
                    >
                  </tr>
                </thead>
                <tbody>
                  {#each resultslipbet as rec}
                    <tr>
                      <td
                        NOWRAP
                        style="text-align:center;vertical-align:top;font-size:12px;{rec.backgroundstatus}"
                        >{rec.status}</td
                      >
                      <td
                        on:click={() => {
                          showDetail(rec.idinvoice, rec.periode);
                        }}
                        NOWRAP
                        style="text-decoration:underline;cursor:pointer;text-align:center;vertical-align:top;font-size:12px;color:white;"
                        >{rec.periode}</td
                      >
                      <td
                        NOWRAP
                        style="text-align:center;vertical-align:top;font-size:12px;color:white;"
                        >{rec.tglkeluaran}</td
                      >

                      <td
                        NOWRAP
                        style="text-align:right;vertical-align:top;font-size:12px;color:#fc0;"
                      >
                        {new Intl.NumberFormat().format(rec.totallose)}
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
          id="pills-result"
          role="tabpanel"
          aria-labelledby="pills-result-tab"
        >
          <PanelFull
            header={false}
            footer={false}
            body_style="padding:0px;margin:0px;background:#121212;border:1px solid #0e0c13;height:710px;"
          >
            <slot:template slot="body">
              <table class="table table-dark table-striped">
                <thead>
                  <tr>
                    <th
                      width="10%"
                      style="text-align:center;vertical-align:top;background:#303030;font-size:13px;border-bottom:none;"
                      NOWRAP>NO</th
                    >
                    <th
                      width="20%"
                      style="text-align:center;vertical-align:top;background:#303030;font-size:13px;border-bottom:none;"
                      NOWRAP>TANGGAL</th
                    >
                    <th
                      width="20%"
                      style="text-align:center;vertical-align:top;background:#303030;font-size:13px;border-bottom:none;"
                      NOWRAP>PERIODE</th
                    >
                    <th
                      width="*"
                      style="text-align:center;vertical-align:top;background:#303030;font-size:13px;border-bottom:none;"
                      NOWRAP>HASIL</th
                    >
                  </tr>
                </thead>
                <tbody>
                  {#each resulttogel as rec}
                    <tr>
                      <td
                        NOWRAP
                        style="text-align:center;vertical-align:top;font-size:12px;color:white;"
                        >{rec.no}</td
                      >
                      <td
                        NOWRAP
                        style="text-align:center;vertical-align:top;font-size:12px;color:white;"
                        >{rec.date}</td
                      >
                      <td
                        NOWRAP
                        style="text-align:center;vertical-align:top;font-size:12px;color:white;"
                        >{rec.periode}</td
                      >
                      <td
                        NOWRAP
                        style="text-align:center;vertical-align:top;font-size:12px;color:#fc0;"
                        >{rec.result}</td
                      >
                    </tr>
                  {/each}
                </tbody>
              </table>
            </slot:template>
          </PanelFull>
        </div>
        <div
          class="tab-pane fade"
          id="pills-bukumimpi"
          role="tabpanel"
          aria-labelledby="pills-result-tab"
        >
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
                            <span style="color:#fc0;font-size:14px;"
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
                            <span style="color:#fc0;font-size:14px;"
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
                            <span style="color:#fc0;font-size:14px;"
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
                            <span style="color:#fc0;font-size:14px;"
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
        </div>
      </div>
    </Col>
  {:else}
    <div style="margin-top:40px;">&nbsp;</div>
    <ul
      class="nav nav-pills mb-3"
      id="pills-tab"
      role="tablist"
      style="background-color: #181818;"
    >
      <li class="nav-item" role="presentation">
        <button
          style="font-size:10px;"
          class="nav-link active"
          id="pills-keranjang-tab"
          data-bs-toggle="pill"
          data-bs-target="#pills-keranjang"
          type="button"
          role="tab"
          aria-controls="pills-keranjang"
          aria-selected="true">KERANJANG</button
        >
      </li>
      <li class="nav-item" role="presentation">
        <button
          style="font-size:10px;"
          class="nav-link"
          id="pills-pasangan-tab"
          data-bs-toggle="pill"
          data-bs-target="#pills-pasangan"
          type="button"
          role="tab"
          aria-controls="pills-pasangan"
          aria-selected="true">PASANGAN</button
        >
      </li>
      <li
        on:click={() => {
          handleClickTab("SLIP");
        }}
        class="nav-item"
        role="presentation"
      >
        <button
          style="font-size:10px;"
          class="nav-link"
          id="pills-profile-tab"
          data-bs-toggle="pill"
          data-bs-target="#pills-slip"
          type="button"
          role="tab"
          aria-controls="pills-slip"
          aria-selected="false">INVOICE</button
        >
      </li>
      <li
        on:click={() => {
          handleClickTab("RESULT");
        }}
        class="nav-item"
        role="presentation"
      >
        <button
          style="font-size:10px;"
          class="nav-link"
          id="pills-contact-tab"
          data-bs-toggle="pill"
          data-bs-target="#pills-result"
          type="button"
          role="tab"
          aria-controls="pills-result"
          aria-selected="false">RESULT</button
        >
      </li>
      <li
        on:click={() => {
          handleClickTab("BUKUMIMPI");
        }}
        class="nav-item"
        role="presentation"
      >
        <button
          style="font-size:10px;"
          class="nav-link"
          id="pills-contact-tab"
          data-bs-toggle="pill"
          data-bs-target="#pills-bukumimpi"
          type="button"
          role="tab"
          aria-controls="pills-bukumimpi"
          aria-selected="false">BUKU</button
        >
      </li>
    </ul>
    <div class="tab-content" id="pills-tabContent">
      <div
        class="tab-pane fade show active"
        id="pills-keranjang"
        role="tabpanel"
        aria-labelledby="pills-keranjang-tab"
      >
        <br />
        <b style="font-size: 13px;">Pilih Permainan Dibawah Ini : </b>
        <select
          on:change={handleSelect}
          style="background-color: #323030;color:white;border:1px solid #323030;"
          class="form-control"
        >
          <option value="4-3-2">4D/3D/2D</option>
          <option value="colok">COLOK</option>
          <option value="5050">50-50</option>
          <option value="kombinasi">MACAU / KOMBINASI</option>
          <option value="dasar">DASAR</option>
          <option value="shio">SHIO</option>
        </select>
        <div class="clear-fix" />
        <br />
        {#if permainan == "4-3-2"}
          <Form432d
            on:handleInvoice={handleInvoice}
            {idcomppasaran}
            {idtrxkeluaran}
            {client_token}
            {client_company}
            {client_username}
            {client_timezone}
            {client_ipaddress}
            {client_device}
            {pasaran_name}
            {pasaran_code}
            {pasaran_periode}
            {permainan_title}
          />
        {/if}
        {#if permainan == "colok"}
          <Formcolok
            on:handleInvoice={handleInvoice}
            {idcomppasaran}
            {idtrxkeluaran}
            {client_token}
            {client_company}
            {client_username}
            {client_timezone}
            {client_ipaddress}
            {client_device}
            {pasaran_name}
            {pasaran_code}
            {pasaran_periode}
            {permainan_title}
          />
        {/if}
        {#if permainan == "5050"}
          <Form5050
            on:handleInvoice={handleInvoice}
            {idcomppasaran}
            {idtrxkeluaran}
            {client_token}
            {client_company}
            {client_username}
            {client_timezone}
            {client_ipaddress}
            {client_device}
            {pasaran_name}
            {pasaran_code}
            {pasaran_periode}
            {permainan_title}
          />
        {/if}
        {#if permainan == "kombinasi"}
          <Formkombinasi
            on:handleInvoice={handleInvoice}
            {idcomppasaran}
            {idtrxkeluaran}
            {client_token}
            {client_company}
            {client_username}
            {client_timezone}
            {client_ipaddress}
            {client_device}
            {pasaran_name}
            {pasaran_code}
            {pasaran_periode}
            {permainan_title}
          />
        {/if}
        {#if permainan == "dasar"}
          <Formdasar
            on:handleInvoice={handleInvoice}
            {idcomppasaran}
            {idtrxkeluaran}
            {client_token}
            {client_company}
            {client_username}
            {client_timezone}
            {client_ipaddress}
            {client_device}
            {pasaran_name}
            {pasaran_code}
            {pasaran_periode}
            {permainan_title}
          />
        {/if}
        {#if permainan == "shio"}
          <Formshio
            on:handleInvoice={handleInvoice}
            {idcomppasaran}
            {idtrxkeluaran}
            {client_token}
            {client_company}
            {client_username}
            {client_timezone}
            {client_ipaddress}
            {client_device}
            {pasaran_name}
            {pasaran_code}
            {pasaran_periode}
            {permainan_title}
          />
        {/if}
      </div>
      <div
        class="tab-pane"
        id="pills-pasangan"
        role="tabpanel"
        aria-labelledby="pills-pasangan-tab"
      >
        <PanelFull
          header={true}
          footer={true}
          header_style="padding:0px;margin:0px;"
          body_style="padding:0px;margin:0px;background:#121212;border:1px solid #0e0c13;height:450px;"
        >
          <slot:template slot="header" />
          <slot:template slot="body">
            <table class="table table-dark table-striped">
              <thead>
                <tr>
                  <th
                    width="10%"
                    style="text-align:center;vertical-align:top;background:#303030;font-size:11px;border-bottom:none;"
                    NOWRAP>NOMOR</th
                  >
                  <th
                    width="10%"
                    style="text-align:center;vertical-align:top;background:#303030;font-size:11px;border-bottom:none;"
                    NOWRAP>PERMINAN</th
                  >
                  <th
                    width="20%"
                    style="text-align:right;vertical-align:top;background:#303030;font-size:11px;border-bottom:none;"
                    NOWRAP>BET</th
                  >
                  <th
                    width="20%"
                    style="text-align:right;vertical-align:top;background:#303030;font-size:11px;border-bottom:none;"
                    NOWRAP>KEI(%)</th
                  >
                  <th
                    width="20%"
                    style="text-align:right;vertical-align:top;background:#303030;font-size:11px;border-bottom:none;"
                    NOWRAP>DIS(%)</th
                  >
                  <th
                    width="20%"
                    style="text-align:right;vertical-align:top;background:#303030;font-size:11px;border-bottom:none;"
                    NOWRAP>BAYAR</th
                  >
                </tr>
              </thead>
              <tbody>
                {#each resultinvoice as rec}
                  <tr>
                    <td
                      style="text-align:center;vertical-align:top;font-size:12px;color:#fc0;"
                      NOWRAP>{rec.nomor}</td
                    >
                    <td
                      style="text-align:center;vertical-align:top;font-size:12px;color:#fc0;"
                      NOWRAP>{rec.permainan}</td
                    >
                    <td
                      style="text-align:right;vertical-align:top;font-size:12px;color:#fc0;"
                      NOWRAP>{new Intl.NumberFormat().format(rec.bet)}</td
                    >
                    <td
                      style="text-align:right;vertical-align:top;font-size:12px;color:#fc0;"
                      NOWRAP>{rec.kei.toFixed(1)}</td
                    >
                    <td
                      style="text-align:right;vertical-align:top;font-size:12px;color:#fc0;"
                      NOWRAP>{rec.diskon.toFixed(1)}</td
                    >
                    <td
                      style="text-align:right;vertical-align:top;font-size:12px;color:#fc0;"
                      NOWRAP>{new Intl.NumberFormat().format(rec.bayar)}</td
                    >
                  </tr>
                {/each}
              </tbody>
            </table>
          </slot:template>
          <slot:template slot="footer">
            <table class="table" style="background:none;">
              <tr>
                <td style="text-align:right;color:white;font-size:11px;"
                  >TOTAL BET</td
                >
                <td style="text-align:right;color:white;font-size:11px;">:</td>
                <td style="text-align:right;color:#fc0;font-size:11px;"
                  >{new Intl.NumberFormat().format(totalbet_invoice)}</td
                >
              </tr>
              <tr>
                <td style="text-align:right;color:white;font-size:11px;"
                  >TOTAL BAYAR</td
                >
                <td style="text-align:right;color:white;font-size:11px;">:</td>
                <td style="text-align:right;color:#fc0;font-size:11px;"
                  >{new Intl.NumberFormat().format(totalbayar_invoice)}</td
                >
              </tr>
            </table>
          </slot:template>
        </PanelFull>
      </div>
      <div
        class="tab-pane "
        id="pills-slip"
        role="tabpanel"
        aria-labelledby="pills-slip-tab"
      >
        <PanelFull
          header={false}
          footer={false}
          body_style="padding:0px;margin:0px;background:#121212;border:1px solid #0e0c13;height:450px;"
        >
          <slot:template slot="body">
            <table class="table table-dark ">
              <thead>
                <tr>
                  <th
                    width="10%"
                    style="text-align:center;vertical-align:top;background:#303030;font-size:11px;border-bottom:none;"
                    NOWRAP>STATUS</th
                  >
                  <th
                    width="20%"
                    style="text-align:center;vertical-align:top;background:#303030;font-size:11px;border-bottom:none;"
                    NOWRAP>PERIODE</th
                  >
                  <th
                    width="20%"
                    style="text-align:center;vertical-align:top;background:#303030;font-size:11px;border-bottom:none;"
                    NOWRAP>TANGGAL</th
                  >

                  <th
                    width="*"
                    style="text-align:right;vertical-align:top;background:#303030;font-size:11px;border-bottom:none;"
                    NOWRAP>WINLOSE</th
                  >
                </tr>
              </thead>
              <tbody>
                {#each resultslipbet as rec}
                  <tr>
                    <td
                      NOWRAP
                      style="text-align:center;vertical-align:top;font-size:11px;{rec.backgroundstatus}"
                      >{rec.status}</td
                    >
                    <td
                      on:click={() => {
                        showDetail(rec.idinvoice, rec.periode);
                      }}
                      NOWRAP
                      style="text-decoration:underline;cursor:pointer;text-align:center;vertical-align:top;font-size:11px;color:white;"
                      >{rec.periode}</td
                    >
                    <td
                      NOWRAP
                      style="text-align:center;vertical-align:top;font-size:11px;color:white;"
                      >{rec.tglkeluaran}</td
                    >

                    <td
                      NOWRAP
                      style="text-align:right;vertical-align:top;font-size:11px;color:#fc0;"
                    >
                      {new Intl.NumberFormat().format(rec.totallose)}
                    </td>
                  </tr>
                {/each}
              </tbody>
            </table>
          </slot:template>
        </PanelFull>
      </div>
      <div
        class="tab-pane "
        id="pills-result"
        role="tabpanel"
        aria-labelledby="pills-result-tab"
      >
        <PanelFull
          header={false}
          footer={false}
          body_style="padding:0px;margin:0px;background:#121212;border:1px solid #0e0c13;height:450px;"
        >
          <slot:template slot="body">
            <table class="table table-dark table-striped">
              <thead>
                <tr>
                  <th
                    width="10%"
                    style="text-align:center;vertical-align:top;background:#303030;font-size:12px;border-bottom:none;"
                    NOWRAP>NO</th
                  >
                  <th
                    width="20%"
                    style="text-align:center;vertical-align:top;background:#303030;font-size:12px;border-bottom:none;"
                    NOWRAP>TANGGAL</th
                  >
                  <th
                    width="20%"
                    style="text-align:center;vertical-align:top;background:#303030;font-size:12px;border-bottom:none;"
                    NOWRAP>PERIODE</th
                  >
                  <th
                    width="*"
                    style="text-align:center;vertical-align:top;background:#303030;font-size:12px;border-bottom:none;"
                    NOWRAP>HASIL</th
                  >
                </tr>
              </thead>
              <tbody>
                {#each resulttogel as rec}
                  <tr>
                    <td
                      NOWRAP
                      style="text-align:center;vertical-align:top;font-size:11px;color:white;"
                      >{rec.no}</td
                    >
                    <td
                      NOWRAP
                      style="text-align:center;vertical-align:top;font-size:11px;color:white;"
                      >{rec.date}</td
                    >
                    <td
                      NOWRAP
                      style="text-align:center;vertical-align:top;font-size:11px;color:white;"
                      >{rec.periode}</td
                    >
                    <td
                      NOWRAP
                      style="text-align:center;vertical-align:top;font-size:11px;color:#fc0;"
                      >{rec.result}</td
                    >
                  </tr>
                {/each}
              </tbody>
            </table>
          </slot:template>
        </PanelFull>
      </div>
      <div
        class="tab-pane "
        id="pills-bukumimpi"
        role="tabpanel"
        aria-labelledby="pills-result-tab"
      >
        <ul
          class="nav nav-pills mb-3"
          id="pills-tab"
          role="tablist"
          style="background-color: #181818;"
        >
          <li class="nav-item" role="presentation">
            <button
              style="font-size:10px;"
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
              style="font-size:10px;"
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
              style="font-size:10px;"
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
              style="font-size:10px;"
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
                          style="text-align:center;vertical-align:top;font-size:12px;color:#fc0;"
                          >{rec.bukumimpi_tipe}</td
                        >
                        <td
                          width="*"
                          style="text-align:left;vertical-align:top;font-size:12px;color:#8b8989;"
                          >{rec.bukumimpi_nama}
                          <br />
                          <span style="color:#fc0;font-size:12px;"
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
                          style="text-align:center;vertical-align:top;font-size:12px;color:#fc0;"
                          >{rec.bukumimpi_tipe}</td
                        >
                        <td
                          width="*"
                          style="text-align:left;vertical-align:top;font-size:12px;color:#8b8989;"
                          >{rec.bukumimpi_nama}
                          <br />
                          <span style="color:#fc0;font-size:12px;"
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
                          style="text-align:center;vertical-align:top;font-size:12px;color:#fc0;"
                          >{rec.bukumimpi_tipe}</td
                        >
                        <td
                          width="*"
                          style="text-align:left;vertical-align:top;font-size:12px;color:#8b8989;"
                          >{rec.bukumimpi_nama}
                          <br />
                          <span style="color:#fc0;font-size:12px;"
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
                          style="text-align:center;vertical-align:top;font-size:12px;color:#fc0;"
                          >{rec.bukumimpi_tipe}</td
                        >
                        <td
                          width="*"
                          style="text-align:left;vertical-align:top;font-size:12px;color:#8b8989;"
                          >{rec.bukumimpi_nama}
                          <br />
                          <span style="color:#fc0;font-size:12px;"
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
      </div>
    </div>
  {/if}
  <Modal modal_id={"modalSlipdetail"} modal_size={"modal-dialog-centered"}>
    <slot:template slot="header">
      <div class="float-end">
        <img
          style="cursor:pointer;"
          data-bs-dismiss="modal"
          aria-label="Close"
          width="25"
          src="https://i.ibb.co/9yNF25L/outline-close-white-48dp.png"
          alt=""
        />
      </div>
      <h5 class="modal-title" id="exampleModalLabel">
        PASARAN - {defailheader}
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
              style="text-align:left;vertical-align:top;background:#303030;font-size:{modal_table_fontsize_header};border-bottom:none;"
              >GAME</th
            >
            <th
              width="50%"
              style="text-align:right;vertical-align:top;background:#303030;font-size:{modal_table_fontsize_header};border-bottom:none;"
              >BAYAR</th
            >
            <th
              width="25%"
              style="text-align:right;vertical-align:top;background:#303030;font-size:{modal_table_fontsize_header};border-bottom:none;"
              >MENANG</th
            >
          </tr>
        </thead>
        <tbody>
          <tr>
            <td
              style="text-align:center;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >1</td
            >
            <td
              NOWRAP
              style="text-align:left;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >4D</td
            >
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(total4d_bayar)}
            </td>
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalwin_4d)}
            </td>
          </tr>
          <tr>
            <td
              style="text-align:center;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >2</td
            >
            <td
              NOWRAP
              style="text-align:left;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >3D</td
            >
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(total3d_bayar)}
            </td>
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalwin_3d)}
            </td>
          </tr>
          <tr>
            <td
              style="text-align:center;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >3</td
            >
            <td
              NOWRAP
              style="text-align:left;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >2D</td
            >
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(total2d_bayar)}
            </td>
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalwin_2d)}
            </td>
          </tr>
          <tr>
            <td
              style="text-align:center;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >4</td
            >
            <td
              NOWRAP
              style="text-align:left;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >COLOK BEBAS</td
            >
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalcolokbebas_bayar)}
            </td>
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalwin_colokbebas)}
            </td>
          </tr>
          <tr>
            <td
              style="text-align:center;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >5</td
            >
            <td
              NOWRAP
              style="text-align:left;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >COLOK MACAU</td
            >
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalcolokmacau_bayar)}
            </td>
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalwin_colokmacau)}
            </td>
          </tr>
          <tr>
            <td
              style="text-align:center;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >6</td
            >
            <td
              NOWRAP
              style="text-align:left;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >COLOK NAGA</td
            >
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalcoloknaga_bayar)}
            </td>
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalwin_coloknaga)}
            </td>
          </tr>
          <tr>
            <td
              style="text-align:center;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >7</td
            >
            <td
              NOWRAP
              style="text-align:left;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >COLOK JITU</td
            >
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalcolokjitu_bayar)}
            </td>
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalwin_colokjitu)}
            </td>
          </tr>
          <tr>
            <td
              style="text-align:center;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >8</td
            >
            <td
              NOWRAP
              style="text-align:left;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >50 - 50 UMUM</td
            >
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(total5050umum_bayar)}
            </td>
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalwin_5050umum)}
            </td>
          </tr>
          <tr>
            <td
              style="text-align:center;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >9</td
            >
            <td
              NOWRAP
              style="text-align:left;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >50 - 50 SPECIAL</td
            >
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(total5050special_bayar)}
            </td>
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalwin_5050special)}
            </td>
          </tr>
          <tr>
            <td
              style="text-align:center;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >10</td
            >
            <td
              NOWRAP
              style="text-align:left;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >50 - 50 KOMBINASI</td
            >
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(total5050kombinasi_bayar)}
            </td>
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalwin_5050kombinasi)}
            </td>
          </tr>
          <tr>
            <td
              style="text-align:center;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >11</td
            >
            <td
              NOWRAP
              style="text-align:left;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >MACAU / KOMBINASI</td
            >
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalmacaukombinasi_bayar)}
            </td>
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalwin_macaukombinasi)}
            </td>
          </tr>
          <tr>
            <td
              style="text-align:center;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >12</td
            >
            <td
              NOWRAP
              style="text-align:left;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >DASAR</td
            >
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totaldasar_bayar)}
            </td>
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalwin_dasar)}
            </td>
          </tr>
          <tr>
            <td
              style="text-align:center;vertical-align:top;font-size:{modal_table_fontsize_body};f"
              >13</td
            >
            <td
              NOWRAP
              style="text-align:left;vertical-align:top;font-size:{modal_table_fontsize_body};"
              >SHIO</td
            >
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalshio_bayar)}
            </td>
            <td
              style="text-align:right;vertical-align:top;color:#fc0;font-size:{modal_table_fontsize_body};"
            >
              {new Intl.NumberFormat().format(totalwin_shio)}
            </td>
          </tr>
        </tbody>
      </table>
    </slot:template>
    <slot:template slot="footer">
      <table class="table" width="100%">
        <tr>
          <td
            NOWRAP
            width="50%"
            style="text-align:right;vertical-align:top;color:white;font-size:{modal_table_fontsize_body};"
            >TOTAL BAYAR</td
          >
          <td
            NOWRAP
            width="1%"
            style="text-align:center;vertical-align:top;color:white;font-size:{modal_table_fontsize_body};"
            >:</td
          >
          <td
            NOWRAP
            width="*"
            style="text-align:right;vertical-align:top;color:white;font-size:{modal_table_fontsize_body};color:#fc0;"
            >{new Intl.NumberFormat().format(subtotal_bayar)}</td
          >
        </tr>
        <tr>
          <td
            NOWRAP
            style="text-align:right;vertical-align:top;color:white;font-size:12px;"
            >TOTAL WINNER</td
          >
          <td
            NOWRAP
            style="text-align:center;vertical-align:top;color:white;font-size:12px;"
            >:</td
          >
          <td
            NOWRAP
            style="text-align:right;vertical-align:top;color:white;font-size:12px;color:#fc0;"
            >{new Intl.NumberFormat().format(subtotal_winner)}</td
          >
        </tr>
        <tr>
          <td
            NOWRAP
            style="text-align:right;vertical-align:top;color:white;font-size:12px;"
            >WINLOSE</td
          >
          <td
            NOWRAP
            style="text-align:center;vertical-align:top;color:white;font-size:12px;"
            >:</td
          >
          <td
            NOWRAP
            style="text-align:right;vertical-align:top;color:white;font-size:12px;color:#fc0;"
            >{new Intl.NumberFormat().format(total_winlose)}</td
          >
        </tr>
      </table>
    </slot:template>
  </Modal>
  <div class="clearfix" />
  <style>
    #stream::-webkit-scrollbar {
      width: 0.3em;
    }

    #stream::-webkit-scrollbar-track {
      -webkit-box-shadow: inset 0 0 6px rgba(0, 0, 0, 0.3);
    }

    #stream::-webkit-scrollbar-thumb {
      background-color: #505768;
      outline: 1px solid slategrey;
    }
    .table-dark {
      --bs-table-bg: #212529;
      --bs-table-striped-bg: #1e1e1e;
      --bs-table-striped-color: #fff;
      --bs-table-active-bg: #373b3e;
      --bs-table-active-color: #fff;
      --bs-table-hover-bg: #323539;
      --bs-table-hover-color: #fff;
      color: #fff;
      border-color: #373b3e;
    }
  </style>
{:else if statuspasaran == "OFFLINE"}
  <div style="margin-bottom:10px;margin-left: -10px;">
    <center>
      <div style="cursor: pointer;">
        <a
          href="/?token={client_token}&agent={client_company}"
          title="nuketoto"
        >
          Back to Home
        </a>
      </div>
    </center>
  </div>
{/if}
<Loader2 cssstyle={css_loader} />
