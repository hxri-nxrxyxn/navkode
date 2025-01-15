import { post } from "./fetch.js"

const query = `
[out:json];
area["name"="Kozhikode"];
(
  node["restaurant"](area);
  way["restaurant"](area);
  relation["restaurant"](area);
);
out body;
`;

// Overpass API endpoint
const overpassUrl = 'https://overpass-api.de/api/interpreter';
  

// Fetch parking data
fetch(overpassUrl, {
  method: 'POST',
  headers: {
    'Content-Type': 'application/x-www-form-urlencoded',
  },
  body: `data=${encodeURIComponent(query)}`,
})
  .then(response => response.json())
  .then(data => {
    console.log(data.elements);
    data.elements.map((element, index) => {

        const doPost = async () => {
            let datab = {
                "lon" : (element.lon).toString(),
                "lat" : (element.lat).toString(),
                "name" : (element.tags.name == undefined ? "" : element.tags.name),
            }
            console.log(element)
            try {
                const response = await fetch("http://127.0.0.1:dsf", {
                  method: "POST",
                  headers: {
                    'Content-Type': 'application/json',
                  },
                //   causes preflight requests -> cors issues
                  body: JSON.stringify(datab),
                });
                if (!response.ok) {
                  throw new Error(`Network response was not ok (${response.status})`);
                }
                return await response.json();
              } catch (error) {
                console.error("Error creating data:", error);
        }
    }
        doPost()
    })
  })
  .catch(error => {
    console.error('Error fetching parking data:', error);
  });
