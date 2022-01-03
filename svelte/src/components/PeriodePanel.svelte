<script>
  import { CardHeader } from "sveltestrap";
  export let pasaran_name;
  export let permainan_title;
  export let pasaran_periode;
  export let pasaran_code;
</script>

<CardHeader style="background:#171717;">
  <div class="row">
    <div class="col">
      <h2
        style="font-weight:bold; padding-top:2px; color:white; font-family: 'Poppins';"
      >
        {pasaran_name}
      </h2>
    </div>
    <div class="col">
      <div class="float-end">
        <h5 style="padding-top: 10px;color:white; font-family: 'Poppins';">
          {permainan_title} | PERIODE : {pasaran_periode + " - " + pasaran_code}
        </h5>
      </div>
    </div>
  </div>
</CardHeader>
