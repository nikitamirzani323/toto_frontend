<script>
  import { Card, CardBody, CardHeader } from "sveltestrap";
</script>

<Card style="background:rgba(0,0,0.78); opacity:.9;">
  <CardBody style="padding:10px;margin:0px;">
    <slot name="body" />
  </CardBody>
</Card>
