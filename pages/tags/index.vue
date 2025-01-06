<template>
  <section>
    <h2>Tags</h2>
    <div v-for="tag in uniqueTags" :key="tag">
      <span><NuxtLink :to="`/tags/${tag}`">{{ tag }} </NuxtLink> - {{ postsByTag[tag].length }}</span>
    </div>
  </section>
</template>

<script lang="ts" setup>
const { data } = await useAsyncData("posts", () =>
	queryContent("posts").only(["title", "tags"]).find(),
);

const uniqueTags = Array.from(new Set(data.value.flatMap((post) => post.tags)));

const postsByTag = uniqueTags.reduce(
	(acc, tag) => {
		acc[tag] = data.value.filter((post) => post.tags.includes(tag));
		return acc;
	},
	{} as Record<string, unknown[]>,
);
</script>

<style>

</style>