<script>
  import { Button, Col } from "sveltestrap";
  import dayjs from "dayjs";
  import Loader2 from "../components/Loader.svelte";
  import Notif from "../components/Notif.svelte";
  import Headerback from "../components/Headerback.svelte";
  import Panel from "../components/Panel.svelte";
  import PanelFull from "../components/Panelfull.svelte";
  import Form432d from "../permainan/Form432d.svelte";
  import Formcolok from "../permainan/Formcolok.svelte";
  import Form5050 from "../permainan/Form5050.svelte";
  import Formkombinasi from "../permainan/Formkombinasi.svelte";
  import Formdasar from "../permainan/Formdasar.svelte";
  import Formshio from "../permainan/Formshio.svelte";

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
  let totalbayar_invoice = 0;
  let totalbet_invoice = 0;
  let searchbukumimpi = "";
  let tipe = "";
  let message_err = "";
  let css_err = "display:none;";
  let tabs = ["4-3-2", "colok", "5050", "kombinasi", "dasar", "shio"];
  let selected = tabs[0];

  const changePermainan = (e) => {
    permainan = e;
    console.log(permainan);
    switch (e) {
      case "4-3-2":
        permainan_title = "4D / 3D / 2D";
        selected = tabs[0];
        break;
      case "colok":
        permainan_title = "COLOK";
        selected = tabs[1];
        break;
      case "5050":
        permainan_title = "50 - 50";
        selected = tabs[2];
        break;
      case "kombinasi":
        permainan_title = "MACAU / KOMBINASI";
        selected = tabs[3];
        break;
      case "dasar":
        permainan_title = "DASAR";
        selected = tabs[4];
        break;
      case "shio":
        permainan_title = "SHIO";
        selected = tabs[5];
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
    if (statuspasaran == "OFFLINE") {
      css_err = "display:inline-block";
      message_err = "Pasaran OFFLINE";
      setTimeout(function () {
        css_err = "display: none;";
      }, 5000);
    }
    invoicebet("all");
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
        pasaran_idtransaction: parseInt(idtrxkeluaran),
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
              bukumimpi_tipe: record[i]["bukumimpi_type"],
              bukumimpi_nama: record[i]["bukumimpi_name"],
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
</script>

{#if statuspasaran == "ONLINE"}
  <Headerback
    {client_username}
    {client_company}
    {client_credit}
    {client_token}
    {client_ipaddress}
    {client_timezone}
    {client_device}
  />
  {#if client_device == "WEBSITE"}
    <Col xs="12" style="padding:0;">
      <Panel>
        <slot:template slot="body">
          <ul class="nav top-menu nav-pills justify-content-center">
            <li class="nav-item">
              <a
                class="nav-link"
                class:active={selected === "4-3-2"}
                on:click={() => {
                  changePermainan("4-3-2");
                }}>4D/3D/2D</a
              >
            </li>
            <li class="nav-item">
              <a
                class="nav-link"
                class:active={selected === "colok"}
                on:click={() => {
                  changePermainan("colok");
                }}>Colok</a
              >
            </li>
            <li class="nav-item">
              <a
                class="nav-link"
                class:active={selected === "5050"}
                on:click={() => {
                  changePermainan("5050");
                }}>50-50</a
              >
            </li>
            <li class="nav-item">
              <a
                class="nav-link"
                class:active={selected === "kombinasi"}
                on:click={() => {
                  changePermainan("kombinasi");
                }}>KOMBINASI</a
              >
            </li>
            <li class="nav-item">
              <a
                class="nav-link"
                class:active={selected === "dasar"}
                on:click={() => {
                  changePermainan("dasar");
                }}>DASAR</a
              >
            </li>
            <li class="nav-item">
              <a
                class="nav-link"
                class:active={selected === "shio"}
                on:click={() => {
                  changePermainan("shio");
                }}>SHIO</a
              >
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
        style="background-color: #171717;padding:10px;"
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
            <slot:template slot="header">
              <div
                class="row"
                style="padding: 20px;font-size: 15px;color: #fff;"
              >
                <div class="col">
                  <span
                    >BET SAYA : <span style="color:#FF9900; font-weight:bold"
                      >{new Intl.NumberFormat().format(totalbet_invoice)}</span
                    ></span
                  >
                </div>
                <div class="col text-end">
                  <span
                    >TOTAL BAYAR : <span style="color:#FF9900; font-weight:bold"
                      >{new Intl.NumberFormat().format(
                        totalbayar_invoice
                      )}</span
                    ></span
                  >
                </div>
              </div>
              <!-- <table class="table" style="background:none;">
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
              </table> -->
            </slot:template>
            <slot:template slot="body">
              <table class="table table-dark table-striped">
                <thead>
                  <tr>
                    <th
                      width="10%"
                      style="text-align:center;vertical-align:top;background:#171717;font-size:13px;border-bottom:none;"
                      NOWRAP>NOMOR</th
                    >
                    <th
                      width="10%"
                      style="text-align:center;vertical-align:top;background:#171717;font-size:13px;border-bottom:none;"
                      NOWRAP>PERMAINAN</th
                    >
                    <th
                      width="20%"
                      style="text-align:right;vertical-align:top;background:#171717;font-size:13px;border-bottom:none;"
                      NOWRAP>BET</th
                    >
                    <th
                      width="20%"
                      style="text-align:right;vertical-align:top;background:#171717;font-size:13px;border-bottom:none;"
                      NOWRAP>KEI(%)</th
                    >
                    <th
                      width="20%"
                      style="text-align:right;vertical-align:top;background:#171717;font-size:13px;border-bottom:none;"
                      NOWRAP>DIS(%)</th
                    >
                    <th
                      width="20%"
                      style="text-align:right;vertical-align:top;background:#171717;font-size:13px;border-bottom:none;"
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
            <slot:template slot="footer" />
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
          aria-label="Permainan"
          class="form-select"
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
          <slot:template slot="header">
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
                    NOWRAP>PERMAINAN</th
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
          <slot:template slot="footer" />
        </PanelFull>
      </div>
    </div>
  {/if}
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
  <Notif message={message_err} css_init={css_err} />
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

<style scoped>
  .nav.top-menu > .nav-item > .nav-link.active,
  .nav.top-menu > .nav-item > .nav-link:hover {
    background: #ff9900;
    border-radius: 4px;
    color: #000;
    cursor: pointer;
  }
  .nav.top-menu > .nav-item {
    font-size: 10pt;
    font-family: "Poppins";
    opacity: 1;
    color: #fff;
    width: 105px;
    text-align: center;
    margin: 0 5px;
  }

  .nav.top-menu > .nav-item > .nav-link,
  .nav.top-menu > .nav-item > .nav-link.active {
    text-decoration: none;
  }
</style>
