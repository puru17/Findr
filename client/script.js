// Map configuration
const MAP_CONFIG = {
    startLocation: [49.257083, -122.953722], // vancouver
    zoom: 13,
    options: {
        inertia: true,
        inertiaDeceleration: 5000,
        inertiaMaxSpeed: Infinity,
        inertiaThreshold: 60,
        keyboard: true,
        keyboardPanDelta: 200
    }
};

// API endpoints
const API = {
    users: '/api/users',
    clicks: '/api/clicks'
};

// Map instance
let map;
let markers = [];

// Initialize map
function initMap() {
    map = L.map('map', MAP_CONFIG.options).setView(MAP_CONFIG.startLocation, MAP_CONFIG.zoom);
    
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">Findr</a> contributors'
    }).addTo(map);

    L.control.scale().addTo(map);
    map.on('click', onMapClick);
}

// Add marker to map
function addMapMarker(location, text) {
    const marker = L.marker(location)
        .addTo(map)
        .bindPopup(text)
        .openPopup();
    markers.push(marker);
    return marker;
}

// Clear all markers
function clearMarkers() {
    markers.forEach(marker => marker.remove());
    markers = [];
}

// Handle map click
function onMapClick(e) {
    console.log("You clicked the map at", e.latlng);
    // TODO: Implement click handling
}

// Show loading state
function setLoading(isLoading) {
    const button = document.getElementById('qry-btn');
    if (isLoading) {
        button.disabled = true;
        button.textContent = 'Loading...';
    } else {
        button.disabled = false;
        button.textContent = 'Populate Map';
    }
}

// Show error message
function showError(message) {
    // TODO: Implement proper error UI
    console.error(message);
    alert(message);
}

// Validate search radius
function validateRadius(radius) {
    const num = parseFloat(radius);
    if (isNaN(num)) {
        throw new Error('Please enter a valid number');
    }
    if (num <= 0) {
        throw new Error('Radius must be greater than 0');
    }
    if (num > 41) {
        throw new Error('Radius must be less than 41 km');
    }
    return num;
}

// Update search radius
async function updateSearchRadius(radius) {
    try {
        const response = await fetch(API.users, {
            method: 'POST',
            headers: {
                'Content-Type': 'text/plain'
            },
            body: radius.toString()
        });

        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.error || 'Failed to update search radius');
        }

        return await response.json();
    } catch (error) {
        showError(error.message);
        throw error;
    }
}

// Fetch and display users
async function getUsers() {
    try {
        setLoading(true);
        clearMarkers();

        const radiusInput = document.getElementById("search-radius");
        const radius = validateRadius(radiusInput.value);
        
        await updateSearchRadius(radius);

        const response = await fetch(API.users);
        if (!response.ok) {
            console.error('Failed to fetch locations');
            return;
        }

        const locations = await response.json();
        console.log("Locations fetched successfully:", locations);

        locations.forEach(location => {
            if (location.dist != null) {
                const markerText = `
                    ID: ${location.id}<br/>
                    Name: ${location.name}<br/>
                    Location: ${location.latitude}, ${location.longitude}<br/>
                    Distance: ${location.dist} metres
                `;
                addMapMarker([location.latitude, location.longitude], markerText);
            }
        });
    } catch (error) {
        console.error('Error fetching users:', error);
    } finally {
        setLoading(false);
    }
}

// Handle login button click
function handleLoginClick() {
    window.location.href = '/login';
}

// Initialize application
document.addEventListener('DOMContentLoaded', () => {
    initMap();
    getUsers();
    
    // Add click handler for login button
    const loginBtn = document.getElementById('login-btn');
    if (loginBtn) {
        loginBtn.addEventListener('click', handleLoginClick);
    }
});

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