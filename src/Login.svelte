<script>
    import { Link } from "svelte-routing";
    import Nav from "./Nav.svelte";
    import { post } from "./fetch";

    let email = $state("");
    let password = $state("");

    const doPost = async () => {
        let data = {
            email: email,
            password: password,
        };
        console.log(data);
        let res = await post("/api/v1/login", data);
        console.log(res);
    };
</script>

<main>
    <Nav />
    <h1>Login with ease</h1>
    <div class="input__box">
        <div class="input__line">
            <label>Email</label>
            <input
                placeholder="Enter your mail address"
                type="email"
                id="email"
                bind:value={email}
            />
        </div>
        <div class="input__line">
            <label>Password</label>
            <input
                placeholder="Enter your password"
                type="password"
                id="password"
                bind:value={password}
            />
        </div>
        <div class="placeholder"></div>
        <button class="solid" onclick={doPost}>LOGIN</button>
    </div>
</main>

<style>
    .input__line {
        margin: 1rem 0;
    }
    label {
        font-size: 0.8rem;
        font-weight: 600;
        margin-left: 1rem;
    }
    input {
        display: block;
        margin: 0.5rem 0;
        width: 100%;
        border: 0;
        background-color: var(--color-shade);
        padding: 1rem;
        border-radius: 1rem;
        border: 2px solid transparent;
    }
    input:focus {
        border: 2px solid var(--color-primary);
    }
    .placeholder {
        height: 40vh;
    }
</style>
