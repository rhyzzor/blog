<template>
  <NuxtLink :to="`/posts/${slug}`">
    <strong><p>{{  date }} - {{ title }}</p> </strong>
    <p>{{ description }}</p>
  </NuxtLink>
</template>

<script lang="ts" setup>
  const props = defineProps({
    title: {
      type: String,
    },
    description: {
      type: String,
    },
    tags: {
      type: Array,
    },
    date: {
      type: String,
    }
  })

  const slug = slugify(props?.title)


</script>

<style>

</style>