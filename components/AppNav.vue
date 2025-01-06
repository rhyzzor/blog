<template>
  <nav>
    <ul>
      <li>
        <NuxtLink to="/tags">Tags</NuxtLink>
      </li>
      <li>
        <a href="https://github.com/rhyzzor" target="_blank">GitHub</a>
      </li>
      <li>
        <a href="https://www.linkedin.com/in/ryanvsouza/" target="_blank">LinkedIn</a>
      </li>
    </ul>
  </nav>
</template>

<script lang="ts" setup>

</script>

<style>
ul {
  display: flex;
  gap: 16px;
}

ul li {
  list-style-type: none;
  display: inline;
}
</style>