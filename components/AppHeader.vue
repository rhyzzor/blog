<template>
  <header>

    <NuxtLink to="/" class="title">
      <h1>Rhyzzor's Blog</h1>
    </NuxtLink>

    <AppNav />
  </header>

</template>


<style>
.title {
  text-decoration: none;
  color: inherit;

  :hover {
    text-decoration: underline;
  }
}

.title h1 {
  font-size: 1.75rem;
}

header {
  display: flex;
  flex-direction: row;
  align-items: center;

  justify-content: space-between;
}
</style>