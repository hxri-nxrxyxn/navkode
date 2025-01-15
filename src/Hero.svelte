<script>
    import { Link } from "svelte-routing";
    const categories = ["Malls", "Monuments"];
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
    <h1>Choose your destination</h1>

    <div class="logo__box"></div>

    <div class="option__box">
        <div onclick={goLeft}>left</div>
        <div class="option__text">
            <h2>{category}</h2>
        </div>
        <div onclick={goRight}>right</div>
    </div>
</main>

<style>
</style>
