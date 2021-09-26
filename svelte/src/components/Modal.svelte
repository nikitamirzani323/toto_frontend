<script>
    export let modal_id = "";
    export let modal_size = "";
</script>
<div class="modal fade" id="{modal_id}" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
    <div class="modal-dialog {modal_size}">
        <div class="modal-content">
            <div class="modal-header" style="border-bottom:none;background: #2e323e;display:block;">
                <slot name="header" />
            </div>
            <div class="modal-body" id="streaminvoice" style="font-size:12px;padding:0px;margin:0px;background:#121212;border:1px solid #0e0c13;height:350px;overflow-y:scroll;color:white;">
                <slot name="body" />
            </div>
        </div>
    </div>
</div>