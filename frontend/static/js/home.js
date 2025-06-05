fetch("http://localhost:8082/products", {
  method: "GET",
  credentials: "include", 
})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error(error));


const mockProducts = [
  {
    id: 1,
    title: "Cosmic Dream #1284",
    artist: "Alex Reynolds",
    image: "https://images.unsplash.com/photo-1543857778-c4a1a569e7bd?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=600&h=800&q=80",
    bid: "2.45 ETH",
    time: "Ends in 2h 45m",
    featured: true
  },
  {
    id: 2,
    title: "Neon Jungle",
    artist: "Maya Chen",
    image: "https://images.unsplash.com/photo-1579546929518-9e396f3cc809?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=600&h=800&q=80",
    bid: "1.78 ETH",
    time: "Ends in 1d 3h"
  },
  {
    id: 3,
    title: "Digital Abyss",
    artist: "Sam Rodriguez",
    image: "https://images.unsplash.com/photo-1618005182384-a83a8bd57fbe?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=600&h=800&q=80",
    bid: "3.21 ETH",
    time: "Ends in 5h 12m",
    featured: true
  },
  {
    id: 4,
    title: "Quantum Waves",
    artist: "Taylor Kim",
    image: "https://images.unsplash.com/photo-1532274402911-5a369e4c4bb5?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=600&h=800&q=80",
    bid: "0.89 ETH",
    time: "Ends in 8h 30m"
  },
  {
    id: 5,
    title: "Pixelated Reality",
    artist: "Jordan Smith",
    image: "https://images.unsplash.com/photo-1517245386807-bb43f82c33c4?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=600&h=800&q=80",
    bid: "4.12 ETH",
    time: "Ends in 3d 4h"
  },
  {
    id: 6,
    title: "Virtual Symphony",
    artist: "Casey Johnson",
    image: "https://images.unsplash.com/photo-1513151233558-d860c5398176?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=600&h=800&q=80",
    bid: "1.55 ETH",
    time: "Ends in 12h 15m"
  }
];

// Mock data for notifications
const mockNotifications = [
  {
    id: 1,
    title: "New Bid Received",
    content: "Your artwork 'Cosmic Dream' received a new bid of 2.75 ETH",
    time: "10 minutes ago",
    icon: "fas fa-gavel"
  },
  {
    id: 2,
    title: "Auction Ending Soon",
    content: "Your auction for 'Neon Jungle' is ending in 2 hours",
    time: "1 hour ago",
    icon: "fas fa-clock"
  },
  {
    id: 3,
    title: "New Follower",
    content: "Alex Reynolds started following your collection",
    time: "3 hours ago",
    icon: "fas fa-user-plus"
  },
  {
    id: 4,
    title: "Payment Received",
    content: "You received 1.2 ETH for your artwork 'Digital Abyss'",
    time: "1 day ago",
    icon: "fas fa-wallet"
  }
];

// DOM Elements
const cardContainer = document.getElementById('cardContainer');
const loadingSpinner = document.getElementById('loadingSpinner');
const notificationIcon = document.getElementById('notificationIcon');
const notificationDropdown = document.getElementById('notificationDropdown');
const notificationList = document.getElementById('notificationList');

// Initialize variables
let currentPage = 1;
const itemsPerPage = 6;
let isLoading = false;

// Simulate API fetch with delay
function fetchProducts(page) {
  return new Promise((resolve) => {
    setTimeout(() => {
      // For demo purposes, we're just using mock data
      // In a real app, you would fetch from an actual API
      const startIndex = (page - 1) * itemsPerPage;
      const endIndex = startIndex + itemsPerPage;
      
      // For infinite scroll demo, we'll just duplicate the mock data
      const allProducts = [...mockProducts];
      for (let i = 0; i < 12; i++) {
        allProducts.push(...mockProducts.map(p => ({...p, id: p.id + (i+1)*10})));
      }
      
      const products = allProducts.slice(startIndex, endIndex);
      resolve(products);
    }, 1000);
  });
}

// Render product cards
function renderProducts(products) {
  products.forEach(product => {
    const card = document.createElement('div');
    card.className = 'photo-card';
    card.innerHTML = `
      <div class="photo-img-container">
        <img src="${product.image}" alt="${product.title}" class="photo-img">
        ${product.featured ? '<div class="photo-badge">Featured</div>' : ''}
      </div>
      <div class="photo-info">
        <div class="photo-title">${product.title}</div>
        <div class="photo-details">
          <span class="label">Current Bid</span>
          <span class="bid"><i class="fas fa-ethereum"></i> ${product.bid}</span>
          <span class="time"><i class="fas fa-clock"></i> ${product.time}</span>
        </div>
      </div>
    `;
    cardContainer.appendChild(card);
  });
}

// Render notifications
function renderNotifications() {
  notificationList.innerHTML = '';
  mockNotifications.forEach(notification => {
    const notificationItem = document.createElement('div');
    notificationItem.className = 'notification-item';
    notificationItem.innerHTML = `
      <div class="notification-icon-small">
        <i class="${notification.icon}"></i>
      </div>
      <div class="notification-content">
        <h4>${notification.title}</h4>
        <p>${notification.content}</p>
        <div class="notification-time">${notification.time}</div>
      </div>
    `;
    notificationList.appendChild(notificationItem);
  });
}

// Load more products (for infinite scroll)
async function loadMoreProducts() {
  if (isLoading) return;
  
  isLoading = true;
  loadingSpinner.style.display = 'block';
  
  try {
    currentPage++;
    const products = await fetchProducts(currentPage);
    
    if (products.length > 0) {
      renderProducts(products);
    } else {
      // No more products to load
      window.removeEventListener('scroll', handleScroll);
    }
  } catch (error) {
    console.error('Error loading products:', error);
  } finally {
    isLoading = false;
    loadingSpinner.style.display = 'none';
  }
}

// Scroll handler for infinite scrolling
function handleScroll() {
  const { scrollTop, scrollHeight, clientHeight } = document.documentElement;
  
  // Load more when 100px from the bottom
  if (scrollTop + clientHeight >= scrollHeight - 100) {
    loadMoreProducts();
  }
}

// Toggle notification dropdown
function toggleNotifications() {
  notificationDropdown.classList.toggle('show');
}

// Close dropdown when clicking outside
document.addEventListener('click', (e) => {
  if (!notificationIcon.contains(e.target)) {
    notificationDropdown.classList.remove('show');
  }
});

// Initialize the page
async function init() {
  // Render initial products
  const initialProducts = await fetchProducts(currentPage);
  renderProducts(initialProducts);
  
  // Render notifications
  renderNotifications();
  
  // Add event listeners
  window.addEventListener('scroll', handleScroll);
  notificationIcon.addEventListener('click', toggleNotifications);
  
  // Add demo functionality to notification badge
  const notificationBadge = document.querySelector('.notification-badge');
  notificationBadge.addEventListener('click', function(e) {
    e.stopPropagation();
    this.textContent = '0';
    this.style.backgroundColor = '#4ecdc4';
    
    setTimeout(() => {
      this.textContent = '3';
      this.style.backgroundColor = '#ff4757';
    }, 3000);
  });
}

// Start the app
init();