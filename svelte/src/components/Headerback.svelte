<script>
  import { Col, Card, CardBody } from "sveltestrap";
  import dayjs from "dayjs";

  export let client_token = "";
  export let client_company = "";
  export let client_username = "";
  export let client_credit = 0;
  export let client_ipaddress = "";
  export let client_timezone = "";
  export let client_device = "";
  let clockmachine = "";
  let display_credit = 0;

  function updateClock() {
    let endtime = dayjs().tz(client_timezone).format("DD MMM YYYY | HH:mm:ss");

    clockmachine = endtime;
  }
  $: {
    setInterval(updateClock, 100);
  }
  display_credit = new Intl.NumberFormat().format(client_credit);
</script>

{#if client_device == "WEBSITE"}
  <nav class="navbar">
    <div class="col">
      <a
        style="margin-left: -10px;"
        class="navbar-brand"
        href="/?token={client_token}&agent={client_company}"
        title="isbtoto"
      >
        <img style="margin-top: 0px;" width="135" src="logo.svg" alt="close" />
      </a>
    </div>
    <div class="col">
      <Card style="border:none;background-color:transparent;text-align:right;">
        <CardBody
          style="
        padding-right: 0;
        margin-right: -10px;"
        >
          <span style="font-size:13px;"
            >Timezone : <span id="style_text"
              >{client_timezone}, {clockmachine} WIB</span
            ></span
          ><br />

          <span style="font-size:15px;color:#fff;"
            >{client_username} ({client_ipaddress})</span
          ><br />
          <span style="font-size:15px;color:#fff;"
            >Saldo : IDR <span id="style_text">{display_credit}</span></span
          ><br />
        </CardBody>
      </Card>
    </div>
  </nav>
{:else}
  <nav class="navbar fixed-top " style="background-color: #2b2a33;">
    <div class="container-fluid">
      <center>
        <a
          style="margin-left: -20px;"
          class="navbar-brand"
          href="/?token={client_token}&agent={client_company}"
          title="isbtoto"
        >
          <img
            style="margin-top: 0px;"
            width="30"
            src="back_home.png"
            alt="close"
          />
        </a>
      </center>
    </div>
  </nav>
{/if}
