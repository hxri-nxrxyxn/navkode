<script>
    import { get } from "./fetch";
    import { Link } from "svelte-routing";
    import Nav from "./Nav.svelte";

    const urlParams = new URLSearchParams(window.location.search);
    const landmarkId = urlParams.get("id");

    let landmarkPromise = get(`/api/v1/landmark/${landmarkId}`);
    let parkingPromise = get("/api/v1/parking");
    let parkingSpots = [];
    let nearestParking = null;
    let minDistance = Infinity;

    const calculateDistance = (lat1, lon1, lat2, lon2) => {
        lat1 = parseFloat(lat1);
        lon1 = parseFloat(lon1);
        lat2 = parseFloat(lat2);
        lon2 = parseFloat(lon2);

        const R = 6371; // Radius of the earth in km
        const dLat = deg2rad(lat2 - lat1);
        const dLon = deg2rad(lon2 - lon1);
        const a =
            Math.sin(dLat / 2) * Math.sin(dLat / 2) +
            Math.cos(deg2rad(lat1)) *
                Math.cos(deg2rad(lat2)) *
                Math.sin(dLon / 2) *
                Math.sin(dLon / 2);
        const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));
        const d = R * c; // Distance in km
        return d;
    };

    const deg2rad = (deg) => {
        return deg * (Math.PI / 180);
    };

    Promise.all([landmarkPromise, parkingPromise])
        .then(([landmarkResponse, parkingResponse]) => {
            parkingSpots = parkingResponse.data;

            if (!parkingSpots || parkingSpots.length === 0) {
                console.log("No parking spots found.");
                return;
            }

            const landmarkLat = parseFloat(landmarkResponse.data.lat);
            const landmarkLon = parseFloat(landmarkResponse.data.lon);

            parkingSpots.forEach((parkingSpot) => {
                const parkingLat = parseFloat(parkingSpot.lat);
                const parkingLon = parseFloat(parkingSpot.lon);

                const distance = calculateDistance(
                    landmarkLat, // Landmark Latitude
                    landmarkLon, // Landmark Longitude
                    parkingLat, // Parking Latitude
                    parkingLon, // Parking Longitude
                );

                if (distance < minDistance) {
                    minDistance = distance;
                    nearestParking = parkingSpot;
                }
            });

            if (nearestParking) {
                console.log("Nearest parking:", nearestParking);
                console.log("Distance:", minDistance);
            }
        })
        .catch((error) => {
            console.error("Error fetching data:", error);
        });
</script>

<main>
    <Nav />

    {#await landmarkPromise}
        <p>Loading landmark data...</p>
    {:then landmarkResponse}
        <h1>
            Here's the shortest path to <Link
                to={`/landmarks/${landmarkResponse.data.id}`}
                >{landmarkResponse.data.name}</Link
            >
        </h1>
        <div class="card">
            <div class="card__text">
                <p>{landmarkResponse.data.name}</p>
                <h2>{landmarkResponse.data.lon} E</h2>
                <h2>{landmarkResponse.data.lat} N</h2>
            </div>
            {#if nearestParking}
                <div class="card__text">
                    <p>Nearest Parking Space</p>
                    <h2>{nearestParking.lat} E</h2>
                    <h2>{nearestParking.lon} N</h2>
                </div>
                <div class="card__text">
                    <p>About</p>
                    <h2>{(minDistance / 1000).toFixed(2)} km</h2>
                </div>
                <div
                    class="card__text"
                    style="display: flex; align-items: center; justify-content: center;"
                >
                    <a
                        href={`https://google.com/maps?q=${nearestParking.lat},${nearestParking.lon}`}
                    >
                        <button>Open Maps</button>
                    </a>
                </div>
            {:else if parkingSpots && parkingSpots.length === 0}
                <p>No parking spots found</p>
            {:else}
                <p>Finding Nearest Parking</p>
            {/if}
        </div>
    {:catch error}
        <p style="color: red;">
            Error loading data: {error.message ||
                "An unexpected error occurred."}
        </p>
    {/await}
</main>

<style>
    .card {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
    }
    .card__text {
        margin: 2rem 0;
    }
    .card button {
        background: var(--color-primary);
        padding: 1rem 2rem;
        border: 0;
        border-radius: 2rem;
        color: white;
        font-family: "Inter";
        text-transform: uppercase;
        font-weight: 600;
    }
    p {
        color: var(--color-light);
    }
</style>
