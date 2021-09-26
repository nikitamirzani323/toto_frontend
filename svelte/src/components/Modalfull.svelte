<script>
    export let modal_id = "";
    export let modal_size = "";
</script>

<div
    class="modal fade"
    id={modal_id}
    tabindex="-1"
    aria-labelledby="exampleModalLabel"
    aria-hidden="true"
>
    <div class="modal-dialog {modal_size}">
        <div class="modal-content">
            <div
                class="modal-header"
                style="border-bottom:none;background: #2e323e;display:block;"
            >
                <slot name="header" />
            </div>
            <div
                data-simplebar
                class="modal-body"
                id="streaminvoice"
                style="font-size:12px;padding:0px;margin:0px;background:#121212;border:1px solid #0e0c13;height:350px;color:white;"
            >
                <slot name="body" />
            </div>
            <div class="modal-footer" style="margin:0px;background: rgb(16, 16, 16);border-top: none;">
                <slot name="footer" />
            </div>
        </div>
    </div>
</div>
<style>
    #streaminvoice::-webkit-scrollbar {
        width: 0.3em;
    }

    #streaminvoice::-webkit-scrollbar-track {
        -webkit-box-shadow: inset 0 0 6px rgba(0, 0, 0, 0.3);
    }

    #streaminvoice::-webkit-scrollbar-thumb {
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
