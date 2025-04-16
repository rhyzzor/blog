import { defineCollection, defineContentConfig, z } from "@nuxt/content";
import { asSeoCollection } from "@nuxtjs/seo/content";

export default defineContentConfig({
	collections: {
		blog: defineCollection(
			asSeoCollection({
				type: "page",
				source: "blog/*.md",
				schema: z.object({
					date: z.date(),
					rawbody: z.string(),
				}),
			}),
		),
	},
});
