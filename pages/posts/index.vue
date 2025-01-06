<template>
  <section>
    <h2>Posts</h2>
      <div v-for="post in data" :key="post._path">
        <Post v-bind="post" />
      </div>
  </section>
</template>

<script lang="ts" setup>
  const { data } = await useAsyncData('posts', () => queryContent('posts').find())

</script>

<style>

</style>