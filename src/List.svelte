<script>
    import { get } from "./fetch";
    import Nav from "./Nav.svelte";
    import Star from "./assets/star.svg";
    import { Link } from "svelte-routing";

    const urlParams = new URLSearchParams(window.location.search);
    const category = urlParams.get("category"); // Get the value of 'category'

    let promise = get(`/api/v1/landmarkCat/${category}`); // Store the promise
</script>

<main>
    <Nav />
    <h1>Top {category}s in Kozhikode</h1>
    {#await promise}
        <p>Loading data...</p>
    {:then data}
        {#if data.data && data.data.length > 0}
            {#each data.data as i}
                <Link to={`/details?id=${i.id}`}>
                    <div class="card">
                        <div class="card__image">
                            <img src={i.preview} alt="" />
                        </div>
                        <div class="card__text">
                            <h2>{i.name}</h2>
                            <p>{i.location}</p>
                            <div class="card__stars">
                                {#each {length: i.star}}
                                    <img src={Star} alt="" />
                                {/each}
                            </div>
                        </div>
                    </div>
                </Link>
            {/each}
        {:else}
            <p>No data found.</p>
        {/if}
    {:catch error}
        <p style="color: red;">Error loading data: {error.message}</p>
    {/await}
</main>

<style>
    .card {
        display: flex;
        margin: 1rem 0;
        border-radius: 1rem;
    }
    .card__image {
        width: 30%;
    }
    .card__image img {
        width: 100%;
        height: 100%;
        border-top-left-radius: 1rem;
        border-bottom-left-radius: 1rem;
    }
    .card__text {
        flex: 1;
        padding: 1rem;
        font-size: 0.8rem;
        background: var(--color-shade);
        border-top-right-radius: 1rem;
        border-bottom-right-radius: 1rem;
    }
    .card__text h2 {
        padding-bottom: 0.2rem;
        font-size: 1.5rem;
        font-weight: 800;
    }
    .card__text p {
        color: var(--color-light);
    }
    .card__stars {
        margin-top: 1rem;
    }
    .card__stars img {
        width: 1rem;
        margin-right: 0.25rem;
        filter: invert(50%) sepia(55%) saturate(4014%) hue-rotate(331deg)
            brightness(100%) contrast(100%);
    }
</style>
