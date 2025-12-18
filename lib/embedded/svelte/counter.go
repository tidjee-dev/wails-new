package svelte

const CounterSvelte = `<script lang="ts">
  let count: number = $state(0);

  const increment = () => count++;
  const reset = () => (count = 0);
</script>

<section class="flex flex-col items-center text-center space-y-6">
  <h3 class="text-lg sm:text-xl text-gray-300 font-medium">
    Click the button below to increment the counter
  </h3>

  <div class="flex flex-col sm:flex-row items-center justify-center gap-6">
    <div class="bg-zinc-800/70 px-6 py-4 rounded-xl shadow-md">
      <p class="text-gray-400 text-sm mb-1">Counter</p>
      <span class="text-3xl font-bold text-white tracking-wide">{count}</span>
    </div>

    <div class="flex items-center gap-4">
      <button
        onclick={increment}
        class="px-5 py-2.5 bg-linear-to-br from-pink-600 via-purple-600 to-indigo-600 text-white font-medium rounded-lg
               hover:from-pink-500 hover:via-purple-500 hover:to-indigo-500 transition-all duration-200 shadow-md hover:shadow-lg"
      >
        Increment
      </button>

      <button
        onclick={reset}
        class="text-sm text-gray-400 hover:text-gray-200 underline-offset-4 hover:underline transition-colors"
      >
        Reset
      </button>
    </div>
  </div>
</section>
`
