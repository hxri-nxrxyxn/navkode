<script>
    import { Link } from "svelte-routing";
    import Chevron from "./assets/angle-small-right.svg";
    import Nav from "./Nav.svelte";
    const categories = [
        "Mall",
        "Beach",
        "Park",
        "Shopping",
        "Restaurant",
        "Historical",
        "Nature",
        "Education",
    ];
    let category = $state("");

    let i = 0;
    category = categories[i];

    const goLeft = () => {
        if (i != 0) {
            console.log(i);
            category = categories[--i];
        }
    };
    const goRight = () => {
        if (i < categories.length - 1) {
            console.log(i);
            category = categories[++i];
        }
    };
</script>

<main>
    <Nav />
    <h1>Choose your destination</h1>

    <div class="logo__box"></div>

    <div class="placeholder"></div>

    <div class="option__box">
        <div onclick={goLeft}>
            <div class="option__image">
                <img src={Chevron} style="transform: scaleX(-1);" />
            </div>
        </div>
        <div class="option__text">
            <h2>{category}</h2>
        </div>
        <div onclick={goRight}>
            <div class="option__image">
                <img src={Chevron} />
            </div>
        </div>
    </div>

    <Link to={`/list?category=${category}`}>
        <button class="solid">Continue</button>
    </Link>
</main>

<style>
    .option__box {
        display: flex;
        justify-content: space-between;
        padding: 2rem 0;
    }
    .option__box div {
        display: flex;
        justify-content: center;
        align-items: center;
    }
    .option__image {
        background: var(--color-shade);
        padding: 0.5rem;
        border-radius: 50%;
    }
    .option__box img {
        height: 40px;
        filter: invert(41%) sepia(1%) saturate(0%) hue-rotate(173deg)
            brightness(94%) contrast(87%);
    }
    .placeholder {
        height: 50vh;
    }
</style>
