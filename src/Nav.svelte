<script>
    import Cursor from "./assets/cursor-finger.svg";
    import House from "./assets/house-crack.svg";
</script>

<main>
    <div class="nav">
        <div class="nav__main">
            <img src={House} alt="" class="super" />
        </div>
        <div class="nav__placeholder">hi</div>
        <div class="nav__sub">
            <img src={Cursor} alt="" />
        </div>
    </div>
</main>

<style>
    .nav {
        display: flex;
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
</style>
