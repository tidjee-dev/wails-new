package svelte

const AppSvelte = `<script lang="ts">
  import Counter from "./components/Counter.svelte";
  import Footer from "./components/Footer.svelte";
  import Help from "./components/Help.svelte";
  import Hero from "./components/Hero.svelte";
  import Separator from "./components/ui/Separator.svelte";
</script>

<div
  class="min-h-screen bg-linear-to-b from-zinc-950 via-zinc-900 to-zinc-950 text-gray-200 flex flex-col"
>
  <main
    class="flex flex-col items-center justify-center px-4 py-16 space-y-12 grow"
  >
    <Hero />

    <Separator />

    <Counter />

    <Separator />

    <Help />
  </main>

  <Footer />
</div>
`
