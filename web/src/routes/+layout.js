export const ssr = false;
export function load({ url, route }) {
	return {
		transPageKey: url.pathname 
	};
};
