var startLocation = [49.257083, -122.953722] // vancouver
var zoom = 13

var mapOptions = {
    inertia: true, // Enable inertia
    inertiaDeceleration: 5000, // Adjust deceleration rate (default: 3000)
    inertiaMaxSpeed: Infinity, // Maximum speed during inertia (default: Infinity)
    // Optional: You might want to limit the zoom level during inertia:
    inertiaThreshold: 60, // Adjust the zoom threshold as needed
    keyboard: true,
    keyboardPanDelta: 200
}
var map = L.map('map', mapOptions).setView(startLocation, zoom);

L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">Findr</a> contributors'
}).addTo(map);

L.control.scale().addTo(map);


function addMapMarker(location, text) {
    L.marker(location).addTo(map)
        .bindPopup(text)
        .openPopup();
}

function onMapClick(e) {
    console.log("You clicked the map at " + e.latlng);

    const url = '/api/clicks'; // API endpoint
    // const data = 'myDATA'; // Data to send

    // postData(url, e.latlng);
}

map.on('click', onMapClick);

async function getUsers(url = '/api/users') {
    var schRadius = document.getElementById("search-radius").value
    // console.log(schRadius)
    postData(url, schRadius)

    try {
        const response = await fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json' // Adjust content type if necessary
            },
            // body: data
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const responseData = await response.json(); // or response.json() if the API returns JSON
        console.log("Get locations Success");
        console.log(responseData);

        for (let i = 0; i < responseData.length; i++) {

            if (responseData[i]['dist'] != null) {
                var markerText = "id: " + responseData[i]['id'] + "<br/>Name: " + responseData[i]['name'] + "<br/>LatLng: " + responseData[i]['latitude'] + ", " + responseData[i]['longitude'] + "<br/>Distance: " + responseData[i]['dist'] + " metres"
                addMapMarker([responseData[i]['latitude'], responseData[i]['longitude']], markerText)
            }
        }
        return responseData;
    } catch (error) {
        console.error('Error:', error);
    }
}


async function postData(url = '', data = '') {
    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'text/plain' // Adjust content type if necessary
            },
            body: data
        });

        if (!response.ok) {
            // Handle HTTP errors (e.g., 404, 500)
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const responseData = await response.text(); // or response.json() if the API returns JSON
        console.log('Success:', responseData);
        return responseData;
    } catch (error) {
        console.error('Error:', error);
    }
}

getUsers()

// let { data: Users, error } = await supabase
// .from('Users')
// .select('*')

// console.log(data)
// console.log(error)


// function adjustMapPadding() {
//     const navbar = document.querySelector('.navbar');
//     const navbarHeight = navbar.offsetHeight; // Get navbar's calculated height
//     const map = document.getElementById('map');
//     const viewportHeight = window.innerHeight; // Get viewport height

//     // Set padding-top to accommodate the navbar
//     map.style.paddingTop = navbarHeight + 'px';

//     // Calculate and set the map's height dynamically
//     const mapHeight = viewportHeight - navbarHeight;
//     map.style.height = mapHeight + 'px';
// }

// // Call the function on initial load and when the window is resized
// window.addEventListener('load', adjustMapPadding);
// // window.addEventListener('resize', adjustMapPadding);