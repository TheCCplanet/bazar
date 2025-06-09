document.addEventListener('DOMContentLoaded', function() {
    const auctionToggle = document.getElementById('auctionToggle');
    const auctionFields = document.getElementById('auctionFields');
    const uploadArea = document.getElementById('uploadArea');
    const fileInput = document.getElementById('fileInput');
    const imagePreviews = document.getElementById('imagePreviews');
    let primaryImageId = null;
    
    // Auction toggle handler
    auctionToggle.addEventListener('change', function() {
      if (this.checked) {
        auctionFields.classList.add('active');
        document.getElementById('startingBid').required = true;
        document.getElementById('auctionDuration').required = true;
      } else {
        auctionFields.classList.remove('active');
        document.getElementById('startingBid').required = false;
        document.getElementById('auctionDuration').required = false;
      }
    });
    
    // Upload area click handler

    uploadArea.addEventListener('click', () => {
      fileInput.click();
    });
    
    // File input change handler
    fileInput.addEventListener('change', handleFileSelect);
    
    // Handle drag and drop
    uploadArea.addEventListener('dragover', (e) => {
      e.preventDefault();
      uploadArea.style.borderColor = 'var(--primary)';
      uploadArea.style.backgroundColor = 'rgba(30, 50, 55, 0.5)';
    });
    
    uploadArea.addEventListener('dragleave', () => {
      uploadArea.style.borderColor = 'rgba(255, 255, 255, 0.2)';
      uploadArea.style.backgroundColor = 'rgba(20, 30, 35, 0.4)';
    });
    
    uploadArea.addEventListener('drop', (e) => {
      e.preventDefault();
      uploadArea.style.borderColor = 'rgba(255, 255, 255, 0.2)';
      uploadArea.style.backgroundColor = 'rgba(20, 30, 35, 0.4)';
      
      if (e.dataTransfer.files.length) {
        fileInput.files = e.dataTransfer.files;
        handleFileSelect();
      }
    });
    
    // Handle file selection
    function handleFileSelect() {
      const files = fileInput.files;
      
      if (files.length > 0) {
        for (let i = 0; i < files.length; i++) {
          const file = files[i];
          
          // Check if file is an image
          if (!file.type.match('image.*')) {
            alert('Please select only image files.');
            continue;
          }
          
          // Check file size (max 10MB)
          if (file.size > 10 * 1024 * 1024) {
            alert('File size exceeds 10MB: ' + file.name);
            continue;
          }
          
          // Create preview
          const reader = new FileReader();
          
          reader.onload = function(e) {
            const imageId = 'img-' + Date.now() + '-' + Math.floor(Math.random() * 1000);
            const isFirstImage = imagePreviews.children.length === 0;
            
            const preview = document.createElement('div');
            preview.className = 'image-preview' + (isFirstImage ? ' primary' : '');
            preview.dataset.id = imageId;
            preview.innerHTML = `
              <img src="${e.target.result}" alt="Preview">
              <div class="primary-badge">Primary</div>
              <div class="actions">
                <button type="button" class="set-primary" title="Set as primary">
                  <i class="fas fa-star"></i>
                </button>
                <button type="button" class="delete" title="Remove image">
                  <i class="fas fa-trash"></i>
                </button>
              </div>
            `;
            
            imagePreviews.appendChild(preview);
            
            // Set first image as primary by default
            if (isFirstImage) {
              primaryImageId = imageId;
            }
            
            // Add event listeners to buttons
            preview.querySelector('.set-primary').addEventListener('click', function(e) {
              e.stopPropagation();
              setPrimaryImage(imageId);
            });
            
            preview.querySelector('.delete').addEventListener('click', function(e) {
              e.stopPropagation();
              preview.remove();
              if (primaryImageId === imageId) {
                if (imagePreviews.children.length > 0) {
                  setPrimaryImage(imagePreviews.children[0].dataset.id);
                } else {
                  primaryImageId = null;
                }
              }
            });
            
            // Click on preview to set as primary
            preview.addEventListener('click', function(e) {
              if (e.target.tagName !== 'BUTTON') {
                setPrimaryImage(imageId);
              }
            });
          };
          
          reader.readAsDataURL(file);
        }
      }
    }
    
    // Set an image as primary
    function setPrimaryImage(imageId) {
      // Remove primary class from all images
      document.querySelectorAll('.image-preview').forEach(preview => {
        preview.classList.remove('primary');
      });
      
      // Add primary class to selected image
      const selectedPreview = document.querySelector(`.image-preview[data-id="${imageId}"]`);
      if (selectedPreview) {
        selectedPreview.classList.add('primary');
        primaryImageId = imageId;
      }
    }
    // Form submission handler
    document.getElementById('productForm').addEventListener('submit', async function(e) {
      e.preventDefault();
      
      // Validate form
      const productName = document.getElementById('productName').value;
      const category = document.getElementById('productCategory').value;
      const price = document.getElementById('productPrice').value;
      
      if (!productName || !category || !price) {
        alert('Please fill in all required fields.');
        return;
      }
      
      // Validate images
      const images = document.querySelectorAll('.image-preview');
      if (images.length < 3) {
        alert('Please upload at least 3 images.');
        return;
      }
      
      if (!primaryImageId) {
        alert('Please select a primary image.');
        return;
      }
      
      // Check auction fields if auction is enabled
      if (auctionToggle.checked) {
        const startingBid = document.getElementById('startingBid').value;
        if (!startingBid) {
          alert('Please enter a starting bid for the auction.');
          return;
        }
      }
      
      // Form data would be sent to server here
      const formData = {
        productName,
        category,
        description: document.getElementById('productDescription').value,
        price,
        isAuction: auctionToggle.checked,
        primaryImage: primaryImageId,
        imageCount: images.length
      };

      
      if (auctionToggle.checked) {
        formData.startingBid = document.getElementById('startingBid').value;
        formData.auctionDuration = document.getElementById('auctionDuration').value;
        formData.reservePrice = document.getElementById('reservePrice').value;
      }
      for (let i = 0 ; i < images.length;i++) {
        this.append('images',images[i])
      }
      const res = await fetch('http://localhost:8082/upload',{
        method: 'POST',
        body: formData
          
        }
      );
      
      console.log('Form submitted:', formData);
      
      // Show success message
      alert('Product uploaded successfully!');
    });
  });