<script>
    import { get } from "./fetch";
    import { Link } from "svelte-routing";
    import Nav from "./Nav.svelte";

    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get("id"); // Get the value of 'category'

    let promise = get(`/api/v1/landmark/${id}`); // Store the promise
    promise.then((p) => console.log(p.data));
</script>

<main>
    <Nav />
    {#await promise}
        <p>Loading data...</p>
    {:then response}
        <h1>Here's the shortest path to {response.data.name}</h1>
        <div class="card">
            <div class="card__image"></div>
            <div class="card__text">
                <h2>{response.data.name}</h2>
                <p>{response.data.id}</p>
                <div class="card__stars">stars</div>
            </div>
        </div>
    {:catch error}
        <p style="color: red;">Error loading data: {error.message}</p>
    {/await}
</main>

<style>
</style>
