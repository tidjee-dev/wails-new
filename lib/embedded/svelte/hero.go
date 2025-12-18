package svelte

const HeroSvelte = `<script>
  import TechIcons from "./TechIcons.svelte";
</script>

<section
  class="flex flex-col items-center justify-center text-center space-y-10"
>
  <h1
    class="text-5xl sm:text-6xl font-extrabold bg-linear-to-r from-orange-400 via-pink-500 to-purple-500 text-transparent bg-clip-text"
  >
    Welcome to Your Svelte App
  </h1>

  <p class="text-gray-300 text-base sm:text-lg max-w-lg">
    A modern starter powered by Svelte, Tailwind, Vite, and Wails.
  </p>

  <TechIcons />
</section>
`
