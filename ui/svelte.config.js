import node from '@sveltejs/adapter-node';
import preprocess from 'svelte-preprocess';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	preprocess: preprocess(),

	kit: {
		adapter: node({
			out: 'build',
			precompress: true,
			env: {
				host: 'HOST',
				port: 'PORT'
			}
		}),

		vite: () => ({
			envPrefix: 'HEROES_'
		})
	}
};

export default config;
