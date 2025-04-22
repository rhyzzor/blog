<template>
  <div class="projects-container">
    <p class="main-description">Welcome to my projects showcase. Here you'll find a collection of my work including
      open-source NPM packages and projects I've developed. Each project represents my approach to solving specific
      problems with clean, maintainable code.</p>

    <section class="projects-section">
      <div class="project-grid">
        <div v-for="project in projects" :key="project.name" class="project-card">
          <div class="project-header">
            <h3>
              <NuxtLink :to="project.link" target="_blank" rel="noopener noreferrer">{{ project.name }}</NuxtLink>
            </h3>
            <span class="project-type-tag" :class="`type-${project.type}`">
              {{ project.type }}
            </span>
          </div>
          <p class="description">{{ project.description }}</p>
          <div class="tech-tags">
            <span v-for="tech in project.technologies" :key="tech" class="tag">{{ tech }}</span>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script lang="ts" setup>
useSeoMeta({
	description:
		"Explore my portfolio of open-source NPM packages and projects including CLI tools, bots, and web applications.",
	ogDescription:
		"Explore my portfolio of open-source NPM packages and projects including CLI tools, bots, and web applications.",
	robots: "index, follow",
});

const projects = [
	{
		name: "translate-table",
		description:
			"A powerful CLI tool for translating PostgreSQL table content to multiple languages.",
		link: "https://www.npmjs.com/package/translate-table",
		type: "npm",
		technologies: ["TypeScript", "PostgreSQL", "CLI"],
	},
	{
		name: "Rhyzzor Bot",
		description:
			"Simplify character creation, track stats, manage inventory, and focus on what matters most - your adventure!",
		link: "https://rpg.rhyzzor.com",
		type: "bot",
		technologies: ["Next.js", "Discord.js", "Turso"],
	},
];
</script>

<style>
.projects-container {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-2x);
}

.main-description {
  color: #444;
}

.project-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: var(--spacing);
}

.project-card {
  display: flex;
  flex-direction: column;
  gap: var(--spacing);
  padding: 1.5rem;
  border-radius: 8px;
  background-color: #f9f9f9;
  transition: transform 0.2s, box-shadow 0.2s;
}

.project-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.project-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.project-card a {
  text-decoration: none;
  color: #3182ce;
}

.project-card a:hover {
  text-decoration: underline;
}

.description {
  color: #4a5568;
  line-height: 1.5;
}

.tech-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin-top: auto;
}

.tag {
  font-size: 0.75rem;
  padding: 0.25rem 0.5rem;
  background-color: #e2e8f0;
  color: #4a5568;
  border-radius: 4px;
}

.project-type-tag {
  font-size: 0.7rem;
  padding: 0.2rem 0.5rem;
  border-radius: 4px;
  font-weight: 500;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.type-npm {
  background-color: #cb3837;
  color: white;
}

.type-web {
  background-color: #3182ce;
  color: white;
}

.type-bot {
  background-color: #5865F2;
  color: white;
}
</style>