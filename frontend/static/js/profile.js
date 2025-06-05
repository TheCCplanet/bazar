const profilePicInput = document.getElementById('profilePic');
const preview = document.getElementById('preview');

profilePicInput.addEventListener('change', (event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = e => {
      preview.src = e.target.result;
      preview.style.display = 'block';
    };
    reader.readAsDataURL(file);
  } else {
    preview.src = '';
    preview.style.display = 'none';
  }
});

// Optional: handle submit
document.getElementById('registerForm').addEventListener('submit', e => {
  e.preventDefault();
  alert("Registration submitted!");
  // You can now send the form data to your backend using fetch/FormData
});
