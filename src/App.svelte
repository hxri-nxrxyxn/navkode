<script>
    import Hero from "./Hero.svelte";
    import List from "./List.svelte";
    import Details from "./Details.svelte";
    import Login from "./Login.svelte";
    import { Route, Router, Link } from "svelte-routing";
</script>

<main>
    <Router>
        <Route path="/" component={Hero} />
        <Route path="/list" component={List} />
        <Route path="/details" component={Details} />
        <Route path="/login" component={Login} />
    </Router>
</main>

<style>
    main {
        padding: 2rem;
    }
</style>
