<script>
  import {
    Button,
    Card,
    CardBody,
    CardHeader,
    TabContent,
    TabPane,
  } from "sveltestrap";
  import Modal from "../components/Modal.svelte";
  import Tablekeranjangdasar from "../permainan/Tablekeranjangdasar.svelte";
  import Loader from "../components/Loader.svelte";
  import { createEventDispatcher } from "svelte";
  import { notifications } from "../components/Noti.svelte";
  import PeriodePanel from "../components/PeriodePanel.svelte";

  export let idcomppasaran = "";
  export let idtrxkeluaran = "";
  export let client_token = "";
  export let client_company = "";
  export let client_username = "";
  export let client_ipaddress = "";
  export let client_device = "";
  export let client_timezone = "";
  export let pasaran_code = "";
  export let pasaran_name = "";
  export let pasaran_periode = 0;
  export let permainan_title = "COLOK";
  let keranjang = [];
  let css_loader = "display:none;";
  let nomor_global = 0;
  let totalkeranjang = 0;
  let group_btn_beli = false;
  let record = "";
  let temp_bulk_error = "";

  let min_bet = 0;
  let max_bet = 0;
  let kei_besar_bet = 0;
  let kei_kecil_bet = 0;
  let kei_genap_bet = 0;
  let kei_ganjil_bet = 0;
  let disc_besar_bet = 0;
  let disc_kecil_bet = 0;
  let disc_genap_bet = 0;
  let disc_ganjil_bet = 0;

  let limit_total = 0;

  let count_line_dasar = 0;
  let count_line_standart = 0;

  let db_formdasar = 0;
  let db_formdasar_standart = 0;

  let dispatch = createEventDispatcher();

  async function inittogel_432d(e) {
    const res = await fetch("/api/inittogel_432d", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        company: client_company,
        pasaran_code: pasaran_code,
        permainan: e,
      }),
    });
    group_btn_beli = true;
    const json = await res.json();
    record = json.record;

    for (var i = 0; i < record.length; i++) {
      min_bet = parseInt(record[i]["min_bet"]);
      max_bet = parseInt(record[i]["max_bet"]);
      kei_besar_bet = parseFloat(record[i]["kei_besar_bet"]);
      kei_kecil_bet = parseFloat(record[i]["kei_kecil_bet"]);
      kei_genap_bet = parseFloat(record[i]["kei_genap_bet"]);
      kei_ganjil_bet = parseFloat(record[i]["kei_ganjil_bet"]);
      disc_besar_bet = parseFloat(record[i]["disc_besar_bet"]);
      disc_kecil_bet = parseFloat(record[i]["disc_kecil_bet"]);
      disc_genap_bet = parseFloat(record[i]["disc_genap_bet"]);
      disc_ganjil_bet = parseFloat(record[i]["disc_ganjil_bet"]);
      limit_total = parseInt(record[i]["limit_total"]);
    }
  }
  async function savetransaksi() {
    if (client_device == "WEBSITE") {
      css_loader =
        "position:absolute;z-index: 1000;left: 50%;top: 35%;display:inline;";
    } else {
      css_loader =
        "position:absolute;z-index: 1000;left: 0%;top: 35%;display:inline;";
    }
    group_btn_beli = false;
    const res = await fetch("/api/savetransaksi", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        transaction: reverseString(
          btoa(
            JSON.stringify({
              pasaran_idtransaction: idtrxkeluaran,
              pasaran_idcomp: idcomppasaran,
              token: client_token,
              company: client_company,
              username: client_username,
              ipaddress: client_ipaddress,
              devicemember: client_device,
              timezone: client_timezone,
              total: totalkeranjang,
              data: keranjang,
            })
          )
        ),
      }),
    });
    const json = await res.json();
    if (json.status == "200") {
      css_loader = "display:none;";
      notifications.push(
        "Data telah berhasil disimpan, Total belanja : " +
          new Intl.NumberFormat().format(totalkeranjang),
        "warning",
        "middle"
      );
      dispatch("handleInvoice", "call");
      reset();
    } else {
      css_loader = "display:none;";
      switch (json.status) {
        case "500":
          group_btn_beli = true;
          notifications.push(json.message);
          break;
        case "400":
          group_btn_beli = true;
          notifications.push(json.message);
          break;
        default:
          notifications.push(json.message);
          break;
      }
    }
  }

  function reverseString(str) {
    return str.split("").reverse().join("");
  }

  function reset() {
    keranjang = [];
    group_btn_beli = true;
    totalkeranjang = 0;
    count_line_dasar = 0;
    count_line_standart = 0;
  }
  inittogel_432d("dasar");
  function addKeranjang(
    nomor,
    game,
    bet,
    diskon_percen,
    diskon,
    bayar,
    win,
    kei_percen,
    kei
  ) {
    let total_data = keranjang.length;
    let flag_data = false;
    for (var i = 0; i < total_data; i++) {
      switch (game) {
        case "MACAU_KOMBINASI":
          if (nomor == keranjang[i].nomor.toString()) {
            let maxtotal_bayarmacaukombinasi = 0;
            for (var j = 0; j < keranjang.length; j++) {
              if ("50_50_UMUM" == keranjang[j].permainan) {
                if (parseInt(nomor) == parseInt(keranjang[j].nomor)) {
                  maxtotal_bayarmacaukombinasi =
                    parseInt(maxtotal_bayarmacaukombinasi) +
                    (parseInt(keranjang[j].bet) + parseInt(bet));
                }
              }
            }
            if (
              parseInt(limittotal_bet_macaukombinasi) <
              parseInt(maxtotal_bayarmacaukombinasi)
            ) {
              temp_bulk_error +=
                "Nomor ini : " +
                nomor +
                " sudah melebihi LIMIT TOTAL MACAU KOMBINASI<br />";
              flag_data = true;
            }
          }
          break;
      }
    }

    if (flag_data == false) {
      nomor_global = nomor_global + 1;
      const data = {
        id: nomor_global,
        nomor,
        permainan: game,
        bet,
        diskon,
        diskonpercen: diskon_percen,
        bayar,
        win,
        kei,
        kei_percen,
      };
      keranjang = [data, ...keranjang];
      count_keranjang();
    }
  }
  const removekeranjang = (e) => {
    keranjang = keranjang.filter(
      (keranjang) => keranjang.id != e.detail.idkeranjang
    );
    totalkeranjang = totalkeranjang - e.detail.bayar;
    count_keranjang();
  };
  const removekeranjangall = (e) => {
    if (keranjang.length > 0) {
      reset();
      count_keranjang();
    } else {
      notifications.push("Tidak ada list transaksi", "", "middle");
    }
  };
  const handleSave = (e) => {
    if (keranjang.length > 0) {
      savetransaksi();
    } else {
      notifications.push("Tidak ada list transaksi", "", "middle");
    }
  };
  function count_keranjang() {
    let count_umum = 0;
    let count_dasar = 0;
    for (let i = 0; i < keranjang.length; i++) {
      switch (keranjang[i].permainan.toString()) {
        case "DASAR":
          count_dasar = count_dasar + 1;
          break;
      }
    }
    count_line_dasar = count_dasar + db_formdasar;
    count_line_standart = count_umum + db_formdasar_standart;
  }

  //DASAR - INIT FORM
  let select_dasar = "";
  let select_dasar_input;
  let bet_dasar = "";

  function form_clear(e) {
    switch (e) {
      case "dasar":
        select_dasar = "";
        select_dasar_input.focus();
        bet_dasar = 0;
        break;
    }
  }

  function formdasar_add() {
    let flag = true;
    let nomor = select_dasar;
    let bet = bet_dasar;
    let diskon = 0;
    let diskonpercen = 0;
    let win = 1;
    let kei = 0;
    let keipersen = 0;
    let bayar = 0;
    let nmgame = "DASAR";
    if (nomor == "") {
      select_dasar_input.focus();
      flag = false;
    }
    if (bet == "") {
      flag = false;
      notifications.push("Amount tidak boleh kosong");
    }
    if (parseInt(bet) < parseInt(min_bet)) {
      bet_dasar = min_bet;
      flag = false;
      notifications.push("Minimal Bet : " + min_bet);
    }
    if (parseInt(bet) > parseInt(max_bet)) {
      bet_dasar = max_bet;
      flag = false;
      notifications.push(" Maximal Bet : " + max_bet);
    }
    if (flag == true) {
      switch (nomor) {
        case "GANJIL":
          kei = kei_ganjil_bet;
          diskon = disc_ganjil_bet;
          break;
        case "GENAP":
          kei = kei_genap_bet;
          diskon = disc_genap_bet;
          break;
        case "BESAR":
          kei = kei_besar_bet;
          diskon = disc_besar_bet;
          break;
        case "KECIL":
          kei = kei_kecil_bet;
          diskon = disc_kecil_bet;
          break;
      }
      keipersen = kei;
      diskonpercen = diskon;
      kei = parseInt(bet) * kei;
      diskon = parseInt(bet) * diskon;
      bayar = parseInt(bet) - parseInt(diskon) - parseInt(kei);
      totalkeranjang = bayar + totalkeranjang;

      addKeranjang(
        nomor,
        nmgame,
        bet_dasar,
        diskonpercen,
        diskon,
        bayar,
        win,
        keipersen,
        kei
      );
      form_clear("dasar");
    }
  }
  const handleTambah = (e) => {
    switch (e) {
      case "dasar":
        if (select_dasar == "" && parseInt(bet_dasar) < min_bet) {
          select_dasar_input.focus();
        } else {
          formdasar_add();
        }
        break;
    }
  };

  const handleKeyboard_number = (e) => {
    let numbera;
    for (let i = 0; i < bet_dasar.length; i++) {
      numbera = parseInt(bet_dasar[i]);
      if (isNaN(numbera)) {
        bet_dasar = "";
      }
    }
  };
  const handleKeyboard_checkenter = (e) => {
    let keyCode = e.which || e.keyCode;
    if (keyCode === 13) {
      formdasar_add();
    }
  };
