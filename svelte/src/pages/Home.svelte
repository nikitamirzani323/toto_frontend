<script>
  import { createEventDispatcher } from "svelte";
  import { Card, CardBody, CardHeader, CardFooter, Col } from "sveltestrap";
  import Header from "../components/Header.svelte";
  import { notifications } from "../components/Noti.svelte";

  export let client_token = "";
  export let client_company = "";
  export let client_username = "";
  export let client_credit = 0;
  export let client_ipaddress = "";
  export let client_timezone = "Asia/Jakarta";
  export let client_device = "";
  export let listkeluaran = [];
  let dispatch = createEventDispatcher();

  const handleClick = (code, name, periode, status) => {
    if (status == "ONLINE") {
      const pasaran = {
        code,
        periode,
        name,
        client_token,
      };
      dispatch("pasaran", pasaran);
    } else {
      notifications.push("PASARAN " + name + " OFFLINE");
    }
  };
</script>

<Header
  {client_username}
  {client_company}
  {client_credit}
  {client_token}
  {client_ipaddress}
  {client_timezone}
  {client_device}
/>
{#each listkeluaran as rec}
  {#if client_device == "WEBSITE"}
    <Col
      xxl="3"
      xl="3"
      lg="6"
      md="6"
      sm="6"
      xs="6"
      style="cursor:pointer;padding:5px;margin:0px;"
    >
      <div
        class="default-block web-block card"
        on:click={() => {
          handleClick(
            rec.pasaran_code,
            rec.pasaran,
            rec.pasaran_periode,
            rec.pasaran_status
          );
        }}
      >
        <div class="card-header">
          <center id="style_text" style="font-size:18px;">
            {#if rec.pasaran_status == "ONLINE"}
              <span class="badge rounded-pill bg-online blink_me"
                >{rec.pasaran_status}</span
              >
            {:else}
              <span class="badge rounded-pill bg-secondary"
                >{rec.pasaran_status}</span
              >
            {/if}
          </center>
        </div>
        <CardBody>
          <center>
            <br />
            <h2 class="head-fonts">{rec.pasaran}</h2>
            <span style="font-size: 12px;">PERIODE : {rec.pasaran_periode}</span
            ><br />
            <span style="font-size: 11px;">{rec.pasaran_tgl} WIB</span>
            {#if rec.pasaran_status == "ONLINE"}
              <br />
              <button class="btn btn-play">Play Now</button>
            {/if}
          </center>
        </CardBody>
      </div>
    </Col>
  {:else}
    <Col
      xxl="6"
      xl="6"
      lg="6"
      md="6"
      sm="6"
      xs="6"
      style="cursor:pointer;padding:5px;margin:0px;"
    >
      <div
        class="default-block mobile-block card"
        style="border:none;"
        on:click={() => {
          handleClick(
            rec.pasaran_code,
            rec.pasaran,
            rec.pasaran_periode,
            rec.pasaran_status
          );
        }}
      >
        <div class="card-header">
          {#if rec.pasaran_status == "ONLINE"}
            <div
              class="blink_me"
              style="margin-top:0px;color:black;background:#c4f750;font-size:12px;font-weight:bold;padding:5px;border-radius:15px;text-align:center;"
            >
              {rec.pasaran_status}
            </div>
          {:else}
            <div
              style="margin-top:0px;background:#f7602e;color:white;font-size:12px;font-weight:bold;padding:5px;border-radius:15px;text-align:center;"
            >
              {rec.pasaran_status}
            </div>
          {/if}
        </div>
        <CardBody>
          <center id="style_text" style="font-size:15px;">
            <h4 class="head-fonts">{rec.pasaran}</h4>
            <span style="font-size: 11px;">
              PERIODE : {rec.pasaran_periode}
            </span>
            <br />
            <span style="font-size: 10px;">{rec.pasaran_tgl} WIB</span>
          </center>
        </CardBody>
      </div>
    </Col>
  {/if}
{/each}
<div class="clearfix" />
<br />

<style>
  .btn-play {
    color: #fff;
    border-color: #ff9900;
    border-radius: 15px;
    margin-top: 3rem;
    padding: 5px 25px;
  }

  .btn-play:hover {
    background-color: #ff9900;
  }
  .bg-online {
    background-color: #009e42 !important;
  }
  .card:hover {
    border: 2px solid #ff9900;
  }
  .card-header {
    border-bottom: none !important;
    padding: 2rem 1rem 0.5rem !important;
  }
  .default-block {
    background-image: url("/bg-button.svg"),
      linear-gradient(180deg, #8e000e 0%, #5f0009 72.69%);
    background-size: cover;
    margin-bottom: 1.5rem;
    margin-top: 1.5rem;
  }

  .web-block {
    width: 300px;
    height: 300px;
    border-radius: 20px;
  }

  .mobile-block {
    width: 163px;
    height: 163px;
    border-radius: 14px;
  }
  .blink_me {
    animation: blinker 1s linear infinite;
  }

  @keyframes blinker {
    50% {
      opacity: 0;
    }
  }
</style>
