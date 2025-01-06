

<template>

  <section class="description">
    <p>
      I'm a software developer from Brazil, currently 22 years old. This blog serves as a personal journal
      where I share my interests and document my journey - both in my career and life. It's a space to reflect on my
      experiences and keep track of everything I've learned and accomplished along the way.
    </p>

    <p>
      This is the <a href="https://github.com/rhyzzor/blog" target="_blank">blog repository</a>, you can see the source code of this blog.
    </p>
  </section>

  <section>
    <h2>Posts</h2>
    <div class="posts">
      <Post v-bind="post" v-for="post in data" :key="post._path"/>
      <NuxtLink to="/posts">
        <p> <strong>See all posts...</strong></p>
      </NuxtLink>
    </div>

  </section>

</template>

<script lang="ts" setup>
const { data } = await useAsyncData("posts", () =>
	queryContent("posts")
		.limit(3)
		.sort({
			date: -1,
		})
		.find(),
);
</script>

<style>
.description {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.posts {
  display: flex;
  gap: 16px;
  flex-direction: column;
}
</style>