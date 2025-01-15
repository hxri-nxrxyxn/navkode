<script>
    let category;
    import { get } from "./fetch";
    const urlParams = new URLSearchParams(window.location.search);
    category = urlParams.get("category"); // Get the value of 'category'
    console.log(category);

    const doGet = async () => {
        const data = await get(`/api/v1/landmark/${category}`);
        console.log(data);
    };
</script>

<main>
    <div class="card">
        <div class="card__image"></div>
        <div class="card__text">
            <h1>hilite mall</h1>
            <p>location</p>
            <div class="card__stars">stars</div>
        </div>

        <button onclick={doGet}>hey</button>
    </div>
</main>
