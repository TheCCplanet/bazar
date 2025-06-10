document.addEventListener('DOMContentLoaded', function() {
  const auctionToggle = document.getElementById('auctionToggle');
  const auctionFields = document.getElementById('auctionFields');
  const uploadArea = document.getElementById('uploadArea');
  const fileInput = document.getElementById('fileInput');
  const imagePreviews = document.getElementById('imagePreviews');
  let primaryImageId = null;

  // آرایه برای نگهداری فایل‌های انتخاب شده
  let selectedFiles = [];

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
      // فایل‌های درگ شده رو به فایل‌های انتخاب شده اضافه کن
      handleFiles(e.dataTransfer.files);
      fileInput.value = ''; // مقدار input رو پاک کن تا دوباره انتخاب بتونی انجام بدی
    }
  });

  // تابع اصلی برای اضافه کردن فایل‌ها
  function handleFileSelect() {
    const files = fileInput.files;
    handleFiles(files);
    fileInput.value = ''; // پاک کردن مقدار input بعد از خواندن فایل‌ها
  }

  // این تابع فایل‌ها را دریافت و به selectedFiles اضافه می‌کند و preview می‌سازد
  function handleFiles(files) {
    for (let i = 0; i < files.length; i++) {
      const file = files[i];

      // چک کردن نوع فایل (فقط عکس)
      if (!file.type.match('image.*')) {
        alert('Please select only image files.');
        continue;
      }

      // چک کردن حجم فایل (حداکثر 10 مگابایت)
      if (file.size > 10 * 1024 * 1024) {
        alert('File size exceeds 10MB: ' + file.name);
        continue;
      }

      // جلوگیری از اضافه شدن فایل تکراری (بر اساس نام و حجم)
      if (selectedFiles.some(f => f.name === file.name && f.size === file.size)) {
        alert('File already selected: ' + file.name);
        continue;
      }

      // اضافه کردن به آرایه
      selectedFiles.push(file);

      // ساخت preview برای عکس
      const reader = new FileReader();

      reader.onload = function(e) {
        const imageId = 'img-' + Date.now() + '-' + Math.floor(Math.random() * 1000);
        const isFirstImage = imagePreviews.children.length === 0;

        const preview = document.createElement('div');
        preview.className = 'image-preview' + (isFirstImage ? ' primary' : '');
        preview.dataset.id = imageId;
        preview.dataset.fileName = file.name;
        preview.dataset.fileSize = file.size;
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

        // اولین عکس به طور پیش‌فرض primary می‌شود
        if (isFirstImage) {
          primaryImageId = imageId;
        }

        // دکمه ست کردن به عنوان primary
        preview.querySelector('.set-primary').addEventListener('click', function(e) {
          e.stopPropagation();
          setPrimaryImage(imageId);
        });

        // دکمه حذف عکس
        preview.querySelector('.delete').addEventListener('click', function(e) {
          e.stopPropagation();
          preview.remove();

          // حذف فایل مرتبط از selectedFiles
          selectedFiles = selectedFiles.filter(f => !(f.name === preview.dataset.fileName && f.size == preview.dataset.fileSize));

          // اگر عکس primary حذف شد، عکس بعدی primary شود یا null شود
          if (primaryImageId === imageId) {
            if (imagePreviews.children.length > 0) {
              setPrimaryImage(imagePreviews.children[0].dataset.id);
            } else {
              primaryImageId = null;
            }
          }
        });

        // کلیک روی preview برای ست کردن به عنوان primary
        preview.addEventListener('click', function(e) {
          if (e.target.tagName !== 'BUTTON') {
            setPrimaryImage(imageId);
          }
        });
      };

      reader.readAsDataURL(file);
    }
  }

  // تابع برای ست کردن عکس primary
  function setPrimaryImage(imageId) {
    document.querySelectorAll('.image-preview').forEach(preview => {
      preview.classList.remove('primary');
    });

    const selectedPreview = document.querySelector(`.image-preview[data-id="${imageId}"]`);
    if (selectedPreview) {
      selectedPreview.classList.add('primary');
      primaryImageId = imageId;
    }
  }

  // ارسال فرم
  document.getElementById('productForm').addEventListener('submit', async function(e) {
    e.preventDefault();

    const productName = document.getElementById('productName').value;
    const category = document.getElementById('productCategory').value;
    const description = document.getElementById('productDescription').value;
    const price = document.getElementById('productPrice').value;
    const quantity = document.getElementById('quantity')?.value || '1';

    if (!productName || !category || !price) {
      alert('Please fill in all required fields.');
      return;
    }

    const images = document.querySelectorAll('.image-preview');
    if (images.length < 3) {
      alert('Please upload at least 3 images.');
      return;
    }

    if (!primaryImageId) {
      alert('Please select a primary image.');
      return;
    }

    if (auctionToggle.checked && !document.getElementById('startingBid').value) {
      alert('Please enter a starting bid for the auction.');
      return;
    }

    const formData = new FormData();

    formData.append('productName', productName);
    formData.append('category', category);
    formData.append('description', description);
    formData.append('price', price);
    formData.append('isAuction', auctionToggle.checked);
    formData.append('primaryImageId', primaryImageId);
    formData.append('imageCount', images.length);
    formData.append('quantity', quantity);

    if (auctionToggle.checked) {
      formData.append('startingBid', document.getElementById('startingBid').value);
      formData.append('auctionDuration', document.getElementById('auctionDuration').value);
      formData.append('reservePrice', document.getElementById('reservePrice').value || '');
    }

    // ارسال فایل‌ها از آرایه selectedFiles به جای fileInput.files
    for (let i = 0; i < selectedFiles.length; i++) {
      formData.append('images', selectedFiles[i]);
    }

    try {
      const res = await fetch('http://localhost:8082/upload', {
        method: 'POST',
        body: formData,
        credentials: 'include'
      });

      if (!res.ok) {
        const err = await res.text();
        throw new Error(`Server responded with error: ${err}`);
      }

      alert('Product uploaded successfully!');
      // بعد از ارسال موفقیت‌آمیز می‌تونی فرم و preview ها رو پاک کنی اگر خواستی:
      // selectedFiles = [];
      // imagePreviews.innerHTML = '';
      // primaryImageId = null;
      // document.getElementById('productForm').reset();
    } catch (error) {
      console.error('Upload failed:', error);
      alert('Failed to upload product.');
    }
  });
});
