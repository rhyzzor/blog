<template>
 <div class="not-found">
    <h2> Page Not Found </h2>
    <p>Sorry, the page you are looking for does not exist.</p>
    <NuxtLink to="/">Go back to home</NuxtLink>
  </div>
</template>

<style>
.not-found {
  display: flex;
  flex-direction: column;
  gap: var(--spacing);
  align-items: center;
  justify-content: center;
  height: 60vh;
}
</style>