</script>

<Loader cssstyle={css_loader} />
{#if client_device == "WEBSITE"}
  <Card color="dark" style="border:1px solid #262424;">
    <PeriodePanel
      {pasaran_name}
      {permainan_title}
      {pasaran_periode}
      {pasaran_code}
    />
    <CardBody style="background:#121212;padding:0px;margin:0px;">
      <div style="margin:5px;">
        <table class="table" style="background:none;width:100%;">
          <tr>
            <td
              width="25%"
              NOWRAP
              style="padding-right:10px;vertical-align: center;"
            >
              <span style="color:#8a8a8a;">TEBAK</span>
              <select
                bind:value={select_dasar}
                bind:this={select_dasar_input}
                style="border:none;background:#303030;color:white;"
                class="form-control"
              >
                <option value="">--Pilih--</option>
                <option value="GANJIL">GANJIL</option>
                <option value="GENAP">GENAP</option>
                <option value="BESAR">BESAR</option>
                <option value="KECIL">KECIL</option>
              </select>
              <span
                class="help-block"
                style="text-align:right;font-size:12px;"
              />
            </td>

            <td
              width="*"
              NOWRAP
              style="padding-right:10px;vertical-align: center;text-align:right;"
            >
              <span style="color:#8a8a8a;"
                >Bet (min : {new Intl.NumberFormat().format(min_bet)} dan max : {new Intl.NumberFormat().format(
                  max_bet
                )})</span
              >
              <input
                bind:value={bet_dasar}
                on:keyup={handleKeyboard_number}
                on:keypress={handleKeyboard_checkenter}
                type="text"
                class="form-control"
                placeholder="Bet"
                style="border:none;background:#303030;color:white;font-size:20px;text-align:right;"
                minlength="3"
                maxlength="7"
                tab_index="0"
              />
              <span style="text-align:right;font-size:12px;color:#8a8a8a;"
                >{new Intl.NumberFormat().format(bet_dasar)}</span
              >
            </td>
            <td width="20%" NOWRAP style="vertical-align: center;">
              <Button
                id="btn2"
                on:click={() => {
                  handleTambah("dasar");
                }}>TAMBAH</Button
              >
            </td>
          </tr>
        </table>
      </div>
    </CardBody>
  </Card>
{:else}
  <Card color="dark" style="border:1px solid #262424;margin:0px;padding:0px;">
    <CardHeader
      style="background:#323030;border-bottom:1px solid #333;border-top: 0 solid #333;"
    >
      <div class="float-end">
        <div
          style="color:white;text-align:right;font-size:12px;font-weight:bold;"
        >
          {pasaran_name}
        </div>
      </div>
      <h1 style="padding:0px;margin:0px;color:white;font-size:12px;">
        {permainan_title}<br />
        PERIODE : {pasaran_periode + " - " + pasaran_code}
      </h1>
    </CardHeader>
    <CardBody style="background:#121212;padding:0px;margin:0px;">
      <div style="margin:5px">
        <table class="table" style="background:none;width:100%;">
          <tr>
            <td
              width="35%"
              NOWRAP
              style="padding-right:10px;vertical-align: center;"
            >
              <span style="color:#8a8a8a;">TEBAK</span>
              <select
                bind:value={select_dasar}
                bind:this={select_dasar_input}
                style="border:none;background:#303030;color:white;"
                class="form-control"
              >
                <option value="">--Pilih--</option>
                <option value="GANJIL">GANJIL</option>
                <option value="GENAP">GENAP</option>
                <option value="BESAR">BESAR</option>
                <option value="KECIL">KECIL</option>
              </select>
              <span
                class="help-block"
                style="text-align:right;font-size:12px;"
              />
            </td>

            <td
              width="*"
              NOWRAP
              style="padding-right:10px;vertical-align: center;text-align:right;"
            >
              <span style="color:#8a8a8a;"
                >Bet (min : {new Intl.NumberFormat().format(min_bet)} dan max : {new Intl.NumberFormat().format(
                  max_bet
                )})</span
              >
              <input
                bind:value={bet_dasar}
                on:keyup={handleKeyboard_number}
                on:keypress={handleKeyboard_checkenter}
                type="text"
                class="form-control"
                placeholder="Bet"
                style="border:none;background:#303030;color:white;font-size:20px;text-align:right;"
                minlength="3"
                maxlength="7"
                tab_index="0"
              />
              <span style="text-align:right;font-size:12px;color:#8a8a8a;"
                >{new Intl.NumberFormat().format(bet_dasar)}</span
              >
            </td>
          </tr>
        </table>
        <Button
          block
          id="btn2"
          on:click={() => {
            handleTambah("dasar");
          }}>TAMBAH</Button
        >
      </div>
    </CardBody>
  </Card>
{/if}
<Modal modal_id={"modalError"} modal_size={"modal-dialog-centered"}>
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
    <h5 class="modal-title" id="exampleModalLabel">ERROR</h5>
  </slot:template>
  <slot:template slot="body">
    {@html temp_bulk_error}
  </slot:template>
</Modal>
<div class="clearfix" />
<br />
<Tablekeranjangdasar
  on:removekeranjang={removekeranjang}
  on:removekeranjangall={removekeranjangall}
  on:handleSave={handleSave}
  {client_device}
  {group_btn_beli}
  {count_line_dasar}
  {count_line_standart}
  {keranjang}
  {totalkeranjang}
  {min_bet}
  {max_bet}
  {kei_besar_bet}
  {kei_kecil_bet}
  {kei_genap_bet}
  {kei_ganjil_bet}
  {disc_besar_bet}
  {disc_kecil_bet}
  {disc_genap_bet}
  {disc_ganjil_bet}
/>

<style>
  input,
  input::-webkit-input-placeholder {
    font-size: 13px;
    color: #8a8a8a;
  }
  input::placeholder {
    color: #8a8a8a;
  }
</style>
