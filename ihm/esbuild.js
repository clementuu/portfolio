import * as esbuild from 'esbuild';
import sveltePlugin from 'esbuild-svelte';

const isDev = process.argv.includes('--dev')

let commonOptions = {
	entryPoints: ['./App.svelte', './moi.svelte', './projets/projets.svelte', './projets/detail.svelte', './competences/competences.svelte', './competences/detail.svelte', './contacts.svelte'],
	bundle: true,
	format: 'esm',
	plugins: [
		sveltePlugin({
			compilerOptions: { customElement: true}
		})
	],
}
let devOptions = {
	...commonOptions,
	outdir: './build',
	banner: {
        //crée un eventlistener qui détecte les modifications du code et actualise la page pour afficher les modifs en direct
		js: "new EventSource('http://127.0.0.1:8088/esbuild').addEventListener('change', () => location.reload())"
	},
	logLevel: 'info'
};
let prodOptions = {
	...commonOptions,
	outdir: './build',
	minify: true,
	logLevel: 'error'
}

if (!isDev) {
	await esbuild.build(prodOptions);
	process.exit(0);
}

let ctx = await esbuild.context(devOptions);
await ctx.watch();
await ctx.serve({
	servedir: './',
	port: 8088,
	host: '127.0.0.1'
});