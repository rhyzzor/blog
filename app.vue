<template>
  <Transition mode="out-in" name="left-to-right">
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </Transition>
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
.left-to-right-leave-active {
  transition: all 0.5s ease-out;
}

.left-to-right-leave-to {
  transform: translateX(100%);
}

.left-to-right-enter-active {
  transition: all 0.5s ease-out;
}

.left-to-right-enter-from {
  transform: translateX(-100%);
}
</style>