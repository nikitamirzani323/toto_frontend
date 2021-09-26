<script>
    import { Col, Card, CardBody } from "sveltestrap";
    import dayjs from "dayjs";
    import utc from "dayjs/plugin/utc";
    import timezone from "dayjs/plugin/timezone";

    dayjs.extend(utc);
    dayjs.extend(timezone);

    export let client_token = "";
    export let client_username = "";
    export let client_credit = 0;
    export let client_ipaddress = "";
    export let client_timezone = "";
    export let client_device = "";

    let display_credit = 0;
    let clockmachine = "";
    function updateClock() {
        let endtime = dayjs()
            .tz(client_timezone)
            .format("DD MMM YYYY | HH:mm:ss");

        clockmachine = endtime;
    }
    $: {
        setInterval(updateClock, 100);
    }
    display_credit = new Intl.NumberFormat().format(client_credit);
</script>

{#if client_device == "WEBSITE"}
    <center style="margin-bottom:5px;">
        <a href="/?token={client_token}" title="isbtoto">
            <img
                id="imglogo"
                alt=""
                style="margin-top:0px;"
                width="120"
                src="https://dataset.b-cdn.net/assets/images/companies/isb388/isb388.png"
            />
        </a>
    </center>
    
    <Col xxl="3" xl="3" lg="4" md="12" sm="12" style="padding:5px;margin:0px;margin-top:30px;">
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
