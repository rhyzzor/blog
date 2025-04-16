<template>
  <section class="description">
    <p>I'm a software developer from Brazil, currently 22 years old. This blog serves as a personal journal where I
      share my interests and document my journey - both in my career and life. It's a space to reflect on my experiences
      and keep track of everything I've learned and accomplished along the way.
    </p>

    <p>
      This is the <a href="https://github.com/rhyzzor/blog" target="_blank">blog repository</a>, you can see the
      source code of this blog.
    </p>
  </section>

  <AppListPosts />
</template>

<script lang="ts" setup>
defineOgImageComponent("NuxtSEO", {
	title: "Rhyzzor's Blog",
	theme: "#FFFFFF",
	colorMode: "dark",
});
</script>

<style>
.description {
  display: flex;
  flex-direction: column;
  gap: var(--spacing);
}
</style>