<script>
    export let modal_id = "";
    export let modal_size = "";
    export let modal_body_height = "height:350px;";
    export let modal_footer_flag = true;
</script>

<div class="modal fade" id={modal_id} tabindex="-1" data-bs-backdrop="false">
    <div class="modal-dialog {modal_size}">
        <div class="modal-content">
            <div class="modal-header" style="width:100%;background: #2e323e;">
                <slot name="header" />
                <button
                    type="button"
                    class="btn-close btn-close-white"
                    data-bs-dismiss="modal"
                    aria-label="Close"
                />
            </div>
            <div
                id="stream"
                data-simplebar
                class="modal-body"
                style="font-size:12px;padding:0px;margin:0px;color:black;background:#121212;border:1px solid #0e0c13;{modal_body_height}"
            >
                <slot name="body" />
            </div>
            {#if modal_footer_flag}
                <div class="modal-footer" style="margin:0px;border-top: none;">
                    <slot name="footer" />
                </div>
            {/if}
        </div>
    </div>
</div>

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
</style>