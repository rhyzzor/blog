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

useSeoMeta({
	description: "A blog to share my interests and document my journey.",
	ogDescription: "A blog to document my journey",
});
</script>