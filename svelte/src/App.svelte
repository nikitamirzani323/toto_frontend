<script>
	import { Row,Col, Container } from "sveltestrap";
	import Home from "./pages/Home.svelte";
	import Permainan from "./pages/Permainan.svelte";
	import Loader from "./components/Loader.svelte";
	import Footer from "./components/Footer.svelte";
	import dayjs from "dayjs";
	import utc from "dayjs/plugin/utc";
	import timezone from "dayjs/plugin/timezone";

	dayjs.extend(utc);
	dayjs.extend(timezone);

	const queryString = window.location.search;
	const urlParams = new URLSearchParams(queryString);
	const token_browser = urlParams.get("token");
	let client_device = "";
	if (token_browser === null) {
		console.log("TOKEN NOT FOUND");
	} else {
		initTimezone();
	}
	if (
		/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(
			navigator.userAgent
		)
	) {
		// true for mobile device
		client_device = "MOBILE";
	} else {
		client_device = "WEBSITE";
	}

	let listkeluaran = [];
	let pasaran_name = "";
	let pasaran_code = "";
	let pasaran_periode = 0;
	let permainan = "4-3-2";
	let client_token = "";
	let client_company = "";
	let client_username = "";
	let client_ipaddress = "";
	let client_timezone = "";
	let client_credit = 0;
	const pasaran = (e) => {
		pasaran_code = e.detail.code;
		pasaran_name = e.detail.name;
		pasaran_periode = e.detail.periode;
	};

	let record = "";
	async function initTimezone() {
		const res = await fetch("https://ipinfo.io/json?token=a50f35c7f0138c");
		if (!res.ok) {
			const message = `An error has occured: ${res.status}`;
			throw new Error(message);
		} else {
			const json = await res.json();
			client_ipaddress = json.ip;
			client_timezone = json.timezone;
			initapp(token_browser);
		}
	}
	async function initapp(e) {
		const res = await fetch("/api/init", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				token: e,
			}),
		});

		const json = await res.json();
		if (json.status === 200) {
			client_token = json.token;
			client_company = json.company;
			client_username = json.developer;
			client_credit = json.credit;
			initPasaran();
		}
	}
	async function initPasaran() {
		const res = await fetch("/api/listpasaran", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				token: client_token,
				company: client_company,
				timezone: client_timezone,
			}),
		});

		const json = await res.json();
		if (json.status == 200) {
			record = json.record;
			if (record != null) {
				for (var i = 0; i < record.length; i++) {
					listkeluaran = [
						...listkeluaran,
						{
							id: record[i]["pasaran_id"],
							pasaran_code: record[i]["pasaran_id"],
							pasaran: record[i]["pasaran_togel"],
							pasaran_periode: record[i]["pasaran_periode"],
							pasaran_tgl: dayjs(record[i]["pasaran_marketclose"])
								.tz(client_timezone)
								.format("DD MMM YYYY | HH:mm:ss"),
							pasaran_status: record[i]["pasaran_status"],
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
</script>

<svelte:head>
	<title>ISBTOTO APP</title>
</svelte:head>
{#if client_device == "WEBSITE"}
	<div class="content" style="margin-top:20px;margin-bottom:50px;">
		<Container>
			{#if token_browser != ""}
				<Row align-items-start>
					{#if pasaran_code != ""}
						{#if client_token != ""}
							<Permainan
								{client_token}
								{client_company}
								{client_username}
								{client_credit}
								{client_ipaddress}
								{client_timezone}
								{client_device}
								{pasaran_code}
								{pasaran_name}
								{pasaran_periode}
								{permainan}
							/>
						{/if}
					{:else if client_token != ""}
						<Home
							{client_token}
							{client_username}
							{client_credit}
							{client_ipaddress}
							{client_timezone}
							{client_device}
							{listkeluaran}
							on:pasaran={pasaran}/>
					{:else}
						<div style="height: 100%;margin:100px 0px 100px 0px;">
							<center>
								<Loader
									cssstyle={"height: 100%;margin:100px 0px 100px 0px;"}
								/>
							</center>
						</div>
					{/if}
				</Row>
			{/if}
		</Container>
	</div>
{:else}
	<Container>
		{#if token_browser != ""}
			<Row align-items-start>
				{#if pasaran_code != ""}
					{#if client_token != ""}
						<Permainan
							{client_token}
							{client_company}
							{client_username}
							{client_credit}
							{client_ipaddress}
							{client_timezone}
							{client_device}
							{pasaran_code}
							{pasaran_name}
							{pasaran_periode}
							{permainan}
						/>
					{/if}
				{:else if client_token != ""}
					<Home
						{client_token}
						{client_username}
						{client_credit}
						{client_ipaddress}
						{client_timezone}
						{client_device}
						{listkeluaran}
						on:pasaran={pasaran}
					/>
				{:else}
					<div style="height: 100%;margin:100px 0px 100px 0px;">
						<center>
							<Loader
								cssstyle={"height: 100%;margin:100px 0px 100px 0px;"}
							/>
						</center>
					</div>
				{/if}
			</Row>
		{/if}
		<Footer />
	</Container>
{/if}
