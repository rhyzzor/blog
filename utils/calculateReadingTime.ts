export default (text: string) => {
	const wordsPerMinute = 200; // Average reading speed
	const words = text.trim().split(/\s+/).length;
	const minutes = Math.ceil(words / wordsPerMinute);
	const seconds = Math.ceil((words % wordsPerMinute) / (wordsPerMinute / 60));

	return minutes > 0 ? `${minutes} min` : `${seconds} sec`;
};
