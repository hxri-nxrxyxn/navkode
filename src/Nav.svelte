<script>
    import Cursor from "./assets/cursor-finger.svg";
    import House from "./assets/house-crack.svg";
    import Logi from "./assets/insert-alt.svg";
</script>

<main>
    <div class="nav">
        <div class="nav__main">
            <img src={House} alt="" class="super" />
        </div>
        <div class="nav__placeholder">hi</div>
        <div class="nav__sub">
            <img src={Cursor} alt="" />
            <img src={Logi} alt="" />
        </div>
    </div>
</main>

<style>
    .nav {
        display: flex !important;
        padding-bottom: 1rem;
    }
    .nav img {
        height: 20px;
        filter: invert(50%) sepia(55%) saturate(014%) hue-rotate(331deg)
            brightness(100%) contrast(100%) opacity(0.3);
    }
    .nav img.super {
        filter: invert(50%) sepia(55%) saturate(4014%) hue-rotate(331deg)
            brightness(100%) contrast(100%);
    }
    .nav__placeholder {
        color: transparent;
        flex: 1;
    }
    .nav__sub img {
        margin-left: 0.5rem;
    }
</style>
