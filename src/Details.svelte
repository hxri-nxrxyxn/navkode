<script>
    import { get } from "./fetch";
    import { Link } from "svelte-routing";

    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get("id"); // Get the value of 'category'

    let promise = get(`/api/v1/landmark/${id}`); // Store the promise
    promise.then((p) => console.log(p.data));
</script>

<main>
    {#await promise}
        <p>Loading data...</p>
    {:then response}
        <p>hey</p>
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
