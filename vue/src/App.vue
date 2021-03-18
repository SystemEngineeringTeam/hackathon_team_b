<template>
  <div id="app">
    <nav class="nav flex flex-wrap items-center justify-between px-4">
      <div class="flex flex-no-shrink items-center mr-6 py-3">
        <router-link to="/">
          <p
            class="font-semibold text-2xl tracking-tight no-underline hover:no-underline"
          >
            {{ topPageName }}
          </p>
        </router-link>
      </div>
      <ul class="md:border-none flex justify-end m-0 w-full md:w-auto">
        <li class="border-t md:border-none">
          <router-link to="/">
            <a
              class="block md:inline-block px-4 py-3 hover:text-red-400 no-underline hover:no-underline"
              >{{ homeVueName }}</a
            >
          </router-link>
        </li>
        <li class="border-t md:border-none"></li>
      </ul>
    </nav>
    <router-view />
  </div>
</template>

<script>
export default {
  data() {
    return {
      topPageName: "hackathon_team_b",
      homeVueName: "検索トップ",
    };
  },
};
</script>

<style>

</style>