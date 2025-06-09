document.addEventListener('DOMContentLoaded', function() {
  // Sidebar toggle
  const toggleBtn = document.querySelector('.toggle-sidebar');
  const sidebar = document.querySelector('.sidebar');
  
  toggleBtn.addEventListener('click', function() {
    sidebar.classList.toggle('collapsed');
  });
  
  // Dashboard navigation
  const menuLinks = document.querySelectorAll('.menu-link');
  const dashboardSections = document.querySelectorAll('.dashboard-section');
  
  menuLinks.forEach(link => {
    link.addEventListener('click', function(e) {
      e.preventDefault();
      
      // Remove active class from all links
      menuLinks.forEach(l => l.classList.remove('active'));
      
      // Add active class to clicked link
      this.classList.add('active');
      
      // Hide all sections
      dashboardSections.forEach(section => {
        section.classList.remove('active');
      });
      
      // Show target section
      const targetSection = this.getAttribute('data-section');
      if (targetSection) {
        document.getElementById(targetSection).classList.add('active');
      }
    });
  });
  
  // Tab navigation for bids section
  const tabs = document.querySelectorAll('.tab');
  tabs.forEach(tab => {
    tab.addEventListener('click', function() {
      tabs.forEach(t => t.classList.remove('active'));
      this.classList.add('active');
    });
  });
  
  // Bid functionality
  const bidButtons = document.querySelectorAll('.bid-btn');
  bidButtons.forEach(button => {
    button.addEventListener('click', function() {
      const input = this.previousElementSibling;
      const bidAmount = parseFloat(input.value);
      
      if (isNaN(bidAmount)) {
        alert('Please enter a valid bid amount');
        return;
      }
      
      const currentBid = parseFloat(this.closest('.bid-section').querySelector('.detail-value').textContent.split(' ')[0]);
      
      if (bidAmount <= currentBid) {
        alert(`Your bid must be higher than the current bid of ${currentBid} ETH`);
        return;
      }
      
      // In a real app, this would be an API call
      const card = this.closest('.order-card');
      card.querySelector('.detail-value').textContent = bidAmount.toFixed(2) + ' ETH';
      card.querySelector('.current-bid strong').textContent = bidAmount.toFixed(2) + ' ETH';
      card.querySelector('.current-bid div:last-child').textContent = 'by @your_username';
      
      // Show success message
      const successMsg = document.createElement('div');
      successMsg.textContent = `Bid of ${bidAmount.toFixed(2)} ETH placed successfully!`;
      successMsg.style.color = 'var(--success)';
      successMsg.style.marginTop = '10px';
      successMsg.style.fontWeight = '500';
      this.parentElement.appendChild(successMsg);
      
      // Disable the button temporarily
      this.disabled = true;
      this.textContent = 'Bid Placed';
      this.style.background = 'var(--success)';
      
      // Reset after 3 seconds
      setTimeout(() => {
        successMsg.remove();
        this.disabled = false;
        this.textContent = 'Bid';
        this.style.background = 'var(--primary)';
      }, 3000);
    });
  });
  
  // Initialize charts
  const salesCtx = document.getElementById('salesChart').getContext('2d');
  const salesChart = new Chart(salesCtx, {
    type: 'line',
    data: {
      labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct'],
      datasets: [{
        label: 'ETH Revenue',
        data: [1.2, 1.8, 2.5, 3.1, 2.7, 3.5, 4.2, 3.8, 4.5, 5.2],
        borderColor: '#4ecdc4',
        backgroundColor: 'rgba(78, 205, 196, 0.1)',
        borderWidth: 2,
        fill: true,
        tension: 0.3
      }]
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        legend: {
          display: false
        }
      },
      scales: {
        y: {
          beginAtZero: true,
          grid: {
            color: 'rgba(255, 255, 255, 0.1)'
          },
          ticks: {
            color: 'rgba(255, 255, 255, 0.7)'
          }
        },
        x: {
          grid: {
            color: 'rgba(255, 255, 255, 0.1)'
          },
          ticks: {
            color: 'rgba(255, 255, 255, 0.7)'
          }
        }
      }
    }
  });
  
  const revenueCtx = document.getElementById('revenueChart').getContext('2d');
  const revenueChart = new Chart(revenueCtx, {
    type: 'doughnut',
    data: {
      labels: ['Digital Art', 'Collectibles', 'Music', 'Photography', 'Domains'],
      datasets: [{
        data: [45, 25, 15, 10, 5],
        backgroundColor: [
          '#4ecdc4',
          '#ff6b6b',
          '#ffe66d',
          '#1a759f',
          '#7e57c2'
        ],
        borderWidth: 0
      }]
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        legend: {
          position: 'right',
          labels: {
            color: 'rgba(255, 255, 255, 0.7)',
            padding: 15
          }
        }
      }
    }
  });
});