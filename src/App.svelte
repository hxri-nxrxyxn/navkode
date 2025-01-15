<script>
  import svelteLogo from './assets/svelte.svg'
  import Hero from './Hero.svelte'
  import List from './List.svelte'
  import {Route, Router, Link} from 'svelte-routing'
</script>

<main>
    <Router>
        <Route path="/" component="{Hero}"/>
        <Route path="/list" component="{List}"/>
    </Router>
</main>

<style></style>
