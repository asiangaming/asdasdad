<script>
	import Router from "svelte-spa-router";
	import { wrap } from "svelte-spa-router/wrap";
	import Navbar from "./components/Navbar.svelte";
	import Dashboard from "./pages/dashboard/Dashboard.svelte";
	import Admin from "./pages/admin/Admin.svelte";
	import Adminrule from "./pages/adminrule/Adminrule.svelte";
	import Sdsb4dday from "./pages/sdsb4dday/Sdsb4dday.svelte";
	import Sdsb4dnight from "./pages/sdsb4dnight/Sdsb4dnight.svelte";
	import Prediksi from "./pages/prediksi/Prediksi.svelte";
	import Login from "./pages/Login.svelte";
	import NotFound from "./pages/NotFound.svelte";
	export let table_header_font;
	export let table_body_font;
	let token = localStorage.getItem("token");
	let routes = "";
	let isNav = false;
	if (token == "" || token == null) {
		routes = {
			"/": wrap({
				component: Login,
			}),
			"*": NotFound,
		};
	} else {
		isNav = true;
		routes = {
			"/": wrap({
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
				component: Dashboard,
			}),
			"/prediksi": wrap({
				component: Prediksi,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/sdsb4dday": wrap({
				component: Sdsb4dday,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/sdsb4dnight": wrap({
				component: Sdsb4dnight,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/admin": wrap({
				component: Admin,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"/adminrule": wrap({
				component: Adminrule,
				props: {
					table_header_font: table_header_font,
					table_body_font: table_body_font,
				},
			}),
			"*": NotFound,
		};
	}
</script>

{#if isNav}
	<Navbar />
{/if}
<Router {routes} />
