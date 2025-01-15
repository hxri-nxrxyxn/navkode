<script>
    import Hero from "./Hero.svelte";
    import List from "./List.svelte";
    import Details from "./Details.svelte";
    import Login from "./Login.svelte";
    import Signup from "./Signup.svelte";
    import { Route, Router, Link } from "svelte-routing";

    import { onMount, onDestroy } from "svelte";
    import { App } from "@capacitor/app";
    import { useRouter } from "svelte-routing";
    const router = useRouter();

    onMount(() => {
        const backButtonHandler = ({ canGoBack }) => {
            if (canGoBack) {
                // Use pop() if using svelte-spa-router:
                pop();
                // OR, if NOT using svelte-spa-router, use browser history:
                // router.history.go(-1); // or router.history.back()
            }
            App.addListener("backButton", backButtonHandler);
            return () => {
                App.removeAllListeners(); // IMPORTANT: Remove listeners on unmount
            };
        };
    });
</script>

<main>
    <Router>
        <Route path="/hero" component={Hero} />
        <Route path="/list" component={List} />
        <Route path="/details" component={Details} />
        <Route path="/login" component={Login} />
        <Route path="/" component={Signup} />
    </Router>
</main>

<style>
    main {
        padding: 2rem;
    }
</style>
