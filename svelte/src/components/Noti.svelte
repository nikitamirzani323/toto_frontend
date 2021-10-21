<script context="module">
  import { writable } from "svelte/store";

  export const notifications = (() => {
    const { update, subscribe } = writable([]);
    const push = (message, messageType, messagePos) => {
      messageType = messageType ? messageType : "danger";
      switch (messagePos) {
        case "middle":
			messagePos = "top-50 start-50 translate-middle"
          break;
        default:
			messagePos = "bottom-0 start-50 translate-middle-x"
          break;
      }
      update((arr) => [...arr, { message, messageType, messagePos }]);
    };
    const pop = () => update((arr) => (arr.shift(), arr));
    return {
      pop,
      push,
      subscribe,
    };
  })();
</script>

<script>
  import { createEventDispatcher } from "svelte";

  export let duration = 3000;
  const dispatch = createEventDispatcher();
  let timeout;
  notifications.subscribe(({ length }) => {
    if (timeout || !length) return;
    dispatch("notify", $notifications[0]);
    timeout = setTimeout(() => {
      timeout = false;
      notifications.pop();
    }, duration);
  });
</script>

{#if $notifications[0]}
  <!-- <div class="notification">
		<div class="alert alert-{$notifications[0].messageType}" role="alert">
			<span></span>
		</div>
	</div> -->

   <div
      class="toast-container position-absolute p-3 {$notifications[0].messagePos} "
      id="toastPlacement"
      data-original-class="toast-container position-absolute p-3"
    >
      <div
        class="toast bg-dark text-{$notifications[0].messageType}  fade show"
      >
        <div class="toast-body border border-light rounded-2 text-center">
          {$notifications[0].message}
        </div>
      </div>
    </div>
{/if}

<style>
	.toast-container {
		z-index: 1080;
	}
</style>