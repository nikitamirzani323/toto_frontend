<script>
    export let client_token = "";
    export let client_device = "";
</script>

{#if client_device == "WEBSITE"}
    <nav class="navbar fixed-top " style="background-color: #e91e63;">
        <div class="container">
            <center>
                <a
                    style="margin-left: -20px;"
                    class="navbar-brand"
                    href="/?token={client_token}"
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
{:else}
    <nav class="navbar fixed-top " style="background-color: #e91e63;">
        <div class="container-fluid">
            <center>
                <a
                    style="margin-left: -20px;"
                    class="navbar-brand"
                    href="/?token={client_token}"
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