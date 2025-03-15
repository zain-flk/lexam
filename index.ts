/**
 * Welcome to Cloudflare Workers! This is your first worker.
 *
 * - Run `npm run dev` in your terminal to start a development server
 * - Open a browser tab at http://localhost:8787/ to see your worker in action
 * - Run `npm run deploy` to publish your worker
 *
 * Bind resources to your worker in `wrangler.jsonc`. After adding bindings, a type definition for the
 * `Env` object can be regenerated with `npm run cf-typegen`.
 *
 * Learn more at https://developers.cloudflare.com/workers/
 */

function cpuIntensiveTask() {
    const startTime = Date.now();
    let count = 0;

    while (Date.now() - startTime < 2000) { // Run for 2 seconds
        count++;
        if (count > 1000) {
            count = 0; 
        }
    }

    return ("Mildly CPU intensive task completed!");
}

export default {
	async fetch(request, env, ctx): Promise<Response> {

		cpuIntensiveTask();
		return new Response(cpuIntensiveTask());
	},
	
} satisfies ExportedHandler<Env>;
