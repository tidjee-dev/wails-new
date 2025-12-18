package svelte

const TechIconsSvelte = `<script>
  import Svelte from "./ui/icons/Svelte.svelte";
  import Tailwind from "./ui/icons/Tailwind.svelte";
  import Vite from "./ui/icons/Vite.svelte";
  import Wails from "./ui/icons/Wails.svelte";

  const techs = [
    { name: "Wails", icon: Wails, url: "https://wails.io" },
    { name: "Vite", icon: Vite, url: "https://vitejs.dev" },
    { name: "Svelte", icon: Svelte, url: "https://svelte.dev" },
    { name: "Tailwind", icon: Tailwind, url: "https://tailwindcss.com" },
  ];
</script>

<section class="flex flex-col items-center justify-center text-gray-200">
  <h2
    class="text-xl font-semibold tracking-wide mb-6 uppercase text-gray-300 underline underline-offset-4"
  >
    Built With
  </h2>

  <div class="flex flex-wrap justify-center gap-10">
    {#each techs as { name, icon: Icon, url }}
      <a
        href={url}
        target="_blank"
        rel="noreferrer"
        class="group flex flex-col items-center transition-transform hover:scale-105"
      >
        <div
          class="p-5 h-24 rounded-2xl bg-zinc-800/80 backdrop-blur-sm
                 transition-all duration-300 shadow-md hover:shadow-[0_0_30px_rgba(255,255,255,0.15)]"
        >
          <Icon />
        </div>
        <span class="mt-3 text-sm text-gray-400 group-hover:text-gray-100">
          {name}
        </span>
      </a>
    {/each}
  </div>
</section>
`
