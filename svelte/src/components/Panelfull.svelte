<script>
    import { Card, CardHeader, CardBody, CardFooter } from "sveltestrap";

    export let header = true;
    export let footer = true;
    export let header_style = "";
    export let body_style = "";
</script>

<Card color="dark" style="border:none;padding:0px;margin:0px;">
    {#if header}
        <CardHeader style={header_style}>
            <slot name="header" />
        </CardHeader>
    {/if}
    <CardBody id="stream" data-simplebar style={body_style}>
        <slot name="body" />
    </CardBody>
    {#if footer}
        <CardFooter style="background:#101010;">
            <slot name="footer" />
        </CardFooter>
    {/if}
</Card>

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
