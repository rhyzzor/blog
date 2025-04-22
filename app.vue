<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>

<script lang="ts" setup>
useHead({
	htmlAttrs: {
		lang: "en",
	},
	link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }],
});
</script>

<style>
.slide-right-enter-active,
.slide-right-leave-active {
  transition: all 0.2s;
}

.slide-right-enter-from {
  opacity: 0;
  transform: translate(-50px, 0);
}

.slide-right-leave-to {
  opacity: 0;
  transform: translate(50px, 0);
}
</style>