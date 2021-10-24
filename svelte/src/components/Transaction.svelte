<script context="module">
  import { writable } from "svelte/store";
  import { of, from } from "rxjs";
  import { ajax } from "rxjs/ajax";
  import {
    map,
    tap,
    concatMap,
    catchError,
    switchMap,
    startWith,
  } from "rxjs/operators";
  import { onMount$ } from "svelte-rx";

  export const transaction$ = (() => {
    const save$ = (data) => {
      ajax({
        url: "/api/savetransaksi",
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: data,
      }).pipe(
        map((response) => console.log("response: ", response)),
        catchError((error) => {
          console.log("error: ", error);
          return of(error);
        })
      );
    };
    // const pop = () => update((arr) => (arr.shift(), arr));
    return {
      save$,
    };

  })();
</script>
