<template>
  <nav class="navbar">
    <ul>
      <li>
        <NuxtLink to="/projects">Projects</NuxtLink>
      </li>
      <li>
        <NuxtLink to="https://github.com/rhyzzor" target="_blank">GitHub</NuxtLink>
      </li>
      <li>
        <NuxtLink to="https://www.linkedin.com/in/ryanvsouza/" target="_blank">LinkedIn</NuxtLink>
      </li>
    </ul>
  </nav>
</template>

<script lang="ts" setup>

</script>

<style>
.navbar ul {
  display: flex;
  gap: var(--spacing);
}

.navbar ul li {
  display: inline;
}
</style>