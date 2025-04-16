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