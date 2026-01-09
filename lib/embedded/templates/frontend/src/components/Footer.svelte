<script lang="ts">
  // import { OpenExternalURL } from "../../wailsjs/go/main/App";

  function open(url: string) {
    window.open(url, "_blank");
    // OpenExternalURL(url);
  }
</script>

<footer
  class="mt-auto bg-zinc-950 border-t border-zinc-800 py-6 text-center text-gray-300 text-sm"
>
  <p>
    View source on
    <button
      type="button"
      class="text-blue-400 hover:text-blue-300 underline transition-colors bg-transparent p-0 border-none cursor-pointer"
      on:click={() => open("https://github.com/tidjee-dev/wails-new")}
    >
      GitHub
    </button>
  </p>

  <p class="mt-2">
    Made with <span class="text-red-500">♥</span> by
    <button
      type="button"
      class="text-blue-400 hover:text-blue-300 underline transition-colors bg-transparent p-0 border-none cursor-pointer"
      on:click={() => open("https://github.com/tidjee-dev")}
    >
      Tidjee
    </button>
    — &copy; {new Date().getFullYear()}
  </p>
</footer>
