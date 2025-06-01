    // Adding hover effects to cards
    document.querySelectorAll('.photo-card').forEach(card => {
      card.addEventListener('mouseenter', () => {
        card.style.transform = 'translateY(-5px)';
        card.style.boxShadow = '0 10px 20px rgba(0, 0, 0, 0.3)';
        card.style.borderColor = '#4ecdc4';
      });
      
      card.addEventListener('mouseleave', () => {
        card.style.transform = '';
        card.style.boxShadow = '';
        card.style.borderColor = '#1a2b2e';
      });
    });
    
    // Notification click handler
    document.querySelector('.notification-icon').addEventListener('click', function() {
      const badge = this.querySelector('.notification-badge');
      badge.textContent = '0';
      badge.style.backgroundColor = '#4ecdc4';
      
      // Reset after 3 seconds
      setTimeout(() => {
        badge.textContent = '3';
        badge.style.backgroundColor = '#ff4757';
      }, 3000);
    });