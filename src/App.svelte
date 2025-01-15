<script>
  import svelteLogo from './assets/svelte.svg'
  import Hero from './Hero.svelte'
  import {Route, Router, Link} from 'svelte-routing'
</script>

<main>
    <Router>
        <Route path="/" component="{Hero}"/>
    </Router>
</main>

<style></style>